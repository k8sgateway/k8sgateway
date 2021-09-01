package check

import (
	"context"
	"errors"
	"fmt"
	"time"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"

	"github.com/hashicorp/go-multierror"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/printers"
	ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	rlopts "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ratelimit"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	CrdNotFoundErr = func(crdName string) error {
		return eris.Errorf("%s CRD has not been registered", crdName)
	}
)

// contains method
func doesNotContain(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return false
		}
	}
	return true
}

func RootCmd(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   constants.CHECK_COMMAND.Use,
		Short: constants.CHECK_COMMAND.Short,
		Long:  "usage: glooctl check [-o FORMAT]",
		RunE: func(cmd *cobra.Command, args []string) error {

			if !opts.Top.Output.IsTable() && !opts.Top.Output.IsJSON() {
				return errors.New("Invalid output type")
			}

			err := CheckResources(opts)

			if err != nil {
				// Not returning error here because this shouldn't propagate as a standard CLI error, which prints usage.
				if opts.Top.Output.IsTable() {
					return err
				}
			} else {
				printers.AppendResponse("", "", fmt.Sprintf("No problems detected.\n"), "", opts.Top.Output)
			}

			CheckMulticlusterResources(opts)

			if opts.Top.Output.IsJSON() {
				printers.PrintChecks()
			}

			return nil
		},
	}
	pflags := cmd.PersistentFlags()
	flagutils.AddOutputFlag(pflags, &opts.Top.Output)
	flagutils.AddNamespaceFlag(pflags, &opts.Metadata.Namespace)
	flagutils.AddExcludecheckFlag(pflags, &opts.Top.CheckName)
	cliutils.ApplyOptions(cmd, optionsFunc)
	return cmd
}

func CheckResources(opts *options.Options) error {
	var multiErr *multierror.Error

	err := checkConnection(opts.Top.Ctx, opts.Metadata.GetNamespace())
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
		return multiErr
	}

	deployments, err := getAndCheckDeployments(opts)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	includePods := doesNotContain(opts.Top.CheckName, "pods")
	if includePods {
		err := checkPods(opts)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	settings, err := getSettings(opts)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	namespaces, err := getNamespaces(opts.Top.Ctx, settings)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	knownUpstreams, err := checkUpstreams(opts, namespaces)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	includeUpstreamGroup := doesNotContain(opts.Top.CheckName, "upstreamgroup")
	if includeUpstreamGroup {
		err := checkUpstreamGroups(opts, namespaces)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	knownAuthConfigs, err := checkAuthConfigs(opts, namespaces)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	knownRateLimitConfigs, err := checkRateLimitConfigs(opts, namespaces)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	knownVirtualHostOptions, err := checkVirtualHostOptions(opts, namespaces)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	knownRouteOptions, err := checkRouteOptions(opts, namespaces)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	includeSecrets := doesNotContain(opts.Top.CheckName, "secrets")
	if includeSecrets {
		err := checkSecrets(opts, namespaces)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	err = checkVirtualServices(opts, namespaces, knownUpstreams, knownAuthConfigs, knownRateLimitConfigs, knownVirtualHostOptions, knownRouteOptions)
	if err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	includeGateway := doesNotContain(opts.Top.CheckName, "gateways")
	if includeGateway {
		err := checkGateways(opts, namespaces)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	includeProxy := doesNotContain(opts.Top.CheckName, "proxies")
	if includeProxy {
		err := checkProxies(opts, namespaces, opts.Metadata.GetNamespace(), deployments)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	includePrometheusStatsCheck := doesNotContain(opts.Top.CheckName, "xds-metrics")
	if includePrometheusStatsCheck {
		err = checkXdsMetrics(opts, opts.Metadata.GetNamespace(), deployments)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	if multiErr.ErrorOrNil() != nil {
		for _, err := range multiErr.Errors {
			printers.AppendResponse("", "", "", fmt.Sprint(err), opts.Top.Output)
		}
	}

	return multiErr.ErrorOrNil()
}

func getAndCheckDeployments(opts *options.Options) (*appsv1.DeploymentList, error) {
	printers.AppendResponse("Checking deployments... ", "", "", "", opts.Top.Output)
	client := helpers.MustKubeClient()
	_, err := client.CoreV1().Namespaces().Get(opts.Top.Ctx, opts.Metadata.GetNamespace(), metav1.GetOptions{})
	if err != nil {
		errMessage := "Gloo namespace does not exist"
		fmt.Println(errMessage)
		return nil, fmt.Errorf(errMessage)
	}
	deployments, err := client.AppsV1().Deployments(opts.Metadata.GetNamespace()).List(opts.Top.Ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if len(deployments.Items) == 0 {
		errMessage := "Gloo is not installed"
		fmt.Println(errMessage)
		return nil, fmt.Errorf(errMessage)
	}
	var multiErr *multierror.Error
	var message string
	setMessage := func(c appsv1.DeploymentCondition) {
		if c.Message != "" {
			message = fmt.Sprintf(" Message: %s", c.Message)
		}
	}

	for _, deployment := range deployments.Items {
		// possible condition types listed at https://godoc.org/k8s.io/api/apps/v1#DeploymentConditionType
		// check for each condition independently because multiple conditions will be True and DeploymentReplicaFailure
		// tends to provide the most explicit error message.
		for _, condition := range deployment.Status.Conditions {
			setMessage(condition)
			if condition.Type == appsv1.DeploymentReplicaFailure && condition.Status == corev1.ConditionTrue {
				err := fmt.Errorf("Deployment %s in namespace %s failed to create pods!%s", deployment.Name, deployment.Namespace, message)
				multiErr = multierror.Append(multiErr, err)
			}
		}

		for _, condition := range deployment.Status.Conditions {
			setMessage(condition)
			if condition.Type == appsv1.DeploymentProgressing && condition.Status != corev1.ConditionTrue {
				err := fmt.Errorf("Deployment %s in namespace %s is not progressing!%s", deployment.Name, deployment.Namespace, message)
				multiErr = multierror.Append(multiErr, err)
			}
		}

		for _, condition := range deployment.Status.Conditions {
			setMessage(condition)
			if condition.Type == appsv1.DeploymentAvailable && condition.Status != corev1.ConditionTrue {
				err := fmt.Errorf("Deployment %s in namespace %s is not available!%s", deployment.Name, deployment.Namespace, message)
				multiErr = multierror.Append(multiErr, err)
			}

		}

		for _, condition := range deployment.Status.Conditions {
			if condition.Type != appsv1.DeploymentAvailable &&
				condition.Type != appsv1.DeploymentReplicaFailure &&
				condition.Type != appsv1.DeploymentProgressing {
				err := fmt.Errorf("Deployment %s has an unhandled deployment condition %s", deployment.Name, condition.Type)
				multiErr = multierror.Append(multiErr, err)
			}
		}
	}
	if multiErr != nil {
		printers.AppendResponse("deployments", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return nil, multiErr
	}
	printers.AppendResponse("deployments", "OK\n", "", "", opts.Top.Output)
	return deployments, nil
}

func checkPods(opts *options.Options) error {
	printers.AppendResponse("Checking pods... ", "", "", "", opts.Top.Output)
	client := helpers.MustKubeClient()
	pods, err := client.CoreV1().Pods(opts.Metadata.GetNamespace()).List(opts.Top.Ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	var multiErr *multierror.Error
	for _, pod := range pods.Items {
		for _, condition := range pod.Status.Conditions {
			var errorToPrint string
			var message string

			if condition.Message != "" {
				message = fmt.Sprintf(" Message: %s", condition.Message)
			}

			// if condition is not met and the pod is not completed
			conditionNotMet := condition.Status != corev1.ConditionTrue && condition.Reason != "PodCompleted"

			// possible condition types listed at https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-conditions
			switch condition.Type {
			case corev1.PodScheduled:
				if conditionNotMet {
					errorToPrint = fmt.Sprintf("Pod %s in namespace %s is not yet scheduled!%s", pod.Name, pod.Namespace, message)
				}
			case corev1.PodReady:
				if conditionNotMet {
					errorToPrint = fmt.Sprintf("Pod %s in namespace %s is not ready!%s", pod.Name, pod.Namespace, message)
				}
			case corev1.PodInitialized:
				if conditionNotMet {
					errorToPrint = fmt.Sprintf("Pod %s in namespace %s is not yet initialized!%s", pod.Name, pod.Namespace, message)
				}
			case corev1.PodReasonUnschedulable:
				if conditionNotMet {
					errorToPrint = fmt.Sprintf("Pod %s in namespace %s is unschedulable!%s", pod.Name, pod.Namespace, message)
				}
			case corev1.ContainersReady:
				if conditionNotMet {
					errorToPrint = fmt.Sprintf("Not all containers in pod %s in namespace %s are ready!%s", pod.Name, pod.Namespace, message)
				}
			default:
				fmt.Printf("Note: Unhandled pod condition %s", condition.Type)
			}

			if errorToPrint != "" {
				multiErr = multierror.Append(multiErr, fmt.Errorf(errorToPrint))
			}
		}
	}
	if multiErr != nil {
		printers.AppendResponse("pods", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return multiErr
	}
	printers.AppendResponse("pods", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return nil
}

func getSettings(opts *options.Options) (*v1.Settings, error) {
	client := helpers.MustNamespacedSettingsClient(opts.Top.Ctx, opts.Metadata.GetNamespace())
	return client.Read(opts.Metadata.GetNamespace(), defaults.SettingsName, clients.ReadOpts{})
}

func getNamespaces(ctx context.Context, settings *v1.Settings) ([]string, error) {
	if settings.GetWatchNamespaces() != nil {
		return settings.GetWatchNamespaces(), nil
	}

	return helpers.GetNamespaces(ctx)
}

func checkUpstreams(opts *options.Options, namespaces []string) ([]string, error) {
	printers.AppendResponse("Checking upstreams... ", "", "", "", opts.Top.Output)
	var knownUpstreams []string
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		upstreams, err := helpers.MustNamespacedUpstreamClient(opts.Top.Ctx, ns).List(ns, clients.ListOpts{})
		if err != nil {
			return nil, err
		}
		for _, upstream := range upstreams {
			if upstream.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected upstream: %s ", renderMetadata(upstream.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", upstream.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
			if upstream.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found upstream with warnings: %s ", renderMetadata(upstream.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", upstream.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
			knownUpstreams = append(knownUpstreams, renderMetadata(upstream.GetMetadata()))
		}
	}
	if multiErr != nil {
		printers.AppendResponse("upstreams", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return nil, multiErr
	}
	printers.AppendResponse("upstreams", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return knownUpstreams, nil
}

func checkUpstreamGroups(opts *options.Options, namespaces []string) error {
	printers.AppendResponse("Checking upstream groups... ", "", "", "", opts.Top.Output)
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		upstreamGroups, err := helpers.MustNamespacedUpstreamGroupClient(opts.Top.Ctx, ns).List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, upstreamGroup := range upstreamGroups {
			if upstreamGroup.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected upstream group: %s ", renderMetadata(upstreamGroup.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", upstreamGroup.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
			if upstreamGroup.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found upstream group with warnings: %s ", renderMetadata(upstreamGroup.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", upstreamGroup.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
		}
	}
	if multiErr != nil {
		printers.AppendResponse("upstream groups", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return multiErr
	}
	printers.AppendResponse("upstream groups", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return nil
}

func checkAuthConfigs(opts *options.Options, namespaces []string) ([]string, error) {
	printers.AppendResponse("Checking auth configs... ", "", "", "", opts.Top.Output)
	var knownAuthConfigs []string
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		authConfigs, err := helpers.MustNamespacedAuthConfigClient(opts.Top.Ctx, ns).List(ns, clients.ListOpts{})
		if err != nil {
			return nil, err
		}
		for _, authConfig := range authConfigs {
			if authConfig.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected auth config: %s ", renderMetadata(authConfig.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", authConfig.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			} else if authConfig.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found auth config with warnings: %s ", renderMetadata(authConfig.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", authConfig.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
			knownAuthConfigs = append(knownAuthConfigs, renderMetadata(authConfig.GetMetadata()))
		}
	}
	if multiErr != nil {
		printers.AppendResponse("auth configs", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return nil, multiErr
	}
	printers.AppendResponse("auth configs", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return knownAuthConfigs, nil
}

func checkRateLimitConfigs(opts *options.Options, namespaces []string) ([]string, error) {
	printers.AppendResponse("Checking rate limit configs... ", "", "", "", opts.Top.Output)
	var knownConfigs []string
	var multiErr *multierror.Error
	for _, ns := range namespaces {

		rlcClient, err := helpers.RateLimitConfigClient(opts.Top.Ctx, []string{ns})
		if err != nil {
			if isCrdNotFoundErr(ratelimit.RateLimitConfigCrd, err) {
				// Just warn. If the CRD is required, the check would have failed on the crashing gloo/gloo-ee pod.
				fmt.Printf("WARN: %s\n", CrdNotFoundErr(ratelimit.RateLimitConfigCrd.KindName).Error())
				return nil, nil
			}
			return nil, err
		}

		configs, err := rlcClient.List(ns, clients.ListOpts{})
		if err != nil {
			return nil, err
		}
		for _, config := range configs {
			if config.Status.GetState() == v1alpha1.RateLimitConfigStatus_REJECTED {
				errMessage := fmt.Sprintf("Found rejected rate limit config: %s ", renderMetadata(config.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", config.Status.GetMessage())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
			knownConfigs = append(knownConfigs, renderMetadata(config.GetMetadata()))
		}
	}

	if multiErr != nil {
		printers.AppendResponse("rate limit configs", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return nil, multiErr
	}

	printers.AppendResponse("rate limit configs", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return knownConfigs, nil
}

func checkVirtualHostOptions(opts *options.Options, namespaces []string) ([]string, error) {
	printers.AppendResponse("Checking VirtualHostOptions... ", "", "", "", opts.Top.Output)
	var knownVhOpts []string
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		vhoptClient, err := helpers.VirtualHostOptionClient(opts.Top.Ctx, []string{ns})
		if err != nil {
			if isCrdNotFoundErr(gatewayv1.VirtualHostOptionCrd, err) {
				// Just warn. If the CRD is required, the check would have failed on the crashing gloo/gloo-ee pod.
				fmt.Printf("WARN: %s\n", CrdNotFoundErr(gatewayv1.VirtualHostOptionCrd.KindName).Error())
				return nil, nil
			}
			return nil, err
		}
		vhOpts, err := vhoptClient.List(ns, clients.ListOpts{})
		if err != nil {
			return nil, err
		}
		for _, vhOpt := range vhOpts {
			if vhOpt.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected VirtualHostOption: %s ", renderMetadata(vhOpt.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", vhOpt.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			} else if vhOpt.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found VirtualHostOption with warnings: %s ", renderMetadata(vhOpt.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", vhOpt.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
			knownVhOpts = append(knownVhOpts, renderMetadata(vhOpt.GetMetadata()))
		}
	}
	if multiErr != nil {
		printers.AppendResponse("VirtualHostOptions", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return nil, multiErr
	}
	printers.AppendResponse("VirtualHostOptions", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return knownVhOpts, nil
}

func checkRouteOptions(opts *options.Options, namespaces []string) ([]string, error) {
	printers.AppendResponse("Checking RouteOptions... ", "", "", "", opts.Top.Output)
	var knownVhOpts []string
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		routeOptionClient, err := helpers.RouteOptionClient(opts.Top.Ctx, []string{ns})
		if err != nil {
			if isCrdNotFoundErr(gatewayv1.RouteOptionCrd, err) {
				// Just warn. If the CRD is required, the check would have failed on the crashing gloo/gloo-ee pod.
				fmt.Printf("WARN: %s\n", CrdNotFoundErr(gatewayv1.RouteOptionCrd.KindName).Error())
				return nil, nil
			}
			return nil, err
		}
		vhOpts, err := routeOptionClient.List(ns, clients.ListOpts{})
		if err != nil {
			return nil, err
		}
		for _, routeOpt := range vhOpts {
			if routeOpt.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected RouteOption: %s ", renderMetadata(routeOpt.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", routeOpt.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			} else if routeOpt.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found RouteOption with warnings: %s ", renderMetadata(routeOpt.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", routeOpt.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, errors.New(errMessage))
			}
			knownVhOpts = append(knownVhOpts, renderMetadata(routeOpt.GetMetadata()))
		}
	}
	if multiErr != nil {
		printers.AppendResponse("RouteOptions", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return nil, multiErr
	}
	printers.AppendResponse("RouteOptions", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return knownVhOpts, nil
}

func checkVirtualServices(opts *options.Options, namespaces, knownUpstreams, knownAuthConfigs, knownRateLimitConfigs, knownVirtualHostOptions, knownRouteOptions []string) error {
	printers.AppendResponse("Checking virtual services... ", "", "", "", opts.Top.Output)
	var multiErr *multierror.Error

	for _, ns := range namespaces {
		virtualServices, err := helpers.MustNamespacedVirtualServiceClient(opts.Top.Ctx, ns).List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, virtualService := range virtualServices {
			if virtualService.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected virtual service: %s ", renderMetadata(virtualService.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", virtualService.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
			if virtualService.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found virtual service with warnings: %s ", renderMetadata(virtualService.GetMetadata()))
				errMessage += fmt.Sprintf("(Reason: %s)", virtualService.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
			for _, route := range virtualService.GetVirtualHost().GetRoutes() {
				if route.GetRouteAction() != nil {
					if route.GetRouteAction().GetSingle() != nil {
						us := route.GetRouteAction().GetSingle()
						if us.GetUpstream() != nil {
							if !cliutils.Contains(knownUpstreams, renderRef(us.GetUpstream())) {
								//TODO warning message if using rejected or warning upstream
								errMessage := fmt.Sprintf("Virtual service references unknown upstream: ")
								errMessage += fmt.Sprintf("(Virtual service: %s", renderMetadata(virtualService.GetMetadata()))
								errMessage += fmt.Sprintf(" | Upstream: %s)", renderRef(us.GetUpstream()))
								multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
							}
						}
					}
				}
			}

			// Check references to auth configs
			isAuthConfigRefValid := func(knownConfigs []string, ref *core.ResourceRef) error {
				// If the virtual service points to a specific, non-existent authconfig, it is not valid.
				if ref != nil && !cliutils.Contains(knownConfigs, renderRef(ref)) {
					//TODO: Virtual service references rejected or warning auth config
					errMessage := fmt.Sprintf("Virtual service references unknown auth config:\n")
					errMessage += fmt.Sprintf("  Virtual service: %s\n", renderMetadata(virtualService.GetMetadata()))
					errMessage += fmt.Sprintf("  Auth Config: %s\n", renderRef(ref))
					return fmt.Errorf(errMessage)
				}
				return nil
			}
			isOptionsRefValid := func(knownOptions []string, refs []*core.ResourceRef) error {
				// If the virtual host points to a specifc, non-existent VirtualHostOption, it is not valid.
				for _, ref := range refs {
					if ref != nil && !cliutils.Contains(knownOptions, renderRef(ref)) {
						errMessage := fmt.Sprintf("Virtual service references unknown VirtualHostOption:\n")
						errMessage += fmt.Sprintf("  Virtual service: %s\n", renderMetadata(virtualService.GetMetadata()))
						errMessage += fmt.Sprintf("  VirtualHostOption: %s\n", renderRef(ref))
						return fmt.Errorf(errMessage)
					}
				}
				return nil
			}
			// Check virtual host options
			if err := isAuthConfigRefValid(knownAuthConfigs, virtualService.GetVirtualHost().GetOptions().GetExtauth().GetConfigRef()); err != nil {
				multiErr = multierror.Append(multiErr, err)
			}
			vhDelegateOptions := virtualService.GetVirtualHost().GetOptionsConfigRefs().GetDelegateOptions()
			if err := isOptionsRefValid(knownVirtualHostOptions, vhDelegateOptions); err != nil {
				multiErr = multierror.Append(multiErr, err)
			}

			// Check route options
			for _, route := range virtualService.GetVirtualHost().GetRoutes() {
				if err := isAuthConfigRefValid(knownAuthConfigs, route.GetOptions().GetExtauth().GetConfigRef()); err != nil {
					multiErr = multierror.Append(multiErr, err)
				}
				if err := isOptionsRefValid(knownRouteOptions, route.GetOptionsConfigRefs().GetDelegateOptions()); err != nil {
					multiErr = multierror.Append(multiErr, err)
				}
				// Check weighted destination options
				for _, weightedDest := range route.GetRouteAction().GetMulti().GetDestinations() {
					if err := isAuthConfigRefValid(knownAuthConfigs, weightedDest.GetOptions().GetExtauth().GetConfigRef()); err != nil {
						multiErr = multierror.Append(multiErr, err)
					}
				}
			}

			// Check references to rate limit configs
			isRateLimitConfigRefValid := func(knownConfigs []string, ref *rlopts.RateLimitConfigRef) error {
				resourceRef := &core.ResourceRef{
					Name:      ref.GetName(),
					Namespace: ref.GetNamespace(),
				}
				if !cliutils.Contains(knownConfigs, renderRef(resourceRef)) {
					//TODO: check if references rate limit config with error or warning
					errMessage := fmt.Sprintf("Virtual service references unknown rate limit config:\n")
					errMessage += fmt.Sprintf("  Virtual service: %s\n", renderMetadata(virtualService.GetMetadata()))
					errMessage += fmt.Sprintf("  Rate Limit Config: %s\n", renderRef(resourceRef))
					return fmt.Errorf(errMessage)
				}
				return nil
			}
			// Check virtual host options
			for _, ref := range virtualService.GetVirtualHost().GetOptions().GetRateLimitConfigs().GetRefs() {
				if err := isRateLimitConfigRefValid(knownRateLimitConfigs, ref); err != nil {
					multiErr = multierror.Append(multiErr, err)
				}
			}
			// Check route options
			for _, route := range virtualService.GetVirtualHost().GetRoutes() {
				for _, ref := range route.GetOptions().GetRateLimitConfigs().GetRefs() {
					if err := isRateLimitConfigRefValid(knownRateLimitConfigs, ref); err != nil {
						multiErr = multierror.Append(multiErr, err)
					}
				}
			}
		}
	}

	if multiErr != nil {
		printers.AppendResponse("virtual services", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return multiErr
	}
	printers.AppendResponse("virtual services", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return nil
}

func checkGateways(opts *options.Options, namespaces []string) error {
	printers.AppendResponse("Checking gateways... ", "", "", "", opts.Top.Output)
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		gateways, err := helpers.MustNamespacedGatewayClient(opts.Top.Ctx, ns).List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, gateway := range gateways {
			if gateway.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected gateway: %s\n", renderMetadata(gateway.GetMetadata()))
				errMessage += fmt.Sprintf("Reason: %s\n", gateway.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
			if gateway.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found gateway with warnings: %s\n", renderMetadata(gateway.GetMetadata()))
				errMessage += fmt.Sprintf("Reason: %s\n", gateway.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
		}
	}

	if multiErr != nil {
		printers.AppendResponse("gateways", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return multiErr
	}

	printers.AppendResponse("gateways", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return nil
}

func checkProxies(opts *options.Options, namespaces []string, glooNamespace string, deployments *appsv1.DeploymentList) error {
	printers.AppendResponse("Checking proxies... ", "", "", "", opts.Top.Output)
	if deployments == nil {
		fmt.Println("Skipping due to an error in checking deployments")
		return fmt.Errorf("proxy check was skipped due to an error in checking deployments")
	}
	var multiErr *multierror.Error
	for _, ns := range namespaces {
		proxies, err := helpers.MustNamespacedProxyClient(opts.Top.Ctx, ns).List(ns, clients.ListOpts{})
		if err != nil {
			return err
		}
		for _, proxy := range proxies {
			if proxy.GetStatus().GetState() == core.Status_Rejected {
				errMessage := fmt.Sprintf("Found rejected proxy: %s\n", renderMetadata(proxy.GetMetadata()))
				errMessage += fmt.Sprintf("Reason: %s\n", proxy.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
			if proxy.GetStatus().GetState() == core.Status_Warning {
				errMessage := fmt.Sprintf("Found proxy with warnings: %s\n", renderMetadata(proxy.GetMetadata()))
				errMessage += fmt.Sprintf("Reason: %s\n", proxy.GetStatus().GetReason())
				multiErr = multierror.Append(multiErr, fmt.Errorf(errMessage))
			}
		}
	}

	if err := checkProxiesPromStats(opts.Top.Ctx, glooNamespace, deployments); err != nil {
		multiErr = multierror.Append(multiErr, err)
	}

	if multiErr != nil {
		printers.AppendResponse("proxies", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return multiErr
	}
	printers.AppendResponse("proxies", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return nil
}

func checkSecrets(opts *options.Options, namespaces []string) error {
	printers.AppendResponse("Checking secrets... ", "", "", "", opts.Top.Output)
	var multiErr *multierror.Error
	client := helpers.MustSecretClientWithOptions(opts.Top.Ctx, 5*time.Second, namespaces)

	for _, ns := range namespaces {
		_, err := client.List(ns, clients.ListOpts{})
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
		// currently this would only find syntax errors
	}
	if multiErr != nil {
		printers.AppendResponse("secrets", fmt.Sprintf("%v Errors!\n", multiErr.Len()), "", "", opts.Top.Output)
		return multiErr
	}
	printers.AppendResponse("secrets", fmt.Sprintf("OK\n"), "", "", opts.Top.Output)
	return nil
}

func renderMetadata(metadata *core.Metadata) string {
	return renderNamespaceName(metadata.GetNamespace(), metadata.GetName())
}

func renderRef(ref *core.ResourceRef) string {
	return renderNamespaceName(ref.GetNamespace(), ref.GetName())
}

func renderNamespaceName(namespace, name string) string {
	return fmt.Sprintf("%s %s", namespace, name)
}

// Checks whether the cluster that the kubeconfig points at is available
// The timeout for the kubernetes client is set to a low value to notify the user of the failure
func checkConnection(ctx context.Context, ns string) error {
	client, err := helpers.GetKubernetesClientWithTimeout(5 * time.Second)
	if err != nil {
		return eris.Wrapf(err, "Could not get kubernetes client")
	}
	_, err = client.CoreV1().Namespaces().Get(ctx, ns, metav1.GetOptions{})
	if err != nil {
		return eris.Wrapf(err, "Could not communicate with kubernetes cluster")
	}
	return nil
}

func isCrdNotFoundErr(crd crd.Crd, err error) bool {
	for {
		if statusErr, ok := err.(*apierrors.StatusError); ok {
			if apierrors.IsNotFound(err) &&
				statusErr.ErrStatus.Details != nil &&
				statusErr.ErrStatus.Details.Kind == crd.Plural {
				return true
			}
			return false
		}

		// This works for "github.com/pkg/errors"-based errors as well
		if wrappedErr := eris.Unwrap(err); wrappedErr != nil {
			err = wrappedErr
			continue
		}
		return false
	}
}

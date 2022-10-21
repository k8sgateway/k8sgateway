package validation

import (
	"context"
	"sync"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	gloov1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/go-utils/hashutils"

	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	errors "github.com/rotisserie/eris"
	utils2 "github.com/solo-io/gloo/pkg/utils"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/translator"
	"github.com/solo-io/gloo/projects/gateway/pkg/utils"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	gloovalidation "github.com/solo-io/gloo/projects/gloo/pkg/validation"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	skprotoutils "github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GatewayGroup = "gateway.solo.io"

// GvkSupportedValidationGatewayResources the current group of resources that can be validated
var GvkSupportedValidationGatewayResources = map[schema.GroupVersionKind]bool{
	v1.GatewayGVK:        true,
	v1.VirtualServiceGVK: true,
	v1.RouteTableGVK:     true,
}

// GvkSupportedDeleteGatewayResources the current group of resources that can be validated
var GvkSupportedDeleteGatewayResources = map[schema.GroupVersionKind]bool{
	v1.VirtualServiceGVK: true,
	v1.RouteTableGVK:     true,
}

type Reports struct {
	Proxies      []*gloov1.Proxy
	ProxyReports *ProxyReports
}

func (r *Reports) GetProxies() []*gloov1.Proxy {
	if r == nil || r.Proxies == nil {
		return []*gloov1.Proxy{}
	}
	return r.Proxies
}

type ProxyReports []*validation.ProxyReport
type UpstreamReports []*validation.ResourceReport

var (
	NotReadyErr             = errors.Errorf("validation is not yet available. Waiting for first snapshot")
	HasNotReceivedFirstSync = eris.New("proxy validation called before the validation server received its first sync of resources")
	unmarshalErrMsg         = "could not unmarshal raw object"
	couldNotRenderProxy     = "could not render proxy"
	failedGlooValidation    = "failed gloo validation"
	failedResourceReports   = "failed resource reports from gloo validation"
	WrappedUnmarshalErr     = func(err error) error {
		return errors.Wrapf(err, unmarshalErrMsg)
	}

	GlooValidationResponseLengthError = func(reports []*gloovalidation.GlooValidationReport) error {
		return errors.Errorf("Expected Gloo validation response to contain 1 report, but contained %d",
			len(reports))
	}

	proxyFailedGlooValidation = func(err error, proxy *gloov1.Proxy) error {
		return errors.Wrapf(err, "failed to validate Proxy [namespace: %s, name: %s] with gloo validation", proxy.GetMetadata().GetNamespace(), proxy.GetMetadata().GetName())
	}

	mValidConfig = utils2.MakeGauge("validation.gateway.solo.io/valid_config",
		"A boolean indicating whether gloo config is valid")

	groupIsNotSupported = func(resource resources.Resource, gvk schema.GroupVersionKind) error {
		return errors.Errorf("failed validating the resoruce [%T] because the group [%s] kind [%s] is not supported", resource, gvk.Group, gvk.Kind)
	}

	glooFailedResourceValidation = func(err error, resource resources.Resource) error {
		return errors.Wrapf(err, "failed to validate resource [%T] with Gloo validation server", resource)
	}
)

const (
	InvalidSnapshotErrMessage = "validation is disabled due to an invalid resource which has been written to storage. " +
		"Please correct any Rejected resources to re-enable validation."
)

var _ Validator = &validator{}

type Validator interface {
	gloov1snap.ApiSyncer
	// ValidateList will validate a list of resources
	ValidateList(ctx context.Context, ul *unstructured.UnstructuredList, dryRun bool) (*Reports, *multierror.Error)
	// ValidateModifiedGvk validate the deletion of a resource.
	ValidateModifiedGvk(ctx context.Context, gvk schema.GroupVersionKind, resource resources.Resource, dryRun bool) (*Reports, error)
	// ValidateDeletedGvk validate the creation or update of a resource.
	ValidateDeletedGvk(ctx context.Context, gvk schema.GroupVersionKind, resource resources.Resource, dryRun bool) error
	// ModificationIsSupported returns whether a resource is supported
	ModificationIsSupported(gvk schema.GroupVersionKind) bool
	// DeletionIsSupported returns whether a deletion of a resource is supported
	DeletionIsSupported(gvk schema.GroupVersionKind) bool
}

type GlooValidatorFunc = func(ctx context.Context, proxy *gloov1.Proxy,
	resource resources.Resource, delete bool,
) ([]*gloovalidation.GlooValidationReport, error)

type validator struct {
	lock              sync.RWMutex
	latestSnapshot    *gloov1snap.ApiSnapshot
	latestSnapshotErr error
	translator        translator.Translator
	// This function replaces a grpc client from when gloo and gateway pods were separate.
	glooValidator                GlooValidatorFunc
	ignoreProxyValidationFailure bool
	allowWarnings                bool
}

type validationOptions struct {
	Ctx         context.Context
	AcquireLock bool
	DryRun      bool
	Delete      bool
	Resource    resources.Resource
	Gvk         schema.GroupVersionKind
}

type ValidatorConfig struct {
	translator                   translator.Translator
	glooValidator                GlooValidatorFunc
	ignoreProxyValidationFailure bool
	allowWarnings                bool
}

func NewValidatorConfig(
	translator translator.Translator,
	glooValidator GlooValidatorFunc,
	ignoreProxyValidationFailure, allowWarnings bool,
) ValidatorConfig {
	return ValidatorConfig{
		glooValidator:                glooValidator,
		translator:                   translator,
		ignoreProxyValidationFailure: ignoreProxyValidationFailure,
		allowWarnings:                allowWarnings,
	}
}

func NewValidator(cfg ValidatorConfig) *validator {
	return &validator{
		glooValidator:                cfg.glooValidator,
		translator:                   cfg.translator,
		ignoreProxyValidationFailure: cfg.ignoreProxyValidationFailure,
		allowWarnings:                cfg.allowWarnings,
	}
}

func (v *validator) ready() bool {
	return v.latestSnapshot != nil
}

func (v *validator) Sync(ctx context.Context, snap *gloov1snap.ApiSnapshot) error {
	v.lock.Lock() // hashing and cloning resources may mutate the object, so we need to lock
	defer v.lock.Unlock()
	if !v.gatewayUpdate(snap) {
		return nil
	}
	snapCopy := snap.Clone()
	gatewaysByProxy := utils.GatewaysByProxyName(snap.Gateways)
	var errs error
	for proxyName, gatewayList := range gatewaysByProxy {
		_, reports := v.translator.Translate(ctx, proxyName, snap, gatewayList)
		validate := reports.ValidateStrict
		if v.allowWarnings {
			validate = reports.Validate
		}
		if err := validate(); err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	// When the pod is first starting (aka the first snapshot is received),
	// set the value of mValidConfig with respect to the translation loop above.
	// Without this, mValidConfig will not be exported on /metrics until a new
	// resource is applied (https://github.com/solo-io/gloo/issues/5949).
	if v.latestSnapshot == nil {
		if errs == nil {
			utils2.MeasureOne(ctx, mValidConfig)
		} else {
			utils2.MeasureZero(ctx, mValidConfig)
		}
	}

	v.latestSnapshotErr = errs
	v.latestSnapshot = &snapCopy

	if errs != nil {
		return errors.Wrapf(errs, InvalidSnapshotErrMessage)
	}

	return nil
}

func (v *validator) ModificationIsSupported(gvk schema.GroupVersionKind) bool {
	// note ModificationIsSupported does not currently support Secrets.  This is
	// because it is only supported if deleting.
	_, supported := GvkSupportedValidationGatewayResources[gvk]
	if !supported {
		_, supported := gloovalidation.GvkToSupportedGlooResources[gvk]
		return supported
	}
	return supported
}

// DeletionIsSupported checks whether the deletion of a particular resources is supported.
func (v *validator) DeletionIsSupported(gvk schema.GroupVersionKind) bool {
	_, supported := GvkSupportedDeleteGatewayResources[gvk]
	if !supported {
		_, supported = gloovalidation.GvkToSupportedDeleteGlooResources[gvk]
		return supported
	}
	return supported
}

func (v *validator) gatewayUpdate(snap *gloov1snap.ApiSnapshot) bool {

	if v.latestSnapshot == nil {
		return true
	}
	//look at the hash of resources that affect the gateway snapshot
	hashFunc := func(snap *gloov1snap.ApiSnapshot) (uint64, error) {
		toHash := append([]interface{}{}, snap.VirtualHostOptions.AsInterfaces()...)
		toHash = append(toHash, snap.VirtualServices.AsInterfaces()...)
		toHash = append(toHash, snap.Gateways.AsInterfaces()...)
		toHash = append(toHash, snap.RouteOptions.AsInterfaces()...)
		toHash = append(toHash, snap.RouteTables.AsInterfaces()...)
		toHash = append(toHash, snap.HttpGateways.AsInterfaces()...)
		toHash = append(toHash, snap.VirtualHostOptions.AsInterfaces()...)
		hash, err := hashutils.HashAllSafe(nil, toHash...)
		if err != nil {
			contextutils.LoggerFrom(context.Background()).DPanic("this error should never happen, as this is safe hasher")
			return 0, errors.New("this error should never happen, as this is safe hasher")
		}
		return hash, nil
	}
	oldHash, oldHashErr := hashFunc(v.latestSnapshot)
	newHash, newHashErr := hashFunc(snap)

	// If we cannot hash then we choose to treat them as different hashes since this is just a performance optimization.
	// In worst case we'd prefer correctness
	hashChanged := oldHash != newHash || oldHashErr != nil || newHashErr != nil
	return hashChanged
}

func (v *validator) validateSnapshotThreadSafe(opts *validationOptions) (
	*Reports,
	error,
) {
	v.lock.Lock()
	defer v.lock.Unlock()

	return v.validateSnapshot(opts)
}

func (v *validator) validateSnapshot(opts *validationOptions) (*Reports, error) {
	// validate that a snapshot can be modified
	// should be called within a lock
	//
	// validation occurs by the following steps:
	//	1. Clone the most recent snapshot
	//	2. Apply the changes to that snapshot clone
	//	3. Validate the generated proxy of that snapshot clone by validating both gateway and gloo translation.
	//		a. we call gloo translation via a passed method, glooValidator
	//	4. If the proxy is valid, we know that the requested mutation is valid. If this request happens
	//		during a dry run, we don't want to actually apply the change, since this will modify the internal
	//		state of the validator, which is shared across requests. Therefore, only if we are not in a dry run,
	//		we apply the mutation.
	ctx := opts.Ctx
	if !v.ready() {
		return nil, NotReadyErr
	}
	ref := opts.Resource.GetMetadata().Ref()
	ctx = contextutils.WithLogger(ctx, "gateway-validator")

	// currently have the other for Gloo resources
	snapshotClone, err := v.copySnapshotNonThreadSafe(ctx, opts.DryRun)
	if err != nil {
		// allow writes if storage is already broken
		return nil, nil
	}

	// verify the mutation against a snapshot clone first, only apply the change to the actual snapshot if this passes
	if opts.Delete {
		snapshotClone.RemoveFromResourceList(opts.Resource)
	} else {
		snapshotClone.UpsertToResourceList(opts.Resource)
	}

	// only set the glooResource if we are validating a gloo resource
	var glooResource resources.Resource
	if opts.Gvk.Group != GatewayGroup {
		glooResource = opts.Resource
	}

	var (
		errs         error
		proxyReports ProxyReports
		proxies      []*gloov1.Proxy
	)
	gatewaysByProxy := utils.GatewaysByProxyName(snapshotClone.Gateways)
	// translate all the proxies
	for proxyName, gatewayList := range gatewaysByProxy {
		proxy, reports := v.translator.Translate(ctx, proxyName, snapshotClone, gatewayList)
		validate := reports.ValidateStrict
		if v.allowWarnings {
			validate = reports.Validate
		}
		if err := validate(); err != nil {
			errs = multierr.Append(errs, errors.Wrapf(err, couldNotRenderProxy))
			continue
		}

		// a nil proxy may have been returned if 0 listeners were created
		if proxy == nil {
			continue
		}

		proxies = append(proxies, proxy)
		// validate the proxy with gloo
		glooReports, err := v.glooValidator(ctx, proxy, glooResource, opts.Delete)
		if err != nil {
			err = errors.Wrapf(err, failedGlooValidation)
			if v.ignoreProxyValidationFailure {
				contextutils.LoggerFrom(ctx).Error(err)
			} else {
				errs = multierr.Append(errs, err)
			}
			continue
		}

		if len(glooReports) != 1 {
			// This was likely caused by a development error
			err := GlooValidationResponseLengthError(glooReports)
			errs = multierr.Append(errs, err)
			continue
		}

		proxyReport := glooReports[0].ProxyReport
		proxyReports = append(proxyReports, proxyReport)
		if err := validationutils.GetProxyError(proxyReport); err != nil {
			errs = multierr.Append(errs, proxyFailedGlooValidation(err, proxy))
			continue
		}
		if warnings := validationutils.GetProxyWarning(proxyReport); !v.allowWarnings && len(warnings) > 0 {
			for _, warning := range warnings {
				errs = multierr.Append(errs, errors.New(warning))
			}
			continue
		}

		_, err = v.getReportsFromGlooValidation(glooReports)
		if err != nil {
			err = errors.Wrapf(err, failedResourceReports)
			errs = multierr.Append(errs, err)
			continue
		}
	}

	if errs != nil {
		contextutils.LoggerFrom(ctx).Debugf("Rejected %T %v: %v", opts.Resource, ref, errs)
		if !opts.DryRun {
			utils2.MeasureZero(ctx, mValidConfig)
		}
		return &Reports{ProxyReports: &proxyReports, Proxies: proxies}, errors.Wrapf(errs,
			"validating %T %v",
			opts.Resource,
			ref)
	}

	contextutils.LoggerFrom(ctx).Debugf("Accepted %T %v", opts.Resource, ref)
	if !opts.DryRun {
		utils2.MeasureOne(ctx, mValidConfig)
	}

	if !opts.DryRun {
		// update internal snapshot to handle race where a lot of resources may be applied at once, before syncer updates
		if opts.Delete {
			v.latestSnapshot.RemoveFromResourceList(opts.Resource)
		} else {
			v.latestSnapshot.UpsertToResourceList(opts.Resource)
		}
	}

	return &Reports{ProxyReports: &proxyReports, Proxies: proxies}, nil
}

// ValidateDeletedGvk will validate a deletion of a resource, as long as it is supported, against the Gateway and Gloo Translations.
func (v *validator) ValidateDeletedGvk(ctx context.Context, gvk schema.GroupVersionKind, resource resources.Resource, dryRun bool) error {
	_, supportedGatewayResource := GvkSupportedDeleteGatewayResources[gvk]
	_, supportedGlooResource := gloovalidation.GvkToSupportedDeleteGlooResources[gvk]
	if supportedGatewayResource || supportedGlooResource {
		_, err := v.validateResource(&validationOptions{Ctx: ctx, Resource: resource, Delete: true, DryRun: dryRun, AcquireLock: true})
		return err
	}
	contextutils.LoggerFrom(ctx).Debugf("unsupported validation for resource delete ref namespace [%s] name [%s] group [%s] kind [%s]", resource.GetMetadata().GetNamespace(), resource.GetMetadata().GetName(), gvk.Group, gvk.Kind)
	return nil
}

// ValidateModifiedGvk will validate a resource, as long as it is supported, against the Gateway and Gloo translations.
// The resource should be updated or created.  Use Validate Delete Gvk for deleted resources.
func (v *validator) ValidateModifiedGvk(ctx context.Context, gvk schema.GroupVersionKind, resource resources.Resource, dryRun bool) (*Reports, error) {
	return v.validateModifiedResource(ctx, gvk, resource, dryRun, true)
}

func (v *validator) validateModifiedResource(ctx context.Context, gvk schema.GroupVersionKind, resource resources.Resource, dryRun, acquireLock bool) (*Reports, error) {
	var reports *Reports
	if v.ModificationIsSupported(gvk) {
		reports, err := v.validateResource(&validationOptions{Ctx: ctx, Resource: resource, Gvk: gvk, Delete: false, DryRun: dryRun, AcquireLock: acquireLock})
		if err != nil {
			return reports, &multierror.Error{Errors: []error{errors.Wrapf(err, "Validating %T failed", resource)}}
		}
		return reports, nil
	}
	return reports, &multierror.Error{Errors: []error{groupIsNotSupported(resource, gvk)}}
}

func (v *validator) ValidateList(ctx context.Context, ul *unstructured.UnstructuredList, dryRun bool) (
	*Reports,
	*multierror.Error,
) {
	var (
		proxies      []*gloov1.Proxy
		proxyReports = ProxyReports{}
		errs         = &multierror.Error{}
	)

	v.lock.Lock()
	defer v.lock.Unlock()
	originalSnapshot := v.latestSnapshot.Clone()

	for _, item := range ul.Items {

		// this will lock
		var itemProxyReports, err = v.processItem(ctx, item)

		errs = multierror.Append(errs, err)
		if itemProxyReports != nil && itemProxyReports.ProxyReports != nil {
			// ok to return final proxy reports as the latest result includes latest proxy calculated
			// for each resource, as we process incrementally, storing new state in memory as we go
			proxyReports = append(proxyReports, *itemProxyReports.ProxyReports...)
			proxies = append(proxies, itemProxyReports.Proxies...)
		}
	}

	if dryRun {
		// to validate the entire list of changes against one another, each item was applied to the latestSnapshot
		// if this is a dry run, latestSnapshot needs to be reset back to its original value without any of the changes
		v.latestSnapshot = &originalSnapshot
	}

	return &Reports{ProxyReports: &proxyReports, Proxies: proxies}, errs
}

func (v *validator) processItem(ctx context.Context, item unstructured.Unstructured) (*Reports, error) {
	// process a single change in a list of changes
	//
	// when calling the specific internal validate method, dryRun and acquireLock are always false:
	// 	dryRun=false: this enables items to be validated against other items in the list
	// 	acquireLock=false: the entire list of changes are called within a single lock
	gv, err := schema.ParseGroupVersion(item.GetAPIVersion())
	if err != nil {
		return &Reports{ProxyReports: &ProxyReports{}}, err
	}

	itemGvk := schema.GroupVersionKind{
		Version: gv.Version,
		Group:   gv.Group,
		Kind:    item.GetKind(),
	}

	jsonBytes, err := item.MarshalJSON()
	if err != nil {
		return &Reports{ProxyReports: &ProxyReports{}}, err
	}

	if newResourceFunc, hit := gloosnapshot.ApiGvkToHashableResource[itemGvk]; hit && v.ModificationIsSupported(itemGvk) {
		resource := newResourceFunc()
		if unmarshalErr := skprotoutils.UnmarshalResource(jsonBytes, resource); unmarshalErr != nil {
			return &Reports{ProxyReports: &ProxyReports{}}, WrappedUnmarshalErr(unmarshalErr)
		}
		return v.validateModifiedResource(ctx, itemGvk, resource, false, false)
	}
	// should not happen
	return &Reports{ProxyReports: &ProxyReports{}}, errors.Errorf("Unknown group/version/kind, %v", itemGvk)
}

// copySnapshotNonThreadSafe will copy the snapshot. If there is an error with the latest snapshot, it will error.
// NOTE: does not perform any lock, and this function is not thread safe. Any read or write to the snapshot needs to be
// done under a lock
func (v *validator) copySnapshotNonThreadSafe(ctx context.Context, dryRun bool) (*gloosnapshot.ApiSnapshot, error) {
	if v.latestSnapshot == nil {
		return nil, HasNotReceivedFirstSync
	}
	if v.latestSnapshotErr != nil {
		if !dryRun {
			utils2.MeasureZero(ctx, mValidConfig)
		}
		contextutils.LoggerFrom(ctx).Errorw(InvalidSnapshotErrMessage, zap.Error(v.latestSnapshotErr))
		return nil, eris.New(InvalidSnapshotErrMessage)
	}
	snapshotClone := v.latestSnapshot.Clone()
	return &snapshotClone, nil
}

func (v *validator) validateResource(opts *validationOptions) (*Reports, error) {
	if opts.AcquireLock {
		return v.validateSnapshotThreadSafe(opts)
	} else {
		return v.validateSnapshot(opts)
	}
}

func GetDelegateRef(delegate *v1.DelegateAction) *core.ResourceRef {
	// handle deprecated route table resource reference format
	// TODO(marco): remove when we remove the deprecated fields from the API
	if delegate.GetNamespace() != "" || delegate.GetName() != "" {
		return &core.ResourceRef{
			Namespace: delegate.GetNamespace(),
			Name:      delegate.GetName(),
		}
	} else if delegate.GetRef() != nil {
		return delegate.GetRef()
	}
	return nil
}

// Converts the GlooValidationService into Reports.
func (v *validator) getReportsFromGlooValidation(reports []*gloovalidation.GlooValidationReport) (
	*Reports,
	error,
) {
	var (
		errs         error
		proxyReports ProxyReports
		proxies      []*gloov1.Proxy
	)
	for _, report := range reports {
		// for resorce, resourceReport
		for resource, reRpt := range report.ResourceReports {
			if err := resourceReportToMultiErr(reRpt.Errors); err != nil {
				errs = multierr.Append(errs, glooFailedResourceValidation(err, resource))
			}
			if warnings := reRpt.Warnings; !v.allowWarnings && len(warnings) > 0 {
				for _, warning := range warnings {
					errs = multierr.Append(errs, errors.New(warning))
				}
			}
		}
		// Append proxies and proxy reports
		if report.Proxy != nil {
			proxies = append(proxies, report.Proxy)
		}
		if proxyReport := report.ProxyReport; proxyReport != nil {
			proxyReports = append(proxyReports, report.ProxyReport)
			if err := validationutils.GetProxyError(proxyReport); err != nil {
				errs = multierr.Append(errs, errors.Wrapf(err, "failed to validate Proxy with Gloo validation server"))
			}
			if warnings := validationutils.GetProxyWarning(proxyReport); !v.allowWarnings && len(warnings) > 0 {
				for _, warning := range warnings {
					errs = multierr.Append(errs, errors.New(warning))
				}
			}
		}
	}
	return &Reports{
		ProxyReports: &proxyReports,
		Proxies:      proxies,
	}, errs
}

func resourceReportToMultiErr(err error) error {
	var multiErr error
	for _, errStr := range getErrors(err) {
		multiErr = multierr.Append(multiErr, errors.New(errStr))
	}
	return multiErr
}

func getErrors(err error) []string {
	if err == nil {
		return []string{}
	}
	switch err := err.(type) {
	case *multierror.Error:
		var errorStrings []string
		for _, e := range err.Errors {
			errorStrings = append(errorStrings, e.Error())
		}
		return errorStrings
	}
	return []string{err.Error()}
}

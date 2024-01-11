---
title: Prepare to upgrade
weight: 10
description: Prepare your environment, review version changes, and review FAQs before you upgrade Gloo Edge.
---

Before you upgrade Gloo Edge, complete the following preparatory steps:
* [Prepare your environment](#prepare), such as upgrading your current version to the latest patch and upgrading any dependencies to the required supported versions. 
* [Review important changes](#review-changes) made to Gloo Edge in version {{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}}1.16, including CRD, Helm, CLI, and feature changes.
* [Review frequently-asked questions](#faqs) about the upgrade process.


## Prepare your environment {#prepare}

Review the following preparatory steps that might be required for your environment.

### Upgrade your current minor version to the latest patch {#current-patch}

Before you upgrade your minor version, first upgrade your current version to the latest patch. For example, if you currently run Gloo Edge Enterprise version {{< readfile file="static/content/version_gee_n-1_oldpatch.md" markdown="true">}}, first upgrade your installation to version {{< readfile file="static/content/version_gee_n-1.md" markdown="true">}}. This ensures that your current environment is up-to-date with any bug fixes or security patches before you begin the minor version upgrade process.

1. Find the latest patch of your minor version by checking the [Open Source changelog]({{% versioned_link_path fromRoot="/reference/changelog/open_source/" %}}) or [Enterprise changelog]({{% versioned_link_path fromRoot="/reference/changelog/enterprise/" %}}).
2. Go to the documentation set for your current minor version. For example, if you currently run Gloo Edge Enterprise version {{< readfile file="static/content/version_gee_n-1_oldpatch.md" markdown="true">}}, use the drop-down menu in the header of this page to select **v{{< readfile file="static/content/version_geoss_n-1_minor.md" markdown="true">}}.x**.
3. Follow the upgrade guide, using the latest patch for your minor version.

### If required, perform incremental minor version updates {#minor-increment}

If you plan to upgrade to a version that is more than one minor version greater than your current version, such as to version {{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}} from {{< readfile file="static/content/version_geoss_n-2_minor.md" markdown="true">}} or older, you must upgrade incrementally. For example, you must first use the upgrade guide in the v{{< readfile file="static/content/version_geoss_n-1_minor.md" markdown="true">}}.x documentation set to upgrade from {{< readfile file="static/content/version_geoss_n-2_minor.md" markdown="true">}} to {{< readfile file="static/content/version_geoss_n-1_minor.md" markdown="true">}}, and then follow the upgrade guide in the v{{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}}.x documentation set to upgrade from {{< readfile file="static/content/version_geoss_n-1_minor.md" markdown="true">}} to {{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}}.

### Upgrade dependencies {#dependencies}

Check that your underlying infrastructure platform, such as Kubernetes, and other dependencies run a version that is supported for {{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}}.

1. Review the [supported versions]({{% versioned_link_path fromRoot="/reference/support/#supported-versions" %}}) for dependencies such as Kubernetes, Helm, and more.
2. Compare the supported versions against the versions you currently use.
3. If necessary, upgrade your dependencies, such as consulting your cluster infrastructure provider to upgrade the version of Kubernetes that your cluster runs.

### Consider settings to avoid downtime {#downtime}

You might deploy Gloo Edge in Kubernetes environments that use the Kubernetes load balancer, or in non-Kubernetes environments. Depending on your setup, you can take additional steps to avoid downtime during the upgrade process.

* **Kubernetes**: Enable [Envoy readiness and liveness probes]({{< versioned_link_path fromRoot="/operations/production_deployment/#enable-health-checks" >}}) during the upgrade. When these probes are set, Kubernetes sends requests only to the healthy Envoy proxy during the upgrade process, which helps to prevent potential downtime. The probes are not enabled in default installations because they can lead to timeouts or other poor getting started experiences. 
* **Non-Kubernetes**: Configure [health checks]({{< versioned_link_path fromRoot="/guides/traffic_management/request_processing/health_checks" >}}) on Envoy. Then, configure your load balancer to leverage these health checks, so that requests stop going to Envoy when it begins draining connections.

## Review version {{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}} changes {#review-changes}

Review the following changes made to Gloo Edge in version {{< readfile file="static/content/version_geoss_latest_minor.md" markdown="true">}}. For some changes, you might be required to complete additional steps during the upgrade process.

### Changelogs

Check the changelogs for the type of Gloo Edge deployment that you have. Focus especially on any **Breaking Changes** that might require a different upgrade procedure. For Gloo Edge Enterprise, you might also review the open source changelogs because most of the proto definitions are open source.
* [Open Source changelogs]({{% versioned_link_path fromRoot="/reference/changelog/open_source/" %}})
* [Enterprise changelogs]({{% versioned_link_path fromRoot="/reference/changelog/enterprise/" %}}): Keep in mind that Gloo Edge Enterprise pulls in Gloo Edge Open Source as a dependency. Although the major and minor version numbers are the same for open source and enterprise, their patch versions often differ. For example, open source might use version `x.y.a` but enterprise uses version `x.y.b`. If you are unfamiliar with these versioning concepts, see [Semantic versioning](https://semver.org/). Because of the differing patch versions, you might notice different output when checking your version with `glooctl version`. For example, your API server might run Gloo Edge Enterprise version {{< readfile file="static/content/version_gee_latest.md" markdown="true">}}, which pulls in Gloo Edge Open Source version {{< readfile file="static/content/version_geoss_latest.md" markdown="true">}} as a dependency.
  ```bash
  ~ > glooctl version
  Client: {"version":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}"}
  Server: {"type":"Gateway","enterprise":true,"kubernetes":...,{"Tag":"{{< readfile file="static/content/version_gee_latest.md" markdown="true">}}","Name":"grpcserver-ee","Registry":"quay.io/solo-io"},...,{"Tag":"{{< readfile file="static/content/version_geoss_latest.md" markdown="true">}}","Name":"discovery","Registry":"quay.io/solo-io"},...}
  ```

{{% notice tip %}}
You can use the changelogs' built-in [comparison tool]({{< versioned_link_path fromRoot="/reference/changelog/open_source/#compareversions" >}}) to compare between your current version and the version that you want to upgrade to.
{{% /notice %}}

### Feature changes {#features}
Review the following summary of important new, deprecated, or removed features.

{{% notice note %}}
The following lists consist of the changes that were initially introduced with the 1.16.0 release. These changes might be backported to earlier versions of Gloo Edge. Additionally, there might be other changes that are introduced in later 1.16 patch releases. For patch release changes, check the [changelogs](#changelogs).
{{% /notice %}}

**New or improved features**:

* **Renewal of HashiCorp Vault tokens**: You can now use the `settings.secretOptions.sources[].vault.aws.leaseIncrement` Helm value to specify the lease increment to use for the HashiCorp token renewal. This value is passed as the 'Increment' parameter to the HashiCorp vault API. For more information about the Vault integration, see [Securing secrets in HashiCorp Vault using AWS IAM Roles for Service Accounts (IRSA)]({{% versioned_link_path fromRoot="/guides/integrations/vault/aws-irsa-auth/" %}}). 
* **Rotate TLS certificates in stages**: Use the `gateway.certGenJob.rotationDuration` Helm value to configure the duration for Gloo Edge to wait during each stage of the certificate renewal process when mTLS is enabled. For more information, see [Cert Rotation]({{% versioned_link_path fromRoot="/guides/security/tls/mtls/#cert-rotation" %}}). 
* **Validate Kubernetes secrets on removal**: Fixed an issue in the validation webhook where Kubernetes secrets were not validated on deletions. You can opt out of secret validation by using the `gateway.validation.webhook.skipDeleteValidationResources` Helm value. 
* **Local rate limiting**: Limit the number of requests to your gateway or upstream services before the requests reach the rate limiting server in your cluster. Local rate limiting can protect your rate limit servers from being overloaded and help optimize their resource utilization. For more information, see [Local rate limiting]({{% versioned_link_path fromRoot="/guides/security/local_rate_limiting/" %}}). 
* **Traffic tapping (alpha)**: Copy contents of HTTP or gRPC requests and responses to an external tap server. Note that traffic tapping is introduced as an alpha feature in Gloo Edge Enterprise. Alpha features are likely to change, are not fully tested, and are not supported for production. For more information, see [Traffic tapping]({{% versioned_link_path fromRoot="/guides/traffic_management/listener_configuration/tap/" %}}). 

<!--
**Deprecated features**:
N/A

**Removed features**:
N/A
-->

### Helm changes {#helm}

Review the following summary of important new, deprecated, or removed Helm fields. For full details, see the [changelogs](#changelogs).

**New and updated Helm fields**:

* `settings.secretOptions.sources[].vault.aws.leaseIncrement`: Specify the lease increment to use for HashiCorp token renewal. This value is passed as the 'Increment' parameter to the HashiCorp vault API.
* `gatewayproxy.proxyName.disableExtauthSidecar`: Disable the extauth sidecar on a gateway proxy when `global.extensions.extAuth.envoySidecar` is set in a Gloo Edge Enterprise installation. The defaut value is `false`. 
* `global.extraCustomResources`: Deploy all custom resources during the Gloo Edge installation. The default value is `false` for Gloo Edge Open Source, and `true` in Gloo Edge Enterprise. 
* `gateway.certGenJob.rotationDuration`: Configure the duration for Gloo Edge to wait during each stage of the certificate renewal process when mTLS is enabled. For more information, see [Cert Rotation]({{% versioned_link_path fromRoot="/guides/security/tls/mtls/#cert-rotation" %}}).
* `gloo.gateway.certGenJob.forceRotation`: Force the renewal of TLS certificates, even if they are not expired yet.
* `gateway.rolloutJob.timeout`: Specifiy the timeout to wait for the resource rollout job to complete. The default value is 120s. 
* `gateway.validation.webhook.skipDeleteValidationResources`: Skip validation when deleting Kubernetes secrets. 

<!--
**Updated Helm fields**:
N/A

**Deprecated Helm fields**:
N/A

**Removed Helm fields**:
N/A
-->

### CRD changes {#crd}

New CRDs are automatically applied to your cluster when performing a `helm install` operation, but are _not_ applied when performing an `helm upgrade` operation. This is a [deliberate design choice](https://helm.sh/docs/topics/charts/#limitations-on-crds) on the part of the Helm maintainers, given the risk associated with changing CRDs. Given this limitation, you must apply new CRDs to the cluster before upgrading. 

Review the following summary of important new, deprecated, or removed CRD updates. For full details, see the [changelogs](#changelogs).

**New and updated CRDs**:

* `ExtAuthConfig`: In the [`OidcAuthorizationCode`]({{% versioned_link_path fromRoot="/reference/api/github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/extauth/v1/extauth.proto.sk/#oidcauthorizationcode" %}}) section, you can now add a [`clientAuthentication`]({{% versioned_link_path fromRoot="/reference/api/github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/extauth/v1/extauth.proto.sk/#clientauthentication" %}}) block to specify the authentication type that you want to use to exchange the access code for the access and ID tokens.
* `Settings`: You can now use the `secretOptions` block to use an AWS IAM IRSA to get access to Vault, instead of specifying an `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`. 

<!--
**Deprecated CRDs**:
N/A

**Removed CRDs**:
N/A
-->

### CLI changes {#cli}

You must upgrade `glooctl` before you upgrade Gloo Edge. Because `glooctl` can create resources in your cluster, such as with `glooctl add route`, you might have errors in Gloo Edge if you create resources with an older version of `glooctl`.

As part of the 1.16 release, no CLI changes were introduced.
<!--
Review the following summary of important new, deprecated, or removed CLI options. For full details, see the [changelogs](#changelogs).

**New CLI commands or options**:

* `glooctl create secret encryptionkey`: [Create encryption secrets]({{% versioned_link_path fromRoot="/reference/cli/glooctl_create_secret_encryptionkey/" %}}), such as to use in the `cipherConfig` field of the `ExtAuthConfig` resource.

**Changed behavior**:-->



## Frequently-asked questions {#faqs}

Review the following frequently-asked questions about the upgrade process. If you still aren't sure about the version upgrade impact, or if your use case doesn't quite fit the standard upgrade path, feel free to post in the `#gloo` or `#gloo-enterprise` channels of our [public Slack](https://slack.solo.io/).

### How do I upgrade Gloo Edge in testing or sandbox environments?

If downtime is not a concern for your use case, you can follow the [Quick upgrade guide]({{< versioned_link_path fromRoot="/operations/upgrading/upgrade_steps" >}}) to update your Gloo Edge installation.

Note that for sandbox or exploratory environments, the easiest way to upgrade is to uninstall Gloo Edge by running `glooctl uninstall --all`. Then, re-install Gloo Edge at the desired version by the following one of the [installation guides]({{< versioned_link_path fromRoot="/installation" >}}).
 
### How do I upgrade Gloo Edge in a production environment, where downtime is unacceptable?

The basic `helm upgrade` process is not suitable for environments in which downtime is unacceptable. Instead, you can follow the [Canary upgrade]({{% versioned_link_path fromRoot="/operations/upgrading/canary/" %}}) guide to deploy multiple version of Gloo Edge to your cluster, and test the upgrade version before uninstalling the existing version.

Additionally, you might need to take steps to account for other factors such as Gloo Edge version changes, probe configurations, and external infrastructure like the load balancer that Gloo Edge uses. Consider setting up [liveness probes and healthchecks](#downtime) in your environment.

### What happens to my Gloo Edge CRs during an upgrade? How do I handle breaking changes?

A typical upgrade of Gloo Edge across minor versions should not cause disruptions to the existing Gloo Edge state. In the case of a breaking change, Solo will communicate through the upgrade guides, changelogs, or other channels if you must make a specific adjustment to perform the upgrade. Note that you can always use the `glooctl debug yaml` command to download the current Gloo Edge state to one large YAML manifest.

### Is the upgrade procedure different if I am not a cluster administrator?

If you are not an administrator of your cluster, you might be unable to create custom resource definitions (CRDs) and other cluster-scoped resources, such as cluster roles and cluster role bindings. If you encounter an error related to these resources, you can disable their creation by including the following setting in your Helm values:
```yaml
global:
  glooRbac:
    create: false
```

Otherwise, you can try performing an installation of Gloo Edge that is scoped to a single namespace by including the following setting in your Helm values:
```yaml
global:
  glooRbac:
    namespaced: true
```

### Why do I get an error about re-creating CRDs when upgrading using `helm install` or `helm upgrade`?

Helm v2 does not manage CRDs well, and is not supported in Gloo Edge. Upgrade to Helm v3, delete the CRDs, and try again.

### Why do I get an error about a `gateway-certgen` job?

The upgrade creates a Kubernetes Job named `gateway-certgen` to generate a certificate for the validation webhook. The job contains the `ttlSecondsAfterFinished` value so that the cluster cleans up the job automatically, but because this setting is still in Alpha, your cluster might ignore this value. In this case, you might have an issue while upgrading in which the upgrade attempts to change the `gateway-certgen` job, but the change fails because the job is immutable. To fix this issue, you can delete the job, which already completed, and re-apply the upgrade.

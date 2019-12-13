#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------

ROOTDIR := $(shell pwd)
OUTPUT_DIR ?= $(ROOTDIR)/_output

# Kind of a hack to make sure _output exists
z := $(shell mkdir -p $(OUTPUT_DIR))

SOURCES := $(shell find . -name "*.go" | grep -v test.go | grep -v '\.\#*')
RELEASE := "true"
ifeq ($(TAGGED_VERSION),)
	# TAGGED_VERSION := $(shell git describe --tags)
	# This doesn't work in CI, need to find another way...
	TAGGED_VERSION := vdev
	RELEASE := "false"
endif
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

LDFLAGS := "-X github.com/solo-io/gloo/pkg/version.Version=$(VERSION)"
GCFLAGS := all="-N -l"

GO_BUILD_FLAGS := GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64

# Passed by cloudbuild
GCLOUD_PROJECT_ID := $(GCLOUD_PROJECT_ID)
BUILD_ID := $(BUILD_ID)

TEST_IMAGE_TAG := test-$(BUILD_ID)
TEST_ASSET_DIR := $(ROOTDIR)/_test
GCR_REPO_PREFIX := gcr.io/$(GCLOUD_PROJECT_ID)


#----------------------------------------------------------------------------------
# Marcos
#----------------------------------------------------------------------------------

# This macro takes a relative path as its only argument and returns all the files
# in the tree rooted at that directory that match the given criteria.
get_sources = $(shell find $(1) -name "*.go" | grep -v test | grep -v generated.go | grep -v mock_)

# If both GCLOUD_PROJECT_ID and BUILD_ID are set, define a function that takes a docker image name
# and returns a '-t' flag that can be passed to 'docker build' to create a tag for a test image.
# If the function is not defined, any attempt at calling if will return nothing (it does not cause en error)
ifneq ($(GCLOUD_PROJECT_ID),)
ifneq ($(BUILD_ID),)
define get_test_tag
	-t $(GCR_REPO_PREFIX)/$(1):$(TEST_IMAGE_TAG)
endef
endif
endif

#----------------------------------------------------------------------------------
# Repo setup
#----------------------------------------------------------------------------------

# https://www.viget.com/articles/two-ways-to-share-git-hooks-with-your-team/
.PHONY: init
init:
	git config core.hooksPath .githooks

.PHONY: update-deps
update-deps:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/gogo/protobuf/protoc-gen-gogo
	mkdir -p $$GOPATH/src/github.com/envoyproxy
	# use a specific commit (c15f2c24fb27b136e722fa912accddd0c8db9dfa) until v0.0.15 is released, as in v0.0.14 the import paths were not yet changed
	cd $$GOPATH/src/github.com/envoyproxy && if [ ! -e protoc-gen-validate ];then git clone https://github.com/envoyproxy/protoc-gen-validate; fi && cd protoc-gen-validate && git fetch && git checkout c15f2c24fb27b136e722fa912accddd0c8db9dfa
	go get -u github.com/cratonica/2goarray
	go get -v -u github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

.PHONY: pin-repos
pin-repos:
	GO111MODULE=on go run ci/pin_repos/pin_repos.go

.PHONY: check-format
check-format:
	NOT_FORMATTED=$$(gofmt -l ./projects/ ./pkg/ ./test/) && if [ -n "$$NOT_FORMATTED" ]; then echo These files are not formatted: $$NOT_FORMATTED; exit 1; fi

check-spelling:
	./ci/spell.sh check
#----------------------------------------------------------------------------------
# Clean
#----------------------------------------------------------------------------------

# Important to clean before pushing new releases. Dockerfiles and binaries may not update properly
.PHONY: clean
clean:
	rm -rf _output
	rm -rf _test
	rm -rf docs/site*
	rm -rf docs/themes
	rm -rf docs/resources
	git clean -f -X install

#----------------------------------------------------------------------------------
# Generated Code and Docs
#----------------------------------------------------------------------------------

.PHONY: generated-code
generated-code: $(OUTPUT_DIR)/.generated-code verify-enterprise-protos update-licenses

# Note: currently we generate CLI docs, but don't push them to the consolidated docs repo (gloo-docs). Instead, the
# Glooctl enterprise docs are pushed from the private repo.
# TODO(EItanya): make mockgen work for gloo
SUBDIRS:=$(shell ls -d -- */ | grep -v vendor)
$(OUTPUT_DIR)/.generated-code:
	find . -name *.sk.md | xargs rm
	rm docs/content/cli/glooctl*; GO111MODULE=on go run projects/gloo/cli/cmd/docs/main.go
	GO111MODULE=on go generate ./...
	gofmt -w $(SUBDIRS)
	goimports -w $(SUBDIRS)
	mkdir -p $(OUTPUT_DIR)
	touch $@

# Make sure that the enterprise API *.pb.go files that are generated but not used in this repo are valid.
.PHONY: verify-enterprise-protos
verify-enterprise-protos:
	@echo Verifying validity of generated enterprise files...
	$(GO_BUILD_FLAGS) GOOS=linux go build projects/gloo/pkg/api/v1/enterprise/verify.go

#----------------------------------------------------------------------------------
# Generate mocks
#----------------------------------------------------------------------------------

# The values in this array are used in a foreach loop to dynamically generate the
# commands in the generate-client-mocks target.
# For each value, the ":" character will be replaced with " " using the subst function,
# thus turning the string into a 3-element array. The n-th element of the array will
# then be selected via the word function
MOCK_RESOURCE_INFO := \
	gloo:artifact:ArtifactClient \
	gloo:endpoint:EndpointClient \
	gloo:proxy:ProxyClient \
	gloo:secret:SecretClient \
	gloo:settings:SettingsClient \
	gloo:upstream:UpstreamClient \
	gateway:gateway:GatewayClient \
	gateway:virtual_service:VirtualServiceClient\
	gateway:route_table:RouteTableClient\

# Use gomock (https://github.com/golang/mock) to generate mocks for our resource clients.
.PHONY: generate-client-mocks
generate-client-mocks:
	@$(foreach INFO, $(MOCK_RESOURCE_INFO), \
		echo Generating mock for $(word 3,$(subst :, , $(INFO)))...; \
		mockgen -destination=projects/$(word 1,$(subst :, , $(INFO)))/pkg/mocks/mock_$(word 2,$(subst :, , $(INFO)))_client.go \
     		-package=mocks \
     		github.com/solo-io/gloo/projects/$(word 1,$(subst :, , $(INFO)))/pkg/api/v1 \
     		$(word 3,$(subst :, , $(INFO))) \
     	;)

#----------------------------------------------------------------------------------
# glooctl
#----------------------------------------------------------------------------------

CLI_DIR=projects/gloo/cli

$(OUTPUT_DIR)/glooctl: $(SOURCES)
	GO111MODULE=on go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(CLI_DIR)/cmd/main.go

$(OUTPUT_DIR)/glooctl-linux-amd64: $(SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(CLI_DIR)/cmd/main.go

$(OUTPUT_DIR)/glooctl-darwin-amd64: $(SOURCES)
	$(GO_BUILD_FLAGS) GOOS=darwin go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(CLI_DIR)/cmd/main.go

$(OUTPUT_DIR)/glooctl-windows-amd64.exe: $(SOURCES)
	$(GO_BUILD_FLAGS) GOOS=windows go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(CLI_DIR)/cmd/main.go


.PHONY: glooctl
glooctl: $(OUTPUT_DIR)/glooctl
.PHONY: glooctl-linux-amd64
glooctl-linux-amd64: $(OUTPUT_DIR)/glooctl-linux-amd64
.PHONY: glooctl-darwin-amd64
glooctl-darwin-amd64: $(OUTPUT_DIR)/glooctl-darwin-amd64
.PHONY: glooctl-windows-amd64
glooctl-windows-amd64: $(OUTPUT_DIR)/glooctl-windows-amd64.exe

.PHONY: build-cli
build-cli: glooctl-linux-amd64 glooctl-darwin-amd64 glooctl-windows-amd64

#----------------------------------------------------------------------------------
# Gateway
#----------------------------------------------------------------------------------

GATEWAY_DIR=projects/gateway
GATEWAY_SOURCES=$(call get_sources,$(GATEWAY_DIR))

$(OUTPUT_DIR)/gateway-linux-amd64: $(GATEWAY_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(GATEWAY_DIR)/cmd/main.go


.PHONY: gateway
gateway: $(OUTPUT_DIR)/gateway-linux-amd64

$(OUTPUT_DIR)/Dockerfile.gateway: $(GATEWAY_DIR)/cmd/Dockerfile
	cp $< $@

gateway-docker: $(OUTPUT_DIR)/gateway-linux-amd64 $(OUTPUT_DIR)/Dockerfile.gateway
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.gateway \
		-t quay.io/solo-io/gateway:$(VERSION) \
		$(call get_test_tag,gateway)

#----------------------------------------------------------------------------------
# Ingress
#----------------------------------------------------------------------------------

INGRESS_DIR=projects/ingress
INGRESS_SOURCES=$(call get_sources,$(INGRESS_DIR))

$(OUTPUT_DIR)/ingress-linux-amd64: $(INGRESS_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(INGRESS_DIR)/cmd/main.go


.PHONY: ingress
ingress: $(OUTPUT_DIR)/ingress-linux-amd64

$(OUTPUT_DIR)/Dockerfile.ingress: $(INGRESS_DIR)/cmd/Dockerfile
	cp $< $@

ingress-docker: $(OUTPUT_DIR)/ingress-linux-amd64 $(OUTPUT_DIR)/Dockerfile.ingress
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.ingress \
		-t quay.io/solo-io/ingress:$(VERSION) \
		$(call get_test_tag,ingress)

#----------------------------------------------------------------------------------
# Access Logger
#----------------------------------------------------------------------------------

ACCESS_LOG_DIR=projects/accesslogger
ACCESS_LOG_SOURCES=$(call get_sources,$(ACCESS_LOG_DIR))

$(OUTPUT_DIR)/access-logger-linux-amd64: $(ACCESS_LOG_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(ACCESS_LOG_DIR)/cmd/main.go


.PHONY: access-logger
access-logger: $(OUTPUT_DIR)/access-logger-linux-amd64

$(OUTPUT_DIR)/Dockerfile.access-logger: $(ACCESS_LOG_DIR)/cmd/Dockerfile
	cp $< $@

access-logger-docker: $(OUTPUT_DIR)/access-logger-linux-amd64 $(OUTPUT_DIR)/Dockerfile.access-logger
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.access-logger \
		-t quay.io/solo-io/access-logger:$(VERSION) \
		$(call get_test_tag,access-logger)

#----------------------------------------------------------------------------------
# Discovery
#----------------------------------------------------------------------------------

DISCOVERY_DIR=projects/discovery
DISCOVERY_SOURCES=$(call get_sources,$(DISCOVERY_DIR))

$(OUTPUT_DIR)/discovery-linux-amd64: $(DISCOVERY_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(DISCOVERY_DIR)/cmd/main.go


.PHONY: discovery
discovery: $(OUTPUT_DIR)/discovery-linux-amd64

$(OUTPUT_DIR)/Dockerfile.discovery: $(DISCOVERY_DIR)/cmd/Dockerfile
	cp $< $@

discovery-docker: $(OUTPUT_DIR)/discovery-linux-amd64 $(OUTPUT_DIR)/Dockerfile.discovery
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.discovery \
		-t quay.io/solo-io/discovery:$(VERSION) \
		$(call get_test_tag,discovery)

#----------------------------------------------------------------------------------
# Gloo
#----------------------------------------------------------------------------------

GLOO_DIR=projects/gloo
GLOO_SOURCES=$(call get_sources,$(GLOO_DIR))

$(OUTPUT_DIR)/gloo-linux-amd64: $(GLOO_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(GLOO_DIR)/cmd/main.go


.PHONY: gloo
gloo: $(OUTPUT_DIR)/gloo-linux-amd64

$(OUTPUT_DIR)/Dockerfile.gloo: $(GLOO_DIR)/cmd/Dockerfile
	cp hack/utils/oss_compliance/third_party_licenses.txt $(OUTPUT_DIR)/third_party_licenses.txt
	cp $< $@

gloo-docker: $(OUTPUT_DIR)/gloo-linux-amd64 $(OUTPUT_DIR)/Dockerfile.gloo
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.gloo \
		-t quay.io/solo-io/gloo:$(VERSION) \
		$(call get_test_tag,gloo)

#----------------------------------------------------------------------------------
# Envoy init (BASE)
#----------------------------------------------------------------------------------

ENVOYINIT_DIR=projects/envoyinit/cmd
ENVOYINIT_SOURCES=$(call get_sources,$(ENVOYINIT_DIR))

$(OUTPUT_DIR)/envoyinit-linux-amd64: $(ENVOYINIT_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(ENVOYINIT_DIR)/main.go

.PHONY: envoyinit
envoyinit: $(OUTPUT_DIR)/envoyinit-linux-amd64


$(OUTPUT_DIR)/Dockerfile.envoyinit: $(ENVOYINIT_DIR)/Dockerfile.envoyinit
	cp $< $@

.PHONY: gloo-envoy-wrapper-docker
gloo-envoy-wrapper-docker: $(OUTPUT_DIR)/envoyinit-linux-amd64 $(OUTPUT_DIR)/Dockerfile.envoyinit
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.envoyinit \
		-t quay.io/solo-io/gloo-envoy-wrapper:$(VERSION) \
		$(call get_test_tag,gloo-envoy-wrapper)

#----------------------------------------------------------------------------------
# Envoy init (WASM)
#----------------------------------------------------------------------------------

ENVOY_WASM_DIR=projects/envoyinit/cmd
ENVOY_WASM_SOURCES=$(call get_sources,$(ENVOY_WASM_DIR))

$(OUTPUT_DIR)/envoywasm-linux-amd64: $(ENVOY_WASM_SOURCES)
	$(GO_BUILD_FLAGS) GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(ENVOY_WASM_DIR)/main.go

.PHONY: envoywasm
envoywasm: $(OUTPUT_DIR)/envoywasm-linux-amd64


$(OUTPUT_DIR)/Dockerfile.envoywasm: $(ENVOY_WASM_DIR)/Dockerfile.envoywasm
	cp $< $@

.PHONY: gloo-envoy-wasm-wrapper-docker
gloo-envoy-wasm-wrapper-docker: $(OUTPUT_DIR)/envoywasm-linux-amd64 $(OUTPUT_DIR)/Dockerfile.envoywasm
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.envoywasm \
		-t quay.io/solo-io/gloo-envoy-wasm-wrapper:$(VERSION) \
		$(call get_test_tag,gloo-envoy-wasm-wrapper)


#----------------------------------------------------------------------------------
# Certgen - Job for creating TLS Secrets in Kubernetes
#----------------------------------------------------------------------------------

CERTGEN_DIR=jobs/certgen/cmd
CERTGEN_SOURCES=$(call get_sources,$(CERTGEN_DIR))

$(OUTPUT_DIR)/certgen-linux-amd64: $(CERTGEN_SOURCES)
	GO111MODULE=on CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) -o $@ $(CERTGEN_DIR)/main.go

.PHONY: certgen
certgen: $(OUTPUT_DIR)/certgen-linux-amd64


$(OUTPUT_DIR)/Dockerfile.certgen: $(CERTGEN_DIR)/Dockerfile
	cp $< $@

.PHONY: certgen-docker
certgen-docker: $(OUTPUT_DIR)/certgen-linux-amd64 $(OUTPUT_DIR)/Dockerfile.certgen
	docker build $(OUTPUT_DIR) -f $(OUTPUT_DIR)/Dockerfile.certgen \
		-t quay.io/solo-io/certgen:$(VERSION) \
		$(call get_test_tag,certgen)


#----------------------------------------------------------------------------------
# Build All
#----------------------------------------------------------------------------------
.PHONY: build
build: gloo glooctl gateway discovery envoyinit certgen ingress

#----------------------------------------------------------------------------------
# Deployment Manifests / Helm
#----------------------------------------------------------------------------------

HELM_SYNC_DIR := $(OUTPUT_DIR)/helm
HELM_DIR := install/helm/gloo

# Creates Chart.yaml and values.yaml. See install/helm/gloo/README.md for more info.
.PHONY: generate-helm-files
generate-helm-files: $(OUTPUT_DIR)/.helm-prepared

$(OUTPUT_DIR)/.helm-prepared:
	GO111MODULE=on go run $(HELM_DIR)/generate.go $(VERSION)
	touch $@

package-chart: generate-helm-files
	mkdir -p $(HELM_SYNC_DIR)/charts
	helm package --destination $(HELM_SYNC_DIR)/charts $(HELM_DIR)
	helm repo index $(HELM_SYNC_DIR)

.PHONY: fetch-helm
fetch-helm:
	gsutil -m rsync -r gs://solo-public-helm/ './_output/helm'

.PHONY: save-helm
save-helm:
ifeq ($(RELEASE),"true")
	gsutil -m rsync -r './_output/helm' gs://solo-public-helm/
endif

#----------------------------------------------------------------------------------
# Build the Gloo Manifests that are published as release assets
#----------------------------------------------------------------------------------

.PHONY: render-manifests
render-manifests: install/gloo-gateway.yaml install/gloo-ingress.yaml install/gloo-knative.yaml

INSTALL_NAMESPACE ?= gloo-system

MANIFEST_OUTPUT = > /dev/null
ifneq ($(BUILD_ID),)
MANIFEST_OUTPUT =
endif

define HELM_VALUES
namespace:
  create: true
crds:
  create: true
endef

# Export as a shell variable, make variables do not play well with multiple lines
export HELM_VALUES
$(OUTPUT_DIR)/release-manifest-values.yaml:
	@echo "$$HELM_VALUES" > $@

install/gloo-gateway.yaml: $(OUTPUT_DIR)/glooctl-linux-amd64 $(OUTPUT_DIR)/release-manifest-values.yaml package-chart
ifeq ($(RELEASE),"true")
	$(OUTPUT_DIR)/glooctl-linux-amd64 install gateway -n $(INSTALL_NAMESPACE) -f $(HELM_SYNC_DIR)/charts/gloo-$(VERSION).tgz \
		--values $(OUTPUT_DIR)/release-manifest-values.yaml --dry-run | tee $@ $(OUTPUT_YAML) $(MANIFEST_OUTPUT)
endif

install/gloo-knative.yaml: $(OUTPUT_DIR)/glooctl-linux-amd64 $(OUTPUT_DIR)/release-manifest-values.yaml package-chart
ifeq ($(RELEASE),"true")
	$(OUTPUT_DIR)/glooctl-linux-amd64 install knative -n $(INSTALL_NAMESPACE) -f $(HELM_SYNC_DIR)/charts/gloo-$(VERSION).tgz \
		--values $(OUTPUT_DIR)/release-manifest-values.yaml --dry-run | tee $@ $(OUTPUT_YAML) $(MANIFEST_OUTPUT)
endif

install/gloo-ingress.yaml: $(OUTPUT_DIR)/glooctl-linux-amd64 $(OUTPUT_DIR)/release-manifest-values.yaml package-chart
ifeq ($(RELEASE),"true")
	$(OUTPUT_DIR)/glooctl-linux-amd64 install ingress -n $(INSTALL_NAMESPACE) -f $(HELM_SYNC_DIR)/charts/gloo-$(VERSION).tgz \
		--values $(OUTPUT_DIR)/release-manifest-values.yaml --dry-run | tee $@ $(OUTPUT_YAML) $(MANIFEST_OUTPUT)
endif

#----------------------------------------------------------------------------------
# Release
#----------------------------------------------------------------------------------
GLOOE_CHANGELOGS_BUCKET=gloo-ee-changelogs

$(OUTPUT_DIR)/gloo-enterprise-version:
	GO111MODULE=on go run hack/find_latest_enterprise_version.go

.PHONY: download-glooe-changelog
download-glooe-changelog: $(OUTPUT_DIR)/gloo-enterprise-version
ifeq ($(RELEASE),"true")
	mkdir -p '../solo-projects/changelog'
	gsutil -m cp -r gs://$(GLOOE_CHANGELOGS_BUCKET)/$(shell cat $(OUTPUT_DIR)/gloo-enterprise-version)/* '../solo-projects/changelog'
endif

ASSETS_ONLY := false

# The code does the proper checking for a TAGGED_VERSION
.PHONY: upload-github-release-assets
upload-github-release-assets: build-cli render-manifests
	GO111MODULE=on go run ci/upload_github_release_assets.go $(ASSETS_ONLY)

.PHONY: publish-docs
publish-docs: generate-helm-files
ifeq ($(RELEASE),"true")
	cd docs && make docker-push-docs \
		VERSION=$(VERSION) \
		TAGGED_VERSION=$(TAGGED_VERSION) \
		GCLOUD_PROJECT_ID=$(GCLOUD_PROJECT_ID) \
		RELEASE=$(RELEASE)
endif

## START Temporary targets while master/v0.20.x branch is supported
# Whenever a new Gloo release is made from the master/v0.20.x branch,
# run `GCLOUD_PROJECT_ID=solo-public TAGGED_VERSION=v0-20-<PARTICULAR_VERSION> make publish-v-20-docs -B`
# to manually to produce a docs image with the following properties:
# - api, cli, changelog content from master/v0.20.x
# - docs content, hosting config from feature-rc1
.PHONY: publish-v-20-docs
publish-v-20-docs: gather-v-20-docs publish-docs

OLD_DOCS_TEMP_DIR := "old_docs_temp_dir"
.PHONY: gather-v-20-docs
gather-v-20-docs:
	# want the changelog files to reflect master
	git checkout master -- changelog/
	rm -rf $(OLD_DOCS_TEMP_DIR)
	mkdir -p $(OLD_DOCS_TEMP_DIR)
	cd $(OLD_DOCS_TEMP_DIR) && git clone git@github.com:solo-io/solo-docs.git
	rm docs/content/cli/glooctl*
	cp $(OLD_DOCS_TEMP_DIR)/solo-docs/gloo/docs/cli/glooctl* docs/content/cli/
	rm -rf docs/content/api/
	cp -r $(OLD_DOCS_TEMP_DIR)/solo-docs/gloo/docs/api/ docs/content/api/
	rm -rf $(OLD_DOCS_TEMP_DIR)
## END Temporary targets while master/v0.20.x branch is supported

#----------------------------------------------------------------------------------
# Docker
#----------------------------------------------------------------------------------
#
#---------
#--------- Push
#---------

DOCKER_IMAGES :=
ifeq ($(RELEASE),"true")
	DOCKER_IMAGES := docker
endif

.PHONY: docker docker-push
docker: discovery-docker gateway-docker gloo-docker \
 		gloo-envoy-wrapper-docker gloo-envoy-wasm-wrapper-docker \
 		certgen-docker ingress-docker access-logger-docker

# Depends on DOCKER_IMAGES, which is set to docker if RELEASE is "true", otherwise empty (making this a no-op).
# This prevents executing the dependent targets if RELEASE is not true, while still enabling `make docker`
# to be used for local testing.
# docker-push is intended to be run by CI
docker-push: $(DOCKER_IMAGES)
ifeq ($(RELEASE),"true")
	docker push quay.io/solo-io/gateway:$(VERSION) && \
	docker push quay.io/solo-io/ingress:$(VERSION) && \
	docker push quay.io/solo-io/discovery:$(VERSION) && \
	docker push quay.io/solo-io/gloo:$(VERSION) && \
	docker push quay.io/solo-io/gloo-envoy-wrapper:$(VERSION) && \
	docker push quay.io/solo-io/gloo-envoy-wasm-wrapper:$(VERSION) && \
	docker push quay.io/solo-io/certgen:$(VERSION) && \
	docker push quay.io/solo-io/access-logger:$(VERSION)
endif

push-kind-images: docker
	kind load docker-image quay.io/solo-io/gateway:$(VERSION) --name $(CLUSTER_NAME)
	kind load docker-image quay.io/solo-io/ingress:$(VERSION) --name $(CLUSTER_NAME)
	kind load docker-image quay.io/solo-io/discovery:$(VERSION) --name $(CLUSTER_NAME)
	kind load docker-image quay.io/solo-io/gloo:$(VERSION) --name $(CLUSTER_NAME)
	kind load docker-image quay.io/solo-io/gloo-envoy-wrapper:$(VERSION) --name $(CLUSTER_NAME)
	kind load docker-image quay.io/solo-io/certgen:$(VERSION) --name $(CLUSTER_NAME)


#----------------------------------------------------------------------------------
# Build assets for Kube2e tests
#----------------------------------------------------------------------------------
#
# The following targets are used to generate the assets on which the kube2e tests rely upon. The following actions are performed:
#
#   1. Push the images to GCR (images have been tagged as $(GCR_REPO_PREFIX)/<image-name>:$(TEST_IMAGE_TAG)
#   2. Generate Gloo value files providing overrides to make the image elements point to GCR
#      - override the repository prefix for all repository names (e.g. quay.io/solo-io/gateway -> gcr.io/solo-public/gateway)
#      - set the tag for each image to TEST_IMAGE_TAG
#   3. Package the Gloo Helm chart to the _test directory (also generate an index file)
#
# The Kube2e tests will use the generated Gloo Chart to install Gloo to the GKE test cluster.

.PHONY: build-test-assets
build-test-assets: push-test-images build-test-chart $(OUTPUT_DIR)/glooctl-linux-amd64 \
 	$(OUTPUT_DIR)/glooctl-darwin-amd64

.PHONY: build-kind-assets
build-kind-assets: push-kind-images build-kind-chart $(OUTPUT_DIR)/glooctl-linux-amd64 \
 	$(OUTPUT_DIR)/glooctl-darwin-amd64

TEST_DOCKER_TARGETS := gateway-docker-test ingress-docker-test discovery-docker-test gloo-docker-test gloo-envoy-wrapper-docker-test certgen-docker-test

.PHONY: push-test-images $(TEST_DOCKER_TARGETS)
push-test-images: $(TEST_DOCKER_TARGETS)

gateway-docker-test: $(OUTPUT_DIR)/gateway-linux-amd64 $(OUTPUT_DIR)/Dockerfile.gateway
	docker push $(GCR_REPO_PREFIX)/gateway:$(TEST_IMAGE_TAG)

ingress-docker-test: $(OUTPUT_DIR)/ingress-linux-amd64 $(OUTPUT_DIR)/Dockerfile.ingress
	docker push $(GCR_REPO_PREFIX)/ingress:$(TEST_IMAGE_TAG)

discovery-docker-test: $(OUTPUT_DIR)/discovery-linux-amd64 $(OUTPUT_DIR)/Dockerfile.discovery
	docker push $(GCR_REPO_PREFIX)/discovery:$(TEST_IMAGE_TAG)

gloo-docker-test: $(OUTPUT_DIR)/gloo-linux-amd64 $(OUTPUT_DIR)/Dockerfile.gloo
	docker push $(GCR_REPO_PREFIX)/gloo:$(TEST_IMAGE_TAG)

gloo-envoy-wrapper-docker-test: $(OUTPUT_DIR)/envoyinit-linux-amd64 $(OUTPUT_DIR)/Dockerfile.envoyinit
	docker push $(GCR_REPO_PREFIX)/gloo-envoy-wrapper:$(TEST_IMAGE_TAG)

certgen-docker-test: $(OUTPUT_DIR)/certgen-linux-amd64 $(OUTPUT_DIR)/Dockerfile.certgen
	docker push $(GCR_REPO_PREFIX)/certgen:$(TEST_IMAGE_TAG)

.PHONY: build-test-chart
build-test-chart:
	mkdir -p $(TEST_ASSET_DIR)
	GO111MODULE=on go run $(HELM_DIR)/generate.go $(TEST_IMAGE_TAG) $(GCR_REPO_PREFIX) "Always"
	helm package --destination $(TEST_ASSET_DIR) $(HELM_DIR)
	helm repo index $(TEST_ASSET_DIR)

.PHONY: build-kind-chart
build-kind-chart:
	mkdir -p $(TEST_ASSET_DIR)
	GO111MODULE=on go run $(HELM_DIR)/generate.go $(VERSION)
	helm package --destination $(TEST_ASSET_DIR) $(HELM_DIR)
	helm repo index $(TEST_ASSET_DIR)


#----------------------------------------------------------------------------------
# Build assets for non-release charts (meant to be invoked on your dev machine)
#----------------------------------------------------------------------------------

# Must be run on your dev machine because TAGGED_VERSION depends on `git describe --tags`, and our CI doesn't clone
# the repo or have any git context

# Sample invocation:
# TAGGED_VERSION=$(git describe --tags) GCLOUD_PROJECT_ID=solo-public make clean fetch-tagged-helm build-tagged-chart save-tagged-helm

.PHONY: build-tagged-chart
build-tagged-chart:
	mkdir -p $(TEST_ASSET_DIR)
	GO111MODULE=on go run $(HELM_DIR)/generate.go $(TAGGED_VERSION) $(GCR_REPO_PREFIX) "Always"
	mkdir -p $(HELM_SYNC_DIR)/charts
	helm package --destination $(HELM_SYNC_DIR)/charts $(HELM_DIR)
	helm repo index $(HELM_SYNC_DIR)

.PHONY: save-tagged-helm
save-tagged-helm:
	gsutil -m rsync -r './_output/helm' gs://solo-public-tagged-helm/

.PHONY: fetch-tagged-helm
fetch-tagged-helm:
	mkdir -p $(HELM_SYNC_DIR)
	gsutil -m rsync -r gs://solo-public-tagged-helm/ './_output/helm'

#----------------------------------------------------------------------------------
# Third Party License Management
#----------------------------------------------------------------------------------
.PHONY: update-licenses
update-licenses:
# TODO(helm3): fix after we completely drop toml parsing in favor of go modules
#	cd hack/utils/oss_compliance && GO111MODULE=on go run main.go

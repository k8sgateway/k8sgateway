// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/waf/waf.proto

package waf

import (
	reflect "reflect"
	sync "sync"

	waf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/waf"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Disable waf on this resource (if omitted defaults to false).
	// If a route/virtual host is configured with WAF, you must explicitly disable its WAF,
	// i.e., it will not inherit the disabled status of its parent
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Custom massage to display if an intervention occurs.
	CustomInterventionMessage string `protobuf:"bytes,2,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	// Add OWASP core rule set
	// if nil will not be added
	CoreRuleSet *CoreRuleSet `protobuf:"bytes,3,opt,name=core_rule_set,json=coreRuleSet,proto3" json:"core_rule_set,omitempty"`
	// Custom rule sets to add. Any subsequent changes to the rules in these files are not automatically updated. To update rules from files, version and update the file name.
	// If you want dynamically updated rules, use the `configMapRuleSets` option instead.
	RuleSets []*waf.RuleSet `protobuf:"bytes,4,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	// Kubernetes configmaps with the rule sets that you want to use.
	// The rules must be in the value of the key-value mappings in the `data` field of the configmap.
	// Subsequent updates to the configmap values are dynamically updated in the configuration.
	ConfigMapRuleSets []*RuleSetFromConfigMap `protobuf:"bytes,8,rep,name=config_map_rule_sets,json=configMapRuleSets,proto3" json:"config_map_rule_sets,omitempty"`
	// Audit Log settings
	AuditLogging *waf.AuditLogging `protobuf:"bytes,5,opt,name=audit_logging,json=auditLogging,proto3" json:"audit_logging,omitempty"`
	// Only process request headers, not buffering the request body
	RequestHeadersOnly bool `protobuf:"varint,6,opt,name=request_headers_only,json=requestHeadersOnly,proto3" json:"request_headers_only,omitempty"`
	// Only process response headers, not buffering the response body
	ResponseHeadersOnly bool `protobuf:"varint,7,opt,name=response_headers_only,json=responseHeadersOnly,proto3" json:"response_headers_only,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Settings) GetCustomInterventionMessage() string {
	if x != nil {
		return x.CustomInterventionMessage
	}
	return ""
}

func (x *Settings) GetCoreRuleSet() *CoreRuleSet {
	if x != nil {
		return x.CoreRuleSet
	}
	return nil
}

func (x *Settings) GetRuleSets() []*waf.RuleSet {
	if x != nil {
		return x.RuleSets
	}
	return nil
}

func (x *Settings) GetConfigMapRuleSets() []*RuleSetFromConfigMap {
	if x != nil {
		return x.ConfigMapRuleSets
	}
	return nil
}

func (x *Settings) GetAuditLogging() *waf.AuditLogging {
	if x != nil {
		return x.AuditLogging
	}
	return nil
}

func (x *Settings) GetRequestHeadersOnly() bool {
	if x != nil {
		return x.RequestHeadersOnly
	}
	return false
}

func (x *Settings) GetResponseHeadersOnly() bool {
	if x != nil {
		return x.ResponseHeadersOnly
	}
	return false
}

type RuleSetFromConfigMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes configmap that has the rule sets as values in the `data` section.
	ConfigmapLocation *core.ResourceRef `protobuf:"bytes,1,opt,name=configmap_location,json=configmapLocation,proto3" json:"configmap_location,omitempty"`
	// The configmap might have multiple key-value pairs in the `data` section, such as when you create the configmap from multiple files.
	// You can use the `dataMapKey` field to select which rules and the order you want them included.
	// If this field is included, only the specified keys are applied in order. Any rules not included are ignored.
	// If this field is not included, all of the rules in the `data` section of the configmap are sorted and applied. The order might differ from their order in the configmap.
	DataMapKeys []string `protobuf:"bytes,2,rep,name=data_map_keys,json=dataMapKeys,proto3" json:"data_map_keys,omitempty"`
}

func (x *RuleSetFromConfigMap) Reset() {
	*x = RuleSetFromConfigMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleSetFromConfigMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleSetFromConfigMap) ProtoMessage() {}

func (x *RuleSetFromConfigMap) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleSetFromConfigMap.ProtoReflect.Descriptor instead.
func (*RuleSetFromConfigMap) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescGZIP(), []int{1}
}

func (x *RuleSetFromConfigMap) GetConfigmapLocation() *core.ResourceRef {
	if x != nil {
		return x.ConfigmapLocation
	}
	return nil
}

func (x *RuleSetFromConfigMap) GetDataMapKeys() []string {
	if x != nil {
		return x.DataMapKeys
	}
	return nil
}

type CoreRuleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional custom settings for the OWASP core rule set.
	// For an example on the configuration options see: https://github.com/SpiderLabs/owasp-modsecurity-crs/blob/v3.2/dev/crs-setup.conf.example
	// The same rules apply to these options as do to the `RuleSet`s. The file option is better if possible.
	//
	// Types that are assignable to CustomSettingsType:
	//	*CoreRuleSet_CustomSettingsString
	//	*CoreRuleSet_CustomSettingsFile
	CustomSettingsType isCoreRuleSet_CustomSettingsType `protobuf_oneof:"CustomSettingsType"`
}

func (x *CoreRuleSet) Reset() {
	*x = CoreRuleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreRuleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreRuleSet) ProtoMessage() {}

func (x *CoreRuleSet) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreRuleSet.ProtoReflect.Descriptor instead.
func (*CoreRuleSet) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescGZIP(), []int{2}
}

func (m *CoreRuleSet) GetCustomSettingsType() isCoreRuleSet_CustomSettingsType {
	if m != nil {
		return m.CustomSettingsType
	}
	return nil
}

func (x *CoreRuleSet) GetCustomSettingsString() string {
	if x, ok := x.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsString); ok {
		return x.CustomSettingsString
	}
	return ""
}

func (x *CoreRuleSet) GetCustomSettingsFile() string {
	if x, ok := x.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsFile); ok {
		return x.CustomSettingsFile
	}
	return ""
}

type isCoreRuleSet_CustomSettingsType interface {
	isCoreRuleSet_CustomSettingsType()
}

type CoreRuleSet_CustomSettingsString struct {
	// String representing the core rule set custom config options
	CustomSettingsString string `protobuf:"bytes,2,opt,name=custom_settings_string,json=customSettingsString,proto3,oneof"`
}

type CoreRuleSet_CustomSettingsFile struct {
	// String representing a file location with core rule set custom config options
	CustomSettingsFile string `protobuf:"bytes,3,opt,name=custom_settings_file,json=customSettingsFile,proto3,oneof"`
}

func (*CoreRuleSet_CustomSettingsString) isCoreRuleSet_CustomSettingsType() {}

func (*CoreRuleSet_CustomSettingsFile) isCoreRuleSet_CustomSettingsType() {}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x77, 0x61, 0x66, 0x2f, 0x77, 0x61, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x77, 0x61, 0x66, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77,
	0x61, 0x66, 0x2f, 0x77, 0x61, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3,
	0x04, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x1b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x77, 0x61, 0x66, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x73, 0x12, 0x5f, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x77, 0x61, 0x66, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x52,
	0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x73, 0x12, 0x5a, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x30,
	0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x79,
	0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x48, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6d, 0x61, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x66, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x6d, 0x61, 0x70, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x61, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x72, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x14, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x54, 0x79, 0x70, 0x65, 0x42, 0x55, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x66, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01,
	0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_goTypes = []interface{}{
	(*Settings)(nil),             // 0: waf.options.gloo.solo.io.Settings
	(*RuleSetFromConfigMap)(nil), // 1: waf.options.gloo.solo.io.RuleSetFromConfigMap
	(*CoreRuleSet)(nil),          // 2: waf.options.gloo.solo.io.CoreRuleSet
	(*waf.RuleSet)(nil),          // 3: envoy.config.filter.http.modsecurity.v2.RuleSet
	(*waf.AuditLogging)(nil),     // 4: envoy.config.filter.http.modsecurity.v2.AuditLogging
	(*core.ResourceRef)(nil),     // 5: core.solo.io.ResourceRef
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_depIdxs = []int32{
	2, // 0: waf.options.gloo.solo.io.Settings.core_rule_set:type_name -> waf.options.gloo.solo.io.CoreRuleSet
	3, // 1: waf.options.gloo.solo.io.Settings.rule_sets:type_name -> envoy.config.filter.http.modsecurity.v2.RuleSet
	1, // 2: waf.options.gloo.solo.io.Settings.config_map_rule_sets:type_name -> waf.options.gloo.solo.io.RuleSetFromConfigMap
	4, // 3: waf.options.gloo.solo.io.Settings.audit_logging:type_name -> envoy.config.filter.http.modsecurity.v2.AuditLogging
	5, // 4: waf.options.gloo.solo.io.RuleSetFromConfigMap.configmap_location:type_name -> core.solo.io.ResourceRef
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleSetFromConfigMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreRuleSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CoreRuleSet_CustomSettingsString)(nil),
		(*CoreRuleSet_CustomSettingsFile)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_waf_waf_proto_depIdxs = nil
}

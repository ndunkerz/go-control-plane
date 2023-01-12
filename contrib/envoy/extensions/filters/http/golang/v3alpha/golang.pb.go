// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: contrib/envoy/extensions/filters/http/golang/v3alpha/golang.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any1 "github.com/golang/protobuf/ptypes/any"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config_MergePolicy int32

const (
	Config_MERGE_VIRTUALHOST_ROUTER_FILTER Config_MergePolicy = 0
	Config_MERGE_VIRTUALHOST_ROUTER        Config_MergePolicy = 1
	Config_OVERRIDE                        Config_MergePolicy = 3
)

// Enum value maps for Config_MergePolicy.
var (
	Config_MergePolicy_name = map[int32]string{
		0: "MERGE_VIRTUALHOST_ROUTER_FILTER",
		1: "MERGE_VIRTUALHOST_ROUTER",
		3: "OVERRIDE",
	}
	Config_MergePolicy_value = map[string]int32{
		"MERGE_VIRTUALHOST_ROUTER_FILTER": 0,
		"MERGE_VIRTUALHOST_ROUTER":        1,
		"OVERRIDE":                        3,
	}
)

func (x Config_MergePolicy) Enum() *Config_MergePolicy {
	p := new(Config_MergePolicy)
	*p = x
	return p
}

func (x Config_MergePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_MergePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_enumTypes[0].Descriptor()
}

func (Config_MergePolicy) Type() protoreflect.EnumType {
	return &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_enumTypes[0]
}

func (x Config_MergePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_MergePolicy.Descriptor instead.
func (Config_MergePolicy) EnumDescriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescGZIP(), []int{0, 0}
}

// [#protodoc-title: golang extension filter]
// Golang :ref:`configuration overview <config_http_filters_golang>`.
// [#extension: envoy.filters.http.golang]
//
// In the below example, we configured the go plugin 'auth' and 'limit' dynamic libraries into
// Envoy, which can avoid rebuilding Envoy.
//
// * Develop go-plugin
//
// We can implement the interface of ``StreamFilter <contrib/golang/filters/http/source/go/pkg/api.StreamFilter>``
// API by the GO language to achieve the effects of Envoy native filter.
//
// The filter based on the APIs implementation ``StreamFilter <contrib/golang/filters/http/source/go/pkg/api.StreamFilter>``
// For details, take a look at the :repo:`/contrib/golang/filters/http/test/test_data/echo`.
//
// Then put the GO plugin source code into the ${OUTPUT}/src/ directory with the name of the plugin
// for GO plugin builds.
// The following examples implement limit and auth GO plugins.
//
// .. code-block:: bash
//
//   $ tree /home/admin/envoy/go-plugins/src/
//     |--auth
//     |   |--config.go
//     |   |--filter.go
//     ---limit
//         |--config.go
//         |--filter.go
//
// * Build go-plugin
//
// Build the Go plugin so by `go_plugin_generate.sh` script, below example the `liblimit.so` and
// `libauth.so` will be generated in the `/home/admin/envoy/go-plugins/` directory.
//
// .. code-block:: bash
//
//   #!/bin/bash
//   if [ $# != 2 ]; then
//      echo "need input the go plugin name"
//      exit 1
//   fi
//
//   PLUGINNAME=$1
//   OUTPUT=/home/admin/envoy/go-plugins/
//   PLUGINSRCDIR=${OUTPUT}/src/${PLUGINNAME}
//   go build --buildmode=c-shared  -v -o $OUTPUT/lib${PLUGINNAME}.so $PLUGINSRCDIR
//
// .. code-block:: bash
//
//   $ go_plugin_generate.sh limit
//   $ go_plugin_generate.sh auth
//
// * Configure go-plugin
//
// Use the http filter of :ref: `golang <envoy.filters.http.golang>` to specify
// :ref: `library` <envoy.filters.http.golang> in ingress and egress to enable the plugin.
//
// Example:
//
// .. code-block:: yaml
//
//   static_resources:
//     listeners:
//       - name: ingress
//         address:
//           socket_address:
//             protocol: TCP
//             address: 0.0.0.0
//             port_value: 8080
//         filter_chains:
//           - filters:
//               - name: envoy.filters.network.http_connection_manager
//               ......
//                   http_filters:
//                     - name: envoy.filters.http.golang
//                       typed_config:
//                         "@type": type.googleapis.com/envoy.extensions.filters.http.golang.v3alpha.Config
//                         library_id: limit-id
//                         library_path: "/home/admin/envoy/go-plugins/liblimit.so"
//                         plugine_name: limit
//                         plugin_config:
//                           "@type": type.googleapis.com/envoy.extensions.filters.http.golang.plugins.limit.v3.Config
//                           xxx1: xx1
//                           xxx2: xx2
//                     - name: envoy.filters.http.header_to_metadata
//                     - name: envoy.filters.http.golang
//                       typed_config:
//                         "@type": type.googleapis.com/envoy.extensions.filters.http.golang.v3alpha.Config
//                         library_id: auth-id
//                         library_path: "/home/admin/envoy/go-plugins/libauth.so"
//                         plugine_name: auth
//                         plugin_config:
//                           "@type": type.googleapis.com/envoy.extensions.filters.http.golang.plugins.auth.v3.Config
//                           xxx1: xx1
//                           xxx2: xx2
//                     - name: envoy.filters.http.router
//       - name: egress
//         address:
//           socket_address:
//             protocol: TCP
//             address: 0.0.0.0
//             port_value: 8081
//         filter_chains:
//           - filters:
//               - name: envoy.filters.network.http_connection_manager
//                   ......
//                   http_filters:
//                     - name: envoy.filters.http.golang
//                       typed_config:
//                         "@type": type.googleapis.com/envoy.extensions.filters.http.golang.v3alpha.Config
//                         library_id: auth-id
//                         library_path: "/home/admin/envoy/go-plugins/libauth.so"
//                         plugine_name: auth
//                         plugin_config:
//                           "@type": type.googleapis.com/envoy.extensions.filters.http.golang.plugins.auth.v3.Config
//                           xxx1: xx1
//                           xxx2: xx2
//                     - name: envoy.filters.http.router
// [#next-free-field: 6]
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// library_id is a unique ID for a dynamic library file, must be unique globally.
	LibraryId string `protobuf:"bytes,1,opt,name=library_id,json=libraryId,proto3" json:"library_id,omitempty"`
	// Dynamic library implementing the interface of
	// ``StreamFilter <contrib/golang/filters/http/source/go/pkg/api.StreamFilter>``.
	// [#comment:TODO(wangfakang): Support for downloading libraries from remote repositories.]
	LibraryPath string `protobuf:"bytes,2,opt,name=library_path,json=libraryPath,proto3" json:"library_path,omitempty"`
	// plugin_name is the name of the go plugin, which needs to be consistent with the name
	// registered in http::RegisterHttpFilterConfigFactory.
	PluginName string `protobuf:"bytes,3,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// plugin_config is the configuration of the go plugin, note that this configuration is
	// only parsed in the go plugin.
	PluginConfig *any1.Any `protobuf:"bytes,4,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
	// merge_policy is the merge policy configured by the go plugin.
	// go plugin configuration supports three dimensions: the virtual host’s typed_per_filter_config,
	// the route’s typed_per_filter_config or filter's config.
	// The meanings are as follows:
	// MERGE_VIRTUALHOST_ROUTER_FILTER: pass all configuration into go plugin.
	// MERGE_VIRTUALHOST_ROUTER: pass Virtual-Host and Router configuration into go plugin.
	// OVERRIDE: override according to Router > Virtual_host > Filter priority and pass the
	// configuration to the go plugin.
	MergePolicy Config_MergePolicy `protobuf:"varint,5,opt,name=merge_policy,json=mergePolicy,proto3,enum=envoy.extensions.filters.http.golang.v3alpha.Config_MergePolicy" json:"merge_policy,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLibraryId() string {
	if x != nil {
		return x.LibraryId
	}
	return ""
}

func (x *Config) GetLibraryPath() string {
	if x != nil {
		return x.LibraryPath
	}
	return ""
}

func (x *Config) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *Config) GetPluginConfig() *any1.Any {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

func (x *Config) GetMergePolicy() Config_MergePolicy {
	if x != nil {
		return x.MergePolicy
	}
	return Config_MERGE_VIRTUALHOST_ROUTER_FILTER
}

type RouterPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Example
	//
	// .. code-block:: yaml
	//
	//   typed_per_filter_config:
	//     envoy.filters.http.golang:
	//       "@type": type.googleapis.com/envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute
	//       plugins_config:
	//         plugin1:
	//          disabled: true
	//
	// Types that are assignable to Override:
	//	*RouterPlugin_Disabled
	//	*RouterPlugin_Config
	Override isRouterPlugin_Override `protobuf_oneof:"override"`
}

func (x *RouterPlugin) Reset() {
	*x = RouterPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterPlugin) ProtoMessage() {}

func (x *RouterPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterPlugin.ProtoReflect.Descriptor instead.
func (*RouterPlugin) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescGZIP(), []int{1}
}

func (m *RouterPlugin) GetOverride() isRouterPlugin_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (x *RouterPlugin) GetDisabled() bool {
	if x, ok := x.GetOverride().(*RouterPlugin_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (x *RouterPlugin) GetConfig() *any1.Any {
	if x, ok := x.GetOverride().(*RouterPlugin_Config); ok {
		return x.Config
	}
	return nil
}

type isRouterPlugin_Override interface {
	isRouterPlugin_Override()
}

type RouterPlugin_Disabled struct {
	// [#not-implemented-hide:]
	// Disable the filter for this particular vhost or route.
	// If disabled is specified in multiple per-filter-configs, the most specific one will be used.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type RouterPlugin_Config struct {
	// The config field is used to setting per-route plugin config.
	Config *any1.Any `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

func (*RouterPlugin_Disabled) isRouterPlugin_Override() {}

func (*RouterPlugin_Config) isRouterPlugin_Override() {}

type ConfigsPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// plugins_config is the configuration of the go plugin at the per-router, and
	// key is the name of the go plugin.
	// Example
	//
	// .. code-block:: yaml
	//
	//   typed_per_filter_config:
	//     envoy.filters.http.golang:
	//       "@type": type.googleapis.com/envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute
	//       plugins_config:
	//         plugin1:
	//          disabled: true
	//         plugin2:
	//          config:
	//            "@type": type.googleapis.com/golang.http.plugin2
	//            xxx: xxx
	PluginsConfig map[string]*RouterPlugin `protobuf:"bytes,1,rep,name=plugins_config,json=pluginsConfig,proto3" json:"plugins_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigsPerRoute) Reset() {
	*x = ConfigsPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigsPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigsPerRoute) ProtoMessage() {}

func (x *ConfigsPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigsPerRoute.ProtoReflect.Descriptor instead.
func (*ConfigsPerRoute) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigsPerRoute) GetPluginsConfig() map[string]*RouterPlugin {
	if x != nil {
		return x.PluginsConfig
	}
	return nil
}

var File_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDesc = []byte{
	0x0a, 0x41, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76,
	0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x26, 0x0a, 0x0a, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6d, 0x0a, 0x0c, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x5e, 0x0a, 0x0b, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x52, 0x47, 0x45,
	0x5f, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x52, 0x4f, 0x55,
	0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18,
	0x4d, 0x45, 0x52, 0x47, 0x45, 0x5f, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c, 0x48, 0x4f, 0x53,
	0x54, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56,
	0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x10, 0x03, 0x22, 0x76, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x6a,
	0x02, 0x08, 0x01, 0x48, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x0f, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0x88, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x50, 0x65, 0x72, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x77, 0x0a, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x7c, 0x0a,
	0x12, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x50, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xb0, 0x01, 0x0a, 0x3a,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0b, 0x47, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescData = file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_goTypes = []interface{}{
	(Config_MergePolicy)(0), // 0: envoy.extensions.filters.http.golang.v3alpha.Config.MergePolicy
	(*Config)(nil),          // 1: envoy.extensions.filters.http.golang.v3alpha.Config
	(*RouterPlugin)(nil),    // 2: envoy.extensions.filters.http.golang.v3alpha.RouterPlugin
	(*ConfigsPerRoute)(nil), // 3: envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute
	nil,                     // 4: envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute.PluginsConfigEntry
	(*any1.Any)(nil),        // 5: google.protobuf.Any
}
var file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_depIdxs = []int32{
	5, // 0: envoy.extensions.filters.http.golang.v3alpha.Config.plugin_config:type_name -> google.protobuf.Any
	0, // 1: envoy.extensions.filters.http.golang.v3alpha.Config.merge_policy:type_name -> envoy.extensions.filters.http.golang.v3alpha.Config.MergePolicy
	5, // 2: envoy.extensions.filters.http.golang.v3alpha.RouterPlugin.config:type_name -> google.protobuf.Any
	4, // 3: envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute.plugins_config:type_name -> envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute.PluginsConfigEntry
	2, // 4: envoy.extensions.filters.http.golang.v3alpha.ConfigsPerRoute.PluginsConfigEntry.value:type_name -> envoy.extensions.filters.http.golang.v3alpha.RouterPlugin
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_init() }
func file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_init() {
	if File_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterPlugin); i {
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
		file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigsPerRoute); i {
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
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RouterPlugin_Disabled)(nil),
		(*RouterPlugin_Config)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_depIdxs,
		EnumInfos:         file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_enumTypes,
		MessageInfos:      file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto = out.File
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_http_golang_v3alpha_golang_proto_depIdxs = nil
}
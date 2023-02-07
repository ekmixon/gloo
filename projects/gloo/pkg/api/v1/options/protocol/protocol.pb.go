// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/protocol/protocol.proto

package protocol

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Action to take when Envoy receives client request with header names containing underscore
// characters.
// Underscore character is allowed in header names by the RFC-7230 and this behavior is implemented
// as a security measure due to systems that treat '_' and '-' as interchangeable. Envoy by default allows client request headers with underscore
// characters.
type HttpProtocolOptions_HeadersWithUnderscoresAction int32

const (
	// Allow headers with underscores. This is the default behavior.
	HttpProtocolOptions_ALLOW HttpProtocolOptions_HeadersWithUnderscoresAction = 0
	// Reject client request. HTTP/1 requests are rejected with the 400 status. HTTP/2 requests
	// end with the stream reset. The "httpN.requests_rejected_with_underscores_in_headers" counter
	// is incremented for each rejected request.
	HttpProtocolOptions_REJECT_REQUEST HttpProtocolOptions_HeadersWithUnderscoresAction = 1
	// Drop the header with name containing underscores. The header is dropped before the filter chain is
	// invoked and as such filters will not see dropped headers. The
	// "httpN.dropped_headers_with_underscores" is incremented for each dropped header.
	HttpProtocolOptions_DROP_HEADER HttpProtocolOptions_HeadersWithUnderscoresAction = 2
)

// Enum value maps for HttpProtocolOptions_HeadersWithUnderscoresAction.
var (
	HttpProtocolOptions_HeadersWithUnderscoresAction_name = map[int32]string{
		0: "ALLOW",
		1: "REJECT_REQUEST",
		2: "DROP_HEADER",
	}
	HttpProtocolOptions_HeadersWithUnderscoresAction_value = map[string]int32{
		"ALLOW":          0,
		"REJECT_REQUEST": 1,
		"DROP_HEADER":    2,
	}
)

func (x HttpProtocolOptions_HeadersWithUnderscoresAction) Enum() *HttpProtocolOptions_HeadersWithUnderscoresAction {
	p := new(HttpProtocolOptions_HeadersWithUnderscoresAction)
	*p = x
	return p
}

func (x HttpProtocolOptions_HeadersWithUnderscoresAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpProtocolOptions_HeadersWithUnderscoresAction) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_enumTypes[0].Descriptor()
}

func (HttpProtocolOptions_HeadersWithUnderscoresAction) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_enumTypes[0]
}

func (x HttpProtocolOptions_HeadersWithUnderscoresAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpProtocolOptions_HeadersWithUnderscoresAction.Descriptor instead.
func (HttpProtocolOptions_HeadersWithUnderscoresAction) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescGZIP(), []int{0, 0}
}

type HttpProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The idle timeout for connections. The idle timeout is defined as the
	// period in which there are no active requests. When the
	// idle timeout is reached the connection will be closed. If the connection is an HTTP/2
	// downstream connection a drain sequence will occur prior to closing the connection, see
	// :ref:`drain_timeout
	// <envoy_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.drain_timeout>`.
	// Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
	// If not specified, this defaults to 1 hour. To disable idle timeouts explicitly set this to 0.
	//
	// .. warning::
	//   Disabling this timeout has a highly likelihood of yielding connection leaks due to lost TCP
	//   FIN packets, etc.
	IdleTimeout *duration.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// The maximum number of headers. If unconfigured, the default
	// maximum number of request headers allowed is 100. Requests that exceed this limit will receive
	// a 431 response for HTTP/1.x and cause a stream reset for HTTP/2.
	MaxHeadersCount uint32 `protobuf:"varint,2,opt,name=max_headers_count,json=maxHeadersCount,proto3" json:"max_headers_count,omitempty"`
	// Total duration to keep alive an HTTP request/response stream. If the time limit is reached the stream will be
	// reset independent of any other timeouts. If not specified, this value is not set.
	MaxStreamDuration *duration.Duration `protobuf:"bytes,3,opt,name=max_stream_duration,json=maxStreamDuration,proto3" json:"max_stream_duration,omitempty"`
	// Action to take when a client request with a header name containing underscore characters is received.
	// If this setting is not specified, the value defaults to ALLOW.
	// Note: upstream responses are not affected by this setting.
	HeadersWithUnderscoresAction HttpProtocolOptions_HeadersWithUnderscoresAction `protobuf:"varint,4,opt,name=headers_with_underscores_action,json=headersWithUnderscoresAction,proto3,enum=protocol.options.gloo.solo.io.HttpProtocolOptions_HeadersWithUnderscoresAction" json:"headers_with_underscores_action,omitempty"`
}

func (x *HttpProtocolOptions) Reset() {
	*x = HttpProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolOptions) ProtoMessage() {}

func (x *HttpProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolOptions.ProtoReflect.Descriptor instead.
func (*HttpProtocolOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *HttpProtocolOptions) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *HttpProtocolOptions) GetMaxHeadersCount() uint32 {
	if x != nil {
		return x.MaxHeadersCount
	}
	return 0
}

func (x *HttpProtocolOptions) GetMaxStreamDuration() *duration.Duration {
	if x != nil {
		return x.MaxStreamDuration
	}
	return nil
}

func (x *HttpProtocolOptions) GetHeadersWithUnderscoresAction() HttpProtocolOptions_HeadersWithUnderscoresAction {
	if x != nil {
		return x.HeadersWithUnderscoresAction
	}
	return HttpProtocolOptions_ALLOW
}

type Http1ProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enables trailers for HTTP/1. By default the HTTP/1 codec drops proxied trailers.
	// Note: Trailers must also be enabled at the gateway level in order for this option to take effect.
	EnableTrailers bool `protobuf:"varint,1,opt,name=enable_trailers,json=enableTrailers,proto3" json:"enable_trailers,omitempty"`
	// Types that are assignable to HeaderFormat:
	//	*Http1ProtocolOptions_ProperCaseHeaderKeyFormat
	//	*Http1ProtocolOptions_PreserveCaseHeaderKeyFormat
	HeaderFormat isHttp1ProtocolOptions_HeaderFormat `protobuf_oneof:"header_format"`
	// Allows invalid HTTP messaging. When this option is false, then Envoy will terminate
	// HTTP/1.1 connections upon receiving an invalid HTTP message. However,
	// when this option is true, then Envoy will leave the HTTP/1.1 connection
	// open where possible.
	// If set, this overrides any HCM :ref:`stream_error_on_invalid_http_messaging
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.stream_error_on_invalid_http_message>`.
	OverrideStreamErrorOnInvalidHttpMessage *wrappers.BoolValue `protobuf:"bytes,2,opt,name=override_stream_error_on_invalid_http_message,json=overrideStreamErrorOnInvalidHttpMessage,proto3" json:"override_stream_error_on_invalid_http_message,omitempty"`
}

func (x *Http1ProtocolOptions) Reset() {
	*x = Http1ProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http1ProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http1ProtocolOptions) ProtoMessage() {}

func (x *Http1ProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http1ProtocolOptions.ProtoReflect.Descriptor instead.
func (*Http1ProtocolOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *Http1ProtocolOptions) GetEnableTrailers() bool {
	if x != nil {
		return x.EnableTrailers
	}
	return false
}

func (m *Http1ProtocolOptions) GetHeaderFormat() isHttp1ProtocolOptions_HeaderFormat {
	if m != nil {
		return m.HeaderFormat
	}
	return nil
}

func (x *Http1ProtocolOptions) GetProperCaseHeaderKeyFormat() bool {
	if x, ok := x.GetHeaderFormat().(*Http1ProtocolOptions_ProperCaseHeaderKeyFormat); ok {
		return x.ProperCaseHeaderKeyFormat
	}
	return false
}

func (x *Http1ProtocolOptions) GetPreserveCaseHeaderKeyFormat() bool {
	if x, ok := x.GetHeaderFormat().(*Http1ProtocolOptions_PreserveCaseHeaderKeyFormat); ok {
		return x.PreserveCaseHeaderKeyFormat
	}
	return false
}

func (x *Http1ProtocolOptions) GetOverrideStreamErrorOnInvalidHttpMessage() *wrappers.BoolValue {
	if x != nil {
		return x.OverrideStreamErrorOnInvalidHttpMessage
	}
	return nil
}

type isHttp1ProtocolOptions_HeaderFormat interface {
	isHttp1ProtocolOptions_HeaderFormat()
}

type Http1ProtocolOptions_ProperCaseHeaderKeyFormat struct {
	// Formats the RESPONSE HEADER by proper casing words: the first character and any character following
	// a special character will be capitalized if it's an alpha character. For example,
	// "content-type" becomes "Content-Type", and "foo$b#$are" becomes "Foo$B#$Are".
	// Note that while this results in most headers following conventional casing, certain headers
	// are not covered. For example, the "TE" header will be formatted as "Te".
	ProperCaseHeaderKeyFormat bool `protobuf:"varint,22,opt,name=proper_case_header_key_format,json=properCaseHeaderKeyFormat,proto3,oneof"`
}

type Http1ProtocolOptions_PreserveCaseHeaderKeyFormat struct {
	// Generates configuration for a stateful formatter extension that allows using received headers to
	// affect the output of encoding headers. Specifically: preserving RESPONSE HEADER case during proxying.
	PreserveCaseHeaderKeyFormat bool `protobuf:"varint,31,opt,name=preserve_case_header_key_format,json=preserveCaseHeaderKeyFormat,proto3,oneof"`
}

func (*Http1ProtocolOptions_ProperCaseHeaderKeyFormat) isHttp1ProtocolOptions_HeaderFormat() {}

func (*Http1ProtocolOptions_PreserveCaseHeaderKeyFormat) isHttp1ProtocolOptions_HeaderFormat() {}

type Http2ProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `Maximum concurrent streams <https://httpwg.org/specs/rfc7540.html#rfc.section.5.1.2>`_
	// allowed for peer on one HTTP/2 connection. Valid values range from 1 to 2147483647 (2^31 - 1)
	// and defaults to 2147483647.
	//
	// For upstream connections, this also limits how many streams Envoy will initiate concurrently
	// on a single connection. If the limit is reached, Envoy may queue requests or establish
	// additional connections (as allowed per circuit breaker limits).
	//
	// This acts as an upper bound: Envoy will lower the max concurrent streams allowed on a given
	// connection based on upstream settings. Config dumps will reflect the configured upper bound,
	// not the per-connection negotiated limits.
	MaxConcurrentStreams *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	// `Initial stream-level flow-control window
	// <https://httpwg.org/specs/rfc7540.html#rfc.section.6.9.2>`_ size. Valid values range from 65535
	// (2^16 - 1, HTTP/2 default) to 2147483647 (2^31 - 1, HTTP/2 maximum) and defaults to 268435456
	// (256 * 1024 * 1024).
	//
	// NOTE: 65535 is the initial window size from HTTP/2 spec. We only support increasing the default
	// window size now, so it's also the minimum.
	//
	// This field also acts as a soft limit on the number of bytes Envoy will buffer per-stream in the
	// HTTP/2 codec buffers. Once the buffer reaches this pointer, watermark callbacks will fire to
	// stop the flow of data to the codec buffers.
	InitialStreamWindowSize *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=initial_stream_window_size,json=initialStreamWindowSize,proto3" json:"initial_stream_window_size,omitempty"`
	// Similar to *initial_stream_window_size*, but for connection-level flow-control
	// window. Currently, this has the same minimum/maximum/default as *initial_stream_window_size*.
	InitialConnectionWindowSize *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=initial_connection_window_size,json=initialConnectionWindowSize,proto3" json:"initial_connection_window_size,omitempty"`
	// Allows invalid HTTP messaging and headers. When this option is disabled (default), then
	// the whole HTTP/2 connection is terminated upon receiving invalid HEADERS frame. However,
	// when this option is enabled, only the offending stream is terminated.
	//
	// This overrides any HCM :ref:`stream_error_on_invalid_http_messaging
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.stream_error_on_invalid_http_message>`
	//
	// See `RFC7540, sec. 8.1 <https://tools.ietf.org/html/rfc7540#section-8.1>`_ for details.
	OverrideStreamErrorOnInvalidHttpMessage *wrappers.BoolValue `protobuf:"bytes,14,opt,name=override_stream_error_on_invalid_http_message,json=overrideStreamErrorOnInvalidHttpMessage,proto3" json:"override_stream_error_on_invalid_http_message,omitempty"`
}

func (x *Http2ProtocolOptions) Reset() {
	*x = Http2ProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http2ProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http2ProtocolOptions) ProtoMessage() {}

func (x *Http2ProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http2ProtocolOptions.ProtoReflect.Descriptor instead.
func (*Http2ProtocolOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *Http2ProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxConcurrentStreams
	}
	return nil
}

func (x *Http2ProtocolOptions) GetInitialStreamWindowSize() *wrappers.UInt32Value {
	if x != nil {
		return x.InitialStreamWindowSize
	}
	return nil
}

func (x *Http2ProtocolOptions) GetInitialConnectionWindowSize() *wrappers.UInt32Value {
	if x != nil {
		return x.InitialConnectionWindowSize
	}
	return nil
}

func (x *Http2ProtocolOptions) GetOverrideStreamErrorOnInvalidHttpMessage() *wrappers.BoolValue {
	if x != nil {
		return x.OverrideStreamErrorOnInvalidHttpMessage
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb3, 0x03, 0x0a, 0x13, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x61, 0x78,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x96,
	0x01, 0x0a, 0x1f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1c, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x1c, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x48,
	0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x02, 0x22, 0xd8, 0x02, 0x0a, 0x14, 0x48, 0x74, 0x74, 0x70,
	0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x1d, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x43, 0x61, 0x73, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x46, 0x0a,
	0x1f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x1b, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x43, 0x61, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x7a, 0x0a, 0x2d, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x27, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x6e, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x22, 0xa4, 0x03, 0x0a, 0x14, 0x48, 0x74, 0x74, 0x70, 0x32, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x16, 0x6d,
	0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x43, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12,
	0x59, 0x0a, 0x1a, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x17, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x61, 0x0a, 0x1e, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x7a, 0x0a,
	0x2d, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x27, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x6e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x48, 0x74,
	0x74, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4f, 0x5a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0xc0, 0xf5,
	0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_goTypes = []interface{}{
	(HttpProtocolOptions_HeadersWithUnderscoresAction)(0), // 0: protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction
	(*HttpProtocolOptions)(nil),                           // 1: protocol.options.gloo.solo.io.HttpProtocolOptions
	(*Http1ProtocolOptions)(nil),                          // 2: protocol.options.gloo.solo.io.Http1ProtocolOptions
	(*Http2ProtocolOptions)(nil),                          // 3: protocol.options.gloo.solo.io.Http2ProtocolOptions
	(*duration.Duration)(nil),                             // 4: google.protobuf.Duration
	(*wrappers.BoolValue)(nil),                            // 5: google.protobuf.BoolValue
	(*wrappers.UInt32Value)(nil),                          // 6: google.protobuf.UInt32Value
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_depIdxs = []int32{
	4, // 0: protocol.options.gloo.solo.io.HttpProtocolOptions.idle_timeout:type_name -> google.protobuf.Duration
	4, // 1: protocol.options.gloo.solo.io.HttpProtocolOptions.max_stream_duration:type_name -> google.protobuf.Duration
	0, // 2: protocol.options.gloo.solo.io.HttpProtocolOptions.headers_with_underscores_action:type_name -> protocol.options.gloo.solo.io.HttpProtocolOptions.HeadersWithUnderscoresAction
	5, // 3: protocol.options.gloo.solo.io.Http1ProtocolOptions.override_stream_error_on_invalid_http_message:type_name -> google.protobuf.BoolValue
	6, // 4: protocol.options.gloo.solo.io.Http2ProtocolOptions.max_concurrent_streams:type_name -> google.protobuf.UInt32Value
	6, // 5: protocol.options.gloo.solo.io.Http2ProtocolOptions.initial_stream_window_size:type_name -> google.protobuf.UInt32Value
	6, // 6: protocol.options.gloo.solo.io.Http2ProtocolOptions.initial_connection_window_size:type_name -> google.protobuf.UInt32Value
	5, // 7: protocol.options.gloo.solo.io.Http2ProtocolOptions.override_stream_error_on_invalid_http_message:type_name -> google.protobuf.BoolValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolOptions); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http1ProtocolOptions); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http2ProtocolOptions); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Http1ProtocolOptions_ProperCaseHeaderKeyFormat)(nil),
		(*Http1ProtocolOptions_PreserveCaseHeaderKeyFormat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_protocol_protocol_proto_depIdxs = nil
}

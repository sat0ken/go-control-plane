// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3alpha/health_check.proto

package envoy_config_core_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
	HealthStatus_DEGRADED  HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0}
}

type HealthCheck struct {
	Timeout               *duration.Duration    `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval              *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	InitialJitter         *duration.Duration    `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	IntervalJitter        *duration.Duration    `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	IntervalJitterPercent uint32                `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	UnhealthyThreshold    *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold      *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	AltPort               *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	ReuseConnection       *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker                isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	NoTrafficInterval            *duration.Duration          `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	UnhealthyInterval            *duration.Duration          `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	UnhealthyEdgeInterval        *duration.Duration          `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	HealthyEdgeInterval          *duration.Duration          `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	EventLogPath                 string                      `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	AlwaysLogHealthCheckFailures bool                        `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	TlsOptions                   *HealthCheck_TlsOptions     `protobuf:"bytes,21,opt,name=tls_options,json=tlsOptions,proto3" json:"tls_options,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

func (m *HealthCheck) GetTlsOptions() *HealthCheck_TlsOptions {
	if m != nil {
		return m.TlsOptions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

type HealthCheck_HttpHealthCheck struct {
	Host                             string                  `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                             string                  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                             *HealthCheck_Payload    `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                          *HealthCheck_Payload    `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	HiddenEnvoyDeprecatedServiceName string                  `protobuf:"bytes,5,opt,name=hidden_envoy_deprecated_service_name,json=hiddenEnvoyDeprecatedServiceName,proto3" json:"hidden_envoy_deprecated_service_name,omitempty"` // Deprecated: Do not use.
	RequestHeadersToAdd              []*HeaderValueOption    `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove           []string                `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	HiddenEnvoyDeprecatedUseHttp2    bool                    `protobuf:"varint,7,opt,name=hidden_envoy_deprecated_use_http2,json=hiddenEnvoyDeprecatedUseHttp2,proto3" json:"hidden_envoy_deprecated_use_http2,omitempty"` // Deprecated: Do not use.
	ExpectedStatuses                 []*v3alpha.Int64Range   `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	CodecClientType                  v3alpha.CodecClientType `protobuf:"varint,10,opt,name=codec_client_type,json=codecClientType,proto3,enum=envoy.type.v3alpha.CodecClientType" json:"codec_client_type,omitempty"`
	ServiceNameMatcher               *v3alpha1.StringMatcher `protobuf:"bytes,11,opt,name=service_name_matcher,json=serviceNameMatcher,proto3" json:"service_name_matcher,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                `json:"-"`
	XXX_unrecognized                 []byte                  `json:"-"`
	XXX_sizecache                    int32                   `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_HttpHealthCheck) GetHiddenEnvoyDeprecatedServiceName() string {
	if m != nil {
		return m.HiddenEnvoyDeprecatedServiceName
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_HttpHealthCheck) GetHiddenEnvoyDeprecatedUseHttp2() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedUseHttp2
	}
	return false
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*v3alpha.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetCodecClientType() v3alpha.CodecClientType {
	if m != nil {
		return m.CodecClientType
	}
	return v3alpha.CodecClientType_HTTP1
}

func (m *HealthCheck_HttpHealthCheck) GetServiceNameMatcher() *v3alpha1.StringMatcher {
	if m != nil {
		return m.ServiceNameMatcher
	}
	return nil
}

type HealthCheck_TcpHealthCheck struct {
	Send                 *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HealthCheck_GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type HealthCheck_CustomHealthCheck struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {
}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_CustomHealthCheck) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_HiddenEnvoyDeprecatedConfig)(nil),
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
	}
}

type HealthCheck_TlsOptions struct {
	AlpnProtocols        []string `protobuf:"bytes,1,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_TlsOptions) Reset()         { *m = HealthCheck_TlsOptions{} }
func (m *HealthCheck_TlsOptions) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TlsOptions) ProtoMessage()    {}
func (*HealthCheck_TlsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9ea1e0e451c71fe, []int{0, 6}
}

func (m *HealthCheck_TlsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TlsOptions.Unmarshal(m, b)
}
func (m *HealthCheck_TlsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TlsOptions.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TlsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TlsOptions.Merge(m, src)
}
func (m *HealthCheck_TlsOptions) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TlsOptions.Size(m)
}
func (m *HealthCheck_TlsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TlsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TlsOptions proto.InternalMessageInfo

func (m *HealthCheck_TlsOptions) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3alpha.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.core.v3alpha.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.config.core.v3alpha.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.config.core.v3alpha.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.config.core.v3alpha.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.config.core.v3alpha.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.config.core.v3alpha.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.config.core.v3alpha.HealthCheck.CustomHealthCheck")
	proto.RegisterType((*HealthCheck_TlsOptions)(nil), "envoy.config.core.v3alpha.HealthCheck.TlsOptions")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3alpha/health_check.proto", fileDescriptor_c9ea1e0e451c71fe)
}

var fileDescriptor_c9ea1e0e451c71fe = []byte{
	// 1508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xdd, 0x6e, 0xdb, 0xc6,
	0x12, 0x16, 0x65, 0xc5, 0x92, 0x46, 0xb6, 0x4c, 0xad, 0xe3, 0x98, 0x56, 0x6c, 0x1f, 0xe5, 0x1c,
	0x07, 0xd1, 0x31, 0x72, 0xa4, 0x13, 0x3b, 0x71, 0x1b, 0x03, 0x45, 0x6b, 0xda, 0x4e, 0xe4, 0xc4,
	0x71, 0x8c, 0xb5, 0xdc, 0x20, 0xc8, 0x05, 0xbb, 0x26, 0xd7, 0x12, 0x13, 0x9a, 0xcb, 0x2e, 0x57,
	0x4a, 0x74, 0x57, 0xf4, 0x2a, 0xe8, 0x65, 0x81, 0xf6, 0xa2, 0x8f, 0x90, 0xf7, 0xe8, 0x13, 0x14,
	0xe8, 0x33, 0xf4, 0x11, 0x0a, 0x5f, 0x15, 0x5c, 0x52, 0xd4, 0x9f, 0x7f, 0x8b, 0xde, 0x91, 0x33,
	0xf3, 0x7d, 0xdf, 0xcc, 0xee, 0xce, 0x0e, 0x09, 0xf7, 0xa9, 0xdb, 0x66, 0x9d, 0xaa, 0xc9, 0xdc,
	0x63, 0xbb, 0x51, 0x35, 0x19, 0xa7, 0xd5, 0xf6, 0x2a, 0x71, 0xbc, 0x26, 0xa9, 0x36, 0x29, 0x71,
	0x44, 0xd3, 0x30, 0x9b, 0xd4, 0x7c, 0x57, 0xf1, 0x38, 0x13, 0x0c, 0xcd, 0xc9, 0xe8, 0x4a, 0x18,
	0x5d, 0x09, 0xa2, 0x2b, 0x51, 0x74, 0x71, 0xe9, 0x7c, 0xa2, 0x23, 0xe2, 0xd3, 0x90, 0xa0, 0x78,
	0x2f, 0x8c, 0x12, 0x1d, 0x8f, 0x56, 0x4f, 0x88, 0x30, 0x9b, 0x94, 0xc7, 0x61, 0xbe, 0xe0, 0xb6,
	0xdb, 0x88, 0x02, 0x17, 0xfa, 0x02, 0xe3, 0x84, 0x84, 0xf0, 0x22, 0xf7, 0xe2, 0x19, 0x6e, 0x4e,
	0xdc, 0x46, 0x57, 0x67, 0xae, 0xc1, 0x58, 0xc3, 0xa1, 0x55, 0xf9, 0x76, 0xd4, 0x3a, 0xae, 0x12,
	0xb7, 0xd3, 0x85, 0x0e, 0xbb, 0xac, 0x16, 0x27, 0xc2, 0x66, 0x6e, 0xe4, 0x9f, 0x1f, 0xf6, 0xfb,
	0x82, 0xb7, 0x4c, 0x71, 0x1e, 0xfa, 0x3d, 0x27, 0x9e, 0x47, 0xb9, 0x1f, 0xf9, 0xef, 0xb4, 0x2c,
	0x8f, 0x54, 0x89, 0xeb, 0x32, 0x21, 0x49, 0xfd, 0x6a, 0x9b, 0x72, 0xdf, 0x66, 0x6e, 0xaf, 0xb4,
	0xd9, 0x36, 0x71, 0x6c, 0x8b, 0x08, 0x5a, 0xed, 0x3e, 0x84, 0x8e, 0x7f, 0xff, 0x74, 0x1b, 0x72,
	0x35, 0xb9, 0xe8, 0x9b, 0xc1, 0x9a, 0xa3, 0x2f, 0x21, 0x2d, 0xec, 0x13, 0xca, 0x5a, 0x42, 0x53,
	0x4a, 0x4a, 0x39, 0xb7, 0x32, 0x57, 0x09, 0xd5, 0x2b, 0x5d, 0xf5, 0xca, 0x56, 0x94, 0xbb, 0x0e,
	0xa7, 0x7a, 0xfa, 0x93, 0x92, 0xca, 0x28, 0xcb, 0x09, 0xdc, 0x45, 0xa1, 0x0d, 0xc8, 0xd8, 0xae,
	0xa0, 0xbc, 0x4d, 0x1c, 0x2d, 0x79, 0x1d, 0x86, 0x18, 0x86, 0xbe, 0x82, 0xbc, 0xed, 0xda, 0xc2,
	0x26, 0x8e, 0xf1, 0xd6, 0x16, 0x82, 0x72, 0xed, 0xe6, 0x25, 0x44, 0x78, 0x32, 0x02, 0x3c, 0x93,
	0xf1, 0x48, 0x87, 0xa9, 0x2e, 0x5b, 0x97, 0x62, 0xec, 0x32, 0x8a, 0x7c, 0x17, 0x11, 0x71, 0xac,
	0xc1, 0xec, 0x10, 0x87, 0xe1, 0x51, 0x6e, 0x52, 0x57, 0x68, 0xa8, 0xa4, 0x94, 0x27, 0xf1, 0xcc,
	0x20, 0x60, 0x3f, 0x74, 0xa2, 0x57, 0x30, 0xdd, 0x72, 0xc3, 0x73, 0xdc, 0x31, 0x44, 0x93, 0x53,
	0xbf, 0xc9, 0x1c, 0x4b, 0x4b, 0x49, 0xfd, 0xf9, 0x11, 0xfd, 0xc3, 0x1d, 0x57, 0xac, 0xae, 0x7c,
	0x4d, 0x9c, 0x16, 0xd5, 0x33, 0xa7, 0xfa, 0x8d, 0x1f, 0x94, 0xa4, 0xaa, 0x60, 0x14, 0x53, 0xd4,
	0xbb, 0x0c, 0xe8, 0x00, 0x0a, 0xa3, 0xb4, 0x37, 0xae, 0x45, 0xab, 0x8e, 0x90, 0x7e, 0x06, 0x19,
	0xe2, 0x08, 0xc3, 0x63, 0x5c, 0x68, 0xe3, 0x97, 0x73, 0xe1, 0x34, 0x71, 0xc4, 0x3e, 0xe3, 0x02,
	0x6d, 0x83, 0xca, 0x69, 0xcb, 0xa7, 0x86, 0xc9, 0x5c, 0x97, 0x9a, 0xc1, 0x12, 0x6a, 0x69, 0x49,
	0x50, 0x1c, 0x21, 0xd0, 0x19, 0x73, 0x42, 0xf8, 0x94, 0xc4, 0x6c, 0xc6, 0x10, 0x64, 0x41, 0x21,
	0x68, 0x31, 0xa3, 0xbf, 0xf1, 0xb5, 0x8c, 0xe4, 0x59, 0xab, 0x9c, 0xdb, 0xf9, 0x95, 0xbe, 0x23,
	0x5b, 0xa9, 0x09, 0xe1, 0xf5, 0xbd, 0xd7, 0x12, 0x78, 0xaa, 0x39, 0x68, 0x42, 0x04, 0x54, 0x61,
	0x0e, 0x89, 0x64, 0xa5, 0xc8, 0xa3, 0x2b, 0x8a, 0xd4, 0xcd, 0x21, 0x8d, 0xbc, 0x18, 0xb0, 0x04,
	0x85, 0x34, 0xb8, 0x67, 0x0e, 0x6a, 0xe4, 0xae, 0x55, 0xc8, 0x53, 0xee, 0x99, 0x43, 0x85, 0x34,
	0x06, 0x4d, 0xe8, 0x2d, 0x4c, 0x9b, 0x2d, 0x5f, 0xb0, 0x93, 0x41, 0x9d, 0x49, 0xa9, 0xf3, 0xf9,
	0x15, 0x75, 0x36, 0x25, 0xc3, 0xa0, 0x52, 0xc1, 0x1c, 0x36, 0xa2, 0x03, 0x98, 0x76, 0x99, 0x21,
	0x38, 0x39, 0x3e, 0xb6, 0x4d, 0x23, 0x6e, 0xea, 0x89, 0xcb, 0x9a, 0x3a, 0x38, 0x6e, 0x9f, 0x94,
	0xe4, 0x72, 0x02, 0x17, 0x5c, 0x56, 0x0f, 0xe1, 0x3b, 0xdd, 0xde, 0xc6, 0xd0, 0x3b, 0xda, 0x3d,
	0xce, 0xfc, 0x35, 0x38, 0x63, 0x78, 0xcc, 0xf9, 0x06, 0x66, 0x7b, 0x9c, 0xd4, 0x6a, 0xd0, 0x1e,
	0xf1, 0xd4, 0xd5, 0x89, 0x67, 0x62, 0x8e, 0x6d, 0xab, 0x41, 0x63, 0xf2, 0x57, 0x30, 0x73, 0x36,
	0xb5, 0x7a, 0x75, 0xea, 0xe9, 0xb3, 0x88, 0x97, 0x20, 0x4f, 0xdb, 0xd4, 0x15, 0x86, 0xc3, 0x1a,
	0x86, 0x47, 0x44, 0x53, 0x2b, 0x94, 0x94, 0x72, 0x16, 0x4f, 0x48, 0xeb, 0x2e, 0x6b, 0xec, 0x13,
	0xd1, 0x44, 0x4f, 0xa0, 0x44, 0x9c, 0xf7, 0xa4, 0xe3, 0xcb, 0xb0, 0xfe, 0x4d, 0x37, 0x8e, 0x89,
	0xed, 0xb4, 0x38, 0xf5, 0xb5, 0xe9, 0x92, 0x52, 0xce, 0xe0, 0xf9, 0x30, 0x6e, 0x97, 0x35, 0xfa,
	0x36, 0xf1, 0x49, 0x14, 0x83, 0x30, 0xe4, 0x84, 0xe3, 0x1b, 0xcc, 0x93, 0x03, 0x42, 0x9b, 0x91,
	0xc9, 0x3f, 0xb8, 0xea, 0xe1, 0x77, 0xfc, 0x97, 0x21, 0x10, 0x83, 0x88, 0x9f, 0x8b, 0xdf, 0x29,
	0x90, 0xde, 0x27, 0x1d, 0x87, 0x11, 0x0b, 0x2d, 0x40, 0x4a, 0xd0, 0x0f, 0xe1, 0xd0, 0xc8, 0xea,
	0xe9, 0x53, 0x3d, 0xc5, 0x93, 0x25, 0xa5, 0x96, 0xc0, 0xd2, 0x8c, 0x34, 0x18, 0x3f, 0xb2, 0x5d,
	0xc2, 0x3b, 0x72, 0x26, 0x4c, 0xd4, 0x12, 0x38, 0x7a, 0x5f, 0xbf, 0xff, 0xcb, 0xaf, 0x1f, 0x17,
	0xef, 0xc1, 0xdd, 0x30, 0x13, 0xe2, 0xd9, 0x95, 0xf6, 0x4a, 0x98, 0x49, 0x7f, 0x06, 0x91, 0x8c,
	0x9e, 0x87, 0xb4, 0x17, 0x29, 0x8e, 0xfd, 0xa9, 0x2b, 0xc5, 0xdf, 0xc7, 0x61, 0x6a, 0xa8, 0xff,
	0x11, 0x82, 0x54, 0x93, 0xf9, 0x51, 0x2a, 0x58, 0x3e, 0xa3, 0xdb, 0x90, 0x92, 0x4b, 0x9c, 0x1c,
	0x48, 0x0f, 0x4b, 0x23, 0xd2, 0x21, 0xe5, 0x53, 0xd7, 0x8a, 0x46, 0x44, 0xe5, 0x8a, 0x8b, 0x12,
	0xa5, 0x84, 0x25, 0x16, 0xd5, 0x20, 0xcd, 0xa9, 0x49, 0xed, 0x36, 0x8d, 0x6e, 0xfa, 0xeb, 0xd2,
	0x74, 0xe1, 0x08, 0xc3, 0x52, 0xd3, 0xb6, 0x2c, 0xea, 0x1a, 0x92, 0xc0, 0xb0, 0xa8, 0xc7, 0xa9,
	0x49, 0x04, 0xb5, 0x0c, 0x9f, 0xf2, 0xb6, 0x6d, 0x52, 0xc3, 0x25, 0x27, 0x54, 0xde, 0xfc, 0x59,
	0x3d, 0xa9, 0x29, 0xb8, 0x14, 0xc6, 0x6f, 0x07, 0xe1, 0x5b, 0x71, 0xf4, 0x41, 0x18, 0xbc, 0x47,
	0x4e, 0x28, 0x72, 0xe0, 0x16, 0xa7, 0xdf, 0xb6, 0xa8, 0x2f, 0x82, 0x23, 0x64, 0x51, 0xee, 0x1b,
	0x82, 0x19, 0xc4, 0xb2, 0xb4, 0xf1, 0xd2, 0x58, 0x39, 0xb7, 0x72, 0xff, 0xe2, 0x64, 0x2d, 0xca,
	0xe5, 0xf5, 0x1d, 0x6e, 0xbc, 0x9e, 0x3d, 0xd5, 0xc7, 0x7f, 0x54, 0xc6, 0xd4, 0x3f, 0xd2, 0x78,
	0x3a, 0xa2, 0x0d, 0x83, 0xfc, 0x3a, 0xdb, 0xb0, 0x2c, 0xf4, 0x18, 0xe6, 0xce, 0x50, 0xe3, 0xf4,
	0x84, 0xb5, 0xa9, 0x96, 0x29, 0x8d, 0x95, 0xb3, 0xf8, 0xd6, 0x30, 0x0e, 0x4b, 0x2f, 0xda, 0x85,
	0x3b, 0xe7, 0x15, 0x1f, 0xcc, 0x9a, 0xe0, 0x5e, 0x5f, 0x91, 0x63, 0x26, 0x23, 0x2b, 0x5f, 0x38,
	0xb3, 0xf2, 0x43, 0x9f, 0x06, 0x67, 0x62, 0x05, 0x3d, 0x87, 0x02, 0xfd, 0xe0, 0x51, 0x53, 0xae,
	0x9d, 0x20, 0xa2, 0xe5, 0x53, 0x5f, 0xcb, 0xca, 0x8a, 0x17, 0xa3, 0x8a, 0x83, 0xaf, 0xb9, 0xb8,
	0xd4, 0x1d, 0x57, 0xac, 0x3d, 0xc4, 0xc1, 0x27, 0x1d, 0x56, 0xbb, 0xc0, 0x83, 0x08, 0x87, 0x5e,
	0x43, 0xc1, 0x64, 0x16, 0x35, 0x0d, 0xd3, 0xb1, 0x83, 0xb6, 0x0d, 0x90, 0x1a, 0x94, 0x94, 0x72,
	0x7e, 0xe5, 0x3f, 0x67, 0x91, 0x6d, 0x06, 0xc1, 0x9b, 0x32, 0xb6, 0xde, 0xf1, 0xc2, 0x29, 0xfc,
	0xbd, 0x9c, 0xc2, 0x53, 0xe6, 0xa0, 0x0b, 0xbd, 0x81, 0x9b, 0xfd, 0x5b, 0x6b, 0x44, 0x5f, 0xa9,
	0xd1, 0xf8, 0xf8, 0x6f, 0x3f, 0x7b, 0xe4, 0x8a, 0x55, 0x0e, 0xe4, 0x07, 0xec, 0x8b, 0xd0, 0x8a,
	0x91, 0xdf, 0xdb, 0xf4, 0xc8, 0xb6, 0xfe, 0x30, 0x68, 0xb0, 0x2a, 0xfc, 0xef, 0xe2, 0x06, 0x1b,
	0x6a, 0xa2, 0xe2, 0x6f, 0x0a, 0xe4, 0x07, 0x67, 0x5e, 0xdc, 0x26, 0xca, 0x3f, 0xd3, 0x26, 0x49,
	0xb9, 0x0f, 0x7f, 0xb7, 0x4d, 0xd6, 0x57, 0x83, 0xb2, 0x2a, 0xd1, 0xbf, 0xc4, 0xb9, 0x65, 0x0d,
	0x96, 0x50, 0x7c, 0x03, 0x2a, 0xa6, 0x96, 0xed, 0xf7, 0x97, 0xa5, 0xc2, 0xd8, 0x3b, 0xda, 0x89,
	0x6e, 0x8b, 0xe0, 0x71, 0xfd, 0x51, 0x40, 0xfd, 0x7f, 0xa8, 0x5c, 0x4c, 0x3d, 0x4c, 0x54, 0xfc,
	0xa8, 0xc0, 0xd4, 0xd0, 0x08, 0x47, 0x77, 0x60, 0x62, 0xa0, 0x69, 0x43, 0x95, 0x5c, 0xdf, 0x36,
	0xa1, 0x79, 0xc8, 0x92, 0x96, 0x68, 0x32, 0x6e, 0x8b, 0xf0, 0x76, 0xcc, 0xe2, 0x9e, 0xe1, 0x8a,
	0xbb, 0x37, 0x24, 0x5b, 0xfc, 0x39, 0x09, 0x85, 0x91, 0x29, 0x1f, 0x5c, 0x82, 0xbd, 0x24, 0xfa,
	0x2e, 0xc1, 0xc0, 0x88, 0x8e, 0x60, 0xf1, 0xbc, 0xce, 0x0b, 0xf7, 0x26, 0xfa, 0x9a, 0x9f, 0x1d,
	0x19, 0x78, 0x07, 0xf2, 0x5f, 0x25, 0xe8, 0xc7, 0x5a, 0x02, 0xdf, 0x3e, 0xb3, 0x23, 0x37, 0x25,
	0x03, 0x7a, 0x0c, 0x13, 0xc1, 0x21, 0x8e, 0x19, 0xc3, 0x0b, 0xf7, 0xe6, 0x08, 0xe3, 0x86, 0xdb,
	0xa9, 0x25, 0x70, 0x4e, 0xc6, 0x86, 0xd0, 0xf5, 0xb5, 0x60, 0x1d, 0x1e, 0x40, 0xf5, 0xe2, 0x75,
	0x18, 0xa9, 0x59, 0x9f, 0x84, 0x5c, 0x28, 0x26, 0xfb, 0xb5, 0x68, 0x01, 0xf4, 0x86, 0x19, 0xba,
	0x0b, 0x79, 0xe2, 0x78, 0xae, 0x21, 0x85, 0x4d, 0xe6, 0xf8, 0x9a, 0x22, 0x6f, 0xa7, 0xc9, 0xc0,
	0xba, 0xdf, 0x35, 0xae, 0x57, 0x03, 0xed, 0x65, 0x28, 0x5f, 0x72, 0xd4, 0x62, 0xde, 0xf5, 0xa5,
	0x00, 0xf0, 0x2f, 0x58, 0xb8, 0x10, 0xa0, 0xcf, 0x40, 0xbe, 0x7f, 0x9e, 0x53, 0x2e, 0x47, 0xda,
	0xb3, 0x54, 0x06, 0xd4, 0xdc, 0xf2, 0x37, 0x30, 0x11, 0xc6, 0x86, 0xf7, 0x0f, 0xca, 0x41, 0xfa,
	0x70, 0xef, 0xf9, 0xde, 0xcb, 0x57, 0x7b, 0x6a, 0x22, 0x78, 0xa9, 0x6d, 0x6f, 0xec, 0xd6, 0x6b,
	0xaf, 0x55, 0x05, 0x4d, 0x42, 0xf6, 0x70, 0xaf, 0xfb, 0x9a, 0x44, 0x13, 0x90, 0xd9, 0xc2, 0x1b,
	0x3b, 0x7b, 0x3b, 0x7b, 0x4f, 0xd5, 0xb1, 0x20, 0xb2, 0xbe, 0xf3, 0x62, 0xfb, 0xe5, 0x61, 0x5d,
	0x4d, 0x49, 0xd7, 0xf6, 0x53, 0xbc, 0xb1, 0xb5, 0xbd, 0xa5, 0xde, 0xd0, 0xbf, 0x80, 0x7b, 0x36,
	0x0b, 0xcf, 0xb8, 0xc7, 0xd9, 0x87, 0xce, 0xf9, 0x8d, 0xa8, 0xab, 0x7d, 0x69, 0xcb, 0x65, 0xd9,
	0x57, 0x8e, 0xc6, 0xe5, 0xa2, 0xad, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x67, 0x02, 0x33,
	0xcd, 0x0f, 0x00, 0x00,
}

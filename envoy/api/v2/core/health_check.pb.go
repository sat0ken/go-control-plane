// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/health_check.proto

package envoy_api_v2_core

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0}
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0}
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 0}
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
	Host                   string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                   string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                   *HealthCheck_Payload   `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                *HealthCheck_Payload   `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	ServiceName            string                 `protobuf:"bytes,5,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"` // Deprecated: Do not use.
	RequestHeadersToAdd    []*HeaderValueOption   `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove []string               `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	UseHttp2               bool                   `protobuf:"varint,7,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"` // Deprecated: Do not use.
	ExpectedStatuses       []*_type.Int64Range    `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	CodecClientType        _type.CodecClientType  `protobuf:"varint,10,opt,name=codec_client_type,json=codecClientType,proto3,enum=envoy.type.CodecClientType" json:"codec_client_type,omitempty"`
	ServiceNameMatcher     *matcher.StringMatcher `protobuf:"bytes,11,opt,name=service_name_matcher,json=serviceNameMatcher,proto3" json:"service_name_matcher,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 1}
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
func (m *HealthCheck_HttpHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
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
func (m *HealthCheck_HttpHealthCheck) GetUseHttp2() bool {
	if m != nil {
		return m.UseHttp2
	}
	return false
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*_type.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetCodecClientType() _type.CodecClientType {
	if m != nil {
		return m.CodecClientType
	}
	return _type.CodecClientType_HTTP1
}

func (m *HealthCheck_HttpHealthCheck) GetServiceNameMatcher() *matcher.StringMatcher {
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 2}
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 3}
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 4}
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
	//	*HealthCheck_CustomHealthCheck_Config
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 5}
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

type HealthCheck_CustomHealthCheck_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_Config) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *HealthCheck_CustomHealthCheck) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_Config); ok {
		return x.Config
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
		(*HealthCheck_CustomHealthCheck_Config)(nil),
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
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 6}
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
	proto.RegisterEnum("envoy.api.v2.core.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v2.core.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.api.v2.core.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.CustomHealthCheck")
	proto.RegisterType((*HealthCheck_TlsOptions)(nil), "envoy.api.v2.core.HealthCheck.TlsOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/health_check.proto", fileDescriptor_b6ca44dd529b90bd)
}

var fileDescriptor_b6ca44dd529b90bd = []byte{
	// 1374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0x41, 0x6f, 0xdb, 0xc6,
	0x12, 0xc7, 0x4d, 0x59, 0xb6, 0xa4, 0x91, 0x2d, 0x53, 0xeb, 0xd8, 0xa6, 0x15, 0xbf, 0x17, 0xe5,
	0xc1, 0x79, 0xf0, 0x0b, 0xf0, 0xa4, 0x42, 0x6e, 0x53, 0xa4, 0x97, 0xc6, 0x94, 0x9d, 0xc8, 0x69,
	0xa2, 0xb8, 0x6b, 0x39, 0x41, 0xd0, 0x02, 0xec, 0x9a, 0x5c, 0x4b, 0x6c, 0x28, 0x2e, 0xbb, 0x5c,
	0x2a, 0xd1, 0xb5, 0xc7, 0x1e, 0x1a, 0xa0, 0x87, 0xa2, 0x9f, 0x21, 0xb7, 0x7e, 0xab, 0x9e, 0x7b,
	0x2a, 0x7c, 0x28, 0x0a, 0x2e, 0x49, 0x89, 0x92, 0x5c, 0xc8, 0x46, 0x6f, 0xdc, 0x99, 0xf9, 0xff,
	0x76, 0x97, 0x3b, 0x33, 0xbb, 0xb0, 0x4b, 0xdd, 0x01, 0x1b, 0xd6, 0x89, 0x67, 0xd7, 0x07, 0x8d,
	0xba, 0xc9, 0x38, 0xad, 0xf7, 0x28, 0x71, 0x44, 0xcf, 0x30, 0x7b, 0xd4, 0x7c, 0x53, 0xf3, 0x38,
	0x13, 0x0c, 0x95, 0x65, 0x54, 0x8d, 0x78, 0x76, 0x6d, 0xd0, 0xa8, 0x85, 0x51, 0x95, 0x9d, 0x59,
	0xe1, 0x39, 0xf1, 0x69, 0x24, 0xa8, 0x6c, 0x44, 0x5e, 0x31, 0xf4, 0x68, 0xbd, 0x27, 0x84, 0x17,
	0x9b, 0xef, 0xa4, 0xcc, 0x7d, 0x22, 0xcc, 0x1e, 0xe5, 0x75, 0x5f, 0x70, 0xdb, 0xed, 0xc6, 0x01,
	0x9b, 0xa9, 0x00, 0x4e, 0xdc, 0x6e, 0xc2, 0xdb, 0xee, 0x32, 0xd6, 0x75, 0x68, 0x5d, 0x8e, 0xce,
	0x83, 0x8b, 0x3a, 0x71, 0x87, 0xb1, 0xeb, 0xdf, 0xd3, 0x2e, 0x2b, 0xe0, 0x44, 0xd8, 0xcc, 0x8d,
	0xfd, 0x3b, 0xd3, 0x7e, 0x5f, 0xf0, 0xc0, 0x14, 0x7f, 0xa7, 0x7e, 0xcb, 0x89, 0xe7, 0x51, 0xee,
	0x27, 0xfe, 0xc0, 0xf2, 0x48, 0x9d, 0xb8, 0x2e, 0x13, 0x12, 0xea, 0xd7, 0xfb, 0x76, 0x97, 0x13,
	0x91, 0x2c, 0x6c, 0x6b, 0x40, 0x1c, 0xdb, 0x22, 0x82, 0xd6, 0x93, 0x8f, 0xc8, 0xf1, 0x9f, 0xf7,
	0x9b, 0x50, 0x6c, 0xc9, 0x3f, 0xd9, 0x0c, 0x7f, 0x24, 0xfa, 0x1c, 0x72, 0xc2, 0xee, 0x53, 0x16,
	0x08, 0x4d, 0xa9, 0x2a, 0x7b, 0xc5, 0xc6, 0x76, 0x2d, 0x9a, 0xba, 0x96, 0x4c, 0x5d, 0x3b, 0x8c,
	0x17, 0xae, 0xc3, 0xa5, 0x9e, 0xfb, 0xa0, 0x64, 0xf3, 0xca, 0xfd, 0x05, 0x9c, 0xa8, 0xd0, 0x01,
	0xe4, 0x6d, 0x57, 0x50, 0x3e, 0x20, 0x8e, 0x96, 0xb9, 0x09, 0x61, 0x24, 0x43, 0x8f, 0xa0, 0x64,
	0xbb, 0xb6, 0xb0, 0x89, 0x63, 0x7c, 0x6b, 0x0b, 0x41, 0xb9, 0x76, 0x6b, 0x0e, 0x08, 0xaf, 0xc6,
	0x82, 0xa7, 0x32, 0x1e, 0xe9, 0xb0, 0x96, 0xd0, 0x12, 0xc4, 0xe2, 0x3c, 0x44, 0x29, 0x51, 0xc4,
	0x8c, 0x07, 0xb0, 0x35, 0xc5, 0x30, 0x3c, 0xca, 0x4d, 0xea, 0x0a, 0x0d, 0x55, 0x95, 0xbd, 0x55,
	0xbc, 0x31, 0x29, 0x38, 0x89, 0x9c, 0xe8, 0x15, 0xac, 0x07, 0x6e, 0x94, 0x9c, 0x43, 0x43, 0xf4,
	0x38, 0xf5, 0x7b, 0xcc, 0xb1, 0xb4, 0xac, 0x9c, 0x7f, 0x67, 0x66, 0xfe, 0xb3, 0x63, 0x57, 0xec,
	0x37, 0x5e, 0x12, 0x27, 0xa0, 0x7a, 0xfe, 0x52, 0x5f, 0xfa, 0x41, 0xc9, 0xa8, 0x0a, 0x46, 0x23,
	0x44, 0x27, 0x21, 0xa0, 0x53, 0x28, 0xcf, 0x62, 0x97, 0x6e, 0x84, 0x55, 0x67, 0xa0, 0x9f, 0x42,
	0x9e, 0x38, 0xc2, 0xf0, 0x18, 0x17, 0xda, 0xf2, 0x7c, 0x16, 0xce, 0x11, 0x47, 0x9c, 0x30, 0x2e,
	0xd0, 0x11, 0xa8, 0x9c, 0x06, 0x3e, 0x35, 0x4c, 0xe6, 0xba, 0xd4, 0x0c, 0x7f, 0xa1, 0x96, 0x93,
	0x80, 0xca, 0x0c, 0x40, 0x67, 0xcc, 0x89, 0xe4, 0x6b, 0x52, 0xd3, 0x1c, 0x49, 0xd0, 0xd7, 0x50,
	0x0e, 0x0b, 0xcf, 0x48, 0x57, 0xb3, 0x96, 0x97, 0x9c, 0x5a, 0x6d, 0xa6, 0x9c, 0x6b, 0xa9, 0x54,
	0xad, 0xb5, 0x84, 0xf0, 0x52, 0xe3, 0xd6, 0x02, 0x5e, 0xeb, 0x4d, 0x9a, 0xd0, 0x6b, 0x50, 0x85,
	0x39, 0x05, 0x2f, 0x48, 0xf8, 0xff, 0xe7, 0xc0, 0x3b, 0xe6, 0x14, 0xbb, 0x24, 0x26, 0x2c, 0xe1,
	0xc2, 0xbb, 0xdc, 0x33, 0x27, 0xd9, 0xc5, 0x6b, 0x2d, 0xfc, 0x09, 0xf7, 0xcc, 0xa9, 0x85, 0x77,
	0x27, 0x4d, 0xe8, 0x1c, 0xd6, 0xcd, 0xc0, 0x17, 0xac, 0x3f, 0xc9, 0x5f, 0x95, 0xfc, 0x8f, 0xe6,
	0xf0, 0x9b, 0x52, 0x39, 0x39, 0x43, 0xd9, 0x9c, 0x36, 0xa2, 0x53, 0x58, 0x77, 0x99, 0x21, 0x38,
	0xb9, 0xb8, 0xb0, 0x4d, 0x63, 0x54, 0xb4, 0x2b, 0xf3, 0x8a, 0x36, 0x4c, 0xa7, 0x0f, 0x4a, 0xe6,
	0xfe, 0x02, 0x2e, 0xbb, 0xac, 0x13, 0xc9, 0x8f, 0x93, 0xda, 0xc5, 0x30, 0x4e, 0xdd, 0x31, 0xb3,
	0x74, 0x03, 0xe6, 0x48, 0x3e, 0x62, 0x7e, 0x05, 0x5b, 0x63, 0x26, 0xb5, 0xba, 0x74, 0x0c, 0x5e,
	0xbb, 0x3e, 0x78, 0x63, 0xc4, 0x38, 0xb2, 0xba, 0x74, 0x04, 0x7f, 0x05, 0x1b, 0x57, 0xa3, 0xd5,
	0xeb, 0xa3, 0xd7, 0xaf, 0x02, 0xef, 0x42, 0x89, 0x0e, 0xa8, 0x2b, 0x0c, 0x87, 0x75, 0x0d, 0x8f,
	0x88, 0x9e, 0x56, 0xae, 0x2a, 0x7b, 0x05, 0xbc, 0x22, 0xad, 0xcf, 0x58, 0xf7, 0x84, 0x88, 0x1e,
	0x7a, 0x0c, 0x55, 0xe2, 0xbc, 0x25, 0x43, 0x5f, 0x86, 0xa5, 0x0f, 0xdb, 0xb8, 0x20, 0xb6, 0x13,
	0x70, 0xea, 0x6b, 0xeb, 0x55, 0x65, 0x2f, 0x8f, 0x77, 0xa2, 0xb8, 0x67, 0xac, 0x9b, 0x3a, 0xc4,
	0xc7, 0x71, 0x0c, 0x7a, 0x0a, 0x45, 0xe1, 0xf8, 0x06, 0xf3, 0x64, 0xf7, 0xd7, 0x36, 0xe4, 0xe2,
	0xff, 0x37, 0x2f, 0xc9, 0x1d, 0xff, 0x45, 0x24, 0xc0, 0x20, 0x46, 0xdf, 0x15, 0x0c, 0xb9, 0x13,
	0x32, 0x74, 0x18, 0xb1, 0xd0, 0xbf, 0x20, 0x2b, 0xe8, 0xbb, 0xe8, 0x2e, 0x28, 0xe8, 0xb9, 0x4b,
	0x3d, 0xcb, 0x33, 0x55, 0xa5, 0xb5, 0x80, 0xa5, 0x19, 0x69, 0xb0, 0x7c, 0x6e, 0xbb, 0x84, 0x0f,
	0x65, 0xab, 0x5f, 0x69, 0x2d, 0xe0, 0x78, 0xac, 0x97, 0x20, 0xe7, 0xc5, 0x8c, 0xc5, 0x3f, 0x74,
	0xa5, 0xf2, 0xf3, 0x12, 0xac, 0x4d, 0x15, 0x2c, 0x42, 0x90, 0xed, 0x31, 0x3f, 0x86, 0x63, 0xf9,
	0x8d, 0x6e, 0x43, 0x56, 0xfe, 0xab, 0xcc, 0xc4, 0x84, 0x58, 0x1a, 0xd1, 0x67, 0x90, 0xf5, 0xa9,
	0x6b, 0xc5, 0xbd, 0xfc, 0xbf, 0x73, 0x76, 0x17, 0xef, 0x01, 0x4b, 0x0d, 0x7a, 0x04, 0x39, 0x4e,
	0x4d, 0x6a, 0x0f, 0x68, 0xdc, 0x8a, 0xaf, 0x2b, 0x4f, 0x64, 0xe8, 0x1e, 0xac, 0xf8, 0x94, 0x0f,
	0x6c, 0x93, 0x1a, 0x2e, 0xe9, 0x53, 0xd9, 0x7a, 0x0b, 0x7a, 0x46, 0x53, 0x70, 0x31, 0xb6, 0xb7,
	0x49, 0x9f, 0x22, 0x0b, 0x36, 0x39, 0xfd, 0x2e, 0xa0, 0xbe, 0x08, 0x8f, 0xd3, 0xa2, 0xdc, 0x37,
	0x04, 0x33, 0x88, 0x65, 0x69, 0xcb, 0xd5, 0xc5, 0xbd, 0x62, 0x63, 0xf7, 0xea, 0x79, 0x2d, 0xca,
	0x65, 0x8b, 0x8c, 0x0e, 0x41, 0x2f, 0x5c, 0xea, 0xcb, 0x3f, 0x29, 0x8b, 0xea, 0x6f, 0x39, 0xbc,
	0x1e, 0xe3, 0xa2, 0x20, 0xbf, 0xc3, 0x0e, 0x2c, 0x0b, 0x3d, 0x84, 0xed, 0x2b, 0x66, 0xe1, 0xb4,
	0xcf, 0x06, 0x54, 0xcb, 0x57, 0x17, 0xf7, 0x0a, 0x78, 0x73, 0x5a, 0x87, 0xa5, 0x17, 0xdd, 0x81,
	0x42, 0xd8, 0xb7, 0xc3, 0x5e, 0xd9, 0x90, 0x2d, 0x3b, 0x2f, 0x37, 0x91, 0x0f, 0x7c, 0x1a, 0x9e,
	0x50, 0x03, 0x35, 0xa1, 0x4c, 0xdf, 0x79, 0xd4, 0x14, 0xd4, 0x32, 0x7c, 0x41, 0x44, 0xe0, 0x53,
	0x5f, 0x2b, 0xc8, 0xc5, 0x6f, 0xc6, 0x8b, 0x0f, 0x5f, 0x3e, 0xb5, 0x63, 0x57, 0x3c, 0xf8, 0x18,
	0x87, 0xcf, 0x1f, 0xac, 0x26, 0x82, 0xd3, 0x38, 0x1e, 0x7d, 0x09, 0x65, 0x93, 0x59, 0xd4, 0x34,
	0x4c, 0xc7, 0x0e, 0xab, 0x20, 0x54, 0x68, 0x50, 0x55, 0xf6, 0x4a, 0x8d, 0xdb, 0x69, 0x48, 0x33,
	0x0c, 0x6a, 0xca, 0x98, 0xce, 0xd0, 0x8b, 0x2e, 0xab, 0xef, 0xe5, 0x65, 0xb5, 0x66, 0x4e, 0xba,
	0xd0, 0x29, 0xdc, 0x4a, 0x1f, 0x80, 0x11, 0x3f, 0xcd, 0xe2, 0xae, 0x7b, 0x37, 0x4d, 0x8d, 0x5d,
	0xb5, 0x53, 0xf9, 0x6a, 0x7b, 0x1e, 0x8d, 0x30, 0x4a, 0x9d, 0x53, 0x6c, 0xab, 0xfc, 0xa8, 0x40,
	0x69, 0xb2, 0xd9, 0x8f, 0xd2, 0x4c, 0xf9, 0x67, 0x69, 0x96, 0x91, 0x7f, 0xec, 0xa6, 0x69, 0x56,
	0xd9, 0x05, 0x15, 0x53, 0xcb, 0xf6, 0xd3, 0x2b, 0x52, 0x61, 0xf1, 0x0d, 0x1d, 0xc6, 0x85, 0x12,
	0x7e, 0x56, 0x30, 0xac, 0x4d, 0x5d, 0x23, 0xe8, 0xee, 0x54, 0x7e, 0x46, 0xd1, 0x13, 0xb9, 0xb9,
	0x03, 0x05, 0x12, 0x88, 0x1e, 0xe3, 0xb6, 0x88, 0x4a, 0xb6, 0x80, 0xc7, 0x86, 0xca, 0xaf, 0x0a,
	0x94, 0x67, 0xee, 0x8e, 0xb0, 0x22, 0xc7, 0xb8, 0x54, 0x45, 0x86, 0x46, 0xf4, 0x09, 0x2c, 0x9b,
	0xcc, 0xbd, 0xb0, 0xbb, 0xf1, 0x5b, 0x6f, 0x6b, 0xa6, 0x5d, 0x9e, 0xca, 0x67, 0x6c, 0x98, 0x61,
	0x61, 0x77, 0x88, 0x82, 0xd1, 0x43, 0x58, 0x09, 0x8f, 0xc9, 0x32, 0x62, 0x71, 0x54, 0xd0, 0xb7,
	0x66, 0xc4, 0x07, 0xee, 0xb0, 0xb5, 0x80, 0x8b, 0x32, 0xb6, 0x29, 0x43, 0xf5, 0x55, 0x28, 0x46,
	0x22, 0x99, 0x51, 0x95, 0x7d, 0x80, 0x71, 0x17, 0x43, 0xf7, 0xa0, 0x44, 0x1c, 0xcf, 0x35, 0x24,
	0xc0, 0x64, 0x8e, 0xaf, 0x29, 0xb2, 0x14, 0x56, 0x43, 0xeb, 0x49, 0x62, 0xd4, 0x37, 0xa0, 0x94,
	0xee, 0xb4, 0x94, 0xcb, 0x1e, 0xf5, 0x34, 0x9b, 0x07, 0xb5, 0x78, 0xff, 0x1b, 0x58, 0x89, 0xb6,
	0x1f, 0xa5, 0x32, 0x2a, 0x42, 0xee, 0xac, 0xfd, 0x45, 0xfb, 0xc5, 0xab, 0xb6, 0xba, 0x10, 0x0e,
	0x5a, 0x47, 0x07, 0xcf, 0x3a, 0xad, 0xd7, 0xaa, 0x82, 0x56, 0xa1, 0x70, 0xd6, 0x4e, 0x86, 0x19,
	0xb4, 0x02, 0xf9, 0x43, 0x7c, 0x70, 0xdc, 0x3e, 0x6e, 0x3f, 0x51, 0x17, 0xc3, 0xc8, 0xce, 0xf1,
	0xf3, 0xa3, 0x17, 0x67, 0x1d, 0x35, 0x2b, 0x5d, 0x47, 0x4f, 0xf0, 0xc1, 0xe1, 0xd1, 0xa1, 0xba,
	0xa4, 0xbf, 0xfc, 0xfd, 0x97, 0x3f, 0xdf, 0x2f, 0xdd, 0x46, 0xdb, 0x51, 0x66, 0x44, 0xdb, 0x89,
	0x32, 0x63, 0xb0, 0x4f, 0x1c, 0xaf, 0x47, 0xe0, 0x8e, 0xcd, 0xa2, 0xbc, 0xf1, 0x38, 0x7b, 0x37,
	0x9c, 0x4d, 0x21, 0x5d, 0x4d, 0x1d, 0x91, 0xdc, 0xd8, 0x89, 0x72, 0xbe, 0x2c, 0xb7, 0xbd, 0xff,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x78, 0x8b, 0x5a, 0x2d, 0x0d, 0x00, 0x00,
}

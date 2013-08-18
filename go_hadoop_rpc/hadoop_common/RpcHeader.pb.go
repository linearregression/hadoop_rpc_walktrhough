// Code generated by protoc-gen-go.
// source: RpcHeader.proto
// DO NOT EDIT!

package hadoop_common

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// *
// RpcKind determine the rpcEngine and the serialization of the rpc request
type RpcKindProto int32

const (
	RpcKindProto_RPC_BUILTIN         RpcKindProto = 0
	RpcKindProto_RPC_WRITABLE        RpcKindProto = 1
	RpcKindProto_RPC_PROTOCOL_BUFFER RpcKindProto = 2
)

var RpcKindProto_name = map[int32]string{
	0: "RPC_BUILTIN",
	1: "RPC_WRITABLE",
	2: "RPC_PROTOCOL_BUFFER",
}
var RpcKindProto_value = map[string]int32{
	"RPC_BUILTIN":         0,
	"RPC_WRITABLE":        1,
	"RPC_PROTOCOL_BUFFER": 2,
}

func (x RpcKindProto) Enum() *RpcKindProto {
	p := new(RpcKindProto)
	*p = x
	return p
}
func (x RpcKindProto) String() string {
	return proto.EnumName(RpcKindProto_name, int32(x))
}
func (x RpcKindProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RpcKindProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcKindProto_value, data, "RpcKindProto")
	if err != nil {
		return err
	}
	*x = RpcKindProto(value)
	return nil
}

type RpcRequestHeaderProto_OperationProto int32

const (
	RpcRequestHeaderProto_RPC_FINAL_PACKET        RpcRequestHeaderProto_OperationProto = 0
	RpcRequestHeaderProto_RPC_CONTINUATION_PACKET RpcRequestHeaderProto_OperationProto = 1
	RpcRequestHeaderProto_RPC_CLOSE_CONNECTION    RpcRequestHeaderProto_OperationProto = 2
)

var RpcRequestHeaderProto_OperationProto_name = map[int32]string{
	0: "RPC_FINAL_PACKET",
	1: "RPC_CONTINUATION_PACKET",
	2: "RPC_CLOSE_CONNECTION",
}
var RpcRequestHeaderProto_OperationProto_value = map[string]int32{
	"RPC_FINAL_PACKET":        0,
	"RPC_CONTINUATION_PACKET": 1,
	"RPC_CLOSE_CONNECTION":    2,
}

func (x RpcRequestHeaderProto_OperationProto) Enum() *RpcRequestHeaderProto_OperationProto {
	p := new(RpcRequestHeaderProto_OperationProto)
	*p = x
	return p
}
func (x RpcRequestHeaderProto_OperationProto) String() string {
	return proto.EnumName(RpcRequestHeaderProto_OperationProto_name, int32(x))
}
func (x RpcRequestHeaderProto_OperationProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RpcRequestHeaderProto_OperationProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcRequestHeaderProto_OperationProto_value, data, "RpcRequestHeaderProto_OperationProto")
	if err != nil {
		return err
	}
	*x = RpcRequestHeaderProto_OperationProto(value)
	return nil
}

type RpcResponseHeaderProto_RpcStatusProto int32

const (
	RpcResponseHeaderProto_SUCCESS RpcResponseHeaderProto_RpcStatusProto = 0
	RpcResponseHeaderProto_ERROR   RpcResponseHeaderProto_RpcStatusProto = 1
	RpcResponseHeaderProto_FATAL   RpcResponseHeaderProto_RpcStatusProto = 2
)

var RpcResponseHeaderProto_RpcStatusProto_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
	2: "FATAL",
}
var RpcResponseHeaderProto_RpcStatusProto_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
	"FATAL":   2,
}

func (x RpcResponseHeaderProto_RpcStatusProto) Enum() *RpcResponseHeaderProto_RpcStatusProto {
	p := new(RpcResponseHeaderProto_RpcStatusProto)
	*p = x
	return p
}
func (x RpcResponseHeaderProto_RpcStatusProto) String() string {
	return proto.EnumName(RpcResponseHeaderProto_RpcStatusProto_name, int32(x))
}
func (x RpcResponseHeaderProto_RpcStatusProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RpcResponseHeaderProto_RpcStatusProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcResponseHeaderProto_RpcStatusProto_value, data, "RpcResponseHeaderProto_RpcStatusProto")
	if err != nil {
		return err
	}
	*x = RpcResponseHeaderProto_RpcStatusProto(value)
	return nil
}

type RpcResponseHeaderProto_RpcErrorCodeProto int32

const (
	// Non-fatal Rpc error - connection left open for future rpc calls
	RpcResponseHeaderProto_ERROR_APPLICATION          RpcResponseHeaderProto_RpcErrorCodeProto = 1
	RpcResponseHeaderProto_ERROR_NO_SUCH_METHOD       RpcResponseHeaderProto_RpcErrorCodeProto = 2
	RpcResponseHeaderProto_ERROR_NO_SUCH_PROTOCOL     RpcResponseHeaderProto_RpcErrorCodeProto = 3
	RpcResponseHeaderProto_ERROR_RPC_SERVER           RpcResponseHeaderProto_RpcErrorCodeProto = 4
	RpcResponseHeaderProto_ERROR_SERIALIZING_RESPONSE RpcResponseHeaderProto_RpcErrorCodeProto = 5
	RpcResponseHeaderProto_ERROR_RPC_VERSION_MISMATCH RpcResponseHeaderProto_RpcErrorCodeProto = 6
	// Fatal Server side Rpc error - connection closed
	RpcResponseHeaderProto_FATAL_UNKNOWN                   RpcResponseHeaderProto_RpcErrorCodeProto = 10
	RpcResponseHeaderProto_FATAL_UNSUPPORTED_SERIALIZATION RpcResponseHeaderProto_RpcErrorCodeProto = 11
	RpcResponseHeaderProto_FATAL_INVALID_RPC_HEADER        RpcResponseHeaderProto_RpcErrorCodeProto = 12
	RpcResponseHeaderProto_FATAL_DESERIALIZING_REQUEST     RpcResponseHeaderProto_RpcErrorCodeProto = 13
	RpcResponseHeaderProto_FATAL_VERSION_MISMATCH          RpcResponseHeaderProto_RpcErrorCodeProto = 14
	RpcResponseHeaderProto_FATAL_UNAUTHORIZED              RpcResponseHeaderProto_RpcErrorCodeProto = 15
)

var RpcResponseHeaderProto_RpcErrorCodeProto_name = map[int32]string{
	1:  "ERROR_APPLICATION",
	2:  "ERROR_NO_SUCH_METHOD",
	3:  "ERROR_NO_SUCH_PROTOCOL",
	4:  "ERROR_RPC_SERVER",
	5:  "ERROR_SERIALIZING_RESPONSE",
	6:  "ERROR_RPC_VERSION_MISMATCH",
	10: "FATAL_UNKNOWN",
	11: "FATAL_UNSUPPORTED_SERIALIZATION",
	12: "FATAL_INVALID_RPC_HEADER",
	13: "FATAL_DESERIALIZING_REQUEST",
	14: "FATAL_VERSION_MISMATCH",
	15: "FATAL_UNAUTHORIZED",
}
var RpcResponseHeaderProto_RpcErrorCodeProto_value = map[string]int32{
	"ERROR_APPLICATION":               1,
	"ERROR_NO_SUCH_METHOD":            2,
	"ERROR_NO_SUCH_PROTOCOL":          3,
	"ERROR_RPC_SERVER":                4,
	"ERROR_SERIALIZING_RESPONSE":      5,
	"ERROR_RPC_VERSION_MISMATCH":      6,
	"FATAL_UNKNOWN":                   10,
	"FATAL_UNSUPPORTED_SERIALIZATION": 11,
	"FATAL_INVALID_RPC_HEADER":        12,
	"FATAL_DESERIALIZING_REQUEST":     13,
	"FATAL_VERSION_MISMATCH":          14,
	"FATAL_UNAUTHORIZED":              15,
}

func (x RpcResponseHeaderProto_RpcErrorCodeProto) Enum() *RpcResponseHeaderProto_RpcErrorCodeProto {
	p := new(RpcResponseHeaderProto_RpcErrorCodeProto)
	*p = x
	return p
}
func (x RpcResponseHeaderProto_RpcErrorCodeProto) String() string {
	return proto.EnumName(RpcResponseHeaderProto_RpcErrorCodeProto_name, int32(x))
}
func (x RpcResponseHeaderProto_RpcErrorCodeProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RpcResponseHeaderProto_RpcErrorCodeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcResponseHeaderProto_RpcErrorCodeProto_value, data, "RpcResponseHeaderProto_RpcErrorCodeProto")
	if err != nil {
		return err
	}
	*x = RpcResponseHeaderProto_RpcErrorCodeProto(value)
	return nil
}

type RpcSaslProto_SaslState int32

const (
	RpcSaslProto_SUCCESS   RpcSaslProto_SaslState = 0
	RpcSaslProto_NEGOTIATE RpcSaslProto_SaslState = 1
	RpcSaslProto_INITIATE  RpcSaslProto_SaslState = 2
	RpcSaslProto_CHALLENGE RpcSaslProto_SaslState = 3
	RpcSaslProto_RESPONSE  RpcSaslProto_SaslState = 4
)

var RpcSaslProto_SaslState_name = map[int32]string{
	0: "SUCCESS",
	1: "NEGOTIATE",
	2: "INITIATE",
	3: "CHALLENGE",
	4: "RESPONSE",
}
var RpcSaslProto_SaslState_value = map[string]int32{
	"SUCCESS":   0,
	"NEGOTIATE": 1,
	"INITIATE":  2,
	"CHALLENGE": 3,
	"RESPONSE":  4,
}

func (x RpcSaslProto_SaslState) Enum() *RpcSaslProto_SaslState {
	p := new(RpcSaslProto_SaslState)
	*p = x
	return p
}
func (x RpcSaslProto_SaslState) String() string {
	return proto.EnumName(RpcSaslProto_SaslState_name, int32(x))
}
func (x RpcSaslProto_SaslState) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *RpcSaslProto_SaslState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcSaslProto_SaslState_value, data, "RpcSaslProto_SaslState")
	if err != nil {
		return err
	}
	*x = RpcSaslProto_SaslState(value)
	return nil
}

type RpcRequestHeaderProto struct {
	RpcKind          *RpcKindProto                         `protobuf:"varint,1,opt,name=rpcKind,enum=hadoop.common.RpcKindProto" json:"rpcKind,omitempty"`
	RpcOp            *RpcRequestHeaderProto_OperationProto `protobuf:"varint,2,opt,name=rpcOp,enum=hadoop.common.RpcRequestHeaderProto_OperationProto" json:"rpcOp,omitempty"`
	CallId           *uint32                               `protobuf:"varint,3,req,name=callId" json:"callId,omitempty"`
	ClientId         []byte                                `protobuf:"bytes,4,req,name=clientId" json:"clientId,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *RpcRequestHeaderProto) Reset()         { *m = RpcRequestHeaderProto{} }
func (m *RpcRequestHeaderProto) String() string { return proto.CompactTextString(m) }
func (*RpcRequestHeaderProto) ProtoMessage()    {}

func (m *RpcRequestHeaderProto) GetRpcKind() RpcKindProto {
	if m != nil && m.RpcKind != nil {
		return *m.RpcKind
	}
	return 0
}

func (m *RpcRequestHeaderProto) GetRpcOp() RpcRequestHeaderProto_OperationProto {
	if m != nil && m.RpcOp != nil {
		return *m.RpcOp
	}
	return 0
}

func (m *RpcRequestHeaderProto) GetCallId() uint32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RpcRequestHeaderProto) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

// *
// Rpc Response Header
// +------------------------------------------------------------------+
// | Rpc total response length in bytes (4 bytes int)                 |
// |  (sum of next two parts)                                         |
// +------------------------------------------------------------------+
// | RpcResponseHeaderProto - serialized delimited ie has len         |
// +------------------------------------------------------------------+
// | if request is successful:                                        |
// |   - RpcResponse -  The actual rpc response  bytes follow         |
// |     the response header                                          |
// |     This response is serialized based on RpcKindProto            |
// | if request fails :                                               |
// |   The rpc response header contains the necessary info            |
// +------------------------------------------------------------------+
//
// Note that rpc response header is also used when connection setup fails.
// Ie the response looks like a rpc response with a fake callId.
type RpcResponseHeaderProto struct {
	CallId              *uint32                                   `protobuf:"varint,1,req,name=callId" json:"callId,omitempty"`
	Status              *RpcResponseHeaderProto_RpcStatusProto    `protobuf:"varint,2,req,name=status,enum=hadoop.common.RpcResponseHeaderProto_RpcStatusProto" json:"status,omitempty"`
	ServerIpcVersionNum *uint32                                   `protobuf:"varint,3,opt,name=serverIpcVersionNum" json:"serverIpcVersionNum,omitempty"`
	ExceptionClassName  *string                                   `protobuf:"bytes,4,opt,name=exceptionClassName" json:"exceptionClassName,omitempty"`
	ErrorMsg            *string                                   `protobuf:"bytes,5,opt,name=errorMsg" json:"errorMsg,omitempty"`
	ErrorDetail         *RpcResponseHeaderProto_RpcErrorCodeProto `protobuf:"varint,6,opt,name=errorDetail,enum=hadoop.common.RpcResponseHeaderProto_RpcErrorCodeProto" json:"errorDetail,omitempty"`
	XXX_unrecognized    []byte                                    `json:"-"`
}

func (m *RpcResponseHeaderProto) Reset()         { *m = RpcResponseHeaderProto{} }
func (m *RpcResponseHeaderProto) String() string { return proto.CompactTextString(m) }
func (*RpcResponseHeaderProto) ProtoMessage()    {}

func (m *RpcResponseHeaderProto) GetCallId() uint32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetStatus() RpcResponseHeaderProto_RpcStatusProto {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetServerIpcVersionNum() uint32 {
	if m != nil && m.ServerIpcVersionNum != nil {
		return *m.ServerIpcVersionNum
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetExceptionClassName() string {
	if m != nil && m.ExceptionClassName != nil {
		return *m.ExceptionClassName
	}
	return ""
}

func (m *RpcResponseHeaderProto) GetErrorMsg() string {
	if m != nil && m.ErrorMsg != nil {
		return *m.ErrorMsg
	}
	return ""
}

func (m *RpcResponseHeaderProto) GetErrorDetail() RpcResponseHeaderProto_RpcErrorCodeProto {
	if m != nil && m.ErrorDetail != nil {
		return *m.ErrorDetail
	}
	return 0
}

type RpcSaslProto struct {
	Version          *uint32                  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	State            *RpcSaslProto_SaslState  `protobuf:"varint,2,req,name=state,enum=hadoop.common.RpcSaslProto_SaslState" json:"state,omitempty"`
	Token            []byte                   `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	Auths            []*RpcSaslProto_SaslAuth `protobuf:"bytes,4,rep,name=auths" json:"auths,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *RpcSaslProto) Reset()         { *m = RpcSaslProto{} }
func (m *RpcSaslProto) String() string { return proto.CompactTextString(m) }
func (*RpcSaslProto) ProtoMessage()    {}

func (m *RpcSaslProto) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *RpcSaslProto) GetState() RpcSaslProto_SaslState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *RpcSaslProto) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RpcSaslProto) GetAuths() []*RpcSaslProto_SaslAuth {
	if m != nil {
		return m.Auths
	}
	return nil
}

type RpcSaslProto_SaslAuth struct {
	Method           *string `protobuf:"bytes,1,req,name=method" json:"method,omitempty"`
	Mechanism        *string `protobuf:"bytes,2,req,name=mechanism" json:"mechanism,omitempty"`
	Protocol         *string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	ServerId         *string `protobuf:"bytes,4,opt,name=serverId" json:"serverId,omitempty"`
	Challenge        []byte  `protobuf:"bytes,5,opt,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpcSaslProto_SaslAuth) Reset()         { *m = RpcSaslProto_SaslAuth{} }
func (m *RpcSaslProto_SaslAuth) String() string { return proto.CompactTextString(m) }
func (*RpcSaslProto_SaslAuth) ProtoMessage()    {}

func (m *RpcSaslProto_SaslAuth) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetMechanism() string {
	if m != nil && m.Mechanism != nil {
		return *m.Mechanism
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func init() {
	proto.RegisterEnum("hadoop.common.RpcKindProto", RpcKindProto_name, RpcKindProto_value)
	proto.RegisterEnum("hadoop.common.RpcRequestHeaderProto_OperationProto", RpcRequestHeaderProto_OperationProto_name, RpcRequestHeaderProto_OperationProto_value)
	proto.RegisterEnum("hadoop.common.RpcResponseHeaderProto_RpcStatusProto", RpcResponseHeaderProto_RpcStatusProto_name, RpcResponseHeaderProto_RpcStatusProto_value)
	proto.RegisterEnum("hadoop.common.RpcResponseHeaderProto_RpcErrorCodeProto", RpcResponseHeaderProto_RpcErrorCodeProto_name, RpcResponseHeaderProto_RpcErrorCodeProto_value)
	proto.RegisterEnum("hadoop.common.RpcSaslProto_SaslState", RpcSaslProto_SaslState_name, RpcSaslProto_SaslState_value)
}

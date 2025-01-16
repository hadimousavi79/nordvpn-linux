// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.6
// source: state.proto

package pb

import (
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

type AppStateError int32

const (
	AppStateError_FAILED_TO_GET_UID AppStateError = 0
)

// Enum value maps for AppStateError.
var (
	AppStateError_name = map[int32]string{
		0: "FAILED_TO_GET_UID",
	}
	AppStateError_value = map[string]int32{
		"FAILED_TO_GET_UID": 0,
	}
)

func (x AppStateError) Enum() *AppStateError {
	p := new(AppStateError)
	*p = x
	return p
}

func (x AppStateError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppStateError) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[0].Descriptor()
}

func (AppStateError) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[0]
}

func (x AppStateError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppStateError.Descriptor instead.
func (AppStateError) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

type ConnectionState int32

const (
	ConnectionState_DISCONNECTED ConnectionState = 0
	ConnectionState_CONNECTING   ConnectionState = 1
	ConnectionState_CONNECTED    ConnectionState = 2
)

// Enum value maps for ConnectionState.
var (
	ConnectionState_name = map[int32]string{
		0: "DISCONNECTED",
		1: "CONNECTING",
		2: "CONNECTED",
	}
	ConnectionState_value = map[string]int32{
		"DISCONNECTED": 0,
		"CONNECTING":   1,
		"CONNECTED":    2,
	}
)

func (x ConnectionState) Enum() *ConnectionState {
	p := new(ConnectionState)
	*p = x
	return p
}

func (x ConnectionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionState) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[1].Descriptor()
}

func (ConnectionState) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[1]
}

func (x ConnectionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionState.Descriptor instead.
func (ConnectionState) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{1}
}

type LoginEventType int32

const (
	LoginEventType_LOGIN  LoginEventType = 0
	LoginEventType_LOGOUT LoginEventType = 1
)

// Enum value maps for LoginEventType.
var (
	LoginEventType_name = map[int32]string{
		0: "LOGIN",
		1: "LOGOUT",
	}
	LoginEventType_value = map[string]int32{
		"LOGIN":  0,
		"LOGOUT": 1,
	}
)

func (x LoginEventType) Enum() *LoginEventType {
	p := new(LoginEventType)
	*p = x
	return p
}

func (x LoginEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[2].Descriptor()
}

func (LoginEventType) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[2]
}

func (x LoginEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginEventType.Descriptor instead.
func (LoginEventType) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2}
}

type UpdateEvent int32

const (
	UpdateEvent_SERVERS_LIST_UPDATE UpdateEvent = 0
	UpdateEvent_ACTUAL_IP_UPDATE    UpdateEvent = 1
)

// Enum value maps for UpdateEvent.
var (
	UpdateEvent_name = map[int32]string{
		0: "SERVERS_LIST_UPDATE",
		1: "ACTUAL_IP_UPDATE",
	}
	UpdateEvent_value = map[string]int32{
		"SERVERS_LIST_UPDATE": 0,
		"ACTUAL_IP_UPDATE":    1,
	}
)

func (x UpdateEvent) Enum() *UpdateEvent {
	p := new(UpdateEvent)
	*p = x
	return p
}

func (x UpdateEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[3].Descriptor()
}

func (UpdateEvent) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[3]
}

func (x UpdateEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateEvent.Descriptor instead.
func (UpdateEvent) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{3}
}

type ConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State          ConnectionState `protobuf:"varint,1,opt,name=state,proto3,enum=pb.ConnectionState" json:"state,omitempty"`
	ServerIp       string          `protobuf:"bytes,2,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	ServerCountry  string          `protobuf:"bytes,3,opt,name=server_country,json=serverCountry,proto3" json:"server_country,omitempty"`
	ServerCity     string          `protobuf:"bytes,4,opt,name=server_city,json=serverCity,proto3" json:"server_city,omitempty"`
	ServerHostname string          `protobuf:"bytes,5,opt,name=server_hostname,json=serverHostname,proto3" json:"server_hostname,omitempty"`
	ServerName     string          `protobuf:"bytes,6,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	IsMeshPeer     bool            `protobuf:"varint,7,opt,name=is_mesh_peer,json=isMeshPeer,proto3" json:"is_mesh_peer,omitempty"`
	ByUser         bool            `protobuf:"varint,8,opt,name=by_user,json=byUser,proto3" json:"by_user,omitempty"`
}

func (x *ConnectionStatus) Reset() {
	*x = ConnectionStatus{}
	mi := &file_state_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionStatus) ProtoMessage() {}

func (x *ConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionStatus.ProtoReflect.Descriptor instead.
func (*ConnectionStatus) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionStatus) GetState() ConnectionState {
	if x != nil {
		return x.State
	}
	return ConnectionState_DISCONNECTED
}

func (x *ConnectionStatus) GetServerIp() string {
	if x != nil {
		return x.ServerIp
	}
	return ""
}

func (x *ConnectionStatus) GetServerCountry() string {
	if x != nil {
		return x.ServerCountry
	}
	return ""
}

func (x *ConnectionStatus) GetServerCity() string {
	if x != nil {
		return x.ServerCity
	}
	return ""
}

func (x *ConnectionStatus) GetServerHostname() string {
	if x != nil {
		return x.ServerHostname
	}
	return ""
}

func (x *ConnectionStatus) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *ConnectionStatus) GetIsMeshPeer() bool {
	if x != nil {
		return x.IsMeshPeer
	}
	return false
}

func (x *ConnectionStatus) GetByUser() bool {
	if x != nil {
		return x.ByUser
	}
	return false
}

type LoginEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type LoginEventType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.LoginEventType" json:"type,omitempty"`
}

func (x *LoginEvent) Reset() {
	*x = LoginEvent{}
	mi := &file_state_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginEvent) ProtoMessage() {}

func (x *LoginEvent) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginEvent.ProtoReflect.Descriptor instead.
func (*LoginEvent) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{1}
}

func (x *LoginEvent) GetType() LoginEventType {
	if x != nil {
		return x.Type
	}
	return LoginEventType_LOGIN
}

type AccountModification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpiresAt *string `protobuf:"bytes,1,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"`
}

func (x *AccountModification) Reset() {
	*x = AccountModification{}
	mi := &file_state_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountModification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountModification) ProtoMessage() {}

func (x *AccountModification) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountModification.ProtoReflect.Descriptor instead.
func (*AccountModification) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2}
}

func (x *AccountModification) GetExpiresAt() string {
	if x != nil && x.ExpiresAt != nil {
		return *x.ExpiresAt
	}
	return ""
}

type AppState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to State:
	//
	//	*AppState_Error
	//	*AppState_ConnectionStatus
	//	*AppState_LoginEvent
	//	*AppState_SettingsChange
	//	*AppState_UpdateEvent
	//	*AppState_AccountModification
	State isAppState_State `protobuf_oneof:"state"`
}

func (x *AppState) Reset() {
	*x = AppState{}
	mi := &file_state_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppState) ProtoMessage() {}

func (x *AppState) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppState.ProtoReflect.Descriptor instead.
func (*AppState) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{3}
}

func (m *AppState) GetState() isAppState_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (x *AppState) GetError() AppStateError {
	if x, ok := x.GetState().(*AppState_Error); ok {
		return x.Error
	}
	return AppStateError_FAILED_TO_GET_UID
}

func (x *AppState) GetConnectionStatus() *ConnectionStatus {
	if x, ok := x.GetState().(*AppState_ConnectionStatus); ok {
		return x.ConnectionStatus
	}
	return nil
}

func (x *AppState) GetLoginEvent() *LoginEvent {
	if x, ok := x.GetState().(*AppState_LoginEvent); ok {
		return x.LoginEvent
	}
	return nil
}

func (x *AppState) GetSettingsChange() *Settings {
	if x, ok := x.GetState().(*AppState_SettingsChange); ok {
		return x.SettingsChange
	}
	return nil
}

func (x *AppState) GetUpdateEvent() UpdateEvent {
	if x, ok := x.GetState().(*AppState_UpdateEvent); ok {
		return x.UpdateEvent
	}
	return UpdateEvent_SERVERS_LIST_UPDATE
}

func (x *AppState) GetAccountModification() *AccountModification {
	if x, ok := x.GetState().(*AppState_AccountModification); ok {
		return x.AccountModification
	}
	return nil
}

type isAppState_State interface {
	isAppState_State()
}

type AppState_Error struct {
	Error AppStateError `protobuf:"varint,1,opt,name=error,proto3,enum=pb.AppStateError,oneof"`
}

type AppState_ConnectionStatus struct {
	ConnectionStatus *ConnectionStatus `protobuf:"bytes,2,opt,name=connection_status,json=connectionStatus,proto3,oneof"`
}

type AppState_LoginEvent struct {
	LoginEvent *LoginEvent `protobuf:"bytes,3,opt,name=login_event,json=loginEvent,proto3,oneof"`
}

type AppState_SettingsChange struct {
	SettingsChange *Settings `protobuf:"bytes,4,opt,name=settings_change,json=settingsChange,proto3,oneof"`
}

type AppState_UpdateEvent struct {
	UpdateEvent UpdateEvent `protobuf:"varint,5,opt,name=update_event,json=updateEvent,proto3,enum=pb.UpdateEvent,oneof"`
}

type AppState_AccountModification struct {
	AccountModification *AccountModification `protobuf:"bytes,6,opt,name=account_modification,json=accountModification,proto3,oneof"`
}

func (*AppState_Error) isAppState_State() {}

func (*AppState_ConnectionStatus) isAppState_State() {}

func (*AppState_LoginEvent) isAppState_State() {}

func (*AppState_SettingsChange) isAppState_State() {}

func (*AppState_UpdateEvent) isAppState_State() {}

func (*AppState_AccountModification) isAppState_State() {}

var File_state_proto protoreflect.FileDescriptor

var file_state_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x70, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x70, 0x65, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x79, 0x55, 0x73, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x48, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x22, 0xf3, 0x02, 0x0a, 0x08,
	0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x14, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2a, 0x26, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f,
	0x47, 0x45, 0x54, 0x5f, 0x55, 0x49, 0x44, 0x10, 0x00, 0x2a, 0x42, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x27, 0x0a,
	0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f,
	0x47, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x2a, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x53,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x43, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f,
	0x6e, 0x6f, 0x72, 0x64, 0x76, 0x70, 0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_proto_rawDescOnce sync.Once
	file_state_proto_rawDescData = file_state_proto_rawDesc
)

func file_state_proto_rawDescGZIP() []byte {
	file_state_proto_rawDescOnce.Do(func() {
		file_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_proto_rawDescData)
	})
	return file_state_proto_rawDescData
}

var file_state_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_state_proto_goTypes = []any{
	(AppStateError)(0),          // 0: pb.AppStateError
	(ConnectionState)(0),        // 1: pb.ConnectionState
	(LoginEventType)(0),         // 2: pb.LoginEventType
	(UpdateEvent)(0),            // 3: pb.UpdateEvent
	(*ConnectionStatus)(nil),    // 4: pb.ConnectionStatus
	(*LoginEvent)(nil),          // 5: pb.LoginEvent
	(*AccountModification)(nil), // 6: pb.AccountModification
	(*AppState)(nil),            // 7: pb.AppState
	(*Settings)(nil),            // 8: pb.Settings
}
var file_state_proto_depIdxs = []int32{
	1, // 0: pb.ConnectionStatus.state:type_name -> pb.ConnectionState
	2, // 1: pb.LoginEvent.type:type_name -> pb.LoginEventType
	0, // 2: pb.AppState.error:type_name -> pb.AppStateError
	4, // 3: pb.AppState.connection_status:type_name -> pb.ConnectionStatus
	5, // 4: pb.AppState.login_event:type_name -> pb.LoginEvent
	8, // 5: pb.AppState.settings_change:type_name -> pb.Settings
	3, // 6: pb.AppState.update_event:type_name -> pb.UpdateEvent
	6, // 7: pb.AppState.account_modification:type_name -> pb.AccountModification
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_state_proto_init() }
func file_state_proto_init() {
	if File_state_proto != nil {
		return
	}
	file_settings_proto_init()
	file_state_proto_msgTypes[2].OneofWrappers = []any{}
	file_state_proto_msgTypes[3].OneofWrappers = []any{
		(*AppState_Error)(nil),
		(*AppState_ConnectionStatus)(nil),
		(*AppState_LoginEvent)(nil),
		(*AppState_SettingsChange)(nil),
		(*AppState_UpdateEvent)(nil),
		(*AppState_AccountModification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_state_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_proto_goTypes,
		DependencyIndexes: file_state_proto_depIdxs,
		EnumInfos:         file_state_proto_enumTypes,
		MessageInfos:      file_state_proto_msgTypes,
	}.Build()
	File_state_proto = out.File
	file_state_proto_rawDesc = nil
	file_state_proto_goTypes = nil
	file_state_proto_depIdxs = nil
}

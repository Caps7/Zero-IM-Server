// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: conversation.proto

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

type FieldType int32

const (
	FieldType_None               FieldType = 0
	FieldType_FieldRecvMsgOpt    FieldType = 1
	FieldType_FieldIsPinned      FieldType = 2
	FieldType_FieldAttachedInfo  FieldType = 3
	FieldType_FieldIsPrivateChat FieldType = 4
	FieldType_FieldGroupAtType   FieldType = 5
	FieldType_FieldIsNotInGroup  FieldType = 6
	FieldType_FieldEx            FieldType = 7
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0: "None",
		1: "FieldRecvMsgOpt",
		2: "FieldIsPinned",
		3: "FieldAttachedInfo",
		4: "FieldIsPrivateChat",
		5: "FieldGroupAtType",
		6: "FieldIsNotInGroup",
		7: "FieldEx",
	}
	FieldType_value = map[string]int32{
		"None":               0,
		"FieldRecvMsgOpt":    1,
		"FieldIsPinned":      2,
		"FieldAttachedInfo":  3,
		"FieldIsPrivateChat": 4,
		"FieldGroupAtType":   5,
		"FieldIsNotInGroup":  6,
		"FieldEx":            7,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_conversation_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_conversation_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{0}
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResp) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *CommonResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type Conversation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserID      string `protobuf:"bytes,1,opt,name=ownerUserID,proto3" json:"ownerUserID,omitempty"`
	ConversationID   string `protobuf:"bytes,2,opt,name=conversationID,proto3" json:"conversationID,omitempty"`
	RecvMsgOpt       int32  `protobuf:"varint,3,opt,name=recvMsgOpt,proto3" json:"recvMsgOpt,omitempty"`
	ConversationType int32  `protobuf:"varint,4,opt,name=conversationType,proto3" json:"conversationType,omitempty"`
	UserID           string `protobuf:"bytes,5,opt,name=userID,proto3" json:"userID,omitempty"`
	GroupID          string `protobuf:"bytes,6,opt,name=groupID,proto3" json:"groupID,omitempty"`
	UnreadCount      int32  `protobuf:"varint,7,opt,name=unreadCount,proto3" json:"unreadCount,omitempty"`
	DraftTextTime    int64  `protobuf:"varint,8,opt,name=draftTextTime,proto3" json:"draftTextTime,omitempty"`
	IsPinned         bool   `protobuf:"varint,9,opt,name=isPinned,proto3" json:"isPinned,omitempty"`
	AttachedInfo     string `protobuf:"bytes,10,opt,name=attachedInfo,proto3" json:"attachedInfo,omitempty"`
	IsPrivateChat    bool   `protobuf:"varint,11,opt,name=isPrivateChat,proto3" json:"isPrivateChat,omitempty"`
	GroupAtType      int32  `protobuf:"varint,12,opt,name=groupAtType,proto3" json:"groupAtType,omitempty"`
	IsNotInGroup     bool   `protobuf:"varint,13,opt,name=isNotInGroup,proto3" json:"isNotInGroup,omitempty"`
	Ex               string `protobuf:"bytes,14,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *Conversation) Reset() {
	*x = Conversation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conversation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conversation) ProtoMessage() {}

func (x *Conversation) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conversation.ProtoReflect.Descriptor instead.
func (*Conversation) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{1}
}

func (x *Conversation) GetOwnerUserID() string {
	if x != nil {
		return x.OwnerUserID
	}
	return ""
}

func (x *Conversation) GetConversationID() string {
	if x != nil {
		return x.ConversationID
	}
	return ""
}

func (x *Conversation) GetRecvMsgOpt() int32 {
	if x != nil {
		return x.RecvMsgOpt
	}
	return 0
}

func (x *Conversation) GetConversationType() int32 {
	if x != nil {
		return x.ConversationType
	}
	return 0
}

func (x *Conversation) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Conversation) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *Conversation) GetUnreadCount() int32 {
	if x != nil {
		return x.UnreadCount
	}
	return 0
}

func (x *Conversation) GetDraftTextTime() int64 {
	if x != nil {
		return x.DraftTextTime
	}
	return 0
}

func (x *Conversation) GetIsPinned() bool {
	if x != nil {
		return x.IsPinned
	}
	return false
}

func (x *Conversation) GetAttachedInfo() string {
	if x != nil {
		return x.AttachedInfo
	}
	return ""
}

func (x *Conversation) GetIsPrivateChat() bool {
	if x != nil {
		return x.IsPrivateChat
	}
	return false
}

func (x *Conversation) GetGroupAtType() int32 {
	if x != nil {
		return x.GroupAtType
	}
	return 0
}

func (x *Conversation) GetIsNotInGroup() bool {
	if x != nil {
		return x.IsNotInGroup
	}
	return false
}

func (x *Conversation) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

type ModifyConversationFieldReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conversation *Conversation `protobuf:"bytes,1,opt,name=conversation,proto3" json:"conversation,omitempty"`
	FieldType    FieldType     `protobuf:"varint,2,opt,name=fieldType,proto3,enum=pb.FieldType" json:"fieldType,omitempty"`
	UserIDList   []string      `protobuf:"bytes,3,rep,name=userIDList,proto3" json:"userIDList,omitempty"`
	OperationID  string        `protobuf:"bytes,4,opt,name=operationID,proto3" json:"operationID,omitempty"`
}

func (x *ModifyConversationFieldReq) Reset() {
	*x = ModifyConversationFieldReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyConversationFieldReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyConversationFieldReq) ProtoMessage() {}

func (x *ModifyConversationFieldReq) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyConversationFieldReq.ProtoReflect.Descriptor instead.
func (*ModifyConversationFieldReq) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{2}
}

func (x *ModifyConversationFieldReq) GetConversation() *Conversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

func (x *ModifyConversationFieldReq) GetFieldType() FieldType {
	if x != nil {
		return x.FieldType
	}
	return FieldType_None
}

func (x *ModifyConversationFieldReq) GetUserIDList() []string {
	if x != nil {
		return x.UserIDList
	}
	return nil
}

func (x *ModifyConversationFieldReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

type ModifyConversationFieldResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp,omitempty"`
}

func (x *ModifyConversationFieldResp) Reset() {
	*x = ModifyConversationFieldResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyConversationFieldResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyConversationFieldResp) ProtoMessage() {}

func (x *ModifyConversationFieldResp) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyConversationFieldResp.ProtoReflect.Descriptor instead.
func (*ModifyConversationFieldResp) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{3}
}

func (x *ModifyConversationFieldResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

var File_conversation_proto protoreflect.FileDescriptor

var file_conversation_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x3e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0xda, 0x03, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x4d, 0x73, 0x67, 0x4f, 0x70,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x4d, 0x73, 0x67,
	0x4f, 0x70, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x72, 0x61, 0x66, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x72, 0x61, 0x66,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50,
	0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50,
	0x69, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x65, 0x78, 0x22, 0xc1, 0x01, 0x0a, 0x1a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x1b, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x2a, 0xa6, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x63, 0x76, 0x4d, 0x73, 0x67,
	0x4f, 0x70, 0x74, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x73,
	0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x03, 0x12,
	0x16, 0x0a, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x05, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x73, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x10,
	0x07, 0x32, 0x71, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x17, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conversation_proto_rawDescOnce sync.Once
	file_conversation_proto_rawDescData = file_conversation_proto_rawDesc
)

func file_conversation_proto_rawDescGZIP() []byte {
	file_conversation_proto_rawDescOnce.Do(func() {
		file_conversation_proto_rawDescData = protoimpl.X.CompressGZIP(file_conversation_proto_rawDescData)
	})
	return file_conversation_proto_rawDescData
}

var file_conversation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_conversation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conversation_proto_goTypes = []interface{}{
	(FieldType)(0),                      // 0: pb.FieldType
	(*CommonResp)(nil),                  // 1: pb.CommonResp
	(*Conversation)(nil),                // 2: pb.Conversation
	(*ModifyConversationFieldReq)(nil),  // 3: pb.ModifyConversationFieldReq
	(*ModifyConversationFieldResp)(nil), // 4: pb.ModifyConversationFieldResp
}
var file_conversation_proto_depIdxs = []int32{
	2, // 0: pb.ModifyConversationFieldReq.conversation:type_name -> pb.Conversation
	0, // 1: pb.ModifyConversationFieldReq.fieldType:type_name -> pb.FieldType
	1, // 2: pb.ModifyConversationFieldResp.commonResp:type_name -> pb.CommonResp
	3, // 3: pb.conversationService.ModifyConversationField:input_type -> pb.ModifyConversationFieldReq
	4, // 4: pb.conversationService.ModifyConversationField:output_type -> pb.ModifyConversationFieldResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conversation_proto_init() }
func file_conversation_proto_init() {
	if File_conversation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conversation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_conversation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conversation); i {
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
		file_conversation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyConversationFieldReq); i {
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
		file_conversation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyConversationFieldResp); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conversation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conversation_proto_goTypes,
		DependencyIndexes: file_conversation_proto_depIdxs,
		EnumInfos:         file_conversation_proto_enumTypes,
		MessageInfos:      file_conversation_proto_msgTypes,
	}.Build()
	File_conversation_proto = out.File
	file_conversation_proto_rawDesc = nil
	file_conversation_proto_goTypes = nil
	file_conversation_proto_depIdxs = nil
}

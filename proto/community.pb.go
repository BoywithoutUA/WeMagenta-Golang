// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: community.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommunityCreation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName        string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	UserAvatar      string `protobuf:"bytes,2,opt,name=userAvatar,proto3" json:"userAvatar,omitempty"`
	CompositionName string `protobuf:"bytes,3,opt,name=compositionName,proto3" json:"compositionName,omitempty"`
	ForWhat         string `protobuf:"bytes,4,opt,name=forWhat,proto3" json:"forWhat,omitempty"`
	Mp3             string `protobuf:"bytes,5,opt,name=mp3,proto3" json:"mp3,omitempty"`
	Likes           int64  `protobuf:"varint,6,opt,name=likes,proto3" json:"likes,omitempty"`
	CreatedTime     uint64 `protobuf:"varint,7,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	Detail          string `protobuf:"bytes,8,opt,name=detail,proto3" json:"detail,omitempty"`
	Id              int64  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	Report          int64  `protobuf:"varint,10,opt,name=report,proto3" json:"report,omitempty"`
	ChineseNote     string `protobuf:"bytes,11,opt,name=chineseNote,proto3" json:"chineseNote,omitempty"`
	ChineseEmotion  string `protobuf:"bytes,12,opt,name=chineseEmotion,proto3" json:"chineseEmotion,omitempty"`
}

func (x *CommunityCreation) Reset() {
	*x = CommunityCreation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityCreation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityCreation) ProtoMessage() {}

func (x *CommunityCreation) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunityCreation.ProtoReflect.Descriptor instead.
func (*CommunityCreation) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{0}
}

func (x *CommunityCreation) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CommunityCreation) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *CommunityCreation) GetCompositionName() string {
	if x != nil {
		return x.CompositionName
	}
	return ""
}

func (x *CommunityCreation) GetForWhat() string {
	if x != nil {
		return x.ForWhat
	}
	return ""
}

func (x *CommunityCreation) GetMp3() string {
	if x != nil {
		return x.Mp3
	}
	return ""
}

func (x *CommunityCreation) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *CommunityCreation) GetCreatedTime() uint64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *CommunityCreation) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *CommunityCreation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommunityCreation) GetReport() int64 {
	if x != nil {
		return x.Report
	}
	return 0
}

func (x *CommunityCreation) GetChineseNote() string {
	if x != nil {
		return x.ChineseNote
	}
	return ""
}

func (x *CommunityCreation) GetChineseEmotion() string {
	if x != nil {
		return x.ChineseEmotion
	}
	return ""
}

type TopCreation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creation      []*CommunityCreation `protobuf:"bytes,1,rep,name=creation,proto3" json:"creation,omitempty"`
	CreationGong  []*CommunityCreation `protobuf:"bytes,2,rep,name=creationGong,proto3" json:"creationGong,omitempty"`
	CreationShang []*CommunityCreation `protobuf:"bytes,3,rep,name=creationShang,proto3" json:"creationShang,omitempty"`
	CreationJue   []*CommunityCreation `protobuf:"bytes,4,rep,name=creationJue,proto3" json:"creationJue,omitempty"`
	CreationZhi   []*CommunityCreation `protobuf:"bytes,5,rep,name=creationZhi,proto3" json:"creationZhi,omitempty"`
	CreationYu    []*CommunityCreation `protobuf:"bytes,6,rep,name=creationYu,proto3" json:"creationYu,omitempty"`
}

func (x *TopCreation) Reset() {
	*x = TopCreation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopCreation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopCreation) ProtoMessage() {}

func (x *TopCreation) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopCreation.ProtoReflect.Descriptor instead.
func (*TopCreation) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{1}
}

func (x *TopCreation) GetCreation() []*CommunityCreation {
	if x != nil {
		return x.Creation
	}
	return nil
}

func (x *TopCreation) GetCreationGong() []*CommunityCreation {
	if x != nil {
		return x.CreationGong
	}
	return nil
}

func (x *TopCreation) GetCreationShang() []*CommunityCreation {
	if x != nil {
		return x.CreationShang
	}
	return nil
}

func (x *TopCreation) GetCreationJue() []*CommunityCreation {
	if x != nil {
		return x.CreationJue
	}
	return nil
}

func (x *TopCreation) GetCreationZhi() []*CommunityCreation {
	if x != nil {
		return x.CreationZhi
	}
	return nil
}

func (x *TopCreation) GetCreationYu() []*CommunityCreation {
	if x != nil {
		return x.CreationYu
	}
	return nil
}

type Bulletin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Bulletin) Reset() {
	*x = Bulletin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bulletin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bulletin) ProtoMessage() {}

func (x *Bulletin) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bulletin.ProtoReflect.Descriptor instead.
func (*Bulletin) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{2}
}

func (x *Bulletin) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PublicCreation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname        string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	CompositionName string `protobuf:"bytes,2,opt,name=compositionName,proto3" json:"compositionName,omitempty"`
	Detail          string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	ForWhat         string `protobuf:"bytes,4,opt,name=forWhat,proto3" json:"forWhat,omitempty"`
	Likes           int64  `protobuf:"varint,5,opt,name=likes,proto3" json:"likes,omitempty"`
	Report          int64  `protobuf:"varint,6,opt,name=report,proto3" json:"report,omitempty"`
	Mp3             string `protobuf:"bytes,7,opt,name=mp3,proto3" json:"mp3,omitempty"`
	Id              int64  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	Avatar          string `protobuf:"bytes,10,opt,name=avatar,proto3" json:"avatar,omitempty"`
	ChineseNote     string `protobuf:"bytes,11,opt,name=chineseNote,proto3" json:"chineseNote,omitempty"`
	ChineseEmotion  string `protobuf:"bytes,12,opt,name=chineseEmotion,proto3" json:"chineseEmotion,omitempty"`
}

func (x *PublicCreation) Reset() {
	*x = PublicCreation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicCreation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCreation) ProtoMessage() {}

func (x *PublicCreation) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCreation.ProtoReflect.Descriptor instead.
func (*PublicCreation) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{4}
}

func (x *PublicCreation) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PublicCreation) GetCompositionName() string {
	if x != nil {
		return x.CompositionName
	}
	return ""
}

func (x *PublicCreation) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *PublicCreation) GetForWhat() string {
	if x != nil {
		return x.ForWhat
	}
	return ""
}

func (x *PublicCreation) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *PublicCreation) GetReport() int64 {
	if x != nil {
		return x.Report
	}
	return 0
}

func (x *PublicCreation) GetMp3() string {
	if x != nil {
		return x.Mp3
	}
	return ""
}

func (x *PublicCreation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PublicCreation) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *PublicCreation) GetChineseNote() string {
	if x != nil {
		return x.ChineseNote
	}
	return ""
}

func (x *PublicCreation) GetChineseEmotion() string {
	if x != nil {
		return x.ChineseEmotion
	}
	return ""
}

type SearchUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname       string            `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender         string            `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Pic            string            `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`
	Birthday       string            `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Motto          string            `protobuf:"bytes,6,opt,name=motto,proto3" json:"motto,omitempty"`
	PublicCreation []*PublicCreation `protobuf:"bytes,5,rep,name=publicCreation,proto3" json:"publicCreation,omitempty"`
}

func (x *SearchUser) Reset() {
	*x = SearchUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUser) ProtoMessage() {}

func (x *SearchUser) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUser.ProtoReflect.Descriptor instead.
func (*SearchUser) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{5}
}

func (x *SearchUser) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SearchUser) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *SearchUser) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *SearchUser) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *SearchUser) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *SearchUser) GetPublicCreation() []*PublicCreation {
	if x != nil {
		return x.PublicCreation
	}
	return nil
}

type SearchUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*SearchUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SearchUserResponse) Reset() {
	*x = SearchUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResponse) ProtoMessage() {}

func (x *SearchUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResponse.ProtoReflect.Descriptor instead.
func (*SearchUserResponse) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUserResponse) GetUsers() []*SearchUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type SearchCreationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creations []*PublicCreation `protobuf:"bytes,1,rep,name=creations,proto3" json:"creations,omitempty"`
}

func (x *SearchCreationResponse) Reset() {
	*x = SearchCreationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCreationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCreationResponse) ProtoMessage() {}

func (x *SearchCreationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCreationResponse.ProtoReflect.Descriptor instead.
func (*SearchCreationResponse) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{7}
}

func (x *SearchCreationResponse) GetCreations() []*PublicCreation {
	if x != nil {
		return x.Creations
	}
	return nil
}

type AttitudeCreation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Opt int32 `protobuf:"varint,2,opt,name=opt,proto3" json:"opt,omitempty"`
}

func (x *AttitudeCreation) Reset() {
	*x = AttitudeCreation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttitudeCreation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttitudeCreation) ProtoMessage() {}

func (x *AttitudeCreation) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttitudeCreation.ProtoReflect.Descriptor instead.
func (*AttitudeCreation) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{8}
}

func (x *AttitudeCreation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttitudeCreation) GetOpt() int32 {
	if x != nil {
		return x.Opt
	}
	return 0
}

type AttitudeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AttitudeResponse) Reset() {
	*x = AttitudeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttitudeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttitudeResponse) ProtoMessage() {}

func (x *AttitudeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttitudeResponse.ProtoReflect.Descriptor instead.
func (*AttitudeResponse) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{9}
}

func (x *AttitudeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_community_proto protoreflect.FileDescriptor

var file_community_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7,
	0x02, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f,
	0x72, 0x57, 0x68, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x72,
	0x57, 0x68, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x70, 0x33, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x45, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73,
	0x65, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcf, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x70,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x6f, 0x6e, 0x67,
	0x12, 0x38, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x65,
	0x12, 0x34, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x68, 0x69, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5a, 0x68, 0x69, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x59, 0x75, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x75, 0x22, 0x1c, 0x0a, 0x08, 0x42, 0x75,
	0x6c, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xba, 0x02,
	0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x6f, 0x72, 0x57, 0x68, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x6f, 0x72, 0x57, 0x68, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x33, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x70, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x45, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x73, 0x65, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbd, 0x01, 0x0a, 0x0a, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x74, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74,
	0x6f, 0x12, 0x37, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x12, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x34, 0x0a, 0x10,
	0x41, 0x74, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f,
	0x70, 0x74, 0x22, 0x24, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa1, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x54, 0x6f, 0x70, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c,
	0x6c, 0x65, 0x74, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e,
	0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x12, 0x37, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x41, 0x74, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x41, 0x74, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_community_proto_rawDescOnce sync.Once
	file_community_proto_rawDescData = file_community_proto_rawDesc
)

func file_community_proto_rawDescGZIP() []byte {
	file_community_proto_rawDescOnce.Do(func() {
		file_community_proto_rawDescData = protoimpl.X.CompressGZIP(file_community_proto_rawDescData)
	})
	return file_community_proto_rawDescData
}

var file_community_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_community_proto_goTypes = []interface{}{
	(*CommunityCreation)(nil),      // 0: CommunityCreation
	(*TopCreation)(nil),            // 1: TopCreation
	(*Bulletin)(nil),               // 2: Bulletin
	(*SearchRequest)(nil),          // 3: SearchRequest
	(*PublicCreation)(nil),         // 4: PublicCreation
	(*SearchUser)(nil),             // 5: SearchUser
	(*SearchUserResponse)(nil),     // 6: SearchUserResponse
	(*SearchCreationResponse)(nil), // 7: SearchCreationResponse
	(*AttitudeCreation)(nil),       // 8: AttitudeCreation
	(*AttitudeResponse)(nil),       // 9: AttitudeResponse
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_community_proto_depIdxs = []int32{
	0,  // 0: TopCreation.creation:type_name -> CommunityCreation
	0,  // 1: TopCreation.creationGong:type_name -> CommunityCreation
	0,  // 2: TopCreation.creationShang:type_name -> CommunityCreation
	0,  // 3: TopCreation.creationJue:type_name -> CommunityCreation
	0,  // 4: TopCreation.creationZhi:type_name -> CommunityCreation
	0,  // 5: TopCreation.creationYu:type_name -> CommunityCreation
	4,  // 6: SearchUser.publicCreation:type_name -> PublicCreation
	5,  // 7: SearchUserResponse.users:type_name -> SearchUser
	4,  // 8: SearchCreationResponse.creations:type_name -> PublicCreation
	10, // 9: Community.GetTop:input_type -> google.protobuf.Empty
	10, // 10: Community.GetBulletin:input_type -> google.protobuf.Empty
	3,  // 11: Community.SearchUserByName:input_type -> SearchRequest
	3,  // 12: Community.SearchCreationByName:input_type -> SearchRequest
	8,  // 13: Community.CreationAttitude:input_type -> AttitudeCreation
	1,  // 14: Community.GetTop:output_type -> TopCreation
	2,  // 15: Community.GetBulletin:output_type -> Bulletin
	6,  // 16: Community.SearchUserByName:output_type -> SearchUserResponse
	7,  // 17: Community.SearchCreationByName:output_type -> SearchCreationResponse
	9,  // 18: Community.CreationAttitude:output_type -> AttitudeResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_community_proto_init() }
func file_community_proto_init() {
	if File_community_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_community_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityCreation); i {
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
		file_community_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopCreation); i {
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
		file_community_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bulletin); i {
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
		file_community_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_community_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicCreation); i {
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
		file_community_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUser); i {
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
		file_community_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserResponse); i {
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
		file_community_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCreationResponse); i {
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
		file_community_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttitudeCreation); i {
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
		file_community_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttitudeResponse); i {
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
			RawDescriptor: file_community_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_community_proto_goTypes,
		DependencyIndexes: file_community_proto_depIdxs,
		MessageInfos:      file_community_proto_msgTypes,
	}.Build()
	File_community_proto = out.File
	file_community_proto_rawDesc = nil
	file_community_proto_goTypes = nil
	file_community_proto_depIdxs = nil
}

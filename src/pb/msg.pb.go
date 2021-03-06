// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.3.0
// source: msg.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// MsgID:1
// 同步玩家本次登录ID 玩家登录后 由Server端主动生成玩家ID发送给客户端
type SyncPid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"` // 服务器返回的玩家ID
}

func (x *SyncPid) Reset() {
	*x = SyncPid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPid) ProtoMessage() {}

func (x *SyncPid) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPid.ProtoReflect.Descriptor instead.
func (*SyncPid) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SyncPid) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// MsgID:2
// 获取房间列表 发起者 client
type ListRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
}

func (x *ListRoom) Reset() {
	*x = ListRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoom) ProtoMessage() {}

func (x *ListRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoom.ProtoReflect.Descriptor instead.
func (*ListRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *ListRoom) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// MsgID:3
// 创建房间 发起者 client
type CreatRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxX     int32  `protobuf:"varint,1,opt,name=MaxX,proto3" json:"MaxX,omitempty"`        // 地图X大小
	CopiesX  int32  `protobuf:"varint,2,opt,name=CopiesX,proto3" json:"CopiesX,omitempty"`  // X轴方向的份数
	MaxY     int32  `protobuf:"varint,3,opt,name=MaxY,proto3" json:"MaxY,omitempty"`        // 地图Y大小
	CopiesY  int32  `protobuf:"varint,4,opt,name=CopiesY,proto3" json:"CopiesY,omitempty"`  // Y轴方向的份数
	Pid      int32  `protobuf:"varint,5,opt,name=Pid,proto3" json:"Pid,omitempty"`          // 房间成员Pid
	RoomName string `protobuf:"bytes,6,opt,name=RoomName,proto3" json:"RoomName,omitempty"` // 房间名称
}

func (x *CreatRoom) Reset() {
	*x = CreatRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatRoom) ProtoMessage() {}

func (x *CreatRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatRoom.ProtoReflect.Descriptor instead.
func (*CreatRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *CreatRoom) GetMaxX() int32 {
	if x != nil {
		return x.MaxX
	}
	return 0
}

func (x *CreatRoom) GetCopiesX() int32 {
	if x != nil {
		return x.CopiesX
	}
	return 0
}

func (x *CreatRoom) GetMaxY() int32 {
	if x != nil {
		return x.MaxY
	}
	return 0
}

func (x *CreatRoom) GetCopiesY() int32 {
	if x != nil {
		return x.CopiesY
	}
	return 0
}

func (x *CreatRoom) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *CreatRoom) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

// MsgID:4
// 加入房间 发起者 client
type JoinRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid      int32  `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	RoomName string `protobuf:"bytes,2,opt,name=RoomName,proto3" json:"RoomName,omitempty"`
}

func (x *JoinRoom) Reset() {
	*x = JoinRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoom) ProtoMessage() {}

func (x *JoinRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoom.ProtoReflect.Descriptor instead.
func (*JoinRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *JoinRoom) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *JoinRoom) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

// MsgID:5
// 聊天 发起者Client
type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tp      int32  `protobuf:"varint,1,opt,name=Tp,proto3" json:"Tp,omitempty"`          // 聊天类型 1 世界 2 房间 3 单人
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"` // 聊天内容
	Pid     int32  `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`        // 聊天目标人群 tp = 1,2 不需要取该值 服务器根据发起者pid自行获取世界和房间的玩家列表 只有当tp = 3 即单人的时候需要获取聊天对象的pid
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *Chat) GetTp() int32 {
	if x != nil {
		return x.Tp
	}
	return 0
}

func (x *Chat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Chat) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// MsgID:6
// 移动 发起者Client
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V float32 `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{5}
}

func (x *Position) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Position) GetV() float32 {
	if x != nil {
		return x.V
	}
	return 0
}

// MsgID:7
// 开始游戏 发起者 client
type Play struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomName string  `protobuf:"bytes,1,opt,name=RoomName,proto3" json:"RoomName,omitempty"`
	Pid      []int32 `protobuf:"varint,2,rep,packed,name=Pid,proto3" json:"Pid,omitempty"`
}

func (x *Play) Reset() {
	*x = Play{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Play) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Play) ProtoMessage() {}

func (x *Play) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Play.ProtoReflect.Descriptor instead.
func (*Play) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{6}
}

func (x *Play) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *Play) GetPid() []int32 {
	if x != nil {
		return x.Pid
	}
	return nil
}

// MsgID:9
// 广播消息 发起者Server
// Tp 消息类型 1 坐标 2 移动之后坐标信息更新 5 地图信息
type BroadCast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"` // 发起者玩家ID
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`   // 消息类型
	// Types that are assignable to Data:
	//	*BroadCast_P
	Data isBroadCast_Data `protobuf_oneof:"Data"`
}

func (x *BroadCast) Reset() {
	*x = BroadCast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadCast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadCast) ProtoMessage() {}

func (x *BroadCast) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadCast.ProtoReflect.Descriptor instead.
func (*BroadCast) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{7}
}

func (x *BroadCast) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *BroadCast) GetTp() int32 {
	if x != nil {
		return x.Tp
	}
	return 0
}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *BroadCast) GetP() *Position {
	if x, ok := x.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

func (*BroadCast_P) isBroadCast_Data() {}

// MsgID:10
// 同步玩家的显示数据
type SyncPlayers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ps []*Player `protobuf:"bytes,1,rep,name=ps,proto3" json:"ps,omitempty"`
}

func (x *SyncPlayers) Reset() {
	*x = SyncPlayers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayers) ProtoMessage() {}

func (x *SyncPlayers) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayers.ProtoReflect.Descriptor instead.
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{8}
}

func (x *SyncPlayers) GetPs() []*Player {
	if x != nil {
		return x.Ps
	}
	return nil
}

// 玩家信息
type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32     `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	P   *Position `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{9}
}

func (x *Player) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Player) GetP() *Position {
	if x != nil {
		return x.P
	}
	return nil
}

// MsgID:11
// 同步房间列表信息
type SyncRoomList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room []*Room `protobuf:"bytes,1,rep,name=room,proto3" json:"room,omitempty"`
}

func (x *SyncRoomList) Reset() {
	*x = SyncRoomList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRoomList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRoomList) ProtoMessage() {}

func (x *SyncRoomList) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRoomList.ProtoReflect.Descriptor instead.
func (*SyncRoomList) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{10}
}

func (x *SyncRoomList) GetRoom() []*Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	// TODO:扩展方向，发送房间大小，到时候决定是发送字符串描述大小还是发送房间的具体尺寸
	RoomName string `protobuf:"bytes,2,opt,name=RoomName,proto3" json:"RoomName,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{11}
}

func (x *Room) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Room) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

// MsgID:12
// 广播聊天信息
type BroadCastChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid     int32  `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`        // 发起者pid
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"` // 聊天内容
}

func (x *BroadCastChat) Reset() {
	*x = BroadCastChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadCastChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadCastChat) ProtoMessage() {}

func (x *BroadCastChat) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadCastChat.ProtoReflect.Descriptor instead.
func (*BroadCastChat) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{12}
}

func (x *BroadCastChat) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *BroadCastChat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x1b, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x78, 0x58,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4d, 0x61, 0x78, 0x58, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x70, 0x69, 0x65, 0x73, 0x58, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43,
	0x6f, 0x70, 0x69, 0x65, 0x73, 0x58, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x78, 0x59, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4d, 0x61, 0x78, 0x59, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x70, 0x69, 0x65, 0x73, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x6f, 0x70,
	0x69, 0x65, 0x73, 0x59, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x38, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x04,
	0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x54, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64,
	0x22, 0x42, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01,
	0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x5a, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x56, 0x22, 0x34, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x09, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x54, 0x70, 0x12, 0x1c, 0x0a, 0x01, 0x50, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x01, 0x50, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x29, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x02, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x02, 0x70, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x01, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x01, 0x50, 0x22, 0x2c, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x22, 0x34, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x05, 0xaa, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_msg_proto_goTypes = []interface{}{
	(*SyncPid)(nil),       // 0: pb.SyncPid
	(*ListRoom)(nil),      // 1: pb.ListRoom
	(*CreatRoom)(nil),     // 2: pb.CreatRoom
	(*JoinRoom)(nil),      // 3: pb.JoinRoom
	(*Chat)(nil),          // 4: pb.Chat
	(*Position)(nil),      // 5: pb.Position
	(*Play)(nil),          // 6: pb.Play
	(*BroadCast)(nil),     // 7: pb.BroadCast
	(*SyncPlayers)(nil),   // 8: pb.SyncPlayers
	(*Player)(nil),        // 9: pb.Player
	(*SyncRoomList)(nil),  // 10: pb.SyncRoomList
	(*Room)(nil),          // 11: pb.Room
	(*BroadCastChat)(nil), // 12: pb.BroadCastChat
}
var file_msg_proto_depIdxs = []int32{
	5,  // 0: pb.BroadCast.P:type_name -> pb.Position
	9,  // 1: pb.SyncPlayers.ps:type_name -> pb.Player
	5,  // 2: pb.Player.P:type_name -> pb.Position
	11, // 3: pb.SyncRoomList.room:type_name -> pb.Room
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPid); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoom); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatRoom); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoom); i {
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
		file_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Play); i {
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
		file_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadCast); i {
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
		file_msg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayers); i {
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
		file_msg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_msg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRoomList); i {
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
		file_msg_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_msg_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadCastChat); i {
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
	file_msg_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*BroadCast_P)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}

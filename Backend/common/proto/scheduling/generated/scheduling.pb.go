// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: scheduling.proto

package scheduling_pb

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

type ApparatusType int32

const (
	ApparatusType_Floor         ApparatusType = 0
	ApparatusType_PommelHorse   ApparatusType = 1
	ApparatusType_StillRings    ApparatusType = 2
	ApparatusType_Vault         ApparatusType = 3
	ApparatusType_ParallelBars  ApparatusType = 4
	ApparatusType_HorizontalBar ApparatusType = 5
	ApparatusType_BalanceBeam   ApparatusType = 6
	ApparatusType_UnevenBars    ApparatusType = 7
)

// Enum value maps for ApparatusType.
var (
	ApparatusType_name = map[int32]string{
		0: "Floor",
		1: "PommelHorse",
		2: "StillRings",
		3: "Vault",
		4: "ParallelBars",
		5: "HorizontalBar",
		6: "BalanceBeam",
		7: "UnevenBars",
	}
	ApparatusType_value = map[string]int32{
		"Floor":         0,
		"PommelHorse":   1,
		"StillRings":    2,
		"Vault":         3,
		"ParallelBars":  4,
		"HorizontalBar": 5,
		"BalanceBeam":   6,
		"UnevenBars":    7,
	}
)

func (x ApparatusType) Enum() *ApparatusType {
	p := new(ApparatusType)
	*p = x
	return p
}

func (x ApparatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApparatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_scheduling_proto_enumTypes[0].Descriptor()
}

func (ApparatusType) Type() protoreflect.EnumType {
	return &file_scheduling_proto_enumTypes[0]
}

func (x ApparatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApparatusType.Descriptor instead.
func (ApparatusType) EnumDescriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{0}
}

type IdMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdMessage) Reset() {
	*x = IdMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdMessage) ProtoMessage() {}

func (x *IdMessage) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdMessage.ProtoReflect.Descriptor instead.
func (*IdMessage) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{0}
}

func (x *IdMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AgeCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinAge int32  `protobuf:"varint,3,opt,name=minAge,proto3" json:"minAge,omitempty"`
	MaxAge int32  `protobuf:"varint,4,opt,name=maxAge,proto3" json:"maxAge,omitempty"`
}

func (x *AgeCategory) Reset() {
	*x = AgeCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgeCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgeCategory) ProtoMessage() {}

func (x *AgeCategory) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgeCategory.ProtoReflect.Descriptor instead.
func (*AgeCategory) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{1}
}

func (x *AgeCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgeCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgeCategory) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *AgeCategory) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

type ContestantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContestantCompId     int32           `protobuf:"varint,2,opt,name=contestantCompId,proto3" json:"contestantCompId,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TeamNumber           int32           `protobuf:"varint,4,opt,name=teamNumber,proto3" json:"teamNumber,omitempty"`
	Organization         string          `protobuf:"bytes,5,opt,name=organization,proto3" json:"organization,omitempty"`
	AgeCategory          string          `protobuf:"bytes,6,opt,name=ageCategory,proto3" json:"ageCategory,omitempty"`
	Location             string          `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	CompetingApparatuses []ApparatusType `protobuf:"varint,8,rep,packed,name=competingApparatuses,proto3,enum=scheduling_pb.ApparatusType" json:"competingApparatuses,omitempty"`
}

func (x *ContestantInfo) Reset() {
	*x = ContestantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContestantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContestantInfo) ProtoMessage() {}

func (x *ContestantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContestantInfo.ProtoReflect.Descriptor instead.
func (*ContestantInfo) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{2}
}

func (x *ContestantInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContestantInfo) GetContestantCompId() int32 {
	if x != nil {
		return x.ContestantCompId
	}
	return 0
}

func (x *ContestantInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContestantInfo) GetTeamNumber() int32 {
	if x != nil {
		return x.TeamNumber
	}
	return 0
}

func (x *ContestantInfo) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *ContestantInfo) GetAgeCategory() string {
	if x != nil {
		return x.AgeCategory
	}
	return ""
}

func (x *ContestantInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ContestantInfo) GetCompetingApparatuses() []ApparatusType {
	if x != nil {
		return x.CompetingApparatuses
	}
	return nil
}

type ScheduleSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session           int32           `protobuf:"varint,1,opt,name=session,proto3" json:"session,omitempty"`
	StartingApparatus ApparatusType   `protobuf:"varint,2,opt,name=startingApparatus,proto3,enum=scheduling_pb.ApparatusType" json:"startingApparatus,omitempty"`
	ContestantInfo    *ContestantInfo `protobuf:"bytes,3,opt,name=contestantInfo,proto3" json:"contestantInfo,omitempty"`
}

func (x *ScheduleSlot) Reset() {
	*x = ScheduleSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleSlot) ProtoMessage() {}

func (x *ScheduleSlot) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleSlot.ProtoReflect.Descriptor instead.
func (*ScheduleSlot) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{3}
}

func (x *ScheduleSlot) GetSession() int32 {
	if x != nil {
		return x.Session
	}
	return 0
}

func (x *ScheduleSlot) GetStartingApparatus() ApparatusType {
	if x != nil {
		return x.StartingApparatus
	}
	return ApparatusType_Floor
}

func (x *ScheduleSlot) GetContestantInfo() *ContestantInfo {
	if x != nil {
		return x.ContestantInfo
	}
	return nil
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slots          []*ScheduleSlot `protobuf:"bytes,2,rep,name=slots,proto3" json:"slots,omitempty"`
	StartingTimes  []int64         `protobuf:"varint,3,rep,packed,name=startingTimes,proto3" json:"startingTimes,omitempty"`
	ApparatusOrder []ApparatusType `protobuf:"varint,4,rep,packed,name=apparatusOrder,proto3,enum=scheduling_pb.ApparatusType" json:"apparatusOrder,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{4}
}

func (x *Schedule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Schedule) GetSlots() []*ScheduleSlot {
	if x != nil {
		return x.Slots
	}
	return nil
}

func (x *Schedule) GetStartingTimes() []int64 {
	if x != nil {
		return x.StartingTimes
	}
	return nil
}

func (x *Schedule) GetApparatusOrder() []ApparatusType {
	if x != nil {
		return x.ApparatusOrder
	}
	return nil
}

type Apparatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type ApparatusType `protobuf:"varint,2,opt,name=type,proto3,enum=scheduling_pb.ApparatusType" json:"type,omitempty"`
}

func (x *Apparatus) Reset() {
	*x = Apparatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apparatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apparatus) ProtoMessage() {}

func (x *Apparatus) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apparatus.ProtoReflect.Descriptor instead.
func (*Apparatus) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{5}
}

func (x *Apparatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Apparatus) GetType() ApparatusType {
	if x != nil {
		return x.Type
	}
	return ApparatusType_Floor
}

type SchedulingParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompetitionId                    string       `protobuf:"bytes,1,opt,name=competitionId,proto3" json:"competitionId,omitempty"`
	StartTime                        int64        `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime                          int64        `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	WarmupRoomAvailable              bool         `protobuf:"varint,4,opt,name=warmupRoomAvailable,proto3" json:"warmupRoomAvailable,omitempty"`
	GeneralWarmupTime                int32        `protobuf:"varint,5,opt,name=generalWarmupTime,proto3" json:"generalWarmupTime,omitempty"`
	WarmupTime                       int32        `protobuf:"varint,6,opt,name=warmupTime,proto3" json:"warmupTime,omitempty"`
	WarmupsPerApparatus              int32        `protobuf:"varint,7,opt,name=warmupsPerApparatus,proto3" json:"warmupsPerApparatus,omitempty"`
	ContestantNumPerApparatus        int32        `protobuf:"varint,8,opt,name=contestantNumPerApparatus,proto3" json:"contestantNumPerApparatus,omitempty"`
	ExecutionTime                    int32        `protobuf:"varint,9,opt,name=executionTime,proto3" json:"executionTime,omitempty"`
	ApparatusRotationTime            int32        `protobuf:"varint,10,opt,name=apparatusRotationTime,proto3" json:"apparatusRotationTime,omitempty"`
	MedalCeremonyAfterOneSessionTime int32        `protobuf:"varint,11,opt,name=medalCeremonyAfterOneSessionTime,proto3" json:"medalCeremonyAfterOneSessionTime,omitempty"`
	FinalMedalCeremonyTime           int32        `protobuf:"varint,12,opt,name=finalMedalCeremonyTime,proto3" json:"finalMedalCeremonyTime,omitempty"`
	HalfApparatusPerSessionMode      bool         `protobuf:"varint,13,opt,name=halfApparatusPerSessionMode,proto3" json:"halfApparatusPerSessionMode,omitempty"`
	ApparatusOrder                   []*Apparatus `protobuf:"bytes,14,rep,name=apparatusOrder,proto3" json:"apparatusOrder,omitempty"`
}

func (x *SchedulingParameters) Reset() {
	*x = SchedulingParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduling_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulingParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulingParameters) ProtoMessage() {}

func (x *SchedulingParameters) ProtoReflect() protoreflect.Message {
	mi := &file_scheduling_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulingParameters.ProtoReflect.Descriptor instead.
func (*SchedulingParameters) Descriptor() ([]byte, []int) {
	return file_scheduling_proto_rawDescGZIP(), []int{6}
}

func (x *SchedulingParameters) GetCompetitionId() string {
	if x != nil {
		return x.CompetitionId
	}
	return ""
}

func (x *SchedulingParameters) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *SchedulingParameters) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SchedulingParameters) GetWarmupRoomAvailable() bool {
	if x != nil {
		return x.WarmupRoomAvailable
	}
	return false
}

func (x *SchedulingParameters) GetGeneralWarmupTime() int32 {
	if x != nil {
		return x.GeneralWarmupTime
	}
	return 0
}

func (x *SchedulingParameters) GetWarmupTime() int32 {
	if x != nil {
		return x.WarmupTime
	}
	return 0
}

func (x *SchedulingParameters) GetWarmupsPerApparatus() int32 {
	if x != nil {
		return x.WarmupsPerApparatus
	}
	return 0
}

func (x *SchedulingParameters) GetContestantNumPerApparatus() int32 {
	if x != nil {
		return x.ContestantNumPerApparatus
	}
	return 0
}

func (x *SchedulingParameters) GetExecutionTime() int32 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *SchedulingParameters) GetApparatusRotationTime() int32 {
	if x != nil {
		return x.ApparatusRotationTime
	}
	return 0
}

func (x *SchedulingParameters) GetMedalCeremonyAfterOneSessionTime() int32 {
	if x != nil {
		return x.MedalCeremonyAfterOneSessionTime
	}
	return 0
}

func (x *SchedulingParameters) GetFinalMedalCeremonyTime() int32 {
	if x != nil {
		return x.FinalMedalCeremonyTime
	}
	return 0
}

func (x *SchedulingParameters) GetHalfApparatusPerSessionMode() bool {
	if x != nil {
		return x.HalfApparatusPerSessionMode
	}
	return false
}

func (x *SchedulingParameters) GetApparatusOrder() []*Apparatus {
	if x != nil {
		return x.ApparatusOrder
	}
	return nil
}

var File_scheduling_proto protoreflect.FileDescriptor

var file_scheduling_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x62, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61,
	0x0a, 0x0b, 0x41, 0x67, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78,
	0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67,
	0x65, 0x22, 0xb4, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x67, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x67, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70,
	0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41,
	0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x41,
	0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x45, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52,
	0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e,
	0x61, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x4d, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x70,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xc8, 0x05, 0x0a, 0x14, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x77, 0x61, 0x72, 0x6d,
	0x75, 0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x77, 0x61, 0x72, 0x6d, 0x75, 0x70, 0x52, 0x6f, 0x6f,
	0x6d, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x57, 0x61, 0x72, 0x6d, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x57, 0x61,
	0x72, 0x6d, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x72, 0x6d,
	0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x61,
	0x72, 0x6d, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x77, 0x61, 0x72, 0x6d,
	0x75, 0x70, 0x73, 0x50, 0x65, 0x72, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x77, 0x61, 0x72, 0x6d, 0x75, 0x70, 0x73, 0x50, 0x65,
	0x72, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x19, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x41, 0x70,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61,
	0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x20, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x43, 0x65, 0x72,
	0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x20,
	0x6d, 0x65, 0x64, 0x61, 0x6c, 0x43, 0x65, 0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x4f, 0x6e, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x36, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x61, 0x6c, 0x43, 0x65,
	0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x61, 0x6c, 0x43, 0x65, 0x72, 0x65,
	0x6d, 0x6f, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x68, 0x61, 0x6c, 0x66,
	0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x68,
	0x61, 0x6c, 0x66, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x61, 0x70,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2a, 0x8c, 0x01, 0x0a,
	0x0d, 0x41, 0x70, 0x70, 0x61, 0x72, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x6f, 0x6d,
	0x6d, 0x65, 0x6c, 0x48, 0x6f, 0x72, 0x73, 0x65, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74,
	0x69, 0x6c, 0x6c, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65,
	0x6c, 0x42, 0x61, 0x72, 0x73, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x6f, 0x72, 0x69, 0x7a,
	0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x42, 0x61, 0x72, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x61, 0x6d, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x55,
	0x6e, 0x65, 0x76, 0x65, 0x6e, 0x42, 0x61, 0x72, 0x73, 0x10, 0x07, 0x32, 0xb2, 0x01, 0x0a, 0x11,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x00,
	0x42, 0x16, 0x5a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduling_proto_rawDescOnce sync.Once
	file_scheduling_proto_rawDescData = file_scheduling_proto_rawDesc
)

func file_scheduling_proto_rawDescGZIP() []byte {
	file_scheduling_proto_rawDescOnce.Do(func() {
		file_scheduling_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduling_proto_rawDescData)
	})
	return file_scheduling_proto_rawDescData
}

var file_scheduling_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scheduling_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_scheduling_proto_goTypes = []interface{}{
	(ApparatusType)(0),           // 0: scheduling_pb.ApparatusType
	(*IdMessage)(nil),            // 1: scheduling_pb.IdMessage
	(*AgeCategory)(nil),          // 2: scheduling_pb.AgeCategory
	(*ContestantInfo)(nil),       // 3: scheduling_pb.ContestantInfo
	(*ScheduleSlot)(nil),         // 4: scheduling_pb.ScheduleSlot
	(*Schedule)(nil),             // 5: scheduling_pb.Schedule
	(*Apparatus)(nil),            // 6: scheduling_pb.Apparatus
	(*SchedulingParameters)(nil), // 7: scheduling_pb.SchedulingParameters
}
var file_scheduling_proto_depIdxs = []int32{
	0, // 0: scheduling_pb.ContestantInfo.competingApparatuses:type_name -> scheduling_pb.ApparatusType
	0, // 1: scheduling_pb.ScheduleSlot.startingApparatus:type_name -> scheduling_pb.ApparatusType
	3, // 2: scheduling_pb.ScheduleSlot.contestantInfo:type_name -> scheduling_pb.ContestantInfo
	4, // 3: scheduling_pb.Schedule.slots:type_name -> scheduling_pb.ScheduleSlot
	0, // 4: scheduling_pb.Schedule.apparatusOrder:type_name -> scheduling_pb.ApparatusType
	0, // 5: scheduling_pb.Apparatus.type:type_name -> scheduling_pb.ApparatusType
	6, // 6: scheduling_pb.SchedulingParameters.apparatusOrder:type_name -> scheduling_pb.Apparatus
	7, // 7: scheduling_pb.SchedulingService.GenerateSchedule:input_type -> scheduling_pb.SchedulingParameters
	1, // 8: scheduling_pb.SchedulingService.GetByCompetitionId:input_type -> scheduling_pb.IdMessage
	5, // 9: scheduling_pb.SchedulingService.GenerateSchedule:output_type -> scheduling_pb.Schedule
	5, // 10: scheduling_pb.SchedulingService.GetByCompetitionId:output_type -> scheduling_pb.Schedule
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_scheduling_proto_init() }
func file_scheduling_proto_init() {
	if File_scheduling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdMessage); i {
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
		file_scheduling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgeCategory); i {
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
		file_scheduling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContestantInfo); i {
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
		file_scheduling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleSlot); i {
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
		file_scheduling_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_scheduling_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apparatus); i {
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
		file_scheduling_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulingParameters); i {
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
			RawDescriptor: file_scheduling_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduling_proto_goTypes,
		DependencyIndexes: file_scheduling_proto_depIdxs,
		EnumInfos:         file_scheduling_proto_enumTypes,
		MessageInfos:      file_scheduling_proto_msgTypes,
	}.Build()
	File_scheduling_proto = out.File
	file_scheduling_proto_rawDesc = nil
	file_scheduling_proto_goTypes = nil
	file_scheduling_proto_depIdxs = nil
}

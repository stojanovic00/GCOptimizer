// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: application.proto

package application_pb

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

type IdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdResponse) Reset() {
	*x = IdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponse) ProtoMessage() {}

func (x *IdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponse.ProtoReflect.Descriptor instead.
func (*IdResponse) Descriptor() ([]byte, []int) {
	return file_application_proto_rawDescGZIP(), []int{0}
}

func (x *IdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Country      string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City         string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Street       string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	StreetNumber string `protobuf:"bytes,5,opt,name=streetNumber,proto3" json:"streetNumber,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_application_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetStreetNumber() string {
	if x != nil {
		return x.StreetNumber
	}
	return ""
}

type SportsOrganisation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                          string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber                    string   `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	ContactPersonFullName          string   `protobuf:"bytes,5,opt,name=contactPersonFullName,proto3" json:"contactPersonFullName,omitempty"`
	CompetitionOrganisingPrivilege bool     `protobuf:"varint,6,opt,name=competitionOrganisingPrivilege,proto3" json:"competitionOrganisingPrivilege,omitempty"`
	Address                        *Address `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *SportsOrganisation) Reset() {
	*x = SportsOrganisation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SportsOrganisation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SportsOrganisation) ProtoMessage() {}

func (x *SportsOrganisation) ProtoReflect() protoreflect.Message {
	mi := &file_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SportsOrganisation.ProtoReflect.Descriptor instead.
func (*SportsOrganisation) Descriptor() ([]byte, []int) {
	return file_application_proto_rawDescGZIP(), []int{2}
}

func (x *SportsOrganisation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SportsOrganisation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SportsOrganisation) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SportsOrganisation) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SportsOrganisation) GetContactPersonFullName() string {
	if x != nil {
		return x.ContactPersonFullName
	}
	return ""
}

func (x *SportsOrganisation) GetCompetitionOrganisingPrivilege() bool {
	if x != nil {
		return x.CompetitionOrganisingPrivilege
	}
	return false
}

func (x *SportsOrganisation) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_application_proto protoreflect.FileDescriptor

var file_application_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x62, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xa1, 0x02, 0x0a, 0x12, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x46, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x74, 0x0a, 0x12, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x1a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62,
	0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x17, 0x5a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_application_proto_rawDescOnce sync.Once
	file_application_proto_rawDescData = file_application_proto_rawDesc
)

func file_application_proto_rawDescGZIP() []byte {
	file_application_proto_rawDescOnce.Do(func() {
		file_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_proto_rawDescData)
	})
	return file_application_proto_rawDescData
}

var file_application_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_application_proto_goTypes = []interface{}{
	(*IdResponse)(nil),         // 0: application_pb.IdResponse
	(*Address)(nil),            // 1: application_pb.Address
	(*SportsOrganisation)(nil), // 2: application_pb.SportsOrganisation
}
var file_application_proto_depIdxs = []int32{
	1, // 0: application_pb.SportsOrganisation.address:type_name -> application_pb.Address
	2, // 1: application_pb.ApplicationService.RegisterSportsOrganisation:input_type -> application_pb.SportsOrganisation
	0, // 2: application_pb.ApplicationService.RegisterSportsOrganisation:output_type -> application_pb.IdResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_application_proto_init() }
func file_application_proto_init() {
	if File_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdResponse); i {
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
		file_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SportsOrganisation); i {
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
			RawDescriptor: file_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_proto_goTypes,
		DependencyIndexes: file_application_proto_depIdxs,
		MessageInfos:      file_application_proto_msgTypes,
	}.Build()
	File_application_proto = out.File
	file_application_proto_rawDesc = nil
	file_application_proto_goTypes = nil
	file_application_proto_depIdxs = nil
}
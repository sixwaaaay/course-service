// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: school.proto

package school

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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	ServiceTime   int64  `protobuf:"varint,3,opt,name=service_time,json=serviceTime,proto3" json:"service_time,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *BaseResp) GetServiceTime() int64 {
	if x != nil {
		return x.ServiceTime
	}
	return 0
}

type School struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *School) Reset() {
	*x = School{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *School) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*School) ProtoMessage() {}

func (x *School) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use School.ProtoReflect.Descriptor instead.
func (*School) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{1}
}

func (x *School) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *School) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MGetSchoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchoolIds []int32 `protobuf:"varint,1,rep,packed,name=school_ids,json=schoolIds,proto3" json:"school_ids,omitempty"`
}

func (x *MGetSchoolRequest) Reset() {
	*x = MGetSchoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MGetSchoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetSchoolRequest) ProtoMessage() {}

func (x *MGetSchoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetSchoolRequest.ProtoReflect.Descriptor instead.
func (*MGetSchoolRequest) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{2}
}

func (x *MGetSchoolRequest) GetSchoolIds() []int32 {
	if x != nil {
		return x.SchoolIds
	}
	return nil
}

type MGetSchoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schools  []*School `protobuf:"bytes,1,rep,name=schools,proto3" json:"schools,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,2,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *MGetSchoolResponse) Reset() {
	*x = MGetSchoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MGetSchoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetSchoolResponse) ProtoMessage() {}

func (x *MGetSchoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetSchoolResponse.ProtoReflect.Descriptor instead.
func (*MGetSchoolResponse) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{3}
}

func (x *MGetSchoolResponse) GetSchools() []*School {
	if x != nil {
		return x.Schools
	}
	return nil
}

func (x *MGetSchoolResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type QueryIdByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryIdByNameRequest) Reset() {
	*x = QueryIdByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryIdByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryIdByNameRequest) ProtoMessage() {}

func (x *QueryIdByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryIdByNameRequest.ProtoReflect.Descriptor instead.
func (*QueryIdByNameRequest) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{4}
}

func (x *QueryIdByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryIdByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchoolId int32     `protobuf:"varint,1,opt,name=school_id,json=schoolId,proto3" json:"school_id,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,2,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *QueryIdByNameResponse) Reset() {
	*x = QueryIdByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryIdByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryIdByNameResponse) ProtoMessage() {}

func (x *QueryIdByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryIdByNameResponse.ProtoReflect.Descriptor instead.
func (*QueryIdByNameResponse) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{5}
}

func (x *QueryIdByNameResponse) GetSchoolId() int32 {
	if x != nil {
		return x.SchoolId
	}
	return 0
}

func (x *QueryIdByNameResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_school_proto protoreflect.FileDescriptor

var file_school_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x22, 0x75, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x06, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x11, 0x4d,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x73, 0x22,
	0x6d, 0x0a, 0x12, 0x4d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e,
	0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x52, 0x07, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x12,
	0x2d, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2a,
	0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x15, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32,
	0xa6, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x0a, 0x4d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12,
	0x19, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x49, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_school_proto_rawDescOnce sync.Once
	file_school_proto_rawDescData = file_school_proto_rawDesc
)

func file_school_proto_rawDescGZIP() []byte {
	file_school_proto_rawDescOnce.Do(func() {
		file_school_proto_rawDescData = protoimpl.X.CompressGZIP(file_school_proto_rawDescData)
	})
	return file_school_proto_rawDescData
}

var file_school_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_school_proto_goTypes = []interface{}{
	(*BaseResp)(nil),              // 0: school.BaseResp
	(*School)(nil),                // 1: school.School
	(*MGetSchoolRequest)(nil),     // 2: school.MGetSchoolRequest
	(*MGetSchoolResponse)(nil),    // 3: school.MGetSchoolResponse
	(*QueryIdByNameRequest)(nil),  // 4: school.QueryIdByNameRequest
	(*QueryIdByNameResponse)(nil), // 5: school.QueryIdByNameResponse
}
var file_school_proto_depIdxs = []int32{
	1, // 0: school.MGetSchoolResponse.schools:type_name -> school.School
	0, // 1: school.MGetSchoolResponse.base_resp:type_name -> school.BaseResp
	0, // 2: school.QueryIdByNameResponse.base_resp:type_name -> school.BaseResp
	2, // 3: school.SchoolService.MGetSchool:input_type -> school.MGetSchoolRequest
	4, // 4: school.SchoolService.QueryIdByName:input_type -> school.QueryIdByNameRequest
	3, // 5: school.SchoolService.MGetSchool:output_type -> school.MGetSchoolResponse
	5, // 6: school.SchoolService.QueryIdByName:output_type -> school.QueryIdByNameResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_school_proto_init() }
func file_school_proto_init() {
	if File_school_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_school_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_school_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*School); i {
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
		file_school_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MGetSchoolRequest); i {
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
		file_school_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MGetSchoolResponse); i {
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
		file_school_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryIdByNameRequest); i {
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
		file_school_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryIdByNameResponse); i {
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
			RawDescriptor: file_school_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_school_proto_goTypes,
		DependencyIndexes: file_school_proto_depIdxs,
		MessageInfos:      file_school_proto_msgTypes,
	}.Build()
	File_school_proto = out.File
	file_school_proto_rawDesc = nil
	file_school_proto_goTypes = nil
	file_school_proto_depIdxs = nil
}

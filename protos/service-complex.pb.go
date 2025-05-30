// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: protos/service-complex.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnaryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        float64                `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnaryRequest) Reset() {
	*x = UnaryRequest{}
	mi := &file_protos_service_complex_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRequest) ProtoMessage() {}

func (x *UnaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_complex_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRequest.ProtoReflect.Descriptor instead.
func (*UnaryRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_complex_proto_rawDescGZIP(), []int{0}
}

func (x *UnaryRequest) GetNumber() float64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type BinaryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstNumber   float64                `protobuf:"fixed64,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber  float64                `protobuf:"fixed64,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BinaryRequest) Reset() {
	*x = BinaryRequest{}
	mi := &file_protos_service_complex_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryRequest) ProtoMessage() {}

func (x *BinaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_complex_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryRequest.ProtoReflect.Descriptor instead.
func (*BinaryRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_complex_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryRequest) GetFirstNumber() float64 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *BinaryRequest) GetSecondNumber() float64 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type OperationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperationResponse) Reset() {
	*x = OperationResponse{}
	mi := &file_protos_service_complex_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponse) ProtoMessage() {}

func (x *OperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_complex_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponse.ProtoReflect.Descriptor instead.
func (*OperationResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_complex_proto_rawDescGZIP(), []int{2}
}

func (x *OperationResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_protos_service_complex_proto protoreflect.FileDescriptor

const file_protos_service_complex_proto_rawDesc = "" +
	"\n" +
	"\x1cprotos/service-complex.proto\x12\acomplex\"&\n" +
	"\fUnaryRequest\x12\x16\n" +
	"\x06number\x18\x01 \x01(\x01R\x06number\"W\n" +
	"\rBinaryRequest\x12!\n" +
	"\ffirst_number\x18\x01 \x01(\x01R\vfirstNumber\x12#\n" +
	"\rsecond_number\x18\x02 \x01(\x01R\fsecondNumber\"+\n" +
	"\x11OperationResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06result2\xc4\x02\n" +
	"\x0eComplexCompute\x12;\n" +
	"\x04Sqrt\x12\x15.complex.UnaryRequest\x1a\x1a.complex.OperationResponse\"\x00\x12:\n" +
	"\x03Abs\x12\x15.complex.UnaryRequest\x1a\x1a.complex.OperationResponse\"\x00\x12=\n" +
	"\x05Power\x12\x16.complex.BinaryRequest\x1a\x1a.complex.OperationResponse\"\x00\x12;\n" +
	"\x03Log\x12\x16.complex.BinaryRequest\x1a\x1a.complex.OperationResponse\"\x00\x12=\n" +
	"\x05Round\x12\x16.complex.BinaryRequest\x1a\x1a.complex.OperationResponse\"\x00B:Z8github.com/lubie-koty/rpc-compute-service-complex/protosb\x06proto3"

var (
	file_protos_service_complex_proto_rawDescOnce sync.Once
	file_protos_service_complex_proto_rawDescData []byte
)

func file_protos_service_complex_proto_rawDescGZIP() []byte {
	file_protos_service_complex_proto_rawDescOnce.Do(func() {
		file_protos_service_complex_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_service_complex_proto_rawDesc), len(file_protos_service_complex_proto_rawDesc)))
	})
	return file_protos_service_complex_proto_rawDescData
}

var file_protos_service_complex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_service_complex_proto_goTypes = []any{
	(*UnaryRequest)(nil),      // 0: complex.UnaryRequest
	(*BinaryRequest)(nil),     // 1: complex.BinaryRequest
	(*OperationResponse)(nil), // 2: complex.OperationResponse
}
var file_protos_service_complex_proto_depIdxs = []int32{
	0, // 0: complex.ComplexCompute.Sqrt:input_type -> complex.UnaryRequest
	0, // 1: complex.ComplexCompute.Abs:input_type -> complex.UnaryRequest
	1, // 2: complex.ComplexCompute.Power:input_type -> complex.BinaryRequest
	1, // 3: complex.ComplexCompute.Log:input_type -> complex.BinaryRequest
	1, // 4: complex.ComplexCompute.Round:input_type -> complex.BinaryRequest
	2, // 5: complex.ComplexCompute.Sqrt:output_type -> complex.OperationResponse
	2, // 6: complex.ComplexCompute.Abs:output_type -> complex.OperationResponse
	2, // 7: complex.ComplexCompute.Power:output_type -> complex.OperationResponse
	2, // 8: complex.ComplexCompute.Log:output_type -> complex.OperationResponse
	2, // 9: complex.ComplexCompute.Round:output_type -> complex.OperationResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_service_complex_proto_init() }
func file_protos_service_complex_proto_init() {
	if File_protos_service_complex_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_service_complex_proto_rawDesc), len(file_protos_service_complex_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_complex_proto_goTypes,
		DependencyIndexes: file_protos_service_complex_proto_depIdxs,
		MessageInfos:      file_protos_service_complex_proto_msgTypes,
	}.Build()
	File_protos_service_complex_proto = out.File
	file_protos_service_complex_proto_goTypes = nil
	file_protos_service_complex_proto_depIdxs = nil
}

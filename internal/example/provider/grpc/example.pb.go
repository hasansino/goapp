// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: example.proto

package grpc

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

type Fruit struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Fruit) Reset() {
	*x = Fruit{}
	mi := &file_example_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fruit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fruit) ProtoMessage() {}

func (x *Fruit) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fruit.ProtoReflect.Descriptor instead.
func (*Fruit) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *Fruit) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Fruit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListFruitsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFruitsRequest) Reset() {
	*x = ListFruitsRequest{}
	mi := &file_example_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFruitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFruitsRequest) ProtoMessage() {}

func (x *ListFruitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFruitsRequest.ProtoReflect.Descriptor instead.
func (*ListFruitsRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *ListFruitsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListFruitsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListFruitsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fruits        []*Fruit               `protobuf:"bytes,1,rep,name=fruits,proto3" json:"fruits,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFruitsResponse) Reset() {
	*x = ListFruitsResponse{}
	mi := &file_example_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFruitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFruitsResponse) ProtoMessage() {}

func (x *ListFruitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFruitsResponse.ProtoReflect.Descriptor instead.
func (*ListFruitsResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *ListFruitsResponse) GetFruits() []*Fruit {
	if x != nil {
		return x.Fruits
	}
	return nil
}

type GetFruitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFruitRequest) Reset() {
	*x = GetFruitRequest{}
	mi := &file_example_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFruitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFruitRequest) ProtoMessage() {}

func (x *GetFruitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFruitRequest.ProtoReflect.Descriptor instead.
func (*GetFruitRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{3}
}

func (x *GetFruitRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateFruitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFruitRequest) Reset() {
	*x = CreateFruitRequest{}
	mi := &file_example_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFruitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFruitRequest) ProtoMessage() {}

func (x *CreateFruitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFruitRequest.ProtoReflect.Descriptor instead.
func (*CreateFruitRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFruitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateFruitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fruit         *Fruit                 `protobuf:"bytes,1,opt,name=fruit,proto3" json:"fruit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFruitResponse) Reset() {
	*x = CreateFruitResponse{}
	mi := &file_example_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFruitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFruitResponse) ProtoMessage() {}

func (x *CreateFruitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFruitResponse.ProtoReflect.Descriptor instead.
func (*CreateFruitResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{5}
}

func (x *CreateFruitResponse) GetFruit() *Fruit {
	if x != nil {
		return x.Fruit
	}
	return nil
}

type UpdateFruitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFruitRequest) Reset() {
	*x = UpdateFruitRequest{}
	mi := &file_example_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFruitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFruitRequest) ProtoMessage() {}

func (x *UpdateFruitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFruitRequest.ProtoReflect.Descriptor instead.
func (*UpdateFruitRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFruitRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFruitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateFruitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fruit         *Fruit                 `protobuf:"bytes,1,opt,name=fruit,proto3" json:"fruit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFruitResponse) Reset() {
	*x = UpdateFruitResponse{}
	mi := &file_example_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFruitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFruitResponse) ProtoMessage() {}

func (x *UpdateFruitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFruitResponse.ProtoReflect.Descriptor instead.
func (*UpdateFruitResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateFruitResponse) GetFruit() *Fruit {
	if x != nil {
		return x.Fruit
	}
	return nil
}

type DeleteFruitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFruitRequest) Reset() {
	*x = DeleteFruitRequest{}
	mi := &file_example_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFruitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFruitRequest) ProtoMessage() {}

func (x *DeleteFruitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFruitRequest.ProtoReflect.Descriptor instead.
func (*DeleteFruitRequest) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFruitRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteFruitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFruitResponse) Reset() {
	*x = DeleteFruitResponse{}
	mi := &file_example_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFruitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFruitResponse) ProtoMessage() {}

func (x *DeleteFruitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFruitResponse.ProtoReflect.Descriptor instead.
func (*DeleteFruitResponse) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteFruitResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_example_proto protoreflect.FileDescriptor

const file_example_proto_rawDesc = "" +
	"\n" +
	"\rexample.proto\x12\x04grpc\"+\n" +
	"\x05Fruit\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"A\n" +
	"\x11ListFruitsRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x05R\x06offset\"9\n" +
	"\x12ListFruitsResponse\x12#\n" +
	"\x06fruits\x18\x01 \x03(\v2\v.grpc.FruitR\x06fruits\"!\n" +
	"\x0fGetFruitRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"(\n" +
	"\x12CreateFruitRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"8\n" +
	"\x13CreateFruitResponse\x12!\n" +
	"\x05fruit\x18\x01 \x01(\v2\v.grpc.FruitR\x05fruit\"8\n" +
	"\x12UpdateFruitRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"8\n" +
	"\x13UpdateFruitResponse\x12!\n" +
	"\x05fruit\x18\x01 \x01(\v2\v.grpc.FruitR\x05fruit\"$\n" +
	"\x12DeleteFruitRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"/\n" +
	"\x13DeleteFruitResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xcd\x02\n" +
	"\x0eExampleService\x12?\n" +
	"\n" +
	"ListFruits\x12\x17.grpc.ListFruitsRequest\x1a\x18.grpc.ListFruitsResponse\x12.\n" +
	"\bGetFruit\x12\x15.grpc.GetFruitRequest\x1a\v.grpc.Fruit\x12B\n" +
	"\vCreateFruit\x12\x18.grpc.CreateFruitRequest\x1a\x19.grpc.CreateFruitResponse\x12B\n" +
	"\vUpdateFruit\x12\x18.grpc.UpdateFruitRequest\x1a\x19.grpc.UpdateFruitResponse\x12B\n" +
	"\vDeleteFruit\x12\x18.grpc.DeleteFruitRequest\x1a\x19.grpc.DeleteFruitResponseB:Z8github.com/hasansino/go42/internal/example/provider/grpcb\x06proto3"

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData []byte
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_example_proto_rawDesc), len(file_example_proto_rawDesc)))
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_example_proto_goTypes = []any{
	(*Fruit)(nil),               // 0: grpc.Fruit
	(*ListFruitsRequest)(nil),   // 1: grpc.ListFruitsRequest
	(*ListFruitsResponse)(nil),  // 2: grpc.ListFruitsResponse
	(*GetFruitRequest)(nil),     // 3: grpc.GetFruitRequest
	(*CreateFruitRequest)(nil),  // 4: grpc.CreateFruitRequest
	(*CreateFruitResponse)(nil), // 5: grpc.CreateFruitResponse
	(*UpdateFruitRequest)(nil),  // 6: grpc.UpdateFruitRequest
	(*UpdateFruitResponse)(nil), // 7: grpc.UpdateFruitResponse
	(*DeleteFruitRequest)(nil),  // 8: grpc.DeleteFruitRequest
	(*DeleteFruitResponse)(nil), // 9: grpc.DeleteFruitResponse
}
var file_example_proto_depIdxs = []int32{
	0, // 0: grpc.ListFruitsResponse.fruits:type_name -> grpc.Fruit
	0, // 1: grpc.CreateFruitResponse.fruit:type_name -> grpc.Fruit
	0, // 2: grpc.UpdateFruitResponse.fruit:type_name -> grpc.Fruit
	1, // 3: grpc.ExampleService.ListFruits:input_type -> grpc.ListFruitsRequest
	3, // 4: grpc.ExampleService.GetFruit:input_type -> grpc.GetFruitRequest
	4, // 5: grpc.ExampleService.CreateFruit:input_type -> grpc.CreateFruitRequest
	6, // 6: grpc.ExampleService.UpdateFruit:input_type -> grpc.UpdateFruitRequest
	8, // 7: grpc.ExampleService.DeleteFruit:input_type -> grpc.DeleteFruitRequest
	2, // 8: grpc.ExampleService.ListFruits:output_type -> grpc.ListFruitsResponse
	0, // 9: grpc.ExampleService.GetFruit:output_type -> grpc.Fruit
	5, // 10: grpc.ExampleService.CreateFruit:output_type -> grpc.CreateFruitResponse
	7, // 11: grpc.ExampleService.UpdateFruit:output_type -> grpc.UpdateFruitResponse
	9, // 12: grpc.ExampleService.DeleteFruit:output_type -> grpc.DeleteFruitResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_proto_rawDesc), len(file_example_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: movie.proto

package gen

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

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Director    string `protobuf:"bytes,4,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Metadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Metadata) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

type MovieDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating   float32   `protobuf:"fixed32,1,opt,name=rating,proto3" json:"rating,omitempty"`
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MovieDetails) Reset() {
	*x = MovieDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetails) ProtoMessage() {}

func (x *MovieDetails) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetails.ProtoReflect.Descriptor instead.
func (*MovieDetails) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *MovieDetails) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *MovieDetails) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x4d, 0x0a,
	0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_movie_proto_goTypes = []interface{}{
	(*Metadata)(nil),     // 0: Metadata
	(*MovieDetails)(nil), // 1: MovieDetails
}
var file_movie_proto_depIdxs = []int32{
	0, // 0: MovieDetails.metadata:type_name -> Metadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetails); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}

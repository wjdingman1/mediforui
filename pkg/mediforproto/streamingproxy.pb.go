// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: medifor/v1/streamingproxy.proto

package mediforproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// FileChunk contains a single chunk of a file. It is intended to be part of a
// streaming protocol, so size and offset are implied
type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the file. Should be sent with every chunk. Works best when every file sent is uniquely named.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The (optional) MIME type of the file.
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// This chunk's offset into the file specified by the name above.
	Offset int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// The total size of the file.
	TotalBytes int64 `protobuf:"varint,4,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	// The value of this chunk.
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_streamingproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_streamingproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_medifor_v1_streamingproxy_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileChunk) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *FileChunk) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FileChunk) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *FileChunk) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// DetectionChunk is part of the streaming detection protocol, where detection
// objects are sent (both as request and response) along with the files that go
// with them.
type DetectionChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*DetectionChunk_Detection
	//	*DetectionChunk_FileChunk
	Value isDetectionChunk_Value `protobuf_oneof:"value"`
}

func (x *DetectionChunk) Reset() {
	*x = DetectionChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_streamingproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectionChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectionChunk) ProtoMessage() {}

func (x *DetectionChunk) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_streamingproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectionChunk.ProtoReflect.Descriptor instead.
func (*DetectionChunk) Descriptor() ([]byte, []int) {
	return file_medifor_v1_streamingproxy_proto_rawDescGZIP(), []int{1}
}

func (m *DetectionChunk) GetValue() isDetectionChunk_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *DetectionChunk) GetDetection() *Detection {
	if x, ok := x.GetValue().(*DetectionChunk_Detection); ok {
		return x.Detection
	}
	return nil
}

func (x *DetectionChunk) GetFileChunk() *FileChunk {
	if x, ok := x.GetValue().(*DetectionChunk_FileChunk); ok {
		return x.FileChunk
	}
	return nil
}

type isDetectionChunk_Value interface {
	isDetectionChunk_Value()
}

type DetectionChunk_Detection struct {
	// A Detection proto that contains a request and/or response, depending on
	// what is needed. Note that any chunks provided should cover all necessary
	// inputs/outputs as specified in this proto.
	Detection *Detection `protobuf:"bytes,1,opt,name=detection,proto3,oneof"`
}

type DetectionChunk_FileChunk struct {
	// Chunks of file(s) used for or provided in the detection above. Any files
	// mentioned in the Detection should be provided here as needed. For
	// example, on request all Resource URIs should be represented in these
	// file chunks as file chunk names. On response, all response URIs should
	// be present in a similar way.
	FileChunk *FileChunk `protobuf:"bytes,2,opt,name=file_chunk,json=fileChunk,proto3,oneof"`
}

func (*DetectionChunk_Detection) isDetectionChunk_Value() {}

func (*DetectionChunk_FileChunk) isDetectionChunk_Value() {}

var File_medifor_v1_streamingproxy_proto protoreflect.FileDescriptor

var file_medifor_v1_streamingproxy_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x37, 0x0a, 0x09, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66,
	0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x48, 0x00, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x60, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x4e, 0x0a, 0x0c, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_medifor_v1_streamingproxy_proto_rawDescOnce sync.Once
	file_medifor_v1_streamingproxy_proto_rawDescData = file_medifor_v1_streamingproxy_proto_rawDesc
)

func file_medifor_v1_streamingproxy_proto_rawDescGZIP() []byte {
	file_medifor_v1_streamingproxy_proto_rawDescOnce.Do(func() {
		file_medifor_v1_streamingproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_medifor_v1_streamingproxy_proto_rawDescData)
	})
	return file_medifor_v1_streamingproxy_proto_rawDescData
}

var file_medifor_v1_streamingproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_medifor_v1_streamingproxy_proto_goTypes = []interface{}{
	(*FileChunk)(nil),      // 0: mediforproto.FileChunk
	(*DetectionChunk)(nil), // 1: mediforproto.DetectionChunk
	(*Detection)(nil),      // 2: mediforproto.Detection
}
var file_medifor_v1_streamingproxy_proto_depIdxs = []int32{
	2, // 0: mediforproto.DetectionChunk.detection:type_name -> mediforproto.Detection
	0, // 1: mediforproto.DetectionChunk.file_chunk:type_name -> mediforproto.FileChunk
	1, // 2: mediforproto.StreamingProxy.DetectStream:input_type -> mediforproto.DetectionChunk
	1, // 3: mediforproto.StreamingProxy.DetectStream:output_type -> mediforproto.DetectionChunk
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_medifor_v1_streamingproxy_proto_init() }
func file_medifor_v1_streamingproxy_proto_init() {
	if File_medifor_v1_streamingproxy_proto != nil {
		return
	}
	file_medifor_v1_analytic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_medifor_v1_streamingproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_medifor_v1_streamingproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectionChunk); i {
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
	file_medifor_v1_streamingproxy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DetectionChunk_Detection)(nil),
		(*DetectionChunk_FileChunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_medifor_v1_streamingproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_medifor_v1_streamingproxy_proto_goTypes,
		DependencyIndexes: file_medifor_v1_streamingproxy_proto_depIdxs,
		MessageInfos:      file_medifor_v1_streamingproxy_proto_msgTypes,
	}.Build()
	File_medifor_v1_streamingproxy_proto = out.File
	file_medifor_v1_streamingproxy_proto_rawDesc = nil
	file_medifor_v1_streamingproxy_proto_goTypes = nil
	file_medifor_v1_streamingproxy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamingProxyClient is the client API for StreamingProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamingProxyClient interface {
	DetectStream(ctx context.Context, opts ...grpc.CallOption) (StreamingProxy_DetectStreamClient, error)
}

type streamingProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingProxyClient(cc grpc.ClientConnInterface) StreamingProxyClient {
	return &streamingProxyClient{cc}
}

func (c *streamingProxyClient) DetectStream(ctx context.Context, opts ...grpc.CallOption) (StreamingProxy_DetectStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamingProxy_serviceDesc.Streams[0], "/mediforproto.StreamingProxy/DetectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingProxyDetectStreamClient{stream}
	return x, nil
}

type StreamingProxy_DetectStreamClient interface {
	Send(*DetectionChunk) error
	Recv() (*DetectionChunk, error)
	grpc.ClientStream
}

type streamingProxyDetectStreamClient struct {
	grpc.ClientStream
}

func (x *streamingProxyDetectStreamClient) Send(m *DetectionChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingProxyDetectStreamClient) Recv() (*DetectionChunk, error) {
	m := new(DetectionChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingProxyServer is the server API for StreamingProxy service.
type StreamingProxyServer interface {
	DetectStream(StreamingProxy_DetectStreamServer) error
}

// UnimplementedStreamingProxyServer can be embedded to have forward compatible implementations.
type UnimplementedStreamingProxyServer struct {
}

func (*UnimplementedStreamingProxyServer) DetectStream(StreamingProxy_DetectStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DetectStream not implemented")
}

func RegisterStreamingProxyServer(s *grpc.Server, srv StreamingProxyServer) {
	s.RegisterService(&_StreamingProxy_serviceDesc, srv)
}

func _StreamingProxy_DetectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingProxyServer).DetectStream(&streamingProxyDetectStreamServer{stream})
}

type StreamingProxy_DetectStreamServer interface {
	Send(*DetectionChunk) error
	Recv() (*DetectionChunk, error)
	grpc.ServerStream
}

type streamingProxyDetectStreamServer struct {
	grpc.ServerStream
}

func (x *streamingProxyDetectStreamServer) Send(m *DetectionChunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingProxyDetectStreamServer) Recv() (*DetectionChunk, error) {
	m := new(DetectionChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamingProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mediforproto.StreamingProxy",
	HandlerType: (*StreamingProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DetectStream",
			Handler:       _StreamingProxy_DetectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "medifor/v1/streamingproxy.proto",
}

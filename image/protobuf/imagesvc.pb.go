// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/imagesvc.proto

package image

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FilesList struct {
	FileName             []string `protobuf:"bytes,1,rep,name=fileName,proto3" json:"fileName,omitempty"`
	Tar                  string   `protobuf:"bytes,2,opt,name=tar,proto3" json:"tar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilesList) Reset()         { *m = FilesList{} }
func (m *FilesList) String() string { return proto.CompactTextString(m) }
func (*FilesList) ProtoMessage()    {}
func (*FilesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f6bca5a2823a23, []int{0}
}

func (m *FilesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilesList.Unmarshal(m, b)
}
func (m *FilesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilesList.Marshal(b, m, deterministic)
}
func (m *FilesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilesList.Merge(m, src)
}
func (m *FilesList) XXX_Size() int {
	return xxx_messageInfo_FilesList.Size(m)
}
func (m *FilesList) XXX_DiscardUnknown() {
	xxx_messageInfo_FilesList.DiscardUnknown(m)
}

var xxx_messageInfo_FilesList proto.InternalMessageInfo

func (m *FilesList) GetFileName() []string {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *FilesList) GetTar() string {
	if m != nil {
		return m.Tar
	}
	return ""
}

type ImageUri struct {
	ImageLocation        string   `protobuf:"bytes,1,opt,name=imageLocation,proto3" json:"imageLocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageUri) Reset()         { *m = ImageUri{} }
func (m *ImageUri) String() string { return proto.CompactTextString(m) }
func (*ImageUri) ProtoMessage()    {}
func (*ImageUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f6bca5a2823a23, []int{1}
}

func (m *ImageUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageUri.Unmarshal(m, b)
}
func (m *ImageUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageUri.Marshal(b, m, deterministic)
}
func (m *ImageUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageUri.Merge(m, src)
}
func (m *ImageUri) XXX_Size() int {
	return xxx_messageInfo_ImageUri.Size(m)
}
func (m *ImageUri) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageUri.DiscardUnknown(m)
}

var xxx_messageInfo_ImageUri proto.InternalMessageInfo

func (m *ImageUri) GetImageLocation() string {
	if m != nil {
		return m.ImageLocation
	}
	return ""
}

func init() {
	proto.RegisterType((*FilesList)(nil), "image.FilesList")
	proto.RegisterType((*ImageUri)(nil), "image.ImageUri")
}

func init() { proto.RegisterFile("protobuf/imagesvc.proto", fileDescriptor_d7f6bca5a2823a23) }

var fileDescriptor_d7f6bca5a2823a23 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0x2d, 0x2e, 0x4b, 0xd6, 0x03, 0x8b,
	0x08, 0xb1, 0x82, 0xf9, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99,
	0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x45, 0x4a, 0x96,
	0x5c, 0x9c, 0x6e, 0x99, 0x39, 0xa9, 0xc5, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0x69,
	0x99, 0x39, 0xa9, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x70, 0xbe,
	0x90, 0x00, 0x17, 0x73, 0x49, 0x62, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0xa9,
	0x64, 0xc0, 0xc5, 0xe1, 0x09, 0xb2, 0x21, 0xb4, 0x28, 0x53, 0x48, 0x85, 0x8b, 0x17, 0x6c, 0x9b,
	0x4f, 0x7e, 0x32, 0xd8, 0x78, 0x09, 0x46, 0xb0, 0x3a, 0x54, 0x41, 0xa3, 0x28, 0x2e, 0x6e, 0xb0,
	0x8e, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0x21, 0x6f, 0x2e, 0x2e, 0xb7, 0xd4, 0x92, 0xe4, 0x0c,
	0xb0, 0x98, 0x90, 0x80, 0x1e, 0x58, 0xb1, 0x1e, 0xdc, 0x39, 0x52, 0xfc, 0x50, 0x11, 0x98, 0x2d,
	0x4a, 0x52, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12, 0x51, 0xe2, 0xd7, 0x2f, 0x33, 0x84, 0xfa, 0x16,
	0x6c, 0x92, 0x15, 0xa3, 0x56, 0x12, 0x1b, 0xd8, 0x3f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x87, 0x35, 0x04, 0x8b, 0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImageServerClient is the client API for ImageServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageServerClient interface {
	FetchImage(ctx context.Context, in *FilesList, opts ...grpc.CallOption) (*ImageUri, error)
}

type imageServerClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServerClient(cc grpc.ClientConnInterface) ImageServerClient {
	return &imageServerClient{cc}
}

func (c *imageServerClient) FetchImage(ctx context.Context, in *FilesList, opts ...grpc.CallOption) (*ImageUri, error) {
	out := new(ImageUri)
	err := c.cc.Invoke(ctx, "/image.ImageServer/FetchImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServerServer is the server API for ImageServer service.
type ImageServerServer interface {
	FetchImage(context.Context, *FilesList) (*ImageUri, error)
}

// UnimplementedImageServerServer can be embedded to have forward compatible implementations.
type UnimplementedImageServerServer struct {
}

func (*UnimplementedImageServerServer) FetchImage(ctx context.Context, req *FilesList) (*ImageUri, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchImage not implemented")
}

func RegisterImageServerServer(s *grpc.Server, srv ImageServerServer) {
	s.RegisterService(&_ImageServer_serviceDesc, srv)
}

func _ImageServer_FetchImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServerServer).FetchImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image.ImageServer/FetchImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServerServer).FetchImage(ctx, req.(*FilesList))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImageServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "image.ImageServer",
	HandlerType: (*ImageServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchImage",
			Handler:    _ImageServer_FetchImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/imagesvc.proto",
}

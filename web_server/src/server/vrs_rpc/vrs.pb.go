// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/vrs.proto

/*
Package vrs_rpc is a generated protocol buffer package.

It is generated from these files:
	protos/vrs.proto

It has these top-level messages:
	DetectObjectInImgRequest
	DetectObjectInImgResult
	DetectPersonInImgRequest
	DetectPersonsRequest
	SearchRequest
	SearchResult
	ProcessVideoRequest
	Result
	Empty
	DetectPersonInImgResult
*/
package vrs_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 物体检测请求
type DetectObjectInImgRequest struct {
	Path string `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
}

func (m *DetectObjectInImgRequest) Reset()                    { *m = DetectObjectInImgRequest{} }
func (m *DetectObjectInImgRequest) String() string            { return proto.CompactTextString(m) }
func (*DetectObjectInImgRequest) ProtoMessage()               {}
func (*DetectObjectInImgRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DetectObjectInImgRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// 物体检测结果
type DetectObjectInImgResult struct {
	ClassID   uint32  `protobuf:"varint,1,opt,name=ClassID" json:"ClassID,omitempty"`
	ClassName string  `protobuf:"bytes,2,opt,name=ClassName" json:"ClassName,omitempty"`
	Prob      float32 `protobuf:"fixed32,3,opt,name=Prob" json:"Prob,omitempty"`
	Top       int32   `protobuf:"varint,4,opt,name=Top" json:"Top,omitempty"`
	Left      int32   `protobuf:"varint,5,opt,name=Left" json:"Left,omitempty"`
	Width     int32   `protobuf:"varint,6,opt,name=Width" json:"Width,omitempty"`
	Height    int32   `protobuf:"varint,7,opt,name=Height" json:"Height,omitempty"`
}

func (m *DetectObjectInImgResult) Reset()                    { *m = DetectObjectInImgResult{} }
func (m *DetectObjectInImgResult) String() string            { return proto.CompactTextString(m) }
func (*DetectObjectInImgResult) ProtoMessage()               {}
func (*DetectObjectInImgResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DetectObjectInImgResult) GetClassID() uint32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

func (m *DetectObjectInImgResult) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *DetectObjectInImgResult) GetProb() float32 {
	if m != nil {
		return m.Prob
	}
	return 0
}

func (m *DetectObjectInImgResult) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *DetectObjectInImgResult) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *DetectObjectInImgResult) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *DetectObjectInImgResult) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

// 基于图像人脸的人脸检测请求
type DetectPersonInImgRequest struct {
	Path string `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
}

func (m *DetectPersonInImgRequest) Reset()                    { *m = DetectPersonInImgRequest{} }
func (m *DetectPersonInImgRequest) String() string            { return proto.CompactTextString(m) }
func (*DetectPersonInImgRequest) ProtoMessage()               {}
func (*DetectPersonInImgRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DetectPersonInImgRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// 人脸检测请求
type DetectPersonsRequest struct {
	VideoID uint32 `protobuf:"varint,1,opt,name=VideoID" json:"VideoID,omitempty"`
}

func (m *DetectPersonsRequest) Reset()                    { *m = DetectPersonsRequest{} }
func (m *DetectPersonsRequest) String() string            { return proto.CompactTextString(m) }
func (*DetectPersonsRequest) ProtoMessage()               {}
func (*DetectPersonsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DetectPersonsRequest) GetVideoID() uint32 {
	if m != nil {
		return m.VideoID
	}
	return 0
}

// 搜索请求
type SearchRequest struct {
	ImgPath string `protobuf:"bytes,1,opt,name=ImgPath" json:"ImgPath,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SearchRequest) GetImgPath() string {
	if m != nil {
		return m.ImgPath
	}
	return ""
}

// 搜索结果
type SearchResult struct {
	Score   float32 `protobuf:"fixed32,1,opt,name=Score" json:"Score,omitempty"`
	FrameID uint64  `protobuf:"varint,2,opt,name=FrameID" json:"FrameID,omitempty"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SearchResult) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *SearchResult) GetFrameID() uint64 {
	if m != nil {
		return m.FrameID
	}
	return 0
}

// 视频处理请求
type ProcessVideoRequest struct {
	VideoID uint32 `protobuf:"varint,1,opt,name=VideoID" json:"VideoID,omitempty"`
}

func (m *ProcessVideoRequest) Reset()                    { *m = ProcessVideoRequest{} }
func (m *ProcessVideoRequest) String() string            { return proto.CompactTextString(m) }
func (*ProcessVideoRequest) ProtoMessage()               {}
func (*ProcessVideoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ProcessVideoRequest) GetVideoID() uint32 {
	if m != nil {
		return m.VideoID
	}
	return 0
}

// 普通请求结果
type Result struct {
	Code uint32 `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Result) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 空请求
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 基于图像人脸的人脸检测响应结果
type DetectPersonInImgResult struct {
	PersonID uint32 `protobuf:"varint,1,opt,name=PersonID" json:"PersonID,omitempty"`
	Top      int32  `protobuf:"varint,2,opt,name=Top" json:"Top,omitempty"`
	Right    int32  `protobuf:"varint,3,opt,name=Right" json:"Right,omitempty"`
	Bottom   int32  `protobuf:"varint,4,opt,name=Bottom" json:"Bottom,omitempty"`
	Left     int32  `protobuf:"varint,5,opt,name=Left" json:"Left,omitempty"`
}

func (m *DetectPersonInImgResult) Reset()                    { *m = DetectPersonInImgResult{} }
func (m *DetectPersonInImgResult) String() string            { return proto.CompactTextString(m) }
func (*DetectPersonInImgResult) ProtoMessage()               {}
func (*DetectPersonInImgResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DetectPersonInImgResult) GetPersonID() uint32 {
	if m != nil {
		return m.PersonID
	}
	return 0
}

func (m *DetectPersonInImgResult) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *DetectPersonInImgResult) GetRight() int32 {
	if m != nil {
		return m.Right
	}
	return 0
}

func (m *DetectPersonInImgResult) GetBottom() int32 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *DetectPersonInImgResult) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func init() {
	proto.RegisterType((*DetectObjectInImgRequest)(nil), "vrs.rpc.DetectObjectInImgRequest")
	proto.RegisterType((*DetectObjectInImgResult)(nil), "vrs.rpc.DetectObjectInImgResult")
	proto.RegisterType((*DetectPersonInImgRequest)(nil), "vrs.rpc.DetectPersonInImgRequest")
	proto.RegisterType((*DetectPersonsRequest)(nil), "vrs.rpc.DetectPersonsRequest")
	proto.RegisterType((*SearchRequest)(nil), "vrs.rpc.SearchRequest")
	proto.RegisterType((*SearchResult)(nil), "vrs.rpc.SearchResult")
	proto.RegisterType((*ProcessVideoRequest)(nil), "vrs.rpc.ProcessVideoRequest")
	proto.RegisterType((*Result)(nil), "vrs.rpc.Result")
	proto.RegisterType((*Empty)(nil), "vrs.rpc.Empty")
	proto.RegisterType((*DetectPersonInImgResult)(nil), "vrs.rpc.DetectPersonInImgResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VrsRpc service

type VrsRpcClient interface {
	// 根据路径进行检索，并给出评分
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (VrsRpc_SearchClient, error)
	// 视频处理请求
	ProcessVideo(ctx context.Context, in *ProcessVideoRequest, opts ...grpc.CallOption) (*Result, error)
	// 合并各个视频的索引为全局索引
	Merge(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
}

type vrsRpcClient struct {
	cc *grpc.ClientConn
}

func NewVrsRpcClient(cc *grpc.ClientConn) VrsRpcClient {
	return &vrsRpcClient{cc}
}

func (c *vrsRpcClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (VrsRpc_SearchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_VrsRpc_serviceDesc.Streams[0], c.cc, "/vrs.rpc.VrsRpc/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &vrsRpcSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VrsRpc_SearchClient interface {
	Recv() (*SearchResult, error)
	grpc.ClientStream
}

type vrsRpcSearchClient struct {
	grpc.ClientStream
}

func (x *vrsRpcSearchClient) Recv() (*SearchResult, error) {
	m := new(SearchResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vrsRpcClient) ProcessVideo(ctx context.Context, in *ProcessVideoRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/vrs.rpc.VrsRpc/ProcessVideo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vrsRpcClient) Merge(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/vrs.rpc.VrsRpc/Merge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VrsRpc service

type VrsRpcServer interface {
	// 根据路径进行检索，并给出评分
	Search(*SearchRequest, VrsRpc_SearchServer) error
	// 视频处理请求
	ProcessVideo(context.Context, *ProcessVideoRequest) (*Result, error)
	// 合并各个视频的索引为全局索引
	Merge(context.Context, *Empty) (*Result, error)
}

func RegisterVrsRpcServer(s *grpc.Server, srv VrsRpcServer) {
	s.RegisterService(&_VrsRpc_serviceDesc, srv)
}

func _VrsRpc_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VrsRpcServer).Search(m, &vrsRpcSearchServer{stream})
}

type VrsRpc_SearchServer interface {
	Send(*SearchResult) error
	grpc.ServerStream
}

type vrsRpcSearchServer struct {
	grpc.ServerStream
}

func (x *vrsRpcSearchServer) Send(m *SearchResult) error {
	return x.ServerStream.SendMsg(m)
}

func _VrsRpc_ProcessVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VrsRpcServer).ProcessVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vrs.rpc.VrsRpc/ProcessVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VrsRpcServer).ProcessVideo(ctx, req.(*ProcessVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VrsRpc_Merge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VrsRpcServer).Merge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vrs.rpc.VrsRpc/Merge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VrsRpcServer).Merge(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VrsRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vrs.rpc.VrsRpc",
	HandlerType: (*VrsRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessVideo",
			Handler:    _VrsRpc_ProcessVideo_Handler,
		},
		{
			MethodName: "Merge",
			Handler:    _VrsRpc_Merge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _VrsRpc_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/vrs.proto",
}

// Client API for FaceRpc service

type FaceRpcClient interface {
	// 根据标注数据建立人脸分类器
	BuildClassifer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
	// 在视频中检测人脸数据
	DetectPersonsInVideo(ctx context.Context, in *DetectPersonsRequest, opts ...grpc.CallOption) (*Result, error)
	// 基于图像的人脸检测
	DetectPersonInImg(ctx context.Context, in *DetectPersonInImgRequest, opts ...grpc.CallOption) (FaceRpc_DetectPersonInImgClient, error)
}

type faceRpcClient struct {
	cc *grpc.ClientConn
}

func NewFaceRpcClient(cc *grpc.ClientConn) FaceRpcClient {
	return &faceRpcClient{cc}
}

func (c *faceRpcClient) BuildClassifer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/vrs.rpc.FaceRpc/BuildClassifer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRpcClient) DetectPersonsInVideo(ctx context.Context, in *DetectPersonsRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/vrs.rpc.FaceRpc/DetectPersonsInVideo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRpcClient) DetectPersonInImg(ctx context.Context, in *DetectPersonInImgRequest, opts ...grpc.CallOption) (FaceRpc_DetectPersonInImgClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FaceRpc_serviceDesc.Streams[0], c.cc, "/vrs.rpc.FaceRpc/DetectPersonInImg", opts...)
	if err != nil {
		return nil, err
	}
	x := &faceRpcDetectPersonInImgClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FaceRpc_DetectPersonInImgClient interface {
	Recv() (*DetectPersonInImgResult, error)
	grpc.ClientStream
}

type faceRpcDetectPersonInImgClient struct {
	grpc.ClientStream
}

func (x *faceRpcDetectPersonInImgClient) Recv() (*DetectPersonInImgResult, error) {
	m := new(DetectPersonInImgResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FaceRpc service

type FaceRpcServer interface {
	// 根据标注数据建立人脸分类器
	BuildClassifer(context.Context, *Empty) (*Result, error)
	// 在视频中检测人脸数据
	DetectPersonsInVideo(context.Context, *DetectPersonsRequest) (*Result, error)
	// 基于图像的人脸检测
	DetectPersonInImg(*DetectPersonInImgRequest, FaceRpc_DetectPersonInImgServer) error
}

func RegisterFaceRpcServer(s *grpc.Server, srv FaceRpcServer) {
	s.RegisterService(&_FaceRpc_serviceDesc, srv)
}

func _FaceRpc_BuildClassifer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRpcServer).BuildClassifer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vrs.rpc.FaceRpc/BuildClassifer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRpcServer).BuildClassifer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRpc_DetectPersonsInVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRpcServer).DetectPersonsInVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vrs.rpc.FaceRpc/DetectPersonsInVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRpcServer).DetectPersonsInVideo(ctx, req.(*DetectPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRpc_DetectPersonInImg_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DetectPersonInImgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FaceRpcServer).DetectPersonInImg(m, &faceRpcDetectPersonInImgServer{stream})
}

type FaceRpc_DetectPersonInImgServer interface {
	Send(*DetectPersonInImgResult) error
	grpc.ServerStream
}

type faceRpcDetectPersonInImgServer struct {
	grpc.ServerStream
}

func (x *faceRpcDetectPersonInImgServer) Send(m *DetectPersonInImgResult) error {
	return x.ServerStream.SendMsg(m)
}

var _FaceRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vrs.rpc.FaceRpc",
	HandlerType: (*FaceRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildClassifer",
			Handler:    _FaceRpc_BuildClassifer_Handler,
		},
		{
			MethodName: "DetectPersonsInVideo",
			Handler:    _FaceRpc_DetectPersonsInVideo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DetectPersonInImg",
			Handler:       _FaceRpc_DetectPersonInImg_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/vrs.proto",
}

// Client API for ObjectRpc service

type ObjectRpcClient interface {
	DetectObjectInImg(ctx context.Context, in *DetectObjectInImgRequest, opts ...grpc.CallOption) (ObjectRpc_DetectObjectInImgClient, error)
}

type objectRpcClient struct {
	cc *grpc.ClientConn
}

func NewObjectRpcClient(cc *grpc.ClientConn) ObjectRpcClient {
	return &objectRpcClient{cc}
}

func (c *objectRpcClient) DetectObjectInImg(ctx context.Context, in *DetectObjectInImgRequest, opts ...grpc.CallOption) (ObjectRpc_DetectObjectInImgClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ObjectRpc_serviceDesc.Streams[0], c.cc, "/vrs.rpc.ObjectRpc/DetectObjectInImg", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectRpcDetectObjectInImgClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectRpc_DetectObjectInImgClient interface {
	Recv() (*DetectObjectInImgResult, error)
	grpc.ClientStream
}

type objectRpcDetectObjectInImgClient struct {
	grpc.ClientStream
}

func (x *objectRpcDetectObjectInImgClient) Recv() (*DetectObjectInImgResult, error) {
	m := new(DetectObjectInImgResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ObjectRpc service

type ObjectRpcServer interface {
	DetectObjectInImg(*DetectObjectInImgRequest, ObjectRpc_DetectObjectInImgServer) error
}

func RegisterObjectRpcServer(s *grpc.Server, srv ObjectRpcServer) {
	s.RegisterService(&_ObjectRpc_serviceDesc, srv)
}

func _ObjectRpc_DetectObjectInImg_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DetectObjectInImgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectRpcServer).DetectObjectInImg(m, &objectRpcDetectObjectInImgServer{stream})
}

type ObjectRpc_DetectObjectInImgServer interface {
	Send(*DetectObjectInImgResult) error
	grpc.ServerStream
}

type objectRpcDetectObjectInImgServer struct {
	grpc.ServerStream
}

func (x *objectRpcDetectObjectInImgServer) Send(m *DetectObjectInImgResult) error {
	return x.ServerStream.SendMsg(m)
}

var _ObjectRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vrs.rpc.ObjectRpc",
	HandlerType: (*ObjectRpcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DetectObjectInImg",
			Handler:       _ObjectRpc_DetectObjectInImg_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/vrs.proto",
}

func init() { proto.RegisterFile("protos/vrs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x6c, 0x93, 0x51, 0x5b, 0xca, 0x92, 0x16, 0x13, 0x15, 0x29, 0xf8, 0x14, 0x38,
	0xb8, 0x55, 0x7a, 0x44, 0x02, 0x29, 0x0d, 0x55, 0x2d, 0x51, 0x88, 0x5c, 0x54, 0x2e, 0x5c, 0x1c,
	0x67, 0xea, 0x18, 0xc5, 0xd9, 0xb0, 0xbb, 0xa9, 0xc4, 0x4f, 0xe0, 0xef, 0x70, 0xe3, 0x0f, 0xf1,
	0x3b, 0xd0, 0x8e, 0x3f, 0xe2, 0x62, 0x8b, 0x72, 0x7b, 0x6f, 0xfc, 0x76, 0x66, 0xf7, 0xbd, 0xf5,
	0xc2, 0xfe, 0x5a, 0x70, 0xc5, 0xe5, 0xf1, 0xad, 0x90, 0x1e, 0x41, 0x66, 0x6b, 0x28, 0xd6, 0x91,
	0xeb, 0x81, 0x33, 0x41, 0x85, 0x91, 0xfa, 0x38, 0xfb, 0x8a, 0x91, 0xf2, 0x57, 0x7e, 0x1a, 0x07,
	0xf8, 0x6d, 0x83, 0x52, 0x31, 0x06, 0x9d, 0x69, 0xa8, 0x16, 0x8e, 0x31, 0x30, 0x86, 0xdd, 0x80,
	0xb0, 0xfb, 0xcb, 0x80, 0xa7, 0x0d, 0x0b, 0xe4, 0x66, 0xa9, 0x98, 0x03, 0xf6, 0xd9, 0x32, 0x94,
	0xd2, 0x9f, 0xd0, 0x92, 0xdd, 0xa0, 0xa0, 0xec, 0x08, 0xba, 0x04, 0x3f, 0x84, 0x29, 0x3a, 0x2d,
	0x6a, 0xb7, 0x2d, 0xd0, 0x1c, 0xc1, 0x67, 0x4e, 0x7b, 0x60, 0x0c, 0x5b, 0x01, 0x61, 0xb6, 0x0f,
	0xed, 0x4f, 0x7c, 0xed, 0x74, 0x06, 0xc6, 0xd0, 0x0c, 0x34, 0xd4, 0xaa, 0xf7, 0x78, 0xa3, 0x1c,
	0x93, 0x4a, 0x84, 0x59, 0x0f, 0xcc, 0xcf, 0xc9, 0x5c, 0x2d, 0x1c, 0x8b, 0x8a, 0x19, 0x61, 0x87,
	0x60, 0x5d, 0x60, 0x12, 0x2f, 0x94, 0x63, 0x53, 0x39, 0x67, 0xdb, 0xb3, 0x4e, 0x51, 0x48, 0xbe,
	0xba, 0xf7, 0xac, 0x27, 0xd0, 0xab, 0xea, 0x65, 0xa1, 0x75, 0xc0, 0xbe, 0x4e, 0xe6, 0xc8, 0xb7,
	0xe7, 0xcc, 0xa9, 0xfb, 0x12, 0x76, 0xaf, 0x30, 0x14, 0xd1, 0xa2, 0x22, 0xf5, 0xd3, 0xb8, 0xd2,
	0xb9, 0xa0, 0xee, 0x1b, 0xd8, 0x29, 0xa4, 0x64, 0x5e, 0x0f, 0xcc, 0xab, 0x88, 0x0b, 0x24, 0x5d,
	0x2b, 0xc8, 0x88, 0x5e, 0x7f, 0x2e, 0xc2, 0x14, 0xfd, 0x09, 0xd9, 0xd6, 0x09, 0x0a, 0xea, 0x1e,
	0xc3, 0x93, 0xa9, 0xe0, 0x11, 0x4a, 0x49, 0xc3, 0xef, 0xdf, 0x9b, 0x07, 0x56, 0x3e, 0x8a, 0x41,
	0xe7, 0x8c, 0xcf, 0x31, 0x17, 0x10, 0xd6, 0x7e, 0x5f, 0xca, 0x38, 0xcf, 0x46, 0x43, 0xd7, 0x06,
	0xf3, 0x5d, 0xba, 0x56, 0xdf, 0xdd, 0x1f, 0x65, 0xe4, 0x77, 0x7c, 0xa3, 0x56, 0x7d, 0x78, 0x98,
	0x17, 0x8b, 0x79, 0x25, 0x2f, 0x22, 0x6c, 0x6d, 0x23, 0xec, 0x81, 0x19, 0x50, 0x2e, 0xed, 0x2c,
	0x2e, 0x22, 0x3a, 0xae, 0x31, 0x57, 0x8a, 0xa7, 0x79, 0xda, 0x39, 0x6b, 0x0a, 0x7c, 0xf4, 0xd3,
	0x00, 0xeb, 0x5a, 0xc8, 0x60, 0x1d, 0xb1, 0xd7, 0x60, 0x65, 0x06, 0xb2, 0x43, 0x2f, 0xbf, 0xcd,
	0xde, 0x1d, 0xf3, 0xfb, 0x07, 0xb5, 0xba, 0xde, 0xb3, 0xfb, 0xe0, 0xc4, 0x60, 0x6f, 0x61, 0xa7,
	0xea, 0x1e, 0x3b, 0x2a, 0xa5, 0x0d, 0xa6, 0xf6, 0x1f, 0x95, 0x5f, 0x8b, 0x16, 0xec, 0x15, 0x98,
	0x97, 0x28, 0x62, 0x64, 0x7b, 0xe5, 0x37, 0x72, 0xab, 0x41, 0x3b, 0xfa, 0x6d, 0x80, 0x7d, 0x1e,
	0x46, 0xa8, 0x77, 0x7d, 0x0a, 0x7b, 0xe3, 0x4d, 0xb2, 0x9c, 0xd3, 0xed, 0x4f, 0x6e, 0x50, 0xfc,
	0x47, 0x03, 0x76, 0xf1, 0xd7, 0x45, 0xf4, 0x57, 0xd9, 0xae, 0x9f, 0x97, 0xd2, 0xa6, 0x7b, 0xda,
	0xd4, 0xe9, 0x0b, 0x3c, 0xae, 0x45, 0xc9, 0x5e, 0x34, 0xb6, 0xa9, 0xfe, 0x1e, 0xfd, 0xc1, 0xbf,
	0x24, 0x85, 0xab, 0xa3, 0x04, 0xba, 0xd9, 0xab, 0xa0, 0x4f, 0x5a, 0x8e, 0xaa, 0x3c, 0x14, 0xb5,
	0x51, 0xf5, 0x57, 0xa7, 0x36, 0xaa, 0xf6, 0xce, 0xe8, 0x51, 0xe3, 0x67, 0x70, 0x80, 0xf3, 0x8d,
	0x17, 0x47, 0x9b, 0x42, 0x2e, 0x51, 0xdc, 0xa2, 0x98, 0x1a, 0x33, 0x8b, 0x9e, 0xb8, 0xd3, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x51, 0xf1, 0xa0, 0xa0, 0xf6, 0x04, 0x00, 0x00,
}
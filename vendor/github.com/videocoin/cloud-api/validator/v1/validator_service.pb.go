// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/v1/validator_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/videocoin/cloud-api/emitter/v1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidateProofRequest struct {
	StreamId              string           `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	StreamContractAddress string           `protobuf:"bytes,2,opt,name=stream_contract_address,json=streamContractAddress,proto3" json:"stream_contract_address,omitempty"`
	ProfileId             []byte           `protobuf:"bytes,3,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	ChunkId               []byte           `protobuf:"bytes,4,opt,name=chunk_id,json=chunkId,proto3" json:"chunk_id,omitempty"`
	SubmitProofTx         string           `protobuf:"bytes,5,opt,name=submit_proof_tx,json=submitProofTx,proto3" json:"submit_proof_tx,omitempty"`
	SubmitProofTxStatus   v1.ReceiptStatus `protobuf:"varint,6,opt,name=submit_proof_tx_status,json=submitProofTxStatus,proto3,enum=cloud.api.emitter.v1.ReceiptStatus" json:"submit_proof_tx_status,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *ValidateProofRequest) Reset()         { *m = ValidateProofRequest{} }
func (m *ValidateProofRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateProofRequest) ProtoMessage()    {}
func (*ValidateProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b1a82488864da9d, []int{0}
}
func (m *ValidateProofRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateProofRequest.Unmarshal(m, b)
}
func (m *ValidateProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateProofRequest.Marshal(b, m, deterministic)
}
func (m *ValidateProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateProofRequest.Merge(m, src)
}
func (m *ValidateProofRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateProofRequest.Size(m)
}
func (m *ValidateProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateProofRequest proto.InternalMessageInfo

func (m *ValidateProofRequest) GetStreamId() string {
	if m != nil {
		return m.StreamId
	}
	return ""
}

func (m *ValidateProofRequest) GetStreamContractAddress() string {
	if m != nil {
		return m.StreamContractAddress
	}
	return ""
}

func (m *ValidateProofRequest) GetProfileId() []byte {
	if m != nil {
		return m.ProfileId
	}
	return nil
}

func (m *ValidateProofRequest) GetChunkId() []byte {
	if m != nil {
		return m.ChunkId
	}
	return nil
}

func (m *ValidateProofRequest) GetSubmitProofTx() string {
	if m != nil {
		return m.SubmitProofTx
	}
	return ""
}

func (m *ValidateProofRequest) GetSubmitProofTxStatus() v1.ReceiptStatus {
	if m != nil {
		return m.SubmitProofTxStatus
	}
	return v1.ReceiptStatusUnknown
}

type ValidateProofResponse struct {
	ValidateProofTx       string           `protobuf:"bytes,1,opt,name=validate_proof_tx,json=validateProofTx,proto3" json:"validate_proof_tx,omitempty"`
	ValidateProofTxStatus v1.ReceiptStatus `protobuf:"varint,2,opt,name=validate_proof_tx_status,json=validateProofTxStatus,proto3,enum=cloud.api.emitter.v1.ReceiptStatus" json:"validate_proof_tx_status,omitempty"`
	ScrapProofTx          string           `protobuf:"bytes,3,opt,name=scrap_proof_tx,json=scrapProofTx,proto3" json:"scrap_proof_tx,omitempty"`
	ScrapProofTxStatus    v1.ReceiptStatus `protobuf:"varint,4,opt,name=scrap_proof_tx_status,json=scrapProofTxStatus,proto3,enum=cloud.api.emitter.v1.ReceiptStatus" json:"scrap_proof_tx_status,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *ValidateProofResponse) Reset()         { *m = ValidateProofResponse{} }
func (m *ValidateProofResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateProofResponse) ProtoMessage()    {}
func (*ValidateProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b1a82488864da9d, []int{1}
}
func (m *ValidateProofResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateProofResponse.Unmarshal(m, b)
}
func (m *ValidateProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateProofResponse.Marshal(b, m, deterministic)
}
func (m *ValidateProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateProofResponse.Merge(m, src)
}
func (m *ValidateProofResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateProofResponse.Size(m)
}
func (m *ValidateProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateProofResponse proto.InternalMessageInfo

func (m *ValidateProofResponse) GetValidateProofTx() string {
	if m != nil {
		return m.ValidateProofTx
	}
	return ""
}

func (m *ValidateProofResponse) GetValidateProofTxStatus() v1.ReceiptStatus {
	if m != nil {
		return m.ValidateProofTxStatus
	}
	return v1.ReceiptStatusUnknown
}

func (m *ValidateProofResponse) GetScrapProofTx() string {
	if m != nil {
		return m.ScrapProofTx
	}
	return ""
}

func (m *ValidateProofResponse) GetScrapProofTxStatus() v1.ReceiptStatus {
	if m != nil {
		return m.ScrapProofTxStatus
	}
	return v1.ReceiptStatusUnknown
}

func init() {
	proto.RegisterType((*ValidateProofRequest)(nil), "cloud.api.validator.v1.ValidateProofRequest")
	proto.RegisterType((*ValidateProofResponse)(nil), "cloud.api.validator.v1.ValidateProofResponse")
}

func init() {
	proto.RegisterFile("validator/v1/validator_service.proto", fileDescriptor_5b1a82488864da9d)
}

var fileDescriptor_5b1a82488864da9d = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x33, 0xd3, 0x52, 0x1a, 0xab, 0x7f, 0x18, 0x52, 0x86, 0x00, 0x51, 0x14, 0x2a, 0x14,
	0x21, 0xe2, 0x51, 0x8a, 0x04, 0x6b, 0x60, 0x95, 0x1d, 0x9a, 0x56, 0x11, 0x42, 0x48, 0x23, 0xc7,
	0x76, 0xa6, 0x16, 0xc9, 0xdc, 0xc1, 0xf6, 0x58, 0x59, 0xf3, 0x22, 0x2c, 0x79, 0x15, 0x96, 0x3c,
	0x02, 0xca, 0x93, 0xa0, 0xd8, 0x93, 0xe6, 0x87, 0x2e, 0xb2, 0x9b, 0x7b, 0xce, 0xf1, 0xf5, 0xa7,
	0xeb, 0x3b, 0xe8, 0xc2, 0xd2, 0x89, 0xe4, 0xd4, 0x80, 0x8a, 0x6d, 0x3f, 0xbe, 0x2d, 0x52, 0x2d,
	0x94, 0x95, 0x4c, 0x90, 0x42, 0x81, 0x01, 0x7c, 0xce, 0x26, 0x50, 0x72, 0x42, 0x0b, 0x49, 0x6e,
	0x23, 0xc4, 0xf6, 0x9b, 0xcf, 0x32, 0x80, 0x6c, 0x22, 0x62, 0x5a, 0xc8, 0x98, 0xe6, 0x39, 0x18,
	0x6a, 0x24, 0xe4, 0xda, 0x9f, 0x6a, 0xf6, 0x32, 0x69, 0x6e, 0xca, 0x11, 0x61, 0x30, 0x8d, 0x33,
	0xc8, 0x20, 0x76, 0xf2, 0xa8, 0x1c, 0xbb, 0xca, 0x15, 0xee, 0xab, 0x8a, 0xbf, 0x5b, 0x8b, 0x5b,
	0xc9, 0x05, 0x30, 0x90, 0x79, 0xec, 0x6e, 0xee, 0x2d, 0x2e, 0x10, 0x53, 0x69, 0x8c, 0x70, 0x9c,
	0x4a, 0x30, 0x21, 0x0b, 0xe3, 0x0f, 0x76, 0x7e, 0x85, 0xe8, 0xd1, 0xd0, 0x63, 0x89, 0x4f, 0x0a,
	0x60, 0x9c, 0x88, 0xef, 0xa5, 0xd0, 0x06, 0x3f, 0x45, 0x75, 0x6d, 0x94, 0xa0, 0xd3, 0x54, 0xf2,
	0x28, 0x68, 0x07, 0xdd, 0x7a, 0x72, 0xe8, 0x85, 0x01, 0xc7, 0x6f, 0xd1, 0xe3, 0xca, 0x64, 0x90,
	0x1b, 0x45, 0x99, 0x49, 0x29, 0xe7, 0x4a, 0x68, 0x1d, 0x85, 0x2e, 0xda, 0xf0, 0xf6, 0xc7, 0xca,
	0x7d, 0xef, 0x4d, 0xfc, 0x1c, 0xa1, 0x42, 0xc1, 0x58, 0x4e, 0xc4, 0xa2, 0xeb, 0x5e, 0x3b, 0xe8,
	0x1e, 0x25, 0xf5, 0x4a, 0x19, 0x70, 0xfc, 0x04, 0x1d, 0xb2, 0x9b, 0x32, 0xff, 0xb6, 0x30, 0xf7,
	0x9d, 0x79, 0xdf, 0xd5, 0x03, 0x8e, 0x5f, 0xa2, 0x53, 0x5d, 0x8e, 0xa6, 0xd2, 0xa4, 0xc5, 0x82,
	0x32, 0x35, 0xb3, 0xe8, 0x9e, 0xbb, 0xe9, 0xd8, 0xcb, 0x8e, 0xfd, 0x7a, 0x86, 0x3f, 0xa3, 0xf3,
	0xad, 0x5c, 0xaa, 0x0d, 0x35, 0xa5, 0x8e, 0x0e, 0xda, 0x41, 0xf7, 0xe4, 0xf2, 0x05, 0x59, 0x3d,
	0x47, 0x35, 0x14, 0x62, 0xfb, 0x24, 0xf1, 0x43, 0xb9, 0x72, 0xd1, 0xe4, 0xe1, 0x46, 0x4f, 0x2f,
	0x76, 0x7e, 0x86, 0xa8, 0xb1, 0x35, 0x29, 0x5d, 0x40, 0xae, 0x05, 0x7e, 0x85, 0x1e, 0x54, 0x2f,
	0x2b, 0x56, 0x74, 0x7e, 0x64, 0xa7, 0x76, 0xfd, 0xc4, 0xf5, 0x0c, 0x7f, 0x45, 0xd1, 0x7f, 0xd9,
	0x25, 0x61, 0xb8, 0x3b, 0x61, 0x63, 0xab, 0xaf, 0x97, 0xf1, 0x05, 0x3a, 0xd1, 0x4c, 0xd1, 0x62,
	0x85, 0xb1, 0xe7, 0x30, 0x8e, 0x9c, 0xba, 0x64, 0x18, 0xa2, 0xc6, 0x66, 0x6a, 0x09, 0xb0, 0xbf,
	0x3b, 0x00, 0x5e, 0xef, 0xe8, 0xb5, 0xcb, 0x1f, 0x01, 0x3a, 0x1b, 0x2e, 0x57, 0xfc, 0xca, 0xff,
	0x04, 0x38, 0x47, 0xc7, 0x1b, 0x53, 0xc3, 0xaf, 0xc9, 0xdd, 0x3f, 0x04, 0xb9, 0x6b, 0x0d, 0x9b,
	0xbd, 0x1d, 0xd3, 0xfe, 0x29, 0x3a, 0xb5, 0x0f, 0x67, 0xbf, 0xe7, 0xad, 0xda, 0x9f, 0x79, 0xab,
	0xf6, 0x77, 0xde, 0xaa, 0x7d, 0x09, 0x6d, 0x7f, 0x74, 0xe0, 0x36, 0xfd, 0xcd, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x27, 0x37, 0xc3, 0x92, 0xaf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValidatorServiceClient is the client API for ValidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidatorServiceClient interface {
	ValidateProof(ctx context.Context, in *ValidateProofRequest, opts ...grpc.CallOption) (*ValidateProofResponse, error)
}

type validatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewValidatorServiceClient(cc *grpc.ClientConn) ValidatorServiceClient {
	return &validatorServiceClient{cc}
}

func (c *validatorServiceClient) ValidateProof(ctx context.Context, in *ValidateProofRequest, opts ...grpc.CallOption) (*ValidateProofResponse, error) {
	out := new(ValidateProofResponse)
	err := c.cc.Invoke(ctx, "/cloud.api.validator.v1.ValidatorService/ValidateProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorServiceServer is the server API for ValidatorService service.
type ValidatorServiceServer interface {
	ValidateProof(context.Context, *ValidateProofRequest) (*ValidateProofResponse, error)
}

// UnimplementedValidatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedValidatorServiceServer struct {
}

func (*UnimplementedValidatorServiceServer) ValidateProof(ctx context.Context, req *ValidateProofRequest) (*ValidateProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateProof not implemented")
}

func RegisterValidatorServiceServer(s *grpc.Server, srv ValidatorServiceServer) {
	s.RegisterService(&_ValidatorService_serviceDesc, srv)
}

func _ValidatorService_ValidateProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorServiceServer).ValidateProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.validator.v1.ValidatorService/ValidateProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorServiceServer).ValidateProof(ctx, req.(*ValidateProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ValidatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.validator.v1.ValidatorService",
	HandlerType: (*ValidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateProof",
			Handler:    _ValidatorService_ValidateProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validator/v1/validator_service.proto",
}

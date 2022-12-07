// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/nft_transfer/v1/packet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// NonFungibleTokenPacketData defines a struct for the packet payload
// See NonFungibleTokenPacketData spec:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-721-nft-transfer#data-structures
type NonFungibleTokenPacketData struct {
	// the class_id of tokens to be transferred
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// the class_uri of tokens to be transferred
	ClassUri string `protobuf:"bytes,2,opt,name=class_uri,json=classUri,proto3" json:"class_uri,omitempty"`
	// the non fungible tokens to be transferred
	TokenIds []string `protobuf:"bytes,3,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	// the non fungible tokens's uri to be transferred
	TokenUris []string `protobuf:"bytes,4,rep,name=token_uris,json=tokenUris,proto3" json:"token_uris,omitempty"`
	// the non fungible tokens's data to be transferred
	TokenData [][]byte `protobuf:"bytes,5,rep,name=token_data,json=tokenData,proto3" json:"token_data,omitempty"`
	// the sender address
	Sender string `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,7,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (m *NonFungibleTokenPacketData) Reset()         { *m = NonFungibleTokenPacketData{} }
func (m *NonFungibleTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*NonFungibleTokenPacketData) ProtoMessage()    {}
func (*NonFungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f82fdc932b824013, []int{0}
}
func (m *NonFungibleTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonFungibleTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonFungibleTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonFungibleTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonFungibleTokenPacketData.Merge(m, src)
}
func (m *NonFungibleTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *NonFungibleTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonFungibleTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_NonFungibleTokenPacketData proto.InternalMessageInfo

func (m *NonFungibleTokenPacketData) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetClassUri() string {
	if m != nil {
		return m.ClassUri
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetTokenIds() []string {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

func (m *NonFungibleTokenPacketData) GetTokenUris() []string {
	if m != nil {
		return m.TokenUris
	}
	return nil
}

func (m *NonFungibleTokenPacketData) GetTokenData() [][]byte {
	if m != nil {
		return m.TokenData
	}
	return nil
}

func (m *NonFungibleTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func init() {
	proto.RegisterType((*NonFungibleTokenPacketData)(nil), "ibc.applications.nft_transfer.v1.NonFungibleTokenPacketData")
}

func init() {
	proto.RegisterFile("ibc/applications/nft_transfer/v1/packet.proto", fileDescriptor_f82fdc932b824013)
}

var fileDescriptor_f82fdc932b824013 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x18, 0x85, 0xeb, 0xaf, 0x1f, 0xfd, 0xb1, 0x98, 0x32, 0x20, 0x53, 0x84, 0x15, 0xb1, 0xd0, 0xa5,
	0x89, 0x2a, 0xae, 0x00, 0x84, 0x90, 0xba, 0x20, 0x84, 0xe8, 0xc2, 0x12, 0xf9, 0xaf, 0xe5, 0xa5,
	0xc1, 0x8e, 0x6c, 0x27, 0x12, 0x77, 0xc1, 0x65, 0x31, 0x76, 0x64, 0x44, 0xc9, 0x8d, 0xa0, 0x38,
	0x05, 0x65, 0x3c, 0x7e, 0x1e, 0xbd, 0x3e, 0x3a, 0x78, 0x01, 0x5c, 0xa4, 0xac, 0x28, 0x72, 0x10,
	0xcc, 0x83, 0xd1, 0x2e, 0xd5, 0x1b, 0x9f, 0x79, 0xcb, 0xb4, 0xdb, 0x28, 0x9b, 0x56, 0xcb, 0xb4,
	0x60, 0x62, 0xa7, 0x7c, 0x52, 0x58, 0xe3, 0x4d, 0x14, 0x03, 0x17, 0x49, 0x5f, 0x4f, 0xfa, 0x7a,
	0x52, 0x2d, 0x2f, 0x1a, 0x84, 0x67, 0xf7, 0x46, 0xdf, 0x95, 0x7a, 0x0b, 0x3c, 0x57, 0x4f, 0x66,
	0xa7, 0xf4, 0x43, 0x38, 0x71, 0xcb, 0x3c, 0x8b, 0x4e, 0xf1, 0x44, 0xe4, 0xcc, 0xb9, 0x0c, 0x24,
	0x41, 0x31, 0x9a, 0x4f, 0x1f, 0xc7, 0x21, 0xaf, 0x64, 0x74, 0x86, 0xa7, 0x1d, 0x2a, 0x2d, 0x90,
	0x7f, 0x81, 0x75, 0xee, 0xda, 0x42, 0x0b, 0x7d, 0x7b, 0x2a, 0x03, 0xe9, 0xc8, 0x30, 0x1e, 0xb6,
	0x30, 0x3c, 0xac, 0xa4, 0x8b, 0xce, 0x31, 0xee, 0x60, 0x69, 0xc1, 0x91, 0xff, 0x81, 0x76, 0xfa,
	0xda, 0x42, 0x0f, 0x4b, 0xe6, 0x19, 0x39, 0x8a, 0x87, 0xf3, 0xe3, 0x03, 0x0e, 0x95, 0x4e, 0xf0,
	0xc8, 0x29, 0x2d, 0x95, 0x25, 0xa3, 0xf0, 0xe9, 0x21, 0x45, 0x33, 0x3c, 0xb1, 0x4a, 0x28, 0xa8,
	0x94, 0x25, 0xe3, 0xae, 0xce, 0x6f, 0xbe, 0xb9, 0xfe, 0xac, 0x29, 0xda, 0xd7, 0x14, 0x7d, 0xd7,
	0x14, 0x7d, 0x34, 0x74, 0xb0, 0x6f, 0xe8, 0xe0, 0xab, 0xa1, 0x83, 0xe7, 0xcb, 0x2d, 0xf8, 0x97,
	0x92, 0x27, 0xc2, 0xbc, 0xa5, 0x1c, 0x98, 0x7e, 0x05, 0xc5, 0xa0, 0x1d, 0x75, 0xf1, 0x37, 0xaa,
	0x7f, 0x2f, 0x94, 0xe3, 0xa3, 0xb0, 0xe8, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x6a,
	0x89, 0x75, 0x82, 0x01, 0x00, 0x00,
}

func (m *NonFungibleTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonFungibleTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonFungibleTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenData) > 0 {
		for iNdEx := len(m.TokenData) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenData[iNdEx])
			copy(dAtA[i:], m.TokenData[iNdEx])
			i = encodeVarintPacket(dAtA, i, uint64(len(m.TokenData[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TokenUris) > 0 {
		for iNdEx := len(m.TokenUris) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenUris[iNdEx])
			copy(dAtA[i:], m.TokenUris[iNdEx])
			i = encodeVarintPacket(dAtA, i, uint64(len(m.TokenUris[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TokenIds) > 0 {
		for iNdEx := len(m.TokenIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenIds[iNdEx])
			copy(dAtA[i:], m.TokenIds[iNdEx])
			i = encodeVarintPacket(dAtA, i, uint64(len(m.TokenIds[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClassUri) > 0 {
		i -= len(m.ClassUri)
		copy(dAtA[i:], m.ClassUri)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.ClassUri)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NonFungibleTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.ClassUri)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if len(m.TokenIds) > 0 {
		for _, s := range m.TokenIds {
			l = len(s)
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	if len(m.TokenUris) > 0 {
		for _, s := range m.TokenUris {
			l = len(s)
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	if len(m.TokenData) > 0 {
		for _, b := range m.TokenData {
			l = len(b)
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NonFungibleTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NonFungibleTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonFungibleTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIds = append(m.TokenIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenUris", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenUris = append(m.TokenUris, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenData = append(m.TokenData, make([]byte, postIndex-iNdEx))
			copy(m.TokenData[len(m.TokenData)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)

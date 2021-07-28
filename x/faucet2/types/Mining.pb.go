// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: faucet2/Mining.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Mining struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id       uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Minter   string `protobuf:"bytes,3,opt,name=Minter,json=minter,proto3" json:"Minter,omitempty"`
	LastTime string `protobuf:"bytes,4,opt,name=LastTime,json=lastTime,proto3" json:"LastTime,omitempty"`
	Total    string `protobuf:"bytes,5,opt,name=Total,json=total,proto3" json:"Total,omitempty"`
}

func (m *Mining) Reset()         { *m = Mining{} }
func (m *Mining) String() string { return proto.CompactTextString(m) }
func (*Mining) ProtoMessage()    {}
func (*Mining) Descriptor() ([]byte, []int) {
	return fileDescriptor_9425b12eb6e8692f, []int{0}
}
func (m *Mining) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mining) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mining.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mining) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mining.Merge(m, src)
}
func (m *Mining) XXX_Size() int {
	return m.Size()
}
func (m *Mining) XXX_DiscardUnknown() {
	xxx_messageInfo_Mining.DiscardUnknown(m)
}

var xxx_messageInfo_Mining proto.InternalMessageInfo

func (m *Mining) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Mining) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mining) GetMinter() string {
	if m != nil {
		return m.Minter
	}
	return ""
}

func (m *Mining) GetLastTime() string {
	if m != nil {
		return m.LastTime
	}
	return ""
}

func (m *Mining) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func init() {
	proto.RegisterType((*Mining)(nil), "jack139.artchain.faucet2.Mining")
}

func init() { proto.RegisterFile("faucet2/Mining.proto", fileDescriptor_9425b12eb6e8692f) }

var fileDescriptor_9425b12eb6e8692f = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4b, 0x2c, 0x4d,
	0x4e, 0x2d, 0x31, 0xd2, 0xf7, 0xcd, 0xcc, 0xcb, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0xc8, 0x4a, 0x4c, 0xce, 0x36, 0x34, 0xb6, 0xd4, 0x4b, 0x2c, 0x2a, 0x49, 0xce, 0x48,
	0xcc, 0xcc, 0xd3, 0x83, 0x2a, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1,
	0x20, 0xea, 0x95, 0x6a, 0xb8, 0xd8, 0x20, 0xfa, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13,
	0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x3e, 0x2e, 0xa6,
	0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x31, 0xb0, 0x9e,
	0x92, 0xd4, 0x22, 0x09, 0x66, 0xb0, 0x42, 0xb6, 0x5c, 0x30, 0x4f, 0x48, 0x8a, 0x8b, 0xc3, 0x27,
	0xb1, 0xb8, 0x24, 0x24, 0x33, 0x37, 0x55, 0x82, 0x05, 0x2c, 0xc3, 0x91, 0x03, 0xe5, 0x0b, 0x89,
	0x70, 0xb1, 0x86, 0xe4, 0x97, 0x24, 0xe6, 0x48, 0xb0, 0x82, 0x25, 0x58, 0x4b, 0x40, 0x1c, 0x27,
	0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x7a, 0x49, 0x1f, 0xe6, 0x25, 0xfd, 0x0a, 0x7d,
	0x98, 0xdf, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x7e, 0x31, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x0c, 0xd9, 0x67, 0x13, 0x01, 0x00, 0x00,
}

func (m *Mining) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mining) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mining) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Total) > 0 {
		i -= len(m.Total)
		copy(dAtA[i:], m.Total)
		i = encodeVarintMining(dAtA, i, uint64(len(m.Total)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LastTime) > 0 {
		i -= len(m.LastTime)
		copy(dAtA[i:], m.LastTime)
		i = encodeVarintMining(dAtA, i, uint64(len(m.LastTime)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Minter) > 0 {
		i -= len(m.Minter)
		copy(dAtA[i:], m.Minter)
		i = encodeVarintMining(dAtA, i, uint64(len(m.Minter)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintMining(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMining(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMining(dAtA []byte, offset int, v uint64) int {
	offset -= sovMining(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mining) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMining(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovMining(uint64(m.Id))
	}
	l = len(m.Minter)
	if l > 0 {
		n += 1 + l + sovMining(uint64(l))
	}
	l = len(m.LastTime)
	if l > 0 {
		n += 1 + l + sovMining(uint64(l))
	}
	l = len(m.Total)
	if l > 0 {
		n += 1 + l + sovMining(uint64(l))
	}
	return n
}

func sovMining(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMining(x uint64) (n int) {
	return sovMining(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mining) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMining
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
			return fmt.Errorf("proto: Mining: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mining: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMining
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
				return ErrInvalidLengthMining
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMining
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMining
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
				return ErrInvalidLengthMining
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Minter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMining
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
				return ErrInvalidLengthMining
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMining
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
				return ErrInvalidLengthMining
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMining
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Total = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMining(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMining
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
func skipMining(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMining
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
					return 0, ErrIntOverflowMining
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
					return 0, ErrIntOverflowMining
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
				return 0, ErrInvalidLengthMining
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMining
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMining
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMining        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMining          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMining = fmt.Errorf("proto: unexpected end of group")
)

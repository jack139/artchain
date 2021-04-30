// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inventory/review.proto

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

type Review struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	RecType      string `protobuf:"bytes,3,opt,name=recType,proto3" json:"recType,omitempty"`
	ItemId       string `protobuf:"bytes,4,opt,name=itemId,proto3" json:"itemId,omitempty"`
	ReviewerId   string `protobuf:"bytes,5,opt,name=reviewerId,proto3" json:"reviewerId,omitempty"`
	ReviewDetail string `protobuf:"bytes,6,opt,name=reviewDetail,proto3" json:"reviewDetail,omitempty"`
	ReviewDate   string `protobuf:"bytes,7,opt,name=reviewDate,proto3" json:"reviewDate,omitempty"`
	UpCount      string `protobuf:"bytes,8,opt,name=upCount,proto3" json:"upCount,omitempty"`
	DownCount    string `protobuf:"bytes,9,opt,name=downCount,proto3" json:"downCount,omitempty"`
	Status       string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_405643e905c6084c, []int{0}
}
func (m *Review) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Review.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return m.Size()
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Review) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Review) GetRecType() string {
	if m != nil {
		return m.RecType
	}
	return ""
}

func (m *Review) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *Review) GetReviewerId() string {
	if m != nil {
		return m.ReviewerId
	}
	return ""
}

func (m *Review) GetReviewDetail() string {
	if m != nil {
		return m.ReviewDetail
	}
	return ""
}

func (m *Review) GetReviewDate() string {
	if m != nil {
		return m.ReviewDate
	}
	return ""
}

func (m *Review) GetUpCount() string {
	if m != nil {
		return m.UpCount
	}
	return ""
}

func (m *Review) GetDownCount() string {
	if m != nil {
		return m.DownCount
	}
	return ""
}

func (m *Review) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Review)(nil), "jack139.artchain.inventory.Review")
}

func init() { proto.RegisterFile("inventory/review.proto", fileDescriptor_405643e905c6084c) }

var fileDescriptor_405643e905c6084c = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x1c, 0xc4, 0xeb, 0x7c, 0xfd, 0x52, 0x62, 0x21, 0x06, 0x0b, 0x55, 0x56, 0x85, 0xac, 0xaa, 0x53,
	0x17, 0x62, 0xa1, 0x4e, 0xac, 0x50, 0x09, 0x75, 0x8d, 0x98, 0xd8, 0xdc, 0xc4, 0x4a, 0x0d, 0x34,
	0x8e, 0x9c, 0x7f, 0x5a, 0xf2, 0x16, 0x8c, 0x3c, 0x12, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0x8a,
	0x9d, 0x92, 0xb2, 0xe5, 0x77, 0x77, 0xf9, 0xeb, 0x74, 0xc6, 0x63, 0x95, 0xed, 0x64, 0x06, 0xda,
	0x54, 0xdc, 0xc8, 0x9d, 0x92, 0xfb, 0x30, 0x37, 0x1a, 0x34, 0x99, 0x3c, 0x8b, 0xf8, 0xe5, 0x66,
	0x71, 0x1b, 0x0a, 0x03, 0xf1, 0x46, 0xa8, 0x2c, 0xfc, 0x0d, 0x4e, 0x2e, 0x53, 0x9d, 0x6a, 0x1b,
	0xe3, 0xed, 0x97, 0xfb, 0x63, 0xf6, 0xe1, 0x61, 0x3f, 0xb2, 0x27, 0x08, 0xc5, 0xa3, 0xd8, 0x48,
	0x01, 0xda, 0x50, 0x34, 0x45, 0xf3, 0x20, 0x3a, 0x22, 0xb9, 0xc0, 0x9e, 0x4a, 0xa8, 0x37, 0x45,
	0xf3, 0x61, 0xe4, 0xa9, 0xa4, 0x4d, 0x1a, 0x19, 0x3f, 0x56, 0xb9, 0xa4, 0xff, 0x5c, 0xb2, 0x43,
	0x32, 0xc6, 0xbe, 0x02, 0xb9, 0x5d, 0x25, 0x74, 0x68, 0x8d, 0x8e, 0x08, 0xc3, 0xd8, 0x15, 0x95,
	0x66, 0x95, 0xd0, 0xff, 0xd6, 0x3b, 0x51, 0xc8, 0x0c, 0x9f, 0x3b, 0x5a, 0x4a, 0x10, 0xea, 0x95,
	0xfa, 0x36, 0xf1, 0x47, 0xeb, 0x6f, 0x2c, 0x05, 0x48, 0x3a, 0x3a, 0xbd, 0xd1, 0x2a, 0x6d, 0xab,
	0x32, 0xbf, 0xd7, 0x65, 0x06, 0xf4, 0xcc, 0xb5, 0xea, 0x90, 0x5c, 0xe1, 0x20, 0xd1, 0xfb, 0xcc,
	0x79, 0x81, 0xf5, 0x7a, 0xa1, 0xed, 0x5c, 0x80, 0x80, 0xb2, 0xa0, 0xd8, 0x75, 0x76, 0x74, 0xf7,
	0xf0, 0x59, 0x33, 0x74, 0xa8, 0x19, 0xfa, 0xae, 0x19, 0x7a, 0x6f, 0xd8, 0xe0, 0xd0, 0xb0, 0xc1,
	0x57, 0xc3, 0x06, 0x4f, 0xd7, 0xa9, 0x82, 0x4d, 0xb9, 0x0e, 0x63, 0xbd, 0xe5, 0xdd, 0xe2, 0xfc,
	0xb8, 0x38, 0x7f, 0xe3, 0xfd, 0xe3, 0x40, 0x95, 0xcb, 0x62, 0xed, 0xdb, 0xa9, 0x17, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x98, 0x15, 0xe3, 0xdd, 0xb6, 0x01, 0x00, 0x00,
}

func (m *Review) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Review) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Review) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintReview(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DownCount) > 0 {
		i -= len(m.DownCount)
		copy(dAtA[i:], m.DownCount)
		i = encodeVarintReview(dAtA, i, uint64(len(m.DownCount)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.UpCount) > 0 {
		i -= len(m.UpCount)
		copy(dAtA[i:], m.UpCount)
		i = encodeVarintReview(dAtA, i, uint64(len(m.UpCount)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ReviewDate) > 0 {
		i -= len(m.ReviewDate)
		copy(dAtA[i:], m.ReviewDate)
		i = encodeVarintReview(dAtA, i, uint64(len(m.ReviewDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ReviewDetail) > 0 {
		i -= len(m.ReviewDetail)
		copy(dAtA[i:], m.ReviewDetail)
		i = encodeVarintReview(dAtA, i, uint64(len(m.ReviewDetail)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ReviewerId) > 0 {
		i -= len(m.ReviewerId)
		copy(dAtA[i:], m.ReviewerId)
		i = encodeVarintReview(dAtA, i, uint64(len(m.ReviewerId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ItemId) > 0 {
		i -= len(m.ItemId)
		copy(dAtA[i:], m.ItemId)
		i = encodeVarintReview(dAtA, i, uint64(len(m.ItemId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RecType) > 0 {
		i -= len(m.RecType)
		copy(dAtA[i:], m.RecType)
		i = encodeVarintReview(dAtA, i, uint64(len(m.RecType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintReview(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintReview(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReview(dAtA []byte, offset int, v uint64) int {
	offset -= sovReview(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Review) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovReview(uint64(m.Id))
	}
	l = len(m.RecType)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.ItemId)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.ReviewerId)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.ReviewDetail)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.ReviewDate)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.UpCount)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.DownCount)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovReview(uint64(l))
	}
	return n
}

func sovReview(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReview(x uint64) (n int) {
	return sovReview(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Review) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReview
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
			return fmt.Errorf("proto: Review: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Review: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
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
					return ErrIntOverflowReview
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
				return fmt.Errorf("proto: wrong wireType = %d for field RecType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReviewerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReviewerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReviewDetail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReviewDetail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReviewDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReviewDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpCount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownCount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReview
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
				return ErrInvalidLengthReview
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReview
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReview(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReview
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
func skipReview(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReview
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
					return 0, ErrIntOverflowReview
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
					return 0, ErrIntOverflowReview
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
				return 0, ErrInvalidLengthReview
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReview
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReview
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReview        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReview          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReview = fmt.Errorf("proto: unexpected end of group")
)

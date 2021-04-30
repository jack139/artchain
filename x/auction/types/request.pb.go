// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/request.proto

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

type Request struct {
	Creator        string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id             uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	RecType        string `protobuf:"bytes,3,opt,name=recType,proto3" json:"recType,omitempty"`
	ItemId         string `protobuf:"bytes,4,opt,name=itemId,proto3" json:"itemId,omitempty"`
	AuctionHouseId string `protobuf:"bytes,5,opt,name=auctionHouseId,proto3" json:"auctionHouseId,omitempty"`
	SellerId       string `protobuf:"bytes,6,opt,name=SellerId,proto3" json:"SellerId,omitempty"`
	RequestDate    string `protobuf:"bytes,7,opt,name=requestDate,proto3" json:"requestDate,omitempty"`
	ReservePrice   string `protobuf:"bytes,8,opt,name=reservePrice,proto3" json:"reservePrice,omitempty"`
	Status         string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	OpenDate       string `protobuf:"bytes,10,opt,name=openDate,proto3" json:"openDate,omitempty"`
	CloseDate      string `protobuf:"bytes,11,opt,name=closeDate,proto3" json:"closeDate,omitempty"`
	LastDate       string `protobuf:"bytes,12,opt,name=lastDate,proto3" json:"lastDate,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06bed2bd0b09db9, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Request) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetRecType() string {
	if m != nil {
		return m.RecType
	}
	return ""
}

func (m *Request) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *Request) GetAuctionHouseId() string {
	if m != nil {
		return m.AuctionHouseId
	}
	return ""
}

func (m *Request) GetSellerId() string {
	if m != nil {
		return m.SellerId
	}
	return ""
}

func (m *Request) GetRequestDate() string {
	if m != nil {
		return m.RequestDate
	}
	return ""
}

func (m *Request) GetReservePrice() string {
	if m != nil {
		return m.ReservePrice
	}
	return ""
}

func (m *Request) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Request) GetOpenDate() string {
	if m != nil {
		return m.OpenDate
	}
	return ""
}

func (m *Request) GetCloseDate() string {
	if m != nil {
		return m.CloseDate
	}
	return ""
}

func (m *Request) GetLastDate() string {
	if m != nil {
		return m.LastDate
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "jack139.artchain.auction.Request")
}

func init() { proto.RegisterFile("auction/request.proto", fileDescriptor_c06bed2bd0b09db9) }

var fileDescriptor_c06bed2bd0b09db9 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xcb, 0x4a, 0x03, 0x31,
	0x14, 0xed, 0x8c, 0xb5, 0x8f, 0xb4, 0x74, 0x11, 0x54, 0x42, 0x91, 0x50, 0xba, 0x90, 0x82, 0x30,
	0x41, 0xba, 0x72, 0x2b, 0x0a, 0x76, 0x27, 0xd5, 0x95, 0xbb, 0x34, 0x73, 0x69, 0xa3, 0xd3, 0x66,
	0x4c, 0x32, 0x62, 0xff, 0xc2, 0xcf, 0x72, 0xd9, 0xa5, 0x1b, 0x41, 0xda, 0x1f, 0x91, 0x49, 0x32,
	0xbe, 0x76, 0x39, 0x8f, 0x7b, 0xef, 0x09, 0x07, 0x1d, 0xf2, 0x42, 0x58, 0xa9, 0x56, 0x4c, 0xc3,
	0x53, 0x01, 0xc6, 0x26, 0xb9, 0x56, 0x56, 0x61, 0xf2, 0xc0, 0xc5, 0xe3, 0xd9, 0xf8, 0x3c, 0xe1,
	0xda, 0x8a, 0x05, 0x97, 0xab, 0x24, 0xf8, 0xfa, 0x07, 0x73, 0x35, 0x57, 0xce, 0xc4, 0xca, 0x97,
	0xf7, 0x0f, 0x3f, 0x62, 0xd4, 0x9c, 0xfa, 0x0d, 0x98, 0xa0, 0xa6, 0xd0, 0xc0, 0xad, 0xd2, 0x24,
	0x1a, 0x44, 0xa3, 0xf6, 0xb4, 0x82, 0xb8, 0x87, 0x62, 0x99, 0x92, 0x78, 0x10, 0x8d, 0xea, 0xd3,
	0x58, 0xa6, 0xa5, 0x53, 0x83, 0xb8, 0x5b, 0xe7, 0x40, 0xf6, 0xbc, 0x33, 0x40, 0x7c, 0x84, 0x1a,
	0xd2, 0xc2, 0x72, 0x92, 0x92, 0xba, 0x13, 0x02, 0xc2, 0x27, 0xa8, 0x17, 0x82, 0x5c, 0xab, 0xc2,
	0xc0, 0x24, 0x25, 0xfb, 0x4e, 0xff, 0xc7, 0xe2, 0x3e, 0x6a, 0xdd, 0x42, 0x96, 0x81, 0x9e, 0xa4,
	0xa4, 0xe1, 0x1c, 0xdf, 0x18, 0x0f, 0x50, 0x27, 0x7c, 0xf6, 0x92, 0x5b, 0x20, 0x4d, 0x27, 0xff,
	0xa6, 0xf0, 0x10, 0x75, 0x35, 0x18, 0xd0, 0xcf, 0x70, 0xa3, 0xa5, 0x00, 0xd2, 0x72, 0x96, 0x3f,
	0x5c, 0x99, 0xd0, 0x58, 0x6e, 0x0b, 0x43, 0xda, 0x3e, 0xa1, 0x47, 0xe5, 0x65, 0x95, 0xc3, 0xca,
	0xad, 0x46, 0xfe, 0x72, 0x85, 0xf1, 0x31, 0x6a, 0x8b, 0x4c, 0x19, 0x70, 0x62, 0xc7, 0x89, 0x3f,
	0x44, 0x39, 0x99, 0xf1, 0x10, 0xaa, 0xeb, 0x27, 0x2b, 0x7c, 0x71, 0xf5, 0xb6, 0xa5, 0xd1, 0x66,
	0x4b, 0xa3, 0xcf, 0x2d, 0x8d, 0x5e, 0x77, 0xb4, 0xb6, 0xd9, 0xd1, 0xda, 0xfb, 0x8e, 0xd6, 0xee,
	0x4f, 0xe7, 0xd2, 0x2e, 0x8a, 0x59, 0x22, 0xd4, 0x92, 0x85, 0xd2, 0x58, 0x55, 0x1a, 0x7b, 0x61,
	0x55, 0xbd, 0x76, 0x9d, 0x83, 0x99, 0x35, 0x5c, 0x5b, 0xe3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x7e, 0x23, 0x2b, 0xf6, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastDate) > 0 {
		i -= len(m.LastDate)
		copy(dAtA[i:], m.LastDate)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.LastDate)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.CloseDate) > 0 {
		i -= len(m.CloseDate)
		copy(dAtA[i:], m.CloseDate)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.CloseDate)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.OpenDate) > 0 {
		i -= len(m.OpenDate)
		copy(dAtA[i:], m.OpenDate)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.OpenDate)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ReservePrice) > 0 {
		i -= len(m.ReservePrice)
		copy(dAtA[i:], m.ReservePrice)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ReservePrice)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.RequestDate) > 0 {
		i -= len(m.RequestDate)
		copy(dAtA[i:], m.RequestDate)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.RequestDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SellerId) > 0 {
		i -= len(m.SellerId)
		copy(dAtA[i:], m.SellerId)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.SellerId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AuctionHouseId) > 0 {
		i -= len(m.AuctionHouseId)
		copy(dAtA[i:], m.AuctionHouseId)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.AuctionHouseId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ItemId) > 0 {
		i -= len(m.ItemId)
		copy(dAtA[i:], m.ItemId)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ItemId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RecType) > 0 {
		i -= len(m.RecType)
		copy(dAtA[i:], m.RecType)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.RecType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovRequest(uint64(m.Id))
	}
	l = len(m.RecType)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.ItemId)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.AuctionHouseId)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.SellerId)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.RequestDate)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.ReservePrice)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.OpenDate)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.CloseDate)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.LastDate)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
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
					return ErrIntOverflowRequest
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
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
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
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionHouseId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionHouseId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservePrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReservePrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpenDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OpenDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloseDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CloseDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)

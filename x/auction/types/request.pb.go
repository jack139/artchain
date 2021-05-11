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
	SellerId       string `protobuf:"bytes,6,opt,name=SellerId,json=sellerId,proto3" json:"SellerId,omitempty"`
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
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xcb, 0x4a, 0x33, 0x31,
	0x14, 0xee, 0xcc, 0xdf, 0xbf, 0x97, 0xb4, 0x74, 0x11, 0x54, 0x42, 0x91, 0x50, 0xba, 0x90, 0x82,
	0x30, 0x41, 0xba, 0x72, 0x2b, 0x0a, 0x76, 0x27, 0xd5, 0x95, 0xbb, 0x34, 0x73, 0x68, 0xa3, 0xd3,
	0x66, 0x4c, 0x32, 0x62, 0xdf, 0xc2, 0xc7, 0x72, 0xd9, 0xa5, 0x1b, 0x41, 0x3a, 0x2f, 0x22, 0x93,
	0x64, 0xf0, 0xb2, 0xcb, 0x77, 0x39, 0xe7, 0x7c, 0xe1, 0x43, 0x87, 0xbc, 0x10, 0x56, 0xaa, 0x0d,
	0xd3, 0xf0, 0x54, 0x80, 0xb1, 0x49, 0xae, 0x95, 0x55, 0x98, 0x3c, 0x70, 0xf1, 0x78, 0x36, 0x3d,
	0x4f, 0xb8, 0xb6, 0x62, 0xc5, 0xe5, 0x26, 0x09, 0xbe, 0xe1, 0xc1, 0x52, 0x2d, 0x95, 0x33, 0xb1,
	0xea, 0xe5, 0xfd, 0xe3, 0x8f, 0x18, 0xb5, 0xe7, 0x7e, 0x03, 0x26, 0xa8, 0x2d, 0x34, 0x70, 0xab,
	0x34, 0x89, 0x46, 0xd1, 0xa4, 0x3b, 0xaf, 0x21, 0x1e, 0xa0, 0x58, 0xa6, 0x24, 0x1e, 0x45, 0x93,
	0xe6, 0x3c, 0x96, 0x69, 0xe5, 0xd4, 0x20, 0xee, 0xb6, 0x39, 0x90, 0x7f, 0xde, 0x19, 0x20, 0x3e,
	0x42, 0x2d, 0x69, 0x61, 0x3d, 0x4b, 0x49, 0xd3, 0x09, 0x01, 0xe1, 0x13, 0x34, 0x08, 0x41, 0xae,
	0x55, 0x61, 0x60, 0x96, 0x92, 0xff, 0x4e, 0xff, 0xc3, 0xe2, 0x21, 0xea, 0xdc, 0x42, 0x96, 0x81,
	0x9e, 0xa5, 0xa4, 0xe5, 0x1c, 0x1d, 0x13, 0x30, 0x1e, 0xa1, 0x5e, 0xf8, 0xec, 0x25, 0xb7, 0x40,
	0xda, 0x4e, 0xfe, 0x49, 0xe1, 0x31, 0xea, 0x6b, 0x30, 0xa0, 0x9f, 0xe1, 0x46, 0x4b, 0x01, 0xa4,
	0xe3, 0x2c, 0xbf, 0xb8, 0x2a, 0xa1, 0xb1, 0xdc, 0x16, 0x86, 0x74, 0x7d, 0x42, 0x8f, 0xaa, 0xcb,
	0x2a, 0x87, 0x8d, 0x5b, 0x8d, 0xfc, 0xe5, 0x1a, 0xe3, 0x63, 0xd4, 0x15, 0x99, 0x32, 0xe0, 0xc4,
	0x9e, 0x13, 0xbf, 0x89, 0x6a, 0x32, 0xe3, 0x21, 0x54, 0xdf, 0x4f, 0xd6, 0xf8, 0xe2, 0xea, 0x6d,
	0x4f, 0xa3, 0xdd, 0x9e, 0x46, 0x9f, 0x7b, 0x1a, 0xbd, 0x96, 0xb4, 0xb1, 0x2b, 0x69, 0xe3, 0xbd,
	0xa4, 0x8d, 0xfb, 0xd3, 0xa5, 0xb4, 0xab, 0x62, 0x91, 0x08, 0xb5, 0x66, 0xa1, 0x34, 0x56, 0x97,
	0xc6, 0x5e, 0x58, 0x5d, 0xaf, 0xdd, 0xe6, 0x60, 0x16, 0x2d, 0xd7, 0xd6, 0xf4, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0x8c, 0x3f, 0x35, 0xf6, 0x01, 0x00, 0x00,
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

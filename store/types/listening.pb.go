// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/store/v42beta1/listening.proto

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

// StoreKVPair is a KVStore KVPair used for listening to state changes (Sets and Deletes)
// It optionally includes the StoreKey for the originating KVStore and a Boolean flag to distinguish between Sets and
// Deletes
type StoreKVPair struct {
	StoreKey string `protobuf:"bytes,1,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"`
	Delete   bool   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
	Key      []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StoreKVPair) Reset()         { *m = StoreKVPair{} }
func (m *StoreKVPair) String() string { return proto.CompactTextString(m) }
func (*StoreKVPair) ProtoMessage()    {}
func (*StoreKVPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5d350879fe4fecd, []int{0}
}
func (m *StoreKVPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreKVPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreKVPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreKVPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreKVPair.Merge(m, src)
}
func (m *StoreKVPair) XXX_Size() int {
	return m.Size()
}
func (m *StoreKVPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreKVPair.DiscardUnknown(m)
}

var xxx_messageInfo_StoreKVPair proto.InternalMessageInfo

func (m *StoreKVPair) GetStoreKey() string {
	if m != nil {
		return m.StoreKey
	}
	return ""
}

func (m *StoreKVPair) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (m *StoreKVPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StoreKVPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreKVPair)(nil), "cosmos.base.store.v42beta1.StoreKVPair")
}

func init() {
	proto.RegisterFile("cosmos/base/store/v42beta1/listening.proto", fileDescriptor_a5d350879fe4fecd)
}

var fileDescriptor_a5d350879fe4fecd = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2f, 0x33,
	0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0xcb, 0xcc, 0x4b, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x84, 0x28, 0xd5, 0x03, 0x29, 0xd5, 0x03, 0x2b, 0xd5,
	0x83, 0x2a, 0x55, 0xca, 0xe2, 0xe2, 0x0e, 0x06, 0x09, 0x78, 0x87, 0x05, 0x24, 0x66, 0x16, 0x09,
	0x49, 0x73, 0x71, 0x82, 0xe5, 0xe3, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x38, 0xc0, 0x02, 0xde, 0xa9, 0x95, 0x42, 0x62, 0x5c, 0x6c, 0x29, 0xa9, 0x39, 0xa9, 0x25, 0xa9,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x50, 0x9e, 0x90, 0x00, 0x17, 0x33, 0x48, 0x39, 0xb3,
	0x02, 0xa3, 0x06, 0x4f, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a,
	0xc1, 0x02, 0x16, 0x83, 0x70, 0x9c, 0x9c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x4a, 0x23, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xea, 0x2d,
	0x08, 0xa5, 0x5b, 0x9c, 0x92, 0x0d, 0xf5, 0x5c, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8,
	0x47, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xe0, 0xb3, 0x51, 0xfe, 0x00, 0x00, 0x00,
}

func (m *StoreKVPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreKVPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreKVPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintListening(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintListening(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Delete {
		i--
		if m.Delete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.StoreKey) > 0 {
		i -= len(m.StoreKey)
		copy(dAtA[i:], m.StoreKey)
		i = encodeVarintListening(dAtA, i, uint64(len(m.StoreKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintListening(dAtA []byte, offset int, v uint64) int {
	offset -= sovListening(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreKVPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StoreKey)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	if m.Delete {
		n += 2
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovListening(uint64(l))
	}
	return n
}

func sovListening(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozListening(x uint64) (n int) {
	return sovListening(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoreKVPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListening
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
			return fmt.Errorf("proto: StoreKVPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreKVPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Delete = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListening
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
				return ErrInvalidLengthListening
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthListening
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListening(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthListening
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
func skipListening(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListening
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
					return 0, ErrIntOverflowListening
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
					return 0, ErrIntOverflowListening
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
				return 0, ErrInvalidLengthListening
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupListening
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthListening
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthListening        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListening          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupListening = fmt.Errorf("proto: unexpected end of group")
)

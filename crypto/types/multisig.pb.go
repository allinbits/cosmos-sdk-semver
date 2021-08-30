// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crypto/multisig/v43beta1/multisig.proto

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

// MultiSignature wraps the signatures from a multisig.LegacyAminoPubKey.
// See cosmos.tx.v1betata1.ModeInfo.Multi for how to specify which signers
// signed and with which modes.
type MultiSignature struct {
	Signatures       [][]byte `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MultiSignature) Reset()         { *m = MultiSignature{} }
func (m *MultiSignature) String() string { return proto.CompactTextString(m) }
func (*MultiSignature) ProtoMessage()    {}
func (*MultiSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_1177bdf7025769be, []int{0}
}
func (m *MultiSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiSignature.Merge(m, src)
}
func (m *MultiSignature) XXX_Size() int {
	return m.Size()
}
func (m *MultiSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiSignature.DiscardUnknown(m)
}

var xxx_messageInfo_MultiSignature proto.InternalMessageInfo

func (m *MultiSignature) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// CompactBitArray is an implementation of a space efficient bit array.
// This is used to ensure that the encoded data takes up a minimal amount of
// space after proto encoding.
// This is not thread safe, and is not intended for concurrent usage.
type CompactBitArray struct {
	ExtraBitsStored uint32 `protobuf:"varint,1,opt,name=extra_bits_stored,json=extraBitsStored,proto3" json:"extra_bits_stored,omitempty"`
	Elems           []byte `protobuf:"bytes,2,opt,name=elems,proto3" json:"elems,omitempty"`
}

func (m *CompactBitArray) Reset()      { *m = CompactBitArray{} }
func (*CompactBitArray) ProtoMessage() {}
func (*CompactBitArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_1177bdf7025769be, []int{1}
}
func (m *CompactBitArray) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompactBitArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompactBitArray.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompactBitArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactBitArray.Merge(m, src)
}
func (m *CompactBitArray) XXX_Size() int {
	return m.Size()
}
func (m *CompactBitArray) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactBitArray.DiscardUnknown(m)
}

var xxx_messageInfo_CompactBitArray proto.InternalMessageInfo

func (m *CompactBitArray) GetExtraBitsStored() uint32 {
	if m != nil {
		return m.ExtraBitsStored
	}
	return 0
}

func (m *CompactBitArray) GetElems() []byte {
	if m != nil {
		return m.Elems
	}
	return nil
}

func init() {
	proto.RegisterType((*MultiSignature)(nil), "cosmos.crypto.multisig.v43beta1.MultiSignature")
	proto.RegisterType((*CompactBitArray)(nil), "cosmos.crypto.multisig.v43beta1.CompactBitArray")
}

func init() {
	proto.RegisterFile("cosmos/crypto/multisig/v43beta1/multisig.proto", fileDescriptor_1177bdf7025769be)
}

var fileDescriptor_1177bdf7025769be = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0x83, 0x50,
	0x14, 0x85, 0x79, 0x5a, 0x1d, 0x5e, 0xaa, 0x8d, 0xa4, 0x03, 0x71, 0x78, 0x25, 0x9d, 0xd0, 0xa4,
	0x90, 0xc6, 0xc4, 0xa1, 0x9b, 0x74, 0x76, 0xa1, 0x93, 0x2e, 0x0d, 0xd0, 0x17, 0x7c, 0xb1, 0x78,
	0xc9, 0xbb, 0x17, 0x23, 0xff, 0xc2, 0xd1, 0x51, 0xff, 0x8d, 0x23, 0xa3, 0xa3, 0x81, 0x3f, 0x62,
	0xfa, 0x90, 0xa6, 0xd3, 0xbd, 0xe7, 0x9c, 0xef, 0x0e, 0xf7, 0xf0, 0x59, 0x0a, 0x98, 0x03, 0x06,
	0xa9, 0xae, 0x0a, 0x82, 0x20, 0x2f, 0xb7, 0xa4, 0x50, 0x65, 0xc1, 0xeb, 0x3c, 0x91, 0x14, 0xcf,
	0xf7, 0x86, 0x5f, 0x68, 0x20, 0xb0, 0x45, 0x87, 0xfb, 0x1d, 0xee, 0xef, 0xd3, 0x7f, 0xfc, 0x72,
	0x9c, 0x41, 0x06, 0x06, 0x0d, 0x76, 0x5b, 0x77, 0x35, 0xbd, 0xe5, 0xe7, 0xf7, 0x3b, 0x72, 0xa5,
	0xb2, 0x97, 0x98, 0x4a, 0x2d, 0x6d, 0xc1, 0x39, 0xf6, 0x02, 0x1d, 0xe6, 0x1e, 0x7b, 0xc3, 0xe8,
	0xc0, 0x59, 0x0c, 0xea, 0xaf, 0x09, 0x9b, 0x3e, 0xf0, 0xd1, 0x12, 0xf2, 0x22, 0x4e, 0x29, 0x54,
	0x74, 0xa7, 0x75, 0x5c, 0xd9, 0xd7, 0xfc, 0x42, 0xbe, 0x91, 0x8e, 0xd7, 0x89, 0x22, 0x5c, 0x23,
	0x81, 0x96, 0x1b, 0x87, 0xb9, 0xcc, 0x3b, 0x8b, 0x46, 0x26, 0x08, 0x15, 0xe1, 0xca, 0xd8, 0xf6,
	0x98, 0x9f, 0xc8, 0xad, 0xcc, 0xd1, 0x39, 0x72, 0x99, 0x37, 0x8c, 0x3a, 0xb1, 0x18, 0x7c, 0x7c,
	0x4e, 0xac, 0x70, 0xf9, 0xdd, 0x08, 0x56, 0x37, 0x82, 0xfd, 0x36, 0x82, 0xbd, 0xb7, 0xc2, 0xaa,
	0x5b, 0x61, 0xfd, 0xb4, 0xc2, 0x7a, 0xbc, 0xca, 0x14, 0x3d, 0x95, 0x89, 0x9f, 0x42, 0x1e, 0xf4,
	0xe5, 0x98, 0x31, 0xc3, 0xcd, 0x73, 0xdf, 0x13, 0x55, 0x85, 0xc4, 0xe4, 0xd4, 0xbc, 0x77, 0xf3,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xba, 0x52, 0xb5, 0x1f, 0x45, 0x01, 0x00, 0x00,
}

func (m *MultiSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signatures[iNdEx])
			copy(dAtA[i:], m.Signatures[iNdEx])
			i = encodeVarintMultisig(dAtA, i, uint64(len(m.Signatures[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CompactBitArray) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompactBitArray) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompactBitArray) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Elems) > 0 {
		i -= len(m.Elems)
		copy(dAtA[i:], m.Elems)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.Elems)))
		i--
		dAtA[i] = 0x12
	}
	if m.ExtraBitsStored != 0 {
		i = encodeVarintMultisig(dAtA, i, uint64(m.ExtraBitsStored))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMultisig(dAtA []byte, offset int, v uint64) int {
	offset -= sovMultisig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MultiSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, b := range m.Signatures {
			l = len(b)
			n += 1 + l + sovMultisig(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CompactBitArray) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExtraBitsStored != 0 {
		n += 1 + sovMultisig(uint64(m.ExtraBitsStored))
	}
	l = len(m.Elems)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	return n
}

func sovMultisig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMultisig(x uint64) (n int) {
	return sovMultisig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: MultiSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, make([]byte, postIndex-iNdEx))
			copy(m.Signatures[len(m.Signatures)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompactBitArray) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: CompactBitArray: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactBitArray: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraBitsStored", wireType)
			}
			m.ExtraBitsStored = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtraBitsStored |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems[:0], dAtA[iNdEx:postIndex]...)
			if m.Elems == nil {
				m.Elems = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func skipMultisig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMultisig
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
					return 0, ErrIntOverflowMultisig
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
					return 0, ErrIntOverflowMultisig
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
				return 0, ErrInvalidLengthMultisig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMultisig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMultisig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMultisig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMultisig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMultisig = fmt.Errorf("proto: unexpected end of group")
)

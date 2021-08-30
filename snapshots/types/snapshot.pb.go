// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/snapshots/v42beta1/snapshot.proto

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

// Snapshot contains Tendermint state sync snapshot info.
type Snapshot struct {
	Height   uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Format   uint32   `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
	Chunks   uint32   `protobuf:"varint,3,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Hash     []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Metadata Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{0}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Snapshot) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

func (m *Snapshot) GetChunks() uint32 {
	if m != nil {
		return m.Chunks
	}
	return 0
}

func (m *Snapshot) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Snapshot) GetMetadata() Metadata {
	if m != nil {
		return m.Metadata
	}
	return Metadata{}
}

// Metadata contains SDK-specific snapshot metadata.
type Metadata struct {
	ChunkHashes [][]byte `protobuf:"bytes,1,rep,name=chunk_hashes,json=chunkHashes,proto3" json:"chunk_hashes,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetChunkHashes() [][]byte {
	if m != nil {
		return m.ChunkHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*Snapshot)(nil), "cosmos.base.snapshots.v42beta1.Snapshot")
	proto.RegisterType((*Metadata)(nil), "cosmos.base.snapshots.v42beta1.Metadata")
}

func init() {
	proto.RegisterFile("cosmos/base/snapshots/v42beta1/snapshot.proto", fileDescriptor_dd7a3c9b0a19e1ee)
}

var fileDescriptor_dd7a3c9b0a19e1ee = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x1a, 0xaa, 0xc8, 0x0d, 0x8b, 0x85, 0x90, 0x85, 0x84, 0x09, 0x5d, 0xc8, 0xd0,
	0x3a, 0x2a, 0xdc, 0xa0, 0x03, 0x82, 0x81, 0x25, 0x6c, 0x2c, 0xc8, 0x49, 0x4d, 0x5c, 0x55, 0xa9,
	0xa3, 0x3e, 0x17, 0x89, 0x5b, 0x70, 0x15, 0x6e, 0xd1, 0xb1, 0x23, 0x13, 0x42, 0xc9, 0x45, 0x50,
	0x1c, 0x13, 0x31, 0x75, 0xca, 0xfb, 0xbf, 0x7c, 0x4f, 0xcf, 0xfa, 0xf1, 0x24, 0xd7, 0x50, 0x6a,
	0x48, 0x32, 0x01, 0x32, 0x81, 0xb5, 0xa8, 0x40, 0x69, 0x03, 0xc9, 0xdb, 0x2c, 0x93, 0x46, 0xcc,
	0x7a, 0xc2, 0xab, 0x8d, 0x36, 0x9a, 0x5c, 0x74, 0x36, 0x6f, 0x6d, 0xde, 0xdb, 0xdc, 0xd9, 0xe7,
	0xa7, 0x85, 0x2e, 0xb4, 0x35, 0x93, 0x76, 0xea, 0x96, 0xc6, 0x9f, 0x08, 0x07, 0x4f, 0xce, 0x25,
	0x67, 0x78, 0xa8, 0xe4, 0xb2, 0x50, 0x86, 0xa2, 0x08, 0xc5, 0x7e, 0xea, 0x52, 0xcb, 0x5f, 0xf5,
	0xa6, 0x14, 0x86, 0x1e, 0x45, 0x28, 0x3e, 0x49, 0x5d, 0x6a, 0x79, 0xae, 0xb6, 0xeb, 0x15, 0xd0,
	0x41, 0xc7, 0xbb, 0x44, 0x08, 0xf6, 0x95, 0x00, 0x45, 0xfd, 0x08, 0xc5, 0x61, 0x6a, 0x67, 0xf2,
	0x80, 0x83, 0x52, 0x1a, 0xb1, 0x10, 0x46, 0xd0, 0xe3, 0x08, 0xc5, 0xa3, 0x9b, 0x6b, 0x7e, 0xf0,
	0xc1, 0xfc, 0xd1, 0xe9, 0x73, 0x7f, 0xf7, 0x7d, 0xe9, 0xa5, 0xfd, 0xfa, 0x78, 0x8a, 0x83, 0xbf,
	0x7f, 0xe4, 0x0a, 0x87, 0xf6, 0xe8, 0x4b, 0x7b, 0x44, 0x02, 0x45, 0xd1, 0x20, 0x0e, 0xd3, 0x91,
	0x65, 0xf7, 0x16, 0xcd, 0xef, 0x76, 0x35, 0x43, 0xfb, 0x9a, 0xa1, 0x9f, 0x9a, 0xa1, 0x8f, 0x86,
	0x79, 0xfb, 0x86, 0x79, 0x5f, 0x0d, 0xf3, 0x9e, 0x27, 0xc5, 0xd2, 0xa8, 0x6d, 0xc6, 0x73, 0x5d,
	0x26, 0xae, 0xea, 0xee, 0x33, 0x85, 0xc5, 0xea, 0x5f, 0xe1, 0xe6, 0xbd, 0x92, 0x90, 0x0d, 0x6d,
	0x63, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x89, 0x2e, 0x8f, 0x96, 0x01, 0x00, 0x00,
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSnapshot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Chunks != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Chunks))
		i--
		dAtA[i] = 0x18
	}
	if m.Format != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChunkHashes) > 0 {
		for iNdEx := len(m.ChunkHashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChunkHashes[iNdEx])
			copy(dAtA[i:], m.ChunkHashes[iNdEx])
			i = encodeVarintSnapshot(dAtA, i, uint64(len(m.ChunkHashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovSnapshot(uint64(m.Height))
	}
	if m.Format != 0 {
		n += 1 + sovSnapshot(uint64(m.Format))
	}
	if m.Chunks != 0 {
		n += 1 + sovSnapshot(uint64(m.Chunks))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChunkHashes) > 0 {
		for _, b := range m.ChunkHashes {
			l = len(b)
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			m.Chunks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chunks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChunkHashes = append(m.ChunkHashes, make([]byte, postIndex-iNdEx))
			copy(m.ChunkHashes[len(m.ChunkHashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSnapshot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSnapshot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSnapshot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnapshot = fmt.Errorf("proto: unexpected end of group")
)

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/stats/histogram.proto

package stats

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/cockroachdb/cockroach/pkg/sql/types"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// HistogramData encodes the data for a histogram, which captures the
// distribution of values on a specific column.
type HistogramData struct {
	// Value type for the column.
	ColumnType *types.T `protobuf:"bytes,2,opt,name=column_type,json=columnType,proto3" json:"column_type,omitempty"`
	// Histogram buckets. Note that NULL values are excluded from the
	// histogram.
	Buckets []HistogramData_Bucket `protobuf:"bytes,1,rep,name=buckets,proto3" json:"buckets"`
}

func (m *HistogramData) Reset()         { *m = HistogramData{} }
func (m *HistogramData) String() string { return proto.CompactTextString(m) }
func (*HistogramData) ProtoMessage()    {}
func (*HistogramData) Descriptor() ([]byte, []int) {
	return fileDescriptor_histogram_40b21550afd0ea12, []int{0}
}
func (m *HistogramData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistogramData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *HistogramData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistogramData.Merge(dst, src)
}
func (m *HistogramData) XXX_Size() int {
	return m.Size()
}
func (m *HistogramData) XXX_DiscardUnknown() {
	xxx_messageInfo_HistogramData.DiscardUnknown(m)
}

var xxx_messageInfo_HistogramData proto.InternalMessageInfo

type HistogramData_Bucket struct {
	// The estimated number of values that are equal to upper_bound.
	NumEq int64 `protobuf:"varint,1,opt,name=num_eq,json=numEq,proto3" json:"num_eq,omitempty"`
	// The estimated number of values in the bucket (excluding those
	// that are equal to upper_bound). Splitting the count into two
	// makes the histogram effectively equivalent to a histogram with
	// twice as many buckets, with every other bucket containing a
	// single value. This might be particularly advantageous if the
	// histogram algorithm makes sure the top "heavy hitters" (most
	// frequent elements) are bucket boundaries (similar to a
	// compressed histogram).
	NumRange int64 `protobuf:"varint,2,opt,name=num_range,json=numRange,proto3" json:"num_range,omitempty"`
	// The estimated number of distinct values in the bucket (excluding
	// those that are equal to upper_bound). This is a floating point
	// value because it is estimated by distributing the known distinct
	// count for the column among the buckets, in proportion to the number
	// of rows in each bucket. This value is in fact derived from the rest
	// of the data, but is included to avoid re-computing it later.
	DistinctRange float64 `protobuf:"fixed64,4,opt,name=distinct_range,json=distinctRange,proto3" json:"distinct_range,omitempty"`
	// The upper boundary of the bucket. The column values for the upper bound
	// are encoded using the ascending key encoding of the column type.
	UpperBound []byte `protobuf:"bytes,3,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
}

func (m *HistogramData_Bucket) Reset()         { *m = HistogramData_Bucket{} }
func (m *HistogramData_Bucket) String() string { return proto.CompactTextString(m) }
func (*HistogramData_Bucket) ProtoMessage()    {}
func (*HistogramData_Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_histogram_40b21550afd0ea12, []int{0, 0}
}
func (m *HistogramData_Bucket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistogramData_Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *HistogramData_Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistogramData_Bucket.Merge(dst, src)
}
func (m *HistogramData_Bucket) XXX_Size() int {
	return m.Size()
}
func (m *HistogramData_Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_HistogramData_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_HistogramData_Bucket proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HistogramData)(nil), "cockroach.sql.stats.HistogramData")
	proto.RegisterType((*HistogramData_Bucket)(nil), "cockroach.sql.stats.HistogramData.Bucket")
}
func (m *HistogramData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistogramData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Buckets) > 0 {
		for _, msg := range m.Buckets {
			dAtA[i] = 0xa
			i++
			i = encodeVarintHistogram(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ColumnType != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHistogram(dAtA, i, uint64(m.ColumnType.Size()))
		n1, err := m.ColumnType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *HistogramData_Bucket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistogramData_Bucket) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NumEq != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHistogram(dAtA, i, uint64(m.NumEq))
	}
	if m.NumRange != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHistogram(dAtA, i, uint64(m.NumRange))
	}
	if len(m.UpperBound) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHistogram(dAtA, i, uint64(len(m.UpperBound)))
		i += copy(dAtA[i:], m.UpperBound)
	}
	if m.DistinctRange != 0 {
		dAtA[i] = 0x21
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DistinctRange))))
		i += 8
	}
	return i, nil
}

func encodeVarintHistogram(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HistogramData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Buckets) > 0 {
		for _, e := range m.Buckets {
			l = e.Size()
			n += 1 + l + sovHistogram(uint64(l))
		}
	}
	if m.ColumnType != nil {
		l = m.ColumnType.Size()
		n += 1 + l + sovHistogram(uint64(l))
	}
	return n
}

func (m *HistogramData_Bucket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumEq != 0 {
		n += 1 + sovHistogram(uint64(m.NumEq))
	}
	if m.NumRange != 0 {
		n += 1 + sovHistogram(uint64(m.NumRange))
	}
	l = len(m.UpperBound)
	if l > 0 {
		n += 1 + l + sovHistogram(uint64(l))
	}
	if m.DistinctRange != 0 {
		n += 9
	}
	return n
}

func sovHistogram(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHistogram(x uint64) (n int) {
	return sovHistogram(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HistogramData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHistogram
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistogramData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistogramData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buckets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistogram
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHistogram
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buckets = append(m.Buckets, HistogramData_Bucket{})
			if err := m.Buckets[len(m.Buckets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistogram
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHistogram
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ColumnType == nil {
				m.ColumnType = &types.T{}
			}
			if err := m.ColumnType.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHistogram(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHistogram
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
func (m *HistogramData_Bucket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHistogram
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEq", wireType)
			}
			m.NumEq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistogram
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEq |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRange", wireType)
			}
			m.NumRange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistogram
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRange |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperBound", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistogram
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHistogram
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpperBound = append(m.UpperBound[:0], dAtA[iNdEx:postIndex]...)
			if m.UpperBound == nil {
				m.UpperBound = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistinctRange", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.DistinctRange = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipHistogram(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHistogram
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
func skipHistogram(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHistogram
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
					return 0, ErrIntOverflowHistogram
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHistogram
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHistogram
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHistogram
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHistogram(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHistogram = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHistogram   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/stats/histogram.proto", fileDescriptor_histogram_40b21550afd0ea12)
}

var fileDescriptor_histogram_40b21550afd0ea12 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xe3, 0xa6, 0xed, 0xff, 0xe3, 0x50, 0x06, 0x43, 0xa5, 0x10, 0x24, 0x37, 0x42, 0x42,
	0x84, 0xc5, 0x95, 0xca, 0xc8, 0x16, 0x81, 0x04, 0xab, 0xd5, 0x89, 0xa5, 0x4a, 0xd3, 0x28, 0xad,
	0xda, 0xd8, 0x49, 0x6c, 0x0f, 0xdd, 0x39, 0x00, 0x27, 0xe0, 0x3c, 0x1d, 0x3b, 0x76, 0x42, 0x90,
	0x5e, 0x04, 0xd9, 0xa6, 0x03, 0x88, 0x25, 0x8a, 0x1f, 0xbf, 0xef, 0xf3, 0x59, 0x1f, 0x3c, 0x17,
	0xd5, 0x6a, 0x28, 0x64, 0x22, 0xc5, 0x70, 0xbe, 0x10, 0x92, 0xe7, 0x75, 0x52, 0x90, 0xb2, 0xe6,
	0x92, 0xa3, 0xd3, 0x94, 0xa7, 0xcb, 0x9a, 0x27, 0xe9, 0x9c, 0x88, 0x6a, 0x45, 0x4c, 0x28, 0x38,
	0xcb, 0x79, 0xce, 0xcd, 0xfd, 0x50, 0xff, 0xd9, 0x68, 0xd0, 0xd7, 0x16, 0xb9, 0x2e, 0x33, 0x61,
	0xbf, 0x16, 0x5f, 0xbe, 0xb5, 0x60, 0xef, 0xf1, 0x60, 0xbd, 0x4f, 0x64, 0x82, 0x9e, 0xe0, 0xbf,
	0xa9, 0x4a, 0x97, 0x99, 0x14, 0x3e, 0x08, 0xdd, 0xc8, 0x1b, 0xdd, 0x90, 0x3f, 0xa6, 0x90, 0x1f,
	0x25, 0x12, 0x9b, 0x46, 0xdc, 0xde, 0xbc, 0x0f, 0x1c, 0x7a, 0xe8, 0xa3, 0x3b, 0xe8, 0xa5, 0x7c,
	0xa5, 0x0a, 0x36, 0xd1, 0x23, 0xfd, 0x56, 0x08, 0x22, 0x6f, 0x14, 0xfc, 0xd6, 0x65, 0x05, 0xb1,
	0x2f, 0x1a, 0x53, 0x68, 0xe3, 0xe3, 0x75, 0x99, 0x05, 0x2f, 0x00, 0x76, 0xad, 0x16, 0xf5, 0x61,
	0x97, 0xa9, 0x62, 0x92, 0x55, 0x3e, 0x08, 0x41, 0xe4, 0xd2, 0x0e, 0x53, 0xc5, 0x43, 0x85, 0x2e,
	0xe0, 0x91, 0xc6, 0x75, 0xc2, 0x72, 0x2b, 0x77, 0xe9, 0x7f, 0xa6, 0x0a, 0xaa, 0xcf, 0x68, 0x00,
	0x3d, 0x55, 0x96, 0x59, 0x3d, 0x99, 0x72, 0xc5, 0x66, 0xbe, 0x1b, 0x82, 0xe8, 0x98, 0x42, 0x83,
	0x62, 0x4d, 0xd0, 0x15, 0x3c, 0x99, 0x2d, 0x84, 0x5c, 0xb0, 0x54, 0x7e, 0x2b, 0xda, 0x21, 0x88,
	0x00, 0xed, 0x1d, 0xa8, 0xf1, 0xc4, 0xd7, 0x9b, 0x4f, 0xec, 0x6c, 0x1a, 0x0c, 0xb6, 0x0d, 0x06,
	0xbb, 0x06, 0x83, 0x8f, 0x06, 0x83, 0xd7, 0x3d, 0x76, 0xb6, 0x7b, 0xec, 0xec, 0xf6, 0xd8, 0x79,
	0xee, 0x98, 0x85, 0x4c, 0xbb, 0x66, 0xa1, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xd2,
	0x87, 0x1e, 0xaf, 0x01, 0x00, 0x00,
}

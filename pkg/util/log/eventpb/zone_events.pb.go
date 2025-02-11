// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/zone_events.proto

package eventpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

// CommonZoneConfigDetails is common to zone config change events.
type CommonZoneConfigDetails struct {
	// The target object of the zone config change.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:",omitempty"`
	// The applied zone config in YAML format.
	Config string `protobuf:"bytes,2,opt,name=config,proto3" json:",omitempty"`
	// The SQL representation of the applied zone config options.
	Options []string `protobuf:"bytes,3,rep,name=options,proto3" json:",omitempty"`
}

func (m *CommonZoneConfigDetails) Reset()         { *m = CommonZoneConfigDetails{} }
func (m *CommonZoneConfigDetails) String() string { return proto.CompactTextString(m) }
func (*CommonZoneConfigDetails) ProtoMessage()    {}
func (*CommonZoneConfigDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_zone_events_f93c01e91b797e1b, []int{0}
}
func (m *CommonZoneConfigDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonZoneConfigDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *CommonZoneConfigDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonZoneConfigDetails.Merge(dst, src)
}
func (m *CommonZoneConfigDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonZoneConfigDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonZoneConfigDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonZoneConfigDetails proto.InternalMessageInfo

// SetZoneConfig is recorded when a zone config is changed.
type SetZoneConfig struct {
	CommonEventDetails      `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonSQLEventDetails   `protobuf:"bytes,2,opt,name=sql,proto3,embedded=sql" json:""`
	CommonZoneConfigDetails `protobuf:"bytes,3,opt,name=config,proto3,embedded=config" json:""`
}

func (m *SetZoneConfig) Reset()         { *m = SetZoneConfig{} }
func (m *SetZoneConfig) String() string { return proto.CompactTextString(m) }
func (*SetZoneConfig) ProtoMessage()    {}
func (*SetZoneConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_zone_events_f93c01e91b797e1b, []int{1}
}
func (m *SetZoneConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetZoneConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SetZoneConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetZoneConfig.Merge(dst, src)
}
func (m *SetZoneConfig) XXX_Size() int {
	return m.Size()
}
func (m *SetZoneConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SetZoneConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SetZoneConfig proto.InternalMessageInfo

// RemoveZoneConfig is recorded when a zone config is removed.
type RemoveZoneConfig struct {
	CommonEventDetails      `protobuf:"bytes,1,opt,name=common,proto3,embedded=common" json:""`
	CommonSQLEventDetails   `protobuf:"bytes,2,opt,name=sql,proto3,embedded=sql" json:""`
	CommonZoneConfigDetails `protobuf:"bytes,3,opt,name=config,proto3,embedded=config" json:""`
}

func (m *RemoveZoneConfig) Reset()         { *m = RemoveZoneConfig{} }
func (m *RemoveZoneConfig) String() string { return proto.CompactTextString(m) }
func (*RemoveZoneConfig) ProtoMessage()    {}
func (*RemoveZoneConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_zone_events_f93c01e91b797e1b, []int{2}
}
func (m *RemoveZoneConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveZoneConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *RemoveZoneConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveZoneConfig.Merge(dst, src)
}
func (m *RemoveZoneConfig) XXX_Size() int {
	return m.Size()
}
func (m *RemoveZoneConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveZoneConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveZoneConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommonZoneConfigDetails)(nil), "cockroach.util.log.eventpb.CommonZoneConfigDetails")
	proto.RegisterType((*SetZoneConfig)(nil), "cockroach.util.log.eventpb.SetZoneConfig")
	proto.RegisterType((*RemoveZoneConfig)(nil), "cockroach.util.log.eventpb.RemoveZoneConfig")
}
func (m *CommonZoneConfigDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonZoneConfigDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintZoneEvents(dAtA, i, uint64(len(m.Target)))
		i += copy(dAtA[i:], m.Target)
	}
	if len(m.Config) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintZoneEvents(dAtA, i, uint64(len(m.Config)))
		i += copy(dAtA[i:], m.Config)
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *SetZoneConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetZoneConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintZoneEvents(dAtA, i, uint64(m.CommonEventDetails.Size()))
	n1, err := m.CommonEventDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintZoneEvents(dAtA, i, uint64(m.CommonSQLEventDetails.Size()))
	n2, err := m.CommonSQLEventDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintZoneEvents(dAtA, i, uint64(m.CommonZoneConfigDetails.Size()))
	n3, err := m.CommonZoneConfigDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *RemoveZoneConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveZoneConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintZoneEvents(dAtA, i, uint64(m.CommonEventDetails.Size()))
	n4, err := m.CommonEventDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x12
	i++
	i = encodeVarintZoneEvents(dAtA, i, uint64(m.CommonSQLEventDetails.Size()))
	n5, err := m.CommonSQLEventDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x1a
	i++
	i = encodeVarintZoneEvents(dAtA, i, uint64(m.CommonZoneConfigDetails.Size()))
	n6, err := m.CommonZoneConfigDetails.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeVarintZoneEvents(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CommonZoneConfigDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovZoneEvents(uint64(l))
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovZoneEvents(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovZoneEvents(uint64(l))
		}
	}
	return n
}

func (m *SetZoneConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovZoneEvents(uint64(l))
	l = m.CommonSQLEventDetails.Size()
	n += 1 + l + sovZoneEvents(uint64(l))
	l = m.CommonZoneConfigDetails.Size()
	n += 1 + l + sovZoneEvents(uint64(l))
	return n
}

func (m *RemoveZoneConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CommonEventDetails.Size()
	n += 1 + l + sovZoneEvents(uint64(l))
	l = m.CommonSQLEventDetails.Size()
	n += 1 + l + sovZoneEvents(uint64(l))
	l = m.CommonZoneConfigDetails.Size()
	n += 1 + l + sovZoneEvents(uint64(l))
	return n
}

func sovZoneEvents(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozZoneEvents(x uint64) (n int) {
	return sovZoneEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonZoneConfigDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneEvents
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
			return fmt.Errorf("proto: CommonZoneConfigDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonZoneConfigDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneEvents
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
func (m *SetZoneConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneEvents
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
			return fmt.Errorf("proto: SetZoneConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetZoneConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
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
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonSQLEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
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
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonSQLEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonZoneConfigDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
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
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonZoneConfigDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneEvents
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
func (m *RemoveZoneConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneEvents
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
			return fmt.Errorf("proto: RemoveZoneConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveZoneConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
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
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonSQLEventDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
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
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonSQLEventDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonZoneConfigDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneEvents
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
				return ErrInvalidLengthZoneEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommonZoneConfigDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneEvents
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
func skipZoneEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZoneEvents
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
					return 0, ErrIntOverflowZoneEvents
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
					return 0, ErrIntOverflowZoneEvents
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
				return 0, ErrInvalidLengthZoneEvents
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowZoneEvents
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
				next, err := skipZoneEvents(dAtA[start:])
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
	ErrInvalidLengthZoneEvents = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZoneEvents   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("util/log/eventpb/zone_events.proto", fileDescriptor_zone_events_f93c01e91b797e1b)
}

var fileDescriptor_zone_events_f93c01e91b797e1b = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x92, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x33, 0x0d, 0xb4, 0x38, 0x55, 0x91, 0x20, 0x58, 0x0a, 0x4e, 0x4a, 0x16, 0x52, 0x41,
	0x26, 0xd8, 0xde, 0xa0, 0xd5, 0x9d, 0x9b, 0xb6, 0x6e, 0xec, 0x46, 0xda, 0x30, 0x8e, 0xc1, 0x64,
	0x5e, 0x6c, 0xc6, 0x82, 0x9e, 0x42, 0x28, 0x78, 0xa6, 0x2e, 0xbb, 0xec, 0x2a, 0x68, 0xba, 0xeb,
	0x29, 0x64, 0xa6, 0x51, 0x0b, 0xa9, 0xf6, 0x02, 0xee, 0xf2, 0xf8, 0xff, 0xf7, 0xe5, 0x7f, 0x3f,
	0x83, 0x9d, 0x27, 0xe9, 0x07, 0x6e, 0x00, 0xdc, 0x65, 0x63, 0x26, 0x64, 0x34, 0x74, 0x5f, 0x40,
	0xb0, 0x5b, 0x3d, 0xc4, 0x34, 0x1a, 0x81, 0x04, 0xab, 0xea, 0x81, 0xf7, 0x30, 0x82, 0x81, 0x77,
	0x4f, 0x95, 0x9b, 0x06, 0xc0, 0x69, 0xe6, 0xae, 0x1e, 0x72, 0xe0, 0xa0, 0x6d, 0xae, 0xfa, 0x5a,
	0x6d, 0x54, 0x8f, 0x73, 0xd4, 0x75, 0xa0, 0x33, 0x41, 0xf8, 0xa8, 0x0d, 0x61, 0x08, 0xa2, 0x0f,
	0x82, 0xb5, 0x41, 0xdc, 0xf9, 0xfc, 0x82, 0xc9, 0x81, 0x1f, 0xc4, 0xd6, 0x09, 0x2e, 0xca, 0xc1,
	0x88, 0x33, 0x59, 0x41, 0x35, 0x54, 0xdf, 0x69, 0xed, 0x2f, 0x13, 0x1b, 0x9f, 0x41, 0xe8, 0x4b,
	0x16, 0x46, 0xf2, 0xb9, 0x9b, 0xa9, 0xca, 0xe7, 0xe9, 0xc5, 0x4a, 0x61, 0xb3, 0x6f, 0xa5, 0x5a,
	0x75, 0x5c, 0x82, 0x48, 0xfa, 0x20, 0xe2, 0x8a, 0x59, 0x33, 0x37, 0x18, 0xbf, 0x64, 0x67, 0x52,
	0xc0, 0x7b, 0x3d, 0x26, 0x7f, 0x22, 0x59, 0xd7, 0xea, 0x1f, 0x2a, 0xa6, 0xce, 0x52, 0x6e, 0x50,
	0xfa, 0x7b, 0x13, 0x74, 0x75, 0xd0, 0xa5, 0x9a, 0xb2, 0x5b, 0x5a, 0xbb, 0xd3, 0xc4, 0x36, 0x66,
	0x89, 0x8d, 0x96, 0x89, 0x6d, 0x74, 0x33, 0x96, 0xd5, 0xc1, 0x66, 0xfc, 0x18, 0xe8, 0xd8, 0xe5,
	0xc6, 0xf9, 0x76, 0x64, 0xaf, 0x73, 0xf5, 0x07, 0x55, 0xb1, 0xac, 0x9b, 0xef, 0x32, 0x4c, 0x4d,
	0x6d, 0x6e, 0xa7, 0xe6, 0x9a, 0xcf, 0xa7, 0x55, 0xa2, 0xf3, 0x56, 0xc0, 0x07, 0x5d, 0x16, 0xc2,
	0x98, 0xfd, 0x17, 0xb3, 0x5e, 0x4c, 0xeb, 0x74, 0xfa, 0x41, 0x8c, 0x69, 0x4a, 0xd0, 0x2c, 0x25,
	0x68, 0x9e, 0x12, 0xf4, 0x9e, 0x12, 0xf4, 0xba, 0x20, 0xc6, 0x6c, 0x41, 0x8c, 0xf9, 0x82, 0x18,
	0xfd, 0x52, 0x06, 0x1e, 0x16, 0xf5, 0xb3, 0x6f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x1c,
	0x5c, 0xa8, 0x6d, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: orderer/kafka.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// KafkaMessage is a wrapper type for the messages
// that the Kafka-based orderer deals with.
type KafkaMessage struct {
	// Types that are valid to be assigned to Type:
	//	*KafkaMessage_Regular
	//	*KafkaMessage_TimeToCut
	//	*KafkaMessage_Connect
	Type isKafkaMessage_Type `protobuf_oneof:"Type"`
}

func (m *KafkaMessage) Reset()                    { *m = KafkaMessage{} }
func (m *KafkaMessage) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessage) ProtoMessage()               {}
func (*KafkaMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isKafkaMessage_Type interface {
	isKafkaMessage_Type()
}

type KafkaMessage_Regular struct {
	Regular *KafkaMessageRegular `protobuf:"bytes,1,opt,name=regular,oneof"`
}
type KafkaMessage_TimeToCut struct {
	TimeToCut *KafkaMessageTimeToCut `protobuf:"bytes,2,opt,name=time_to_cut,json=timeToCut,oneof"`
}
type KafkaMessage_Connect struct {
	Connect *KafkaMessageConnect `protobuf:"bytes,3,opt,name=connect,oneof"`
}

func (*KafkaMessage_Regular) isKafkaMessage_Type()   {}
func (*KafkaMessage_TimeToCut) isKafkaMessage_Type() {}
func (*KafkaMessage_Connect) isKafkaMessage_Type()   {}

func (m *KafkaMessage) GetType() isKafkaMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *KafkaMessage) GetRegular() *KafkaMessageRegular {
	if x, ok := m.GetType().(*KafkaMessage_Regular); ok {
		return x.Regular
	}
	return nil
}

func (m *KafkaMessage) GetTimeToCut() *KafkaMessageTimeToCut {
	if x, ok := m.GetType().(*KafkaMessage_TimeToCut); ok {
		return x.TimeToCut
	}
	return nil
}

func (m *KafkaMessage) GetConnect() *KafkaMessageConnect {
	if x, ok := m.GetType().(*KafkaMessage_Connect); ok {
		return x.Connect
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KafkaMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KafkaMessage_OneofMarshaler, _KafkaMessage_OneofUnmarshaler, _KafkaMessage_OneofSizer, []interface{}{
		(*KafkaMessage_Regular)(nil),
		(*KafkaMessage_TimeToCut)(nil),
		(*KafkaMessage_Connect)(nil),
	}
}

func _KafkaMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KafkaMessage)
	// Type
	switch x := m.Type.(type) {
	case *KafkaMessage_Regular:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Regular); err != nil {
			return err
		}
	case *KafkaMessage_TimeToCut:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeToCut); err != nil {
			return err
		}
	case *KafkaMessage_Connect:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Connect); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("KafkaMessage.Type has unexpected type %T", x)
	}
	return nil
}

func _KafkaMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KafkaMessage)
	switch tag {
	case 1: // Type.regular
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageRegular)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_Regular{msg}
		return true, err
	case 2: // Type.time_to_cut
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageTimeToCut)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_TimeToCut{msg}
		return true, err
	case 3: // Type.connect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaMessageConnect)
		err := b.DecodeMessage(msg)
		m.Type = &KafkaMessage_Connect{msg}
		return true, err
	default:
		return false, nil
	}
}

func _KafkaMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KafkaMessage)
	// Type
	switch x := m.Type.(type) {
	case *KafkaMessage_Regular:
		s := proto.Size(x.Regular)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KafkaMessage_TimeToCut:
		s := proto.Size(x.TimeToCut)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *KafkaMessage_Connect:
		s := proto.Size(x.Connect)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KafkaMessageRegular wraps a marshalled envelope.
type KafkaMessageRegular struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *KafkaMessageRegular) Reset()                    { *m = KafkaMessageRegular{} }
func (m *KafkaMessageRegular) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessageRegular) ProtoMessage()               {}
func (*KafkaMessageRegular) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// KafkaMessageTimeToCut is used to signal to the orderers
// that it is time to cut block <block_number>.
type KafkaMessageTimeToCut struct {
	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber" json:"block_number,omitempty"`
}

func (m *KafkaMessageTimeToCut) Reset()                    { *m = KafkaMessageTimeToCut{} }
func (m *KafkaMessageTimeToCut) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessageTimeToCut) ProtoMessage()               {}
func (*KafkaMessageTimeToCut) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// KafkaMessageConnect is posted by an orderer upon booting up.
// It is used to prevent the panic that would be caused if we
// were to consume an empty partition. It is ignored by all
// orderers when processing the partition.
type KafkaMessageConnect struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *KafkaMessageConnect) Reset()                    { *m = KafkaMessageConnect{} }
func (m *KafkaMessageConnect) String() string            { return proto.CompactTextString(m) }
func (*KafkaMessageConnect) ProtoMessage()               {}
func (*KafkaMessageConnect) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// LastOffsetPersisted is the encoded value for the Metadata message
// which is encoded in the ORDERER block metadata index for the case
// of the Kafka-based orderer.
type KafkaMetadata struct {
	LastOffsetPersisted int64 `protobuf:"varint,1,opt,name=last_offset_persisted,json=lastOffsetPersisted" json:"last_offset_persisted,omitempty"`
}

func (m *KafkaMetadata) Reset()                    { *m = KafkaMetadata{} }
func (m *KafkaMetadata) String() string            { return proto.CompactTextString(m) }
func (*KafkaMetadata) ProtoMessage()               {}
func (*KafkaMetadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func init() {
	proto.RegisterType((*KafkaMessage)(nil), "orderer.KafkaMessage")
	proto.RegisterType((*KafkaMessageRegular)(nil), "orderer.KafkaMessageRegular")
	proto.RegisterType((*KafkaMessageTimeToCut)(nil), "orderer.KafkaMessageTimeToCut")
	proto.RegisterType((*KafkaMessageConnect)(nil), "orderer.KafkaMessageConnect")
	proto.RegisterType((*KafkaMetadata)(nil), "orderer.KafkaMetadata")
}

func init() { proto.RegisterFile("orderer/kafka.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x3d, 0x6b, 0xf3, 0x30,
	0x14, 0x85, 0x93, 0x37, 0x21, 0xe1, 0x55, 0xd2, 0x45, 0x21, 0xe0, 0xa1, 0x94, 0x36, 0x53, 0x87,
	0x62, 0x41, 0xba, 0x94, 0x4e, 0x25, 0x59, 0x02, 0xa5, 0x1f, 0x08, 0x4f, 0x5d, 0x8c, 0x2c, 0x5f,
	0x3b, 0xc6, 0x1f, 0x32, 0xd2, 0xf5, 0xe0, 0xff, 0xd8, 0x1f, 0x55, 0x2c, 0xcb, 0x50, 0x8a, 0xc9,
	0x78, 0xef, 0x79, 0x0e, 0xe7, 0xe8, 0x8a, 0x6c, 0x94, 0x8e, 0x41, 0x83, 0x66, 0xb9, 0x48, 0x72,
	0xe1, 0xd7, 0x5a, 0xa1, 0xa2, 0x4b, 0xb7, 0xdc, 0x7d, 0x4f, 0xc9, 0xfa, 0xb5, 0x13, 0xde, 0xc0,
	0x18, 0x91, 0x02, 0x7d, 0x22, 0x4b, 0x0d, 0x69, 0x53, 0x08, 0xed, 0x4d, 0x6f, 0xa7, 0xf7, 0xab,
	0xfd, 0xb5, 0xef, 0x58, 0xff, 0x37, 0xc7, 0x7b, 0xe6, 0x34, 0xe1, 0x03, 0x4e, 0x5f, 0xc8, 0x0a,
	0xb3, 0x12, 0x42, 0x54, 0xa1, 0x6c, 0xd0, 0xfb, 0x67, 0xdd, 0x37, 0xa3, 0xee, 0x20, 0x2b, 0x21,
	0x50, 0xc7, 0x06, 0x4f, 0x13, 0xfe, 0x1f, 0x87, 0xa1, 0xcb, 0x96, 0xaa, 0xaa, 0x40, 0xa2, 0x37,
	0xbb, 0x90, 0x7d, 0xec, 0x99, 0x2e, 0xdb, 0xe1, 0x87, 0x05, 0x99, 0x07, 0x6d, 0x0d, 0x3b, 0x46,
	0x36, 0x23, 0x2d, 0xa9, 0x47, 0x96, 0xb5, 0x68, 0x0b, 0x25, 0x62, 0xfb, 0xa8, 0x35, 0x1f, 0xc6,
	0xdd, 0x33, 0xd9, 0x8e, 0x16, 0xa3, 0x77, 0x64, 0x1d, 0x15, 0x4a, 0xe6, 0x61, 0xd5, 0x94, 0x11,
	0xf4, 0xc7, 0x98, 0xf3, 0x95, 0xdd, 0xbd, 0xdb, 0xd5, 0xdf, 0x30, 0x57, 0xeb, 0x42, 0xd8, 0x91,
	0x5c, 0x39, 0x03, 0x8a, 0x58, 0xa0, 0xa0, 0x7b, 0xb2, 0x2d, 0x84, 0xc1, 0x50, 0x25, 0x89, 0x01,
	0x0c, 0x6b, 0xd0, 0x26, 0x33, 0x08, 0xbd, 0x71, 0xc6, 0x37, 0x9d, 0xf8, 0x61, 0xb5, 0xcf, 0x41,
	0x3a, 0xf8, 0x5f, 0x0f, 0x69, 0x86, 0xe7, 0x26, 0xf2, 0xa5, 0x2a, 0xd9, 0xb9, 0xad, 0x41, 0x17,
	0x10, 0xa7, 0xa0, 0x59, 0x22, 0x22, 0x9d, 0x49, 0x66, 0x7f, 0xd8, 0x30, 0x77, 0xb9, 0x68, 0x61,
	0xe7, 0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x89, 0x66, 0xa4, 0x08, 0x02, 0x00, 0x00,
}

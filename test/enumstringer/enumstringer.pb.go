// Code generated by protoc-gen-gogo.
// source: enumstringer.proto
// DO NOT EDIT!

/*
Package enumstringer is a generated protocol buffer package.

It is generated from these files:
	enumstringer.proto

It has these top-level messages:
	NidOptEnum
	NinOptEnum
	NidRepEnum
	NinRepEnum
*/
package enumstringer

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TheTestEnum int32

const (
	TheTestEnum_A TheTestEnum = 0
	TheTestEnum_B TheTestEnum = 1
	TheTestEnum_C TheTestEnum = 2
)

var TheTestEnum_name = map[int32]string{
	0: "A",
	1: "B",
	2: "C",
}
var TheTestEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
	"C": 2,
}

func (x TheTestEnum) Enum() *TheTestEnum {
	p := new(TheTestEnum)
	*p = x
	return p
}
func (x TheTestEnum) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(TheTestEnum_name, int32(x))
}
func (x *TheTestEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TheTestEnum_value, data, "TheTestEnum")
	if err != nil {
		return err
	}
	*x = TheTestEnum(value)
	return nil
}
func (TheTestEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptorEnumstringer, []int{0} }

type NidOptEnum struct {
	Field1           TheTestEnum `protobuf:"varint,1,opt,name=Field1,json=field1,enum=enumstringer.TheTestEnum" json:"Field1"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *NidOptEnum) Reset()                    { *m = NidOptEnum{} }
func (m *NidOptEnum) String() string            { return proto.CompactTextString(m) }
func (*NidOptEnum) ProtoMessage()               {}
func (*NidOptEnum) Descriptor() ([]byte, []int) { return fileDescriptorEnumstringer, []int{0} }

func (m *NidOptEnum) GetField1() TheTestEnum {
	if m != nil {
		return m.Field1
	}
	return TheTestEnum_A
}

type NinOptEnum struct {
	Field1           *TheTestEnum `protobuf:"varint,1,opt,name=Field1,json=field1,enum=enumstringer.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NinOptEnum) Reset()                    { *m = NinOptEnum{} }
func (m *NinOptEnum) String() string            { return proto.CompactTextString(m) }
func (*NinOptEnum) ProtoMessage()               {}
func (*NinOptEnum) Descriptor() ([]byte, []int) { return fileDescriptorEnumstringer, []int{1} }

func (m *NinOptEnum) GetField1() TheTestEnum {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return TheTestEnum_A
}

type NidRepEnum struct {
	Field1           []TheTestEnum `protobuf:"varint,1,rep,name=Field1,json=field1,enum=enumstringer.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NidRepEnum) Reset()                    { *m = NidRepEnum{} }
func (m *NidRepEnum) String() string            { return proto.CompactTextString(m) }
func (*NidRepEnum) ProtoMessage()               {}
func (*NidRepEnum) Descriptor() ([]byte, []int) { return fileDescriptorEnumstringer, []int{2} }

func (m *NidRepEnum) GetField1() []TheTestEnum {
	if m != nil {
		return m.Field1
	}
	return nil
}

type NinRepEnum struct {
	Field1           []TheTestEnum `protobuf:"varint,1,rep,name=Field1,json=field1,enum=enumstringer.TheTestEnum" json:"Field1,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinRepEnum) Reset()                    { *m = NinRepEnum{} }
func (m *NinRepEnum) String() string            { return proto.CompactTextString(m) }
func (*NinRepEnum) ProtoMessage()               {}
func (*NinRepEnum) Descriptor() ([]byte, []int) { return fileDescriptorEnumstringer, []int{3} }

func (m *NinRepEnum) GetField1() []TheTestEnum {
	if m != nil {
		return m.Field1
	}
	return nil
}

func init() {
	proto.RegisterType((*NidOptEnum)(nil), "enumstringer.NidOptEnum")
	proto.RegisterType((*NinOptEnum)(nil), "enumstringer.NinOptEnum")
	proto.RegisterType((*NidRepEnum)(nil), "enumstringer.NidRepEnum")
	proto.RegisterType((*NinRepEnum)(nil), "enumstringer.NinRepEnum")
	proto.RegisterEnum("enumstringer.TheTestEnum", TheTestEnum_name, TheTestEnum_value)
}
func (this *NidOptEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NidOptEnum)
	if !ok {
		that2, ok := that.(NidOptEnum)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NidOptEnum")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NidOptEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NidOptEnum but is not nil && this == nil")
	}
	if this.Field1 != that1.Field1 {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NidOptEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NidOptEnum)
	if !ok {
		that2, ok := that.(NidOptEnum)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != that1.Field1 {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NinOptEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NinOptEnum)
	if !ok {
		that2, ok := that.(NinOptEnum)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NinOptEnum")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NinOptEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NinOptEnum but is not nil && this == nil")
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", *this.Field1, *that1.Field1)
		}
	} else if this.Field1 != nil {
		return fmt.Errorf("this.Field1 == nil && that.Field1 != nil")
	} else if that1.Field1 != nil {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NinOptEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NinOptEnum)
	if !ok {
		that2, ok := that.(NinOptEnum)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NidRepEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NidRepEnum)
	if !ok {
		that2, ok := that.(NidRepEnum)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NidRepEnum")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NidRepEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NidRepEnum but is not nil && this == nil")
	}
	if len(this.Field1) != len(that1.Field1) {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", len(this.Field1), len(that1.Field1))
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return fmt.Errorf("Field1 this[%v](%v) Not Equal that[%v](%v)", i, this.Field1[i], i, that1.Field1[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NidRepEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NidRepEnum)
	if !ok {
		that2, ok := that.(NidRepEnum)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Field1) != len(that1.Field1) {
		return false
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NinRepEnum) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NinRepEnum)
	if !ok {
		that2, ok := that.(NinRepEnum)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NinRepEnum")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NinRepEnum but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NinRepEnum but is not nil && this == nil")
	}
	if len(this.Field1) != len(that1.Field1) {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", len(this.Field1), len(that1.Field1))
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return fmt.Errorf("Field1 this[%v](%v) Not Equal that[%v](%v)", i, this.Field1[i], i, that1.Field1[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *NinRepEnum) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NinRepEnum)
	if !ok {
		that2, ok := that.(NinRepEnum)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Field1) != len(that1.Field1) {
		return false
	}
	for i := range this.Field1 {
		if this.Field1[i] != that1.Field1[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func NewPopulatedNidOptEnum(r randyEnumstringer, easy bool) *NidOptEnum {
	this := &NidOptEnum{}
	this.Field1 = TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

func NewPopulatedNinOptEnum(r randyEnumstringer, easy bool) *NinOptEnum {
	this := &NinOptEnum{}
	if r.Intn(10) != 0 {
		v1 := TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
		this.Field1 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

func NewPopulatedNidRepEnum(r randyEnumstringer, easy bool) *NidRepEnum {
	this := &NidRepEnum{}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Field1 = make([]TheTestEnum, v2)
		for i := 0; i < v2; i++ {
			this.Field1[i] = TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

func NewPopulatedNinRepEnum(r randyEnumstringer, easy bool) *NinRepEnum {
	this := &NinRepEnum{}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.Field1 = make([]TheTestEnum, v3)
		for i := 0; i < v3; i++ {
			this.Field1[i] = TheTestEnum([]int32{0, 1, 2}[r.Intn(3)])
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEnumstringer(r, 2)
	}
	return this
}

type randyEnumstringer interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEnumstringer(r randyEnumstringer) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEnumstringer(r randyEnumstringer) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneEnumstringer(r)
	}
	return string(tmps)
}
func randUnrecognizedEnumstringer(r randyEnumstringer, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEnumstringer(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEnumstringer(dAtA []byte, r randyEnumstringer, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEnumstringer(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateEnumstringer(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateEnumstringer(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEnumstringer(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEnumstringer(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEnumstringer(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEnumstringer(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("enumstringer.proto", fileDescriptorEnumstringer) }

var fileDescriptorEnumstringer = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcd, 0x2b, 0xcd,
	0x2d, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x93, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2,
	0x59, 0xc9, 0x95, 0x8b, 0xcb, 0x2f, 0x33, 0xc5, 0xbf, 0xa0, 0xc4, 0x35, 0xaf, 0x34, 0x57, 0xc8,
	0x9c, 0x8b, 0xcd, 0x2d, 0x33, 0x35, 0x27, 0xc5, 0x50, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48,
	0x52, 0x0f, 0xc5, 0xbe, 0x90, 0x8c, 0xd4, 0x90, 0xd4, 0x62, 0xb0, 0x52, 0x27, 0x96, 0x13, 0xf7,
	0xe4, 0x19, 0x82, 0xd8, 0xd2, 0xc0, 0xca, 0x95, 0xec, 0x41, 0xc6, 0xe4, 0xc1, 0x8c, 0x31, 0x24,
	0xda, 0x18, 0xb8, 0x01, 0x10, 0x77, 0x04, 0xa5, 0x16, 0x60, 0xb8, 0x83, 0x99, 0x74, 0x77, 0xc0,
	0x8c, 0x31, 0x24, 0xda, 0x18, 0x98, 0x01, 0x5a, 0x4a, 0x5c, 0xdc, 0x48, 0xc2, 0x42, 0xac, 0x5c,
	0x8c, 0x8e, 0x02, 0x0c, 0x20, 0xca, 0x49, 0x80, 0x11, 0x44, 0x39, 0x0b, 0x30, 0x39, 0x89, 0x3c,
	0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31,
	0xbe, 0x78, 0x24, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xb1, 0xac, 0x38, 0x9b, 0x01,
	0x00, 0x00,
}

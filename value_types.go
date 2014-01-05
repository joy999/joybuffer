package joybuffer

type Value interface {
	//AsInt() int
	//AsInt8() int8
	//AsInt16() int16
	//AsInt32() int32
	//AsInt64() int64
	//AsUInt() uint
	//AsUInt8() uint8
	//AsUInt16() uint16
	//AsUInt32() uint32
	//AsUInt64() uint64
	//AsFloat32() float32
	//AsFloat64() float64
	//AsString() string
	//AsArray() []Value
	//AsMap() map[string]Value
}

type Int int

type Int8 int8

type Int16 int16

type Int32 int32

type Int64 int64

type UInt uint

type UInt8 uint8

type UInt16 uint16

type UInt32 uint32

type UInt64 uint64

type Float32 float32

type Float64 float64

type String string

type Map map[string]Value

type Array []Value

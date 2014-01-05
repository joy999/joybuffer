package joybuffer

import (
	"fmt"
	"github.com/joy999/utils"
)

const (
	T_INT8  = iota
	T_INT16 = iota
	T_INT32 = iota
	T_INT64 = iota

	T_UINT8  = iota
	T_UINT16 = iota
	T_UINT32 = iota
	T_UINT64 = iota

	T_FLOAT  = iota
	T_DOUBLE = iota

	T_STRING    = iota
	T_VARSTRING = iota
	T_ARRAY     = iota
	T_MAP       = iota
)

func Serialize(v interface{}) string {
	bs := utils.NewByteString()
	serialize(v, bs)
	return string(bs.GetBuff())
}

func serialize(v Value, bs *utils.ByteString) {
	switch val := v.(type) {
	case int:
		bs.WriteInt8(T_INT32)
		bs.WriteInt32(int32(val))

	case int8:
		bs.WriteInt8(T_INT8)
		bs.WriteInt8(val)
	case int16:
		bs.WriteInt8(T_INT16)
		bs.WriteInt16(val)
	case int32:
		bs.WriteInt8(T_INT32)
		bs.WriteInt32(int32(val))
	case int64:
		bs.WriteInt8(T_INT64)
		bs.WriteInt64(val)
	case uint:
		bs.WriteInt8(T_UINT32)
		bs.WriteUInt32(uint32(val))
	case uint8:
		bs.WriteInt8(T_UINT8)
		bs.WriteUInt8(val)
	case uint32:
		bs.WriteInt8(T_UINT32)
		bs.WriteUInt32(uint32(val))
	case uint16:
		bs.WriteInt8(T_UINT16)
		bs.WriteUInt16(val)
	case uint64:
		bs.WriteInt8(T_UINT64)
		bs.WriteUInt64(val)
	case float32:
		bs.WriteInt8(T_FLOAT)
		str := fmt.Sprintf("%f", val)
		bs.WriteVarString(str)
	case float64:
		bs.WriteInt8(T_DOUBLE)
		str := fmt.Sprintf("%f", val)
		bs.WriteVarString(str)
	case string:
		bs.WriteInt8(T_STRING)
		bs.WriteUInt32(uint32(len(val)))
		bs.WriteString(val)
	case Array:
		bs.WriteInt8(T_ARRAY)
		n := len(val)
		bs.WriteUInt32(uint32(n))
		for i := 0; i < n; i++ {
			serialize(val[i], bs)
		}
	case Map:
		bs.WriteInt8(T_MAP)
		n := len(val)
		bs.WriteUInt32(uint32(n))
		for k, v := range val {
			bs.WriteVarString(k)
			serialize(v, bs)
		}
	}
}

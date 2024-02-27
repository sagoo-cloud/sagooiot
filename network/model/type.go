package model

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
	"sagooiot/network/codebin"
	"strings"
)

// DataType 数据类型
type DataType int

const (
	//TypeNONE 空类型
	TypeNONE DataType = iota
	TypeBIT
	TypeBYTE
	TypeWORD
	TypeDWORD
	TypeQWORD
	TypeSHORT
	TypeINTEGER
	TypeLONG
	TypeFLOAT
	TypeDOUBLE
)

// Parse 解析类型
func (dt *DataType) Parse(tp string) error {
	//var *dt DataType
	tp = tp[1 : len(tp)-1]
	//strings.ToLower(tp)
	switch strings.ToLower(tp) {
	case "none":
		*dt = TypeNONE
	case "bit":
		*dt = TypeBIT
	case "byte":
		*dt = TypeBYTE
	case "word":
		fallthrough
	case "uint16":
		*dt = TypeWORD
	case "dword":
		fallthrough
	case "uint32":
		*dt = TypeDWORD
	case "qword":
		fallthrough
	case "uint64":
		*dt = TypeQWORD
	case "short":
		fallthrough
	case "int16":
		*dt = TypeSHORT
	case "integer":
		fallthrough
	case "int32":
		fallthrough
	case "int":
		*dt = TypeINTEGER
	case "long":
		fallthrough
	case "int64":
		*dt = TypeLONG
	case "float":
		*dt = TypeFLOAT
	case "double":
		fallthrough
	case "float64":
		*dt = TypeDOUBLE
	default:
		return fmt.Errorf("Unknown data type: %s ", tp)
	}
	return nil
}

// String 转化成字符串
func (dt *DataType) String() string {
	var str string
	switch *dt {
	case TypeBIT:
		str = "bit"
	case TypeBYTE:
		str = "byte"
	case TypeWORD:
		str = "word"
	case TypeDWORD:
		str = "dword"
	case TypeQWORD:
		str = "qword"
	case TypeSHORT:
		str = "short"
	case TypeINTEGER:
		str = "integer"
	case TypeLONG:
		str = "long"
	case TypeFLOAT:
		str = "float"
	case TypeDOUBLE:
		str = "double"
	default:
		str = "none"
	}
	return str
}

func (dt *DataType) Default() interface{} {
	switch *dt {
	case TypeBIT:
		return false
	case TypeBYTE:
		return byte(0)
	case TypeWORD:
		return uint16(0)
	case TypeDWORD:
		return uint32(0)
	case TypeQWORD:
		return uint64(0)
	case TypeSHORT:
		return int16(0)
	case TypeINTEGER:
		return int32(0)
	case TypeLONG:
		return int64(0)
	case TypeFLOAT:
		return float32(0)
	case TypeDOUBLE:
		return float64(0)
	default:
		return 0
	}
}

func (dt *DataType) Normalize(val interface{}) interface{} {
	switch *dt {
	case TypeBIT:
		return gconv.Bool(val)
	case TypeBYTE:
		return gconv.Uint8(val)
	case TypeWORD:
		return gconv.Uint16(val)
	case TypeDWORD:
		return gconv.Uint32(val)
	case TypeQWORD:
		return gconv.Uint64(val)
	case TypeSHORT:
		return gconv.Int16(val)
	case TypeINTEGER:
		return gconv.Int32(val)
	case TypeLONG:
		return gconv.Int64(val)
	case TypeFLOAT:
		return gconv.Float32(val)
	case TypeDOUBLE:
		return gconv.Float64(val)
	default:
		return 0
	}
}

// Size 宽度
func (dt *DataType) Size() int {
	var s int
	switch *dt {
	case TypeBIT:
		s = 1
	case TypeBYTE:
		s = 1
	case TypeWORD:
		s = 2
	case TypeDWORD:
		s = 4
	case TypeQWORD:
		s = 8
	case TypeSHORT:
		s = 2
	case TypeINTEGER:
		s = 4
	case TypeLONG:
		s = 8
	case TypeFLOAT:
		s = 4
	case TypeDOUBLE:
		s = 8
	default:
		s = 1
	}
	return s
}

// Encode 编码
func (dt *DataType) Encode(value interface{}, le bool, precision int) []byte {
	buf := make([]byte, 8)
	switch *dt {
	case TypeBIT:
		if gconv.Bool(value) {
			buf[0] = 1 //?????
		} else {
			buf[0] = 0
		}
	case TypeBYTE:
		buf[0] = gconv.Uint8(value)
	case TypeWORD:
		var val uint16
		if precision > 0 {
			val = uint16(gconv.Float64(value) * math.Pow10(precision))
		} else {
			val = gconv.Uint16(value)
		}
		if le {
			codebin.WriteUint16LittleEndian(buf, val)
		} else {
			codebin.WriteUint16(buf, val)
		}
	case TypeDWORD:
		var val uint32
		if precision > 0 {
			val = uint32(gconv.Float64(value) * math.Pow10(precision))
		} else {
			val = gconv.Uint32(value)
		}
		if le {
			codebin.WriteUint32LittleEndian(buf, val)
		} else {
			codebin.WriteUint32(buf, val)
		}
	case TypeQWORD:
		var val uint64
		if precision > 0 {
			val = uint64(value.(float64) * math.Pow10(precision))
		} else {
			val = gconv.Uint64(value)
		}
		if le {
			codebin.WriteUint64LittleEndian(buf, val)
		} else {
			codebin.WriteUint64(buf, val)
		}
	case TypeSHORT:
		var val int16
		if precision > 0 {
			val = int16(gconv.Float64(value) * math.Pow10(precision))
		} else {
			val = gconv.Int16(value)
		}
		if le {
			codebin.WriteUint16LittleEndian(buf, uint16(val))
		} else {
			codebin.WriteUint16(buf, uint16(val))
		}
	case TypeINTEGER:
		var val int32
		if precision > 0 {
			val = int32(gconv.Float64(value) * math.Pow10(precision))
		} else {
			val = gconv.Int32(value)
		}
		if le {
			codebin.WriteUint32LittleEndian(buf, uint32(val))
		} else {
			codebin.WriteUint32(buf, uint32(val))
		}
	case TypeLONG:
		var val int64
		if precision > 0 {
			val = int64(gconv.Float64(value) * math.Pow10(precision))
		} else {
			val = gconv.Int64(value)
		}
		if le {
			codebin.WriteUint64LittleEndian(buf, uint64(val))
		} else {
			codebin.WriteUint64(buf, uint64(val))
		}
	case TypeFLOAT:
		val := gconv.Float32(value)
		if le {
			codebin.WriteFloat32LittleEndian(buf, val)
		} else {
			codebin.WriteFloat32(buf, val)
		}
	case TypeDOUBLE:
		val := gconv.Float64(value)
		if le {
			codebin.WriteFloat64LittleEndian(buf, val)
		} else {
			codebin.WriteFloat64(buf, val)
		}
	default:
		//TODO error
	}
	return buf[:dt.Size()]
}

// Decode 解码
func (dt *DataType) Decode(buf []byte, le bool, precision int) (val interface{}, err error) {
	//避免越界访问
	if len(buf) < dt.Size() {
		return nil, fmt.Errorf("长度不够")
	}

	switch *dt {
	case TypeBIT:
		if buf[0] > 0 {
			val = true
		} else {
			val = false
		}
	case TypeBYTE:
		val = buf[0]
	case TypeWORD:
		var value uint16
		if le {
			value = codebin.ParseUint16LittleEndian(buf)
		} else {
			value = codebin.ParseUint16(buf)
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		} else {
			val = value
		}
	case TypeDWORD:
		var value uint32
		if le {
			value = codebin.ParseUint32LittleEndian(buf)
		} else {
			value = codebin.ParseUint32(buf)
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		} else {
			val = value
		}
	case TypeQWORD:
		var value uint64
		if le {
			value = codebin.ParseUint64LittleEndian(buf)
		} else {
			value = codebin.ParseUint64(buf)
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		} else {
			val = value
		}
	case TypeSHORT:
		var value int16
		if le {
			value = int16(codebin.ParseUint16LittleEndian(buf))
		} else {
			value = int16(codebin.ParseUint16(buf))
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		} else {
			val = value
		}
	case TypeINTEGER:
		var value int32
		if le {
			value = int32(codebin.ParseUint32LittleEndian(buf))
		} else {
			value = int32(codebin.ParseUint32(buf))
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		} else {
			val = value
		}
	case TypeLONG:
		var value int64
		if le {
			value = int64(codebin.ParseUint64LittleEndian(buf))
		} else {
			value = int64(codebin.ParseUint64(buf))
		}
		if precision > 0 {
			val = float64(value) / math.Pow10(precision)
		} else {
			val = value
		}
	case TypeFLOAT:
		if le {
			val = codebin.ParseFloat32LittleEndian(buf)
		} else {
			val = codebin.ParseFloat32(buf)
		}
	case TypeDOUBLE:
		if le {
			val = codebin.ParseFloat64LittleEndian(buf)
		} else {
			val = codebin.ParseFloat64(buf)
		}
	default:
		err = errors.New("未知的数据类型")
	}
	return
}

// MarshalJSON 序列化
func (dt *DataType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt.String() + `"`), nil
}

// UnmarshalJSON 解析
func (dt *DataType) UnmarshalJSON(buf []byte) error {
	return dt.Parse(string(buf))
}

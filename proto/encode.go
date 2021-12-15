package proto

import (
	"encoding/binary"
	"math"
	"reflect"

	"github.com/miaogaolin/gofirst/zigzag"
)

// encode 对结构体中的每个字段进行编码
// 如果有符号整数按照 zigzag 编码
// 如果是切片进行递归
func encode(val reflect.Value, fieldNumber uint8) []byte {
	typ := wireType(val)
	switch val.Kind() {
	case reflect.Int64, reflect.Int, reflect.Int32, reflect.Int8, reflect.Int16:
		return append(
			[]byte{encodeWireTag(typ, fieldNumber)},
			zigzag.EncodeInt64(val.Int())...)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return append(
			[]byte{encodeWireTag(typ, fieldNumber)},
			compressUint64(val.Uint())...)
	case reflect.Float32:
		return append(
			[]byte{encodeWireTag(typ, fieldNumber)},
			float32ToByte(float32(val.Float()))...,
		)
	case reflect.String:
		b := []byte(val.String())
		// 添加长度
		b = append(compressUint64(uint64(len(b))), b...)
		return append(
			[]byte{encodeWireTag(typ, fieldNumber)},
			b...,
		)
	case reflect.Slice:
		// 字符串切片和其它类型组装不一样
		// 例如：[]string{"a", "b"}, 结果：34 1 97 34 1 98，每个元素都带有 "(field_number << 3) | wire_type" 和 长度
		// []uint32{3, 4}, 结果：42 2 3 4，只包含一个 "(field_number << 3) | wire_type" 和 切片长度
		var res []byte
		for i := 0; i < val.Len(); i++ {
			if val.Type() == reflect.TypeOf([]string{}) {
				res = append(res, encode(val.Index(i), fieldNumber)...)
			} else {
				res = append(res, encode(val.Index(i), fieldNumber)[1:]...)
			}
		}

		if val.Type() == reflect.TypeOf([]string{}) {
			return res
		} else {
			temp := append(
				[]byte{encodeWireTag(typ, fieldNumber)},
				compressUint64(uint64(len(res)))...)
			res = append(temp, res...)
			return res
		}
	}
	return nil
}

func float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func byteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func byteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

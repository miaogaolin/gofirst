package proto

import (
	"reflect"
)

func decodeWireTag(wireValue byte) (wireType uint8, fieldNumber uint8) {
	fieldNumber, wireType = wireValue>>3, wireValue&7
	return
}

// 将 wire type 和 field number 两个信息组合
func encodeWireTag(wireType uint8, fieldNumber uint8) uint8 {
	return (fieldNumber << 3) | wireType
}

// wireType 根据数据类型获取 wire type
func wireType(val reflect.Value) uint8 {
	switch val.Kind() {
	case reflect.Int64, reflect.Int, reflect.Int32, reflect.Int8, reflect.Int16,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return VarintType
	case reflect.Float32:
		return Fixed32Type
	case reflect.String:
		return BytesType
	case reflect.Slice:
		if val.Type() != reflect.TypeOf([]string{}) {
			return BytesType
		}
	}
	return 0
}

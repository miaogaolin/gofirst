package proto

import (
	"errors"
	"reflect"
	"strconv"
)

type structField struct {
	// Bytes 字段对应的数据
	Bytes    []byte
	Value    reflect.Value
	WireType uint8
}

// fieldDesc 函数，如果 isMarshal 为 true，则序列化 data 结构体的数据。
func fieldDesc(data interface{}, isMarshal bool) (map[uint64]structField, error) {
	v, err := ptrToStruct(data)
	if err != nil {
		return nil, err
	}

	res := make(map[uint64]structField)

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("proto")
		n, _ := strconv.ParseUint(tag, 10, 64)
		if n == 0 {
			continue
		}

		field := structField{
			Value: v.Field(i),
		}
		if isMarshal {
			field.Bytes = encode(v.Field(i), uint8(n))
		}

		res[n] = field

	}

	return res, nil
}

func ptrToStruct(data interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("not struct")
	}
	return v, nil
}

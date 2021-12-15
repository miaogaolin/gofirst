package proto

import (
	"errors"
	"reflect"
)

// Marshal 序列化结构体
func Marshal(data interface{}) ([]byte, error) {
	s, err := fieldDesc(data, true)
	if err != nil {
		return nil, err
	}
	var res []byte
	for _, v := range s {
		res = append(res, v.Bytes...)
	}

	return res, nil
}

func Unmarshal(v interface{}, data []byte) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}

	fields, err := fieldDesc(v, false)
	if err != nil {
		return err
	}

	// 解析字段 number 对应的数据
	err = decode(fields, data)
	if err != nil {
		return err
	}

	return setStructValue(v, fields)
}

package proto

import (
	"reflect"
	"strconv"
)

func decode(fields map[uint64]structField, data []byte) error {
	for i := 0; i < len(data); {
		typ, number := decodeWireTag(data[i])
		field := fields[uint64(number)]
		field.WireType = typ

		var bytes []byte
		switch typ {
		case BytesType:
			length := decompressUint64(data[i+1:])
			end := i + 1 + int(length)
			bytes = data[i+1 : end]
			i = end
		case VarintType:
			for j, b := range data[i+1:] {
				if b&0x80 == 0x80 {
					bytes = append(bytes, b)
				} else {
					i += 1 + j
					break
				}
			}
		case Fixed32Type:
			end := i + 5
			bytes = data[i+1 : end]
			i = end
		case Fixed64Type:
			end := i + 9
			bytes = data[i+1 : end]
			i = end
		}

		field.Bytes = bytes
		fields[uint64(number)] = field
	}
	return nil
}

func setStructValue(v interface{}, desc map[uint64]structField) error {
	// 数据映射到参数 v
	rValue := reflect.ValueOf(v).Elem()

	rType := rValue.Type()
	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)
		tag := field.Tag.Get("proto")
		if tag == "" {
			continue
		}

		number, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return err
		}

		if fieldV, ok := desc[number]; ok {
			fieldValue := rValue.Field(i)
			switch fieldValue.Kind() {
			case reflect.Float32:
				fieldValue.SetFloat(float64(byteToFloat32(fieldV.Bytes)))
			case reflect.Float64:
				fieldValue.SetFloat(byteToFloat64(fieldV.Bytes))
			case reflect.String:
				fieldValue.SetString(string(fieldV.Bytes))

			}
		}
	}

	return nil
}

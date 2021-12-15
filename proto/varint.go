package proto

import "encoding/binary"

// 只按照 7 位压缩，不进行符号位移动
func compressUint64(n uint64) []byte {
	var res []byte
	size := binary.Size(n)
	for i := 0; i < size; i++ {
		if (n & ^uint64(0x7F)) != 0 {
			res = append(res, byte(0x80|(n&0x7F)))
			n = n >> 7
		} else {
			res = append(res, byte(n&0x7F))
			break
		}
	}
	return res
}

func decompressUint64(v []byte) uint64 {
	var res uint64
	for i, offset := 0, 0; i < len(v); i, offset = i+1, offset+7 {
		b := v[i]
		if (b & 0x80) == 0x80 {
			res |= uint64(b&0x7F) << offset
		} else {
			res |= uint64(b) << offset
			break
		}
	}
	return res
}

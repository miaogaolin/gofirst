package zigzag

import "encoding/binary"

func EncodeInt32(n int32) []byte {
	return compress(int32ToZigZag(n))
}

func DecodeInt32(b []byte) int32 {
	return toInt32(decompress(b))
}

func int32ToZigZag(n int32) int32 {
	return (n << 1) ^ (n >> 31)
}

func toInt32(zz int32) int32 {
	return int32(uint32(zz)>>1) ^ -(zz & 1)
}

func compress(zz int32) []byte {
	var res []byte
	size := binary.Size(zz)
	for i := 0; i < size; i++ {
		if (zz & ^0x7F) != 0 {
			res = append(res, byte(0x80|(zz&0x7F)))
			zz = int32(uint32(zz) >> 7)
		} else {
			res = append(res, byte(zz&0x7F))
			break
		}
	}
	return res
}

func decompress(zzByte []byte) int32 {
	var res int32
	for i, offset := 0, 0; i < len(zzByte); i, offset = i+1, offset+7 {
		b := zzByte[i]
		if (b & 0x80) == 0x80 {
			res |= int32(b&0x7F) << offset
		} else {
			res |= int32(b) << offset
			break
		}
	}
	return res
}

package zigzag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntAndZigZag(t *testing.T) {
	assert.Equal(t, int32(-1), toInt32(int32ToZigZag(-1)))
	assert.Equal(t, int32(1), toInt32(int32ToZigZag(1)))

	assert.Equal(t, []byte{1}, compress(int32ToZigZag(-1)))
	assert.Equal(t, []byte{2}, compress(int32ToZigZag(1)))
}

func TestEncodeDecodeInt32(t *testing.T) {
	assert.Equal(t, int32(1), DecodeInt32(EncodeInt32(1)))
	assert.Equal(t, int32(-1), DecodeInt32(EncodeInt32(-1)))
	assert.Equal(t, int32(1000), DecodeInt32(EncodeInt32(1000)))
}

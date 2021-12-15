package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	//Name    string   `proto:"1"`
	//Num     uint32   `proto:"2"`
	Height float32 `proto:"3"`
	//Hobbies []string `proto:"4"`
	//Nums    []uint32 `proto:"5"`
}

func TestMarshal(t *testing.T) {
	data := Example{Height: 52.1}
	res, err := Marshal(&data)
	if assert.Nil(t, err) {
		assert.Equal(t, []byte{29, 102, 102, 80, 66}, res)
	}
}

func TestUnmarshal(t *testing.T) {
	b := []byte{29, 102, 102, 80, 66}
	var data Example
	err := Unmarshal(&data, b)
	if assert.Nil(t, err) {
		assert.Equal(t, data, Example{Height: 52.1})
	}
}

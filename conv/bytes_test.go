package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	bytes := []byte("hello, world")
	assert.Equal(t, "hello, world", Bytes2String(bytes))
	bytes[0] = '0'
	assert.Equal(t, "0ello, world", Bytes2String(bytes))
}

func TestStringToBytes(t *testing.T) {
	str := "hello, world"
	data := String2Bytes(str)
	assert.Equal(t, []byte("hello, world"), data)
}

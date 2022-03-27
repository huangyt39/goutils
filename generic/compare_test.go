package generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	t.Run("test int", func(t *testing.T) {
		var a, b int = 1, 2
		res := Min(a, b)
		assert.Equal(t, 1, res)
	})

	t.Run("test int32", func(t *testing.T) {
		var a, b int32 = 1, 2
		res := Min(a, b)
		assert.Equal(t, int32(1), res)
	})

	t.Run("test int64", func(t *testing.T) {
		var a, b int64 = 1, 2
		res := Min(a, b)
		assert.Equal(t, int64(1), res)
	})

	t.Run("test float32", func(t *testing.T) {
		var a, b float32 = 1.1, 2
		res := Min(a, b)
		assert.Equal(t, float32(1.1), res)
	})

	t.Run("test float64", func(t *testing.T) {
		var a, b float64 = 1.1, 2
		res := Min(a, b)
		assert.Equal(t, float64(1.1), res)
	})

	t.Run("test ~float64", func(t *testing.T) {
		type tmp float64
		var a, b tmp = 1.1, 2
		res := Min(a, b)
		assert.Equal(t, tmp(1.1), res)
	})
}

func TestSumAsInt(t *testing.T) {
	t.Run("test string", func(t *testing.T) {
		var a, b string = "1", "2"
		res, err := SumAsInt(a, b)
		assert.NoError(t, err)
		assert.Equal(t, "3", res)
	})

	t.Run("test ~string", func(t *testing.T) {
		type tmp string
		var a, b tmp = "1", "2"
		res, err := SumAsInt(a, b)
		assert.NoError(t, err)
		assert.Equal(t, tmp("3"), res)
	})

	t.Run("parseInt err", func(t *testing.T) {
		var a, b string = "1.44xx", "2.xx"
		_, err := SumAsInt(a, b)
		assert.NotNil(t, err)
	})
}

func TestSumAsFloat(t *testing.T) {
	t.Run("test string", func(t *testing.T) {
		var a, b string = "1.1", "2.2"
		res, err := SumAsFloat(a, b, 3)
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%0.3f", 3.3), res)
	})

	t.Run("test ~string", func(t *testing.T) {
		type tmp string
		var a, b tmp = "1.1", "2.2"
		res, err := SumAsFloat(a, b, 3)
		assert.NoError(t, err)
		assert.Equal(t, tmp(fmt.Sprintf("%0.3f", 3.3)), res)
	})
	t.Run("parseFloat err", func(t *testing.T) {
		var a, b string = "1xx.44xx", "2.xx"
		_, err := SumAsFloat(a, b, 3)
		assert.NotNil(t, err)
	})
}

package pool

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	t.Run("get & timeout", func(t *testing.T) {
		p := NewPool(2, 5*time.Second)
		c1, err := p.Get()
		assert.NotNil(t, c1)
		assert.NoError(t, err)
		c2, err := p.Get()
		assert.NotNil(t, c2)
		assert.NoError(t, err)
		c3, err := p.Get()
		assert.Nil(t, c3)
		assert.NotNil(t, err)
	})
	t.Run("release", func(t *testing.T) {
		p := NewPool(2, 5*time.Second)
		c1, err := p.Get()
		assert.NotNil(t, c1)
		assert.NoError(t, err)
		c2, err := p.Get()
		assert.NotNil(t, c2)
		assert.NoError(t, err)
		err = p.Release(c2)
		assert.NoError(t, err)
		c3, err := p.Get()
		assert.NotNil(t, c3)
		assert.NoError(t, err)
		err = p.Release(c1)
		assert.NoError(t, err)
		err = p.Release(c3)
		assert.NoError(t, err)
		err = p.Release(&http.Client{})
		assert.NotNil(t, err)
	})
}

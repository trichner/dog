package rand

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewRandom(t *testing.T) {

	seed := []byte{9, 2, 1}
	r1, e1 := NewRandom(seed)
	assert.NoError(t, e1)

	assert.GreaterOrEqual(t, r1.Float64(), float64(0))
	assert.Less(t, r1.Float64(), float64(1))
}

func TestNewRandom_shortSeed(t *testing.T) {

	seed := []byte{1}
	r1, e1 := NewRandom(seed)
	assert.NoError(t, e1)

	r1.Int()
}

func TestNewRandomSource(t *testing.T) {

	r1, e1 := NewRandomSource([]byte{1, 2, 3})
	assert.NoError(t, e1)
	assert.Implements(t, (*rand.Source64)(nil), r1)
}

func TestNewRandomSource_deterministic(t *testing.T) {

	seed := []byte{42, 42, 42, 42, 42}

	r1, e1 := NewRandomSource(seed)
	assert.NoError(t, e1)

	r2, e2 := NewRandomSource(seed)
	assert.NoError(t, e2)

	for i := 0; i < 128; i++ {
		assert.Equal(t, r1.Int63(), r2.Int63())
	}
}

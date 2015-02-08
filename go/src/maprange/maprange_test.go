package maprange

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMapLinear(t *testing.T) {
	r, err := MapLinear(1.0, 0.0, 1.0, 0.0, 100.0)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, r)

	r, err = MapLinear(0.5, 0.0, 1.0, 0.0, 100.0)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, r)

	r, err = MapLinear(0.0, 0.0, 1.0, 10.0, 100.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, r)

	r, err = MapLinear(0.52, 0.0, 1.0, 0.0, 100.0)
	assert.Nil(t, err)
	assert.Equal(t, 52.0, r)

	r, err = MapLinear(0.52, 0.0, 1.0, 1.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, 1.52, r)
}

func TestMapLinearOutOfRange(t *testing.T) {
	r, err := MapLinear(2.0, 0.0, 1.0, 0.0, 100.0)
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, r)

	r, err = MapLinear(-0.5, 0.0, 1.0, 0.0, 100.0)
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, r)
}
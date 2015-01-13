package isbn

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValidISBN10(t *testing.T) {
	r, err := IsValidISBN10("3453270010")
	assert.True(t, r)
	assert.Nil(t, err)

	r, err = IsValidISBN10("3453270011")
	assert.False(t, r)
	assert.Nil(t, err)

	r, err = IsValidISBN10("0747532699")
	assert.True(t, r)
	assert.Nil(t, err)
}

func TestIsValidISBN10_BadFormatting(t *testing.T) {
	r, err := IsValidISBN10("074753269s")
	assert.False(t, r)
	assert.NotNil(t, err)

	r, err = IsValidISBN10("0-747-532-691")
	assert.False(t, r)
	assert.NotNil(t, err)

	r, err = IsValidISBN10("074753269")
	assert.False(t, r)
	assert.NotNil(t, err)
}
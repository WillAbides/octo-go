package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_higherValidSemver(t *testing.T) {
	got, err := higherValidSemver("v0.1.0", "v0.2.0")
	assert.NoError(t, err)
	assert.Equal(t, "v0.2.0", got)
	got, err = higherValidSemver("v0.2.0", "v0.2.0")
	assert.NoError(t, err)
	assert.Equal(t, "v0.2.0", got)
	got, err = higherValidSemver("v0.2.0", "v0.1.0")
	assert.NoError(t, err)
	assert.Equal(t, "v0.2.0", got)
	got, err = higherValidSemver("v0.2.0", "hi")
	assert.NoError(t, err)
	assert.Equal(t, "v0.2.0", got)
	got, err = higherValidSemver("hi", "v0.2.0")
	assert.NoError(t, err)
	assert.Equal(t, "v0.2.0", got)
	got, err = higherValidSemver("hi", "hi")
	assert.Error(t, err)
	assert.Equal(t, "", got)
}

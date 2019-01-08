package vat_test

import (
	"testing"

	"github.com/goenning/vat"
	"github.com/stretchr/testify/assert"
)

func TestIsEU(t *testing.T) {
	assert.True(t, vat.IsEU("DE"))
	assert.True(t, vat.IsEU("GR"))
	assert.True(t, vat.IsEU("GB"))
	assert.True(t, vat.IsEU("IE"))
	assert.True(t, vat.IsEU("PT"))

	assert.False(t, vat.IsEU("EL"))
	assert.False(t, vat.IsEU("UK"))
	assert.False(t, vat.IsEU("US"))
	assert.False(t, vat.IsEU("BR"))
	// assert.False(t, vat.IsEU("CH"))
}

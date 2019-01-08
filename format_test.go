package vat_test

import (
	"testing"

	"github.com/goenning/vat"
	"github.com/stretchr/testify/assert"
)

func TestValidateNumberFormat_Valid(t *testing.T) {
	for _, testCase := range validVATNumberFormat {
		valid, cc := vat.ValidateNumberFormat(testCase.VATNumber)
		assert.True(t, valid)
		assert.Equal(t, cc, testCase.CountryCode)
	}
}

func TestValidateNumberFormat_Invalid(t *testing.T) {
	for _, testCase := range invalidVATNumberFormat {
		valid, cc := vat.ValidateNumberFormat(testCase.VATNumber)
		assert.False(t, valid)
		assert.Equal(t, testCase.CountryCode, cc)
	}
}

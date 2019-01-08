package vat_test

import (
	"testing"

	"github.com/goenning/vat"
	"github.com/stretchr/testify/assert"
)

func TestQuery_ExistentNumbers(t *testing.T) {
	for _, testCase := range existentVATNumbers {
		response, err := vat.Query(testCase.VATNumber)
		assert.Nil(t, err)
		if err == nil {
			assert.Equal(t, testCase.CountryCode, response.CountryCode)
			assert.Equal(t, testCase.Address, response.Address)
			assert.Equal(t, testCase.Name, response.Name)
			assert.True(t, response.IsValid)
		}
	}
}

func TestQuery_InvalidNumberFormat(t *testing.T) {
	for _, testCase := range invalidVATNumberFormat {
		response, err := vat.Query(testCase.VATNumber)
		assert.Nil(t, response)
		assert.Equal(t, vat.ErrInvalidVATNumberFormat, err)
	}
}

package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	LZ77 "github.com/fbonhomm/LZ77/source"
)

func Test_decompress_simple_string_without_duplicate(t *testing.T) {
	var Lz77 = LZ77.Init()
	rawData := Lz77.Decompression(ComSimpleStringWithoutDuplicate)

	assert.Equal(t, len(DecSimpleStringWithoutDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithoutDuplicate, rawData)
}

func Test_decompress_simple_string_with_duplicate(t *testing.T) {
	var Lz77 = LZ77.Init()
	rawData := Lz77.Decompression(ComSimpleStringWithDuplicate)

	assert.Equal(t, len(DecSimpleStringWithDuplicate), len(rawData))
	assert.Equal(t, DecSimpleStringWithDuplicate, rawData)
}
package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	LZ77 "github.com/fbonhomm/LZ77/source"
)

func Test_compress_simple_string_without_duplicate(t *testing.T) {
	var Lz77 = LZ77.Init()
	compressData := Lz77.Compression(DecSimpleStringWithoutDuplicate)

	assert.Equal(t, len(ComSimpleStringWithoutDuplicate), len(compressData))
	assert.Equal(t, ComSimpleStringWithoutDuplicate, compressData)
}

func Test_compress_simple_string_with_duplicate(t *testing.T) {
	var Lz77 = LZ77.Init()
	compressData := Lz77.Compression(DecSimpleStringWithDuplicate)

	assert.Equal(t, len(ComSimpleStringWithDuplicate), len(compressData))
	assert.Equal(t, ComSimpleStringWithDuplicate, compressData)
}
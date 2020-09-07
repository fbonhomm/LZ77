package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	LZ77 "github.com/fbonhomm/LZ77/source"
)

func Test_compress_medium_string(t *testing.T) {
	var Lz77 = LZ77.Init()
	compressData := Lz77.Compression(DecMediumString)

	assert.Equal(t, len(ComMediumString), len(compressData))
	assert.Equal(t, ComMediumString, compressData)
}
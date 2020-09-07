package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	LZ77 "github.com/fbonhomm/LZ77/source"
)

func Test_decompress_medium_string(t *testing.T) {
	var Lz77 = LZ77.Init()
	rawData := Lz77.Decompression(ComMediumString)

	assert.Equal(t, len(DecMediumString), len(rawData))
	assert.Equal(t, DecMediumString, rawData)
}
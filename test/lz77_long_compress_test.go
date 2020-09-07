package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"

	LZ77 "github.com/fbonhomm/LZ77/source"
)

func Test_compress_text(t *testing.T) {
	var Lz77 = LZ77.Init()

	file, err := ioutil.ReadFile("./miserables.txt")

	if err != nil {
		log.Fatal(err)
	}

	compressData := Lz77.Compression(file)
	rawData := Lz77.Decompression(compressData)

	assert.Equal(t, file, rawData)
}
package LZ77

import (
	"bytes"
	"math"
)

func (c *LZ77) match(search, ahead []byte) (position, length int) {
	searchLength := len(search)

	for length = len(ahead); length > 0; length-- {
		position = bytes.LastIndex(search, ahead[:length])

		if position != -1 {
			return searchLength - position, length
		}
	}
	return 0, 0
}

func (c *LZ77) encode(compressData []byte, position, length int, nexCharacter byte) []byte {
	var buf = make([]byte, 3)

	buf[0] = byte(position & 0x00FF)
	buf[1] = byte((position & 0x0F00) >> 8) | byte((length & 0x000F) << 4)
	buf[2] = nexCharacter

	return append(compressData, buf...)
}

func (c *LZ77) Compression(rawData []byte) []byte {
	var rawDataLength = len(rawData)
	var compressData []byte
	var nextCharacter byte

	for idx := 0; idx < rawDataLength; idx++ {
		// whether idx is least than search buffer then 0
		searchIndex := int(math.Max(float64(idx - c.Position), 0))
		// whether idx + look ahead buffer is greater than rawDataLength
		aheadIndex := int(math.Min(float64(idx + c.Length), float64(rawDataLength)))

		position, length := c.match(rawData[searchIndex:idx], rawData[idx:aheadIndex])

		if idx + length >= rawDataLength {
			nextCharacter = 0
		} else {
			nextCharacter = rawData[idx + length]
		}

		compressData = c.encode(compressData, position, length, nextCharacter)

		idx += length
	}

	return compressData
}

package LZ77

func (c *LZ77) decode(buffer []byte, chunk []byte) ([]byte, byte) {
	bufferLength := len(buffer)
	tmp := PutByteToUint32(chunk)

	position := bufferLength - int(tmp & 0x000FFF)
	length := position + int((tmp & 0x00F000) >> 12)
	nextCharacter := byte((tmp & 0xFF0000) >> 16)

	return buffer[position:length], nextCharacter
}

func (c *LZ77) Decompression(compressData []byte) []byte {
	var compressDataLength = len(compressData)
	var rawData []byte

	for idx := 0; idx < compressDataLength; idx += 3 {
		chunk := compressData[idx:(idx+3)]

		data, nextCharacter := c.decode(rawData, chunk)

		rawData = append(rawData, data...)
		rawData = append(rawData, nextCharacter)
	}

	if len(rawData) > 0 && rawData[len(rawData)-1] == 0 {
		return rawData[:len(rawData)-1]
	}
	return rawData
}

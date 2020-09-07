package LZ77

// PutByteToUint32 convert []byte to uint32
func PutByteToUint32(data []byte) uint32 {
	var tmp uint32

	for idx := 0; idx < len(data); idx++ {
		tmp |= uint32(data[idx]) << (8 * idx)
	}

	return tmp
}
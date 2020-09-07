package LZ77

// LZ77 config struct
type LZ77 struct {
	Position		int
	Length			int
}

// Init initialize the config
func Init() LZ77 {
	return LZ77{
		Position: 4095, // 12 bits
		Length: 15, // 4 bits
	}
}
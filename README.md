[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

# LZ77
[EN] Implementation of the LZ77 compression algorithm

[FR] Implémentation de l'algorithme de compression LZ77

## Explanation
 - [English](documentation/explanation.en.md)
 - [Français](documentation/explanation.fr.md)

## Technology
* go [v1.14](https://golang.org/)

## Usage
CLI:
```bash
go test -v ./test
```

CODE:
```go
import "github.com/fbonhomm/LZ77/source"

var lz77 = LZ77.init()

dataCompress := lz77.Compression([]byte("..."))

rawData := lz77.Decompression(dataCompress)
```

## Links
* https://en.wikipedia.org/wiki/LZ77_and_LZ78

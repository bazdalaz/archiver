package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunk string
type BinaryChunks []BinaryChunk

const chunkSize = 8



// Join joins binary chunks to string
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range bcs {
		buf.WriteString(string(chunk))
	}

	return buf.String()
}


// splitByChunks splits string to chunks of given size
// i.g. 010010101010101001010010 -> 01001010 10101010 01010010
func splitByChunks(str string, chunkSize int) BinaryChunks {
	srtLen := utf8.RuneCountInString(str)
	chunksCount := srtLen / chunkSize

	if srtLen%chunkSize != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)
	var buf strings.Builder

	for i, ch := range str {
		buf.WriteString(string(ch))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		lastChank := buf.String()
		lastChank += strings.Repeat("0", chunkSize-len(lastChank))

		res = append(res, BinaryChunk(lastChank))
	}
	return res
}

func NeBinChunks(data []byte) BinaryChunks {

	res := make(BinaryChunks, 0, len(data))

	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}	

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}

	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, 8)

	if err != nil {
		panic(err)
	}

	return byte(num)
}
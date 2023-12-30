package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunk string
type BinaryChunks []BinaryChunk
type hexChunk string
type hexChunks []hexChunk

const chunkSize = 8
const hexChunkSeparator = " "

// ToString converts hexChunks to string
func (hcs hexChunks) ToString() string {
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	default:
		var buf strings.Builder

		buf.WriteString(string(hcs[0]))

		for _, ch := range hcs[1:] {
			buf.WriteString(hexChunkSeparator)
			buf.WriteString(string(ch))
		}
		return buf.String()
	}
}

func (bcs BinaryChunks) ToHex() hexChunks {
	res := make(hexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()
		res = append(res, hexChunk)
	}

	return res

}

func (hcs hexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcs))

	for _, chunk := range hcs {
		binaryChunk := chunk.ToBinary()
		res = append(res, binaryChunk)
	}

	return res
}

// Join joins binary chunks to string
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range bcs {
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func (hc hexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hc), 16, 64) //base 16, 64 bit - chunck size
	if err != nil {
		panic("can't parse hex chunk: " + err.Error())
	}

	res := fmt.Sprintf("%08b", num) //base 2, 8 bit - chunck size

	return BinaryChunk(res)
}


func (bc BinaryChunk) ToHex() hexChunk {
	num, err := strconv.ParseUint(string(bc), 2, 64) //base 2, 64 bit - chunck size
	if err != nil {
		panic(err)
	}

	res := strings.ToUpper(fmt.Sprintf("%X", num))
	if len(res) == 1 {
		res = "0" + res
	}

	return hexChunk(res)
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

func NewHexChunks(str string) hexChunks {
	parts := strings.Split(str, hexChunkSeparator)

	res := make(hexChunks, 0, len(parts))

	for _, part := range parts {
		res = append(res, hexChunk(part))
	}
	return res
}

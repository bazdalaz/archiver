package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunk string
type BinarryChunks []BinaryChunk
type hexChunk string
type hexChunks []hexChunk

const chunkSize = 8

// Encode encodes string to hex implemented by vlc algorithm
func Encode(str string) string {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 0100101001
	bStr := encodeBin(str)

	// split to chuncks (8) bits to bytes: 01001010 10101010 01010010
	chunks := splitByChunks(bStr, chunkSize)


	//return bytes to hex: 0100111 10101011 10100010 -> 4D AB A2
	 return chunks.ToHex().ToString()
	
}
// ToString converts hexChunks to string
func (hcs hexChunks) ToString() string {
	// 20 30 3C
	const sep = " "
	switch len(hcs) {
	case 0:
		return ""
	case 1:
		return string(hcs[0])
	default:
		var buf strings.Builder

		buf.WriteString(string(hcs[0]))

		for _, ch := range hcs[1:] {
			buf.WriteString(sep)
			buf.WriteString(string(ch))
		}
		return buf.String()
	}
}

func (bcs BinarryChunks) ToHex() hexChunks {
	res := make(hexChunks, 0, len(bcs))

	for _, chunk := range bcs {
		hexChunk := chunk.ToHex()
		res = append(res, hexChunk)
	}

	return res
	
}

func (bc BinaryChunk) ToHex() hexChunk {
	num, err := strconv.ParseUint(string(bc), 2, 64) //base 2, 64 bit - chunck size
	if err != nil {
		panic(err)
	}

	res:= strings.ToUpper(fmt.Sprintf("%X", num))
	if len(res) == 1 {
		res = "0" + res
	}

	return hexChunk(res)
}

// prepareText prepares text for encoding:
// changes upper case to ! + lower case letter
// i.g. My name is Ted -> !my name i!s !ted
func prepareText(str string) string {
	var buf strings.Builder //builder is faster than concatenation with + or += because it allocates memory only once

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

// encodeBin encodes text to binary
// i.g. some text -> 0100101001
func encodeBin(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch))
	}

	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character " + string(ch))
	}

	return res

}

func getEncodingTable() encodingTable {
return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

//splitByChunks splits string to chunks of given size
// i.g. 010010101010101001010010 -> 01001010 10101010 01010010
func splitByChunks(str string, chunkSize int) BinarryChunks {
	srtLen := utf8.RuneCountInString(str)
	chunksCount := srtLen / chunkSize

	if srtLen % chunkSize != 0 {
		chunksCount++
	}
	res := make(BinarryChunks, 0, chunksCount)
	var buf strings.Builder

	for i, ch := range str {
		buf.WriteString(string(ch))
		
		if (i + 1) % chunkSize == 0  {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		lastChank := buf.String()
		lastChank += strings.Repeat("0", chunkSize - len(lastChank))

		res = append(res, BinaryChunk(lastChank))
	}
	return res
}
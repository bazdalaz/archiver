package vlc

import (
	"strings"
	"unicode"
)

type EncoderDecoder struct {}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

type encodingTable map[rune]string

// Encode encodes string to hex implemented by vlc algorithm
func (vlc EncoderDecoder) Encode(str string) []byte {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 0100101001
	bStr := encodeBin(str)

	// split to chuncks (8) bits to bytes: 01001010 10101010 01010010
	chunks := splitByChunks(bStr, chunkSize)

	//return bytes to hex: 0100111 10101011 10100010 -> 4D AB A2
	return chunks.Bytes()

}

// Decode decodes string from hex implemented by vlc algorithm
func (vlc EncoderDecoder) Decode(encodedData []byte) string {

	bStr := NeBinChunks(encodedData).Join()
	dTree := getEncodingTable().DecodingTree()

	return exportText(dTree.Decode(bStr))

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

// exportText is opposite to prepareText, it changes ! + lower case letter to upper case
func exportText(str string) string {
	var buf strings.Builder

	var isCapital bool

	for _, ch := range str {
		if isCapital {
			buf.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}

		if ch == '!' {
			isCapital = true

			continue
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

package vlc

import "strings"

func Encode(str string) string {
	// prepare text: M -> !m

	// encode to binary: some text -> 0100101001

	// split to chuncks (8) bits to bytes: 01001010 10101010 01010010

	//bytes to hex: 0100111 10101011 10100010 -> 4D AB A2

	//return hexChunksStr

	return ""
}

// prepareText prepares text for encoding:
// changes upper case to ! + lower case letter
// i.g. My name is Ted -> !my name i!s !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			buf.WriteRune('!')
			buf.WriteRune(r + 32)
		} else {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}
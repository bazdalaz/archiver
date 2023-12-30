package vlc

import (
	"testing"
)

func Test_prepareText(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "empty string",
			str:  "",
			want: "",
		},
		{
			name: "only upper case",
			str:  "MY NAME IS TED",
			want: "!m!y !n!a!m!e !i!s !t!e!d",
		},
		{
			name: "only lower case",
			str:  "my name is ted",
			want: "my name is ted",
		},
		{
			name: "only special characters",
			str:  "!@#$%^&*()_+",
			want: "!@#$%^&*()_+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!ted",
			want: "001000100110100101",
		},
		{
			name: "empty string",
			str:  "",
			want: "",
		},
		{
			name: "only upper case",
			str:  "!m!y !n!a!m!e !i!s !t!e!d",
			want: "001000000011001000000000111001000100000010000110010000000110010001011100100001001001000010111001000100100100010100100000101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {

	tests := []struct {
		name string
		encodedText string
		want string
	}{
		{
			name: "base test",
			encodedText: "20 30 3C 18 77 4A E4 4D 28",
			want: "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.encodedText); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

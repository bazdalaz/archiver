package vlc

import (
	"reflect"
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
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32 ,48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New()
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {

	tests := []struct {
		name string
		encodedData []byte
		want string
	}{
		{
			name: "base test",
			encodedData: []byte{32 ,48, 60, 24, 119, 74, 228, 77, 40},
			want: "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New()
			if got := decoder.Decode(tt.encodedData); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

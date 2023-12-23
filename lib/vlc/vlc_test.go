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

func Test_splitByChunks(t *testing.T) {
	type args struct {
		str       string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinarryChunks
	}{

		{
			name: "base test",
			args: args{
				str:       "001000100110100101",
				chunkSize: 8,
			},
			want: BinarryChunks{"00100010", "01101001", "01000000"},
		},
		{
			name: "empty string",
			args: args{
				str:       "",
				chunkSize: 8,
			},
			want: BinarryChunks{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.str, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinarryChunks
		want hexChunks
	}{
		{
			name: "base test",
			bcs:  BinarryChunks{"00100010", "01101001", "01000000"},
			want: hexChunks{"22", "69", "40"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarryChunks.ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexChunks_ToString(t *testing.T) {
	tests := []struct {
		name string
		hcs  hexChunks
		want string
	}{

		{
			name: "base test",
			hcs:  hexChunks{"22", "69", "40"},
			want: "22 69 40",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToString(); got != tt.want {
				t.Errorf("hexChunks.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str string
		want string
	}{
		{
			name: "base test",
			str: "My name is Ted",
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

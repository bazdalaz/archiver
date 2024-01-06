package vlc

import (
	"reflect"
	"testing"

	"github.com/bazdalaz/archiver/lib/compression/vlc/table"
	"github.com/bazdalaz/archiver/lib/compression/vlc/table/shannon_fano"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := New(shannon_fano.Generator{})
			if got := encoder.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {

	tests := []struct {
		name        string
		encodedData []byte
		want        string
	}{
		{
			name:        "base test",
			encodedData: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want:        "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New(shannon_fano.Generator{})
			if got := decoder.Decode(tt.encodedData); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFile(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name  string
		args  args
		want  table.EncodingTable
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseFile(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

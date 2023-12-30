package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		str       string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{

		{
			name: "base test",
			args: args{
				str:       "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
		{
			name: "empty string",
			args: args{
				str:       "",
				chunkSize: 8,
			},
			want: BinaryChunks{},
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

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"00000000", "00100000", "01000000", "00000000"},
			want: "00000000001000000100000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("BinaryChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "base test",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NeBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NeBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

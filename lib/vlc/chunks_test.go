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

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want hexChunks
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"00100010", "01101001", "01000000"},
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

func TestNewHexChunks(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want hexChunks
	}{
		{
			name: "base test",
			str:  "22 69 40",
			want: hexChunks{"22", "69", "40"},
		},
		{
			name: "empty string",
			str:  "",
			want: hexChunks{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   hexChunk
		want BinaryChunk
	}{
		{
			name: "base test",
			hc:   "01",
			want: "00000001",
		},
		{
			name: "different length",
			hc:   "2",
			want: "00000010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); got != tt.want {
				t.Errorf("hexChunk.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcs  hexChunks
		want BinaryChunks
	}{

		{
			name: "base test",
			hcs:  hexChunks{"00", "20", "40", "00"},
			want: BinaryChunks{"00000000", "00100000", "01000000", "00000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hexChunks.ToBinary() = %v, want %v", got, tt.want)
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

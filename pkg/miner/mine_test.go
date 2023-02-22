package miner

import (
	"reflect"
	"testing"
)

func TestMine(t *testing.T) {
	type args struct {
		prefix     []byte
		difficulty []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "successPOWTestMine",
			args: args{
				prefix:     []byte("FARAWAY"),
				difficulty: []byte("3"),
			},
			want: []byte("FARAWAYAPAxk0SCTXMmspn0Hclw"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mine(tt.args.prefix, tt.args.difficulty); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		str        []byte
		complexity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successTestHash",
			args: args{
				str:        []byte("FARAWAYAPAxk0SCTXMmspn0Hclw"),
				complexity: 3,
			},
			want: true,
		}, {
			name: "failTestHash",
			args: args{
				str:        []byte("FARAWAYAPAxk0SCTXMfspn0Hclw"),
				complexity: 3,
			},
			want: false,
		}, {
			name: "failTestHash",
			args: args{
				str:        []byte("asdfsdfdsfdsfsd"),
				complexity: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.str, tt.args.complexity); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

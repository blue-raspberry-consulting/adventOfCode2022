package day2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test input",
			args: args{file: "testdata/test_in.dat"},
			want: 15,
		},
		{
			name: "input",
			args: args{file: "testdata/in.dat"},
			want: 8890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.args.file)
			assert.Equal(t, tt.want, Part1(bytes.NewReader(file)))
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test input",
			args: args{file: "testdata/test_in.dat"},
			want: 12,
		},
		{
			name: "input",
			args: args{file: "testdata/in.dat"},
			want: 10238,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.args.file)
			assert.Equal(t, tt.want, Part2(bytes.NewReader(file)))
		})
	}
}

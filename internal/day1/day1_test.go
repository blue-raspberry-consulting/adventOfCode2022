package day1

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty input => empty slice",
			args: args{in: strings.NewReader("")},
			want: []int{},
		},
		{
			name: "single input => slice of input",
			args: args{in: strings.NewReader("1000")},
			want: []int{1000},
		},
		{
			name: "two inputs => slice of summed values",
			args: args{in: strings.NewReader("1000\n2000")},
			want: []int{3000},
		},
		{
			name: "two inputs separated by blank line => slice of two inputs",
			args: args{in: strings.NewReader("1000\n\n2000")},
			want: []int{1000, 2000},
		},
		{
			name: "four inputs separated by blank line => slice of two inputs summed",
			args: args{in: strings.NewReader("1000\n2000\n\n3000\n4000")},
			want: []int{3000, 7000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseInput(tt.args.in))
		})
	}
}

func Test_maxCalories(t *testing.T) {
	type args struct {
		calories []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil input => 0",
			want: 0,
		},
		{
			name: "empty input => 0",
			args: args{calories: []int{}},
			want: 0,
		},
		{
			name: "single input => value",
			args: args{calories: []int{1000}},
			want: 1000,
		},
		{
			name: "multiple input => max value",
			args: args{calories: []int{1000, 5000, 2000}},
			want: 5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxCalories(tt.args.calories))
		})
	}
}

func Test_topN(t *testing.T) {
	type args struct {
		calories []int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil input => empty",
			args: args{n: 1},
			want: []int{},
		},
		{
			name: "empty input => empty",
			args: args{calories: []int{}, n: 1},
			want: []int{},
		},
		{
			name: "top 0 => empty",
			args: args{calories: []int{1000}},
			want: []int{},
		},
		{
			name: "top 1 => value",
			args: args{calories: []int{1000}, n: 1},
			want: []int{1000},
		},
		{
			name: "top 3 => top 3 values",
			args: args{calories: []int{1000, 100, 5000, 2000, 200}, n: 3},
			want: []int{5000, 2000, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, topN(tt.args.calories, tt.args.n))
		})
	}
}

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
			want: 24000,
		},
		{
			name: "input",
			args: args{file: "testdata/in.dat"},
			want: 70509,
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
			want: 45000,
		},
		{
			name: "input",
			args: args{file: "testdata/in.dat"},
			want: 208567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.ReadFile(tt.args.file)
			assert.Equal(t, tt.want, Part2(bytes.NewReader(file)))
		})
	}
}

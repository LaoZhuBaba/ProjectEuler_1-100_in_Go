package challenge93

import (
	"reflect"
	"testing"
)

func Test_rp(t *testing.T) {
	type args struct {
		r []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2*((3*4)-1)=22",
			args: args{
				r: []int{3, 4, MULTIPLY, 1, SUBTRACT, 2, MULTIPLY},
			},
			want: 22.0,
		},
		{
			name: "((2+3)*4)-1)=22",
			args: args{
				r: []int{2, 3, ADD, 4, MULTIPLY, 1, SUBTRACT},
			},
			want: 19.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rp(tt.args.r); got != tt.want {
				t.Errorf("rp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rpGroupings(t *testing.T) {
	type args struct {
		numbers   []int
		operators []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []float64
	}{
		{
			name: "1234+-*",
			args: args{
				numbers:   []int{1, 2, 3, 4},
				operators: []int{ADD, SUBTRACT, MULTIPLY},
			},
			wantRet: []float64{1},
		},
		{
			name: "8234/-*",
			args: args{
				numbers:   []int{8, 2, 3, 4},
				operators: []int{DIVIDE, SUBTRACT, MULTIPLY},
			},
			wantRet: []float64{10, 4},
		},
		{
			name: "8264/*+",
			args: args{
				numbers:   []int{8, 2, 6, 4},
				operators: []int{DIVIDE, MULTIPLY, ADD},
			},
			wantRet: []float64{11, 28, 28},
		},
		{
			name: "8264+-+",
			args: args{
				numbers:   []int{8, 2, 6, 4},
				operators: []int{ADD, SUBTRACT, ADD},
			},
			wantRet: []float64{4, 12, 12, 8},
		},
		{
			name: "1258-*+",
			args: args{
				numbers:   []int{1, 2, 5, 8},
				operators: []int{SUBTRACT, MULTIPLY, ADD},
			},
			wantRet: []float64{5, 39, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := rpGroupings(tt.args.numbers, tt.args.operators); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("rpGroupings() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_maxConseqFloat64(t *testing.T) {
	type args struct {
		floats []float64
	}
	tests := []struct {
		name    string
		args    args
		wantMax float64
	}{
		{
			name: "12356",
			args: args{
				floats: []float64{1, 2, 3, 5, 6},
			},
			wantMax: 3.0,
		},
		{
			name: "632518",
			args: args{
				floats: []float64{6, 3, 2, 5, 1, 8},
			},
			wantMax: 3.0,
		},
		{
			name: "empty",
			args: args{
				floats: []float64{},
			},
			wantMax: 0,
		},
		{
			name: "04246",
			args: args{
				floats: []float64{0, 4, 2, 4, 6},
			},
			wantMax: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxConseqFloat64(tt.args.floats); gotMax != tt.wantMax {
				t.Errorf("maxConseqFloat64() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

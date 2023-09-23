package challenge98

import (
	"reflect"
	"testing"
)

func Test_countUniqueChars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ended",
			args: args{"ended"},
			want: 3,
		},
		{
			name: "abacus",
			args: args{"abacus"},
			want: 5,
		},
		{
			name: "empty string",
			args: args{""},
			want: 0,
		},
		{
			name: "one character",
			args: args{"a"},
			want: 1,
		},
		{
			name: "only duplicates",
			args: args{"ddddddddddddddddddddd"},
			want: 1,
		},
		{
			name: "only duplicates with spaces",
			args: args{"ddddd  dddddd dddddd dddd"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueChars(tt.args.s); got != tt.want {
				t.Errorf("countUniqueChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wordToSortedString(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{
			name:  "balloon",
			args:  args{"balloon"},
			wantS: "abllnoo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := wordToSortedString(tt.args.word); gotS != tt.wantS {
				t.Errorf("wordToSortedString() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_findAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name  string
		args  args
		wantM map[string][]string
	}{
		{
			name: "repeated letters",
			args: args{
				words: []string{"tool", "host", "empty", "shot", "loot"},
			},
			wantM: map[string][]string{
				"loot": {"tool", "loot"},
				"host": {"host", "shot"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := findAnagrams(tt.args.words); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("findAnagrams() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_makeTransform(t *testing.T) {
	type args struct {
		ana1 []rune
		ana2 []rune
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
		wantErr bool
	}{
		{
			name: "saw was",
			args: args{
				ana1: []rune{'s', 'a', 'w'},
				ana2: []rune{'w', 'a', 's'},
			},
			wantRet: []int{2, 1, 0},
			wantErr: false,
		},
		{
			name: "creation reaction",
			args: args{
				ana1: []rune{'c', 'r', 'e', 'a', 't', 'i', 'o', 'n'},
				ana2: []rune{'r', 'e', 'a', 'c', 't', 'i', 'o', 'n'},
			},
			wantRet: []int{3, 0, 1, 2, 4, 5, 6, 7},
			wantErr: false,
		},
		{
			name: "saw wase",
			args: args{
				ana1: []rune{'s', 'a', 'w'},
				ana2: []rune{'w', 'a', 's', 'e'},
			},
			wantRet: nil,
			wantErr: true,
		},
		{
			name: "saw wad",
			args: args{
				ana1: []rune{'s', 'a', 'w'},
				ana2: []rune{'w', 'a', 'd'},
			},
			wantRet: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := makeTransform(tt.args.ana1, tt.args.ana2)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeTransform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("makeTransform() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_getSquaresByNumDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		{
			name:    "min max 2 digit squares",
			args:    args{2},
			wantMin: 4,
			wantMax: 9,
		},
		{
			name:    "min max 3 digit squares",
			args:    args{3},
			wantMin: 10,
			wantMax: 31,
		},
		{
			name:    "min max 4 digit squares",
			args:    args{4},
			wantMin: 32,
			wantMax: 99,
		},
		{
			name:    "min max 5 digit squares",
			args:    args{5},
			wantMin: 100,
			wantMax: 316,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := getSquaresByNumDigits(tt.args.n)
			if gotMin != tt.wantMin {
				t.Errorf("getSquaresByNumDigits() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("getSquaresByNumDigits() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_intToSortedString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name        string
		args        args
		wantS       string
		wantNumUniq int
	}{
		{
			name:        "4824->2448",
			args:        args{4824},
			wantS:       "2448",
			wantNumUniq: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotNumUniq := intToSortedString(tt.args.n)
			if gotS != tt.wantS {
				t.Errorf("intToSortedString() gotS = %v, want %v", gotS, tt.wantS)
			}
			if gotNumUniq != tt.wantNumUniq {
				t.Errorf("intToSortedString() gotNumUniq = %v, want %v", gotNumUniq, tt.wantNumUniq)
			}
		})
	}
}

func Test_makeTransformForSquares(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
		wantErr bool
	}{
		{
			name:    "1296-9216",
			args:    args{1296, 9216},
			wantRet: []int{2, 1, 0, 3},
			wantErr: false,
		},
		{
			name:    "9612-1296",
			args:    args{9612, 1296},
			wantRet: []int{2, 3, 0, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := makeTransformForSquares(tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeTransformForSquares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("makeTransformForSquares() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_intSliceToString(t *testing.T) {
	type args struct {
		is []int
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{
			name:  "simple",
			args:  args{[]int{2, 3, 5, 4}},
			wantS: "2354",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := intSliceToString(tt.args.is); gotS != tt.wantS {
				t.Errorf("intSliceToString() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_getPairsInSet(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantPairs [][2]int
	}{
		{
			name:      "2 from 2",
			args:      args{2},
			wantPairs: [][2]int{{0, 1}, {1, 0}},
		},
		{
			name:      "2 from 3",
			args:      args{3},
			wantPairs: [][2]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPairs := getPairsInSet(tt.args.n); !reflect.DeepEqual(gotPairs, tt.wantPairs) {
				t.Errorf("getPairsInSet() = %v, want %v", gotPairs, tt.wantPairs)
			}
		})
	}
}

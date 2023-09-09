package shared

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// allSubsets will be a list of all possible ordered subsets.  E.g.,
// if max is 3 then all Subsets will be:
// [
//  [[0],[0],[2]],
//  [[0,1],[0,2],[1,0],[1,2],[2,0],[2,1]],
//  [[0,1,2],[0,2,1],[1,0,2],[1,2,0],[2,0,1],[2,1,0]]
// ]

func TestAllOrderedCombinations(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{
			name: "max2",
			args: args{
				max: 2,
			},
			wantResult: [][]int{{0}, {1}},
		},
		{
			name: "max3",
			args: args{
				max: 3,
			},
			wantResult: [][]int{{0, 0}, {1, 1}, {2, 2}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}},
		},
		{
			name: "max4",
			args: args{
				max: 4,
			},
			wantResult: [][]int{
				{0, 0, 0}, {1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {0, 0, 1}, {0, 0, 2},
				{0, 0, 3}, {1, 1, 0}, {1, 1, 2}, {1, 1, 3}, {2, 2, 0}, {2, 2, 1},
				{2, 2, 3}, {3, 3, 0}, {3, 3, 1}, {3, 3, 2}, {0, 1, 1}, {0, 2, 2},
				{0, 3, 3}, {1, 0, 0}, {1, 2, 2}, {1, 3, 3}, {2, 0, 0}, {2, 1, 1},
				{2, 3, 3}, {3, 0, 0}, {3, 1, 1}, {3, 2, 2}, {0, 1, 2}, {0, 1, 3},
				{0, 2, 1}, {0, 2, 3}, {0, 3, 1}, {0, 3, 2}, {1, 0, 2}, {1, 0, 3},
				{1, 2, 0}, {1, 2, 3}, {1, 3, 0}, {1, 3, 2}, {2, 0, 1}, {2, 0, 3},
				{2, 1, 0}, {2, 1, 3}, {2, 3, 0}, {2, 3, 1}, {3, 0, 1}, {3, 0, 2},
				{3, 1, 0}, {3, 1, 2}, {3, 2, 0}, {3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AllOrderedCombinations(tt.args.max); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("AllOrderedCombinations() = %#v, want %#v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNMemberSubsets(t *testing.T) {
	type args struct {
		subsetSize int
		setSize    int
	}
	tests := []struct {
		name        string
		args        args
		wantSubsets [][]int
	}{
		{
			name: "subset 1 in set 2",
			args: args{
				subsetSize: 1,
				setSize:    2,
			},
			wantSubsets: [][]int{{0}, {1}},
		},
		{
			name: "subset 2 in set 2",
			args: args{
				subsetSize: 2,
				setSize:    2,
			},
			wantSubsets: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "subset 2 in set 3",
			args: args{
				subsetSize: 2,
				setSize:    3,
			},
			wantSubsets: [][]int{{0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}},
		},
		{
			name: "subset 1 in set 3",
			args: args{
				subsetSize: 1,
				setSize:    3,
			},
			wantSubsets: [][]int{{0}, {1}, {2}},
		},
		{
			name: "subset 3 in set 3",
			args: args{
				subsetSize: 3,
				setSize:    3,
			},
			wantSubsets: [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSubsets := NMemberSubsets(tt.args.subsetSize, tt.args.setSize); !reflect.DeepEqual(gotSubsets, tt.wantSubsets) {
				t.Errorf("NMemberSubsets() = %#v, want %#v", gotSubsets, tt.wantSubsets)
			}
		})
	}
}

func Test_setsWithinSet(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name     string
		args     args
		wantSets [][]int
	}{
		{
			name:     "addends of 2",
			args:     args{2},
			wantSets: [][]int{{1, 1}, {2}},
		},
		{
			name:     "addends of 3",
			args:     args{3},
			wantSets: [][]int{{3}, {2, 1}, {1, 2}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSets := setsWithinSet(tt.args.n); !assert.ElementsMatch(t, gotSets, tt.wantSets) {
				t.Errorf("setsWithinSet() = %v, want %v", gotSets, tt.wantSets)
			}
		})
	}
}

func TestPermutationsWithRepetition(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name     string
		args     args
		wantSets [][]int
	}{
		{
			name:     "k0,2",
			args:     args{0, 2},
			wantSets: [][]int{{}},
		},
		{
			name:     "k1,2",
			args:     args{1, 2},
			wantSets: [][]int{{1}, {0}},
		},
		{
			name:     "k2,2",
			args:     args{2, 2},
			wantSets: [][]int{{0, 0}, {1, 1}, {1, 0}, {0, 1}},
		},
		{
			name:     "k1,3",
			args:     args{1, 3},
			wantSets: [][]int{{0}, {1}, {2}},
		},
		{
			name:     "k2,3",
			args:     args{2, 3},
			wantSets: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "k3,3",
			args: args{3, 3},
			wantSets: [][]int{
				{
					0, 0, 0,
				},
				{
					0, 0, 1,
				},
				{
					0, 0, 2,
				},
				{
					0, 1, 0},
				{
					0, 1, 1,
				},
				{
					0, 1, 2,
				},
				{
					0, 2, 0,
				},
				{
					0, 2, 1,
				},
				{
					0, 2, 2,
				},
				{
					1, 0, 0,
				},
				{
					1, 0, 1,
				},
				{
					1, 0, 2,
				},
				{
					1, 1, 0,
				},
				{
					1, 1, 1,
				},
				{
					1, 1, 2,
				},
				{
					1, 2, 0,
				},
				{
					1, 2, 1,
				},
				{
					1, 2, 2,
				},
				{
					2, 0, 0,
				},
				{
					2, 0, 1,
				},
				{
					2, 0, 2,
				},
				{
					2, 1, 0,
				},
				{
					2, 1, 1,
				},
				{
					2, 1, 2,
				},
				{
					2, 2, 0,
				},
				{
					2, 2, 1,
				},
				{
					2, 2, 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSets := PermutationsWithRepetition(tt.args.k, tt.args.n); !assert.ElementsMatch(t, gotSets, tt.wantSets) {
				t.Errorf("PermutationsWithRepetition() = %#v, want %#v", gotSets, tt.wantSets)
			}
		})
	}
}

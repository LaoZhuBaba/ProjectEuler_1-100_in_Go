package challenge91

import (
	"reflect"
	"testing"
)

func TestVector_simplify(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRv  Vector
		wantErr bool
	}{
		{
			name:    "zero zero returns error",
			fields:  fields{0, 0},
			wantRv:  Vector{0, 0},
			wantErr: true,
		},
		{
			name:    "1 1 returns 1 1",
			fields:  fields{1, 1},
			wantRv:  Vector{1, 1},
			wantErr: false,
		},
		{
			name:    "22 22 returns 1 1",
			fields:  fields{22, 22},
			wantRv:  Vector{1, 1},
			wantErr: false,
		},
		{
			name:    "34 0 returns 1 0",
			fields:  fields{34, 0},
			wantRv:  Vector{1, 0},
			wantErr: false,
		},
		{
			name:    "0 89 returns 0 1",
			fields:  fields{0, 89},
			wantRv:  Vector{0, 1},
			wantErr: false,
		},
		{
			name:    "1 -1 returns 1 -1",
			fields:  fields{1, -1},
			wantRv:  Vector{1, -1},
			wantErr: false,
		},
		{
			name:    "-1 1 returns -1 1",
			fields:  fields{-1, 1},
			wantRv:  Vector{-1, 1},
			wantErr: false,
		},
		{
			name:    "-22 22 returns -1 1",
			fields:  fields{-22, 22},
			wantRv:  Vector{-1, 1},
			wantErr: false,
		},
		{
			name:    "22 -22 returns 1 -1",
			fields:  fields{22, -22},
			wantRv:  Vector{1, -1},
			wantErr: false,
		},
		{
			name:    "-34 0 returns -1 0",
			fields:  fields{-34, 0},
			wantRv:  Vector{-1, 0},
			wantErr: false,
		},
		{
			name:    "0 -89 returns 0 -1",
			fields:  fields{0, -89},
			wantRv:  Vector{0, -1},
			wantErr: false,
		},
		{
			name:    "22 11 returns 2 1",
			fields:  fields{22, 11},
			wantRv:  Vector{2, 1},
			wantErr: false,
		},
		{
			name:    "5 53 returns 5 53",
			fields:  fields{5, 53},
			wantRv:  Vector{5, 53},
			wantErr: false,
		},
		{
			name:    "8 3 returns 8 3",
			fields:  fields{22, 22},
			wantRv:  Vector{1, 1},
			wantErr: false,
		},
		{
			name:    "120 24 returns 5 1",
			fields:  fields{120, 24},
			wantRv:  Vector{5, 1},
			wantErr: false,
		},
		{
			name:    "3 15 returns 1 5",
			fields:  fields{3, 15},
			wantRv:  Vector{1, 5},
			wantErr: false,
		},
		{
			name:    "-22 11 returns -2 1",
			fields:  fields{-22, 11},
			wantRv:  Vector{-2, 1},
			wantErr: false,
		},
		{
			name:    "5 -53 returns 5 -53",
			fields:  fields{5, -53},
			wantRv:  Vector{5, -53},
			wantErr: false,
		},
		{
			name:    "-8 -3 returns -8 -3",
			fields:  fields{-8, -3},
			wantRv:  Vector{-8, -3},
			wantErr: false,
		},
		{
			name:    "120 -24 returns 5 -1",
			fields:  fields{120, -24},
			wantRv:  Vector{5, -1},
			wantErr: false,
		},
		{
			name:    "-3 -15 returns -1 -5",
			fields:  fields{-3, -15},
			wantRv:  Vector{-1, -5},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			gotRv, err := v.simplify()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.simplify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRv, tt.wantRv) {
				t.Errorf("Vector.simplify() = %v, want %v", gotRv, tt.wantRv)
			}
		})
	}
}

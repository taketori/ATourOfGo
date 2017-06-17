package main

import "testing"

func TestVertex_Abs(t *testing.T) {
	tests := []struct {
		name string
		v    Vertex
		want float64
	}{
		// TODO: Add test cases.
		{"1", Vertex{3, 4}, 5},
		{"2", Vertex{0, 4}, 4},
		{"3,", Vertex{3, 0}, 3},
		{"zero", Vertex{0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Abs(); got != tt.want {
				t.Errorf("Vertex.Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVertex_Scale(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		v    *Vertex
		args args
	}{
		// TODO: Add test cases.
		{"", &Vertex{0, 0}, args{0}},
		{"", &Vertex{3, 4}, args{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Scale(tt.args.f)
		})
	}
}

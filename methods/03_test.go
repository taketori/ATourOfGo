package main

import "testing"

func TestMyFloat_Abs(t *testing.T) {
	tests := []struct {
		name string
		f    MyFloat
		want float64
	}{
		// TODO: Add test cases.
		{"plus", 1, 1},
		{"minus", -1, 1},
		{"0", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Abs(); got != tt.want {
				t.Errorf("MyFloat.Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

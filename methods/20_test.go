package main

import "testing"

func TestErrNegativeSqrt_Error(t *testing.T) {
	tests := []struct {
		name string
		e    ErrNegativeSqrt
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("ErrNegativeSqrt.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: args{-1}, want: 0, wantErr: true},
		{name: "", args: args{0}, want: 0, wantErr: false},
		{name: "", args: args{1}, want: 1, wantErr: false},
		{name: "", args: args{2}, want: 1.414213562373095, wantErr: false},
		{name: "", args: args{4}, want: 2, wantErr: false},
		{name: "", args: args{9}, want: 3, wantErr: false},
		{name: "", args: args{16}, want: 4, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sqrt(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sqrt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_rot13(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name  string
		args  args
		wantR byte
	}{
		// TODO: Add test cases.
		{name: "A", args: args{'A'}, wantR: 'N'},
		{name: "N", args: args{'N'}, wantR: 'A'},
		{name: "M", args: args{'M'}, wantR: 'Z'},
		{name: "Z", args: args{'Z'}, wantR: 'M'},
		{name: "a", args: args{'a'}, wantR: 'n'},
		{name: "n", args: args{'n'}, wantR: 'a'},
		{name: "m", args: args{'m'}, wantR: 'z'},
		{name: "z", args: args{'z'}, wantR: 'm'},
		{name: " ", args: args{' '}, wantR: ' '},
		{name: "!", args: args{'!'}, wantR: '!'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := rot13(tt.args.b); gotR != tt.wantR {
				t.Errorf("rot13() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

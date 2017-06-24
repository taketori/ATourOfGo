package main

import "testing"

func TestIPaddr_String(t *testing.T) {
	tests := []struct {
		name string
		ip   IPaddr
		want string
	}{
		// TODO: Add test cases.
		{name: "", ip: IPaddr{0, 0, 0, 0}, want: "0.0.0.0"},
		{name: "", ip: IPaddr{123, 45, 6, 0}, want: "123.45.6.0"},
		{name: "", ip: IPaddr{255, 255, 0, 0}, want: "255.255.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ip.String(); got != tt.want {
				t.Errorf("IPaddr.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

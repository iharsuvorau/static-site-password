package main

import (
	"testing"
)

func TestCalculateHexString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "A",
			args: args{
				s: "hi there",
			},
			want:    "de8fb7f0af8d7d5e3e2abfa2ecc172b21d565effa5c070a65c655b29d0651702",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateHexString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateHexString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateHexString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"fmt"
	"testing"
)

func Test_fact(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fact(tt.args.n)
			fmt.Println(got)
		})
	}
}

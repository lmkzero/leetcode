package main

import "testing"

func Test_changeCoinII(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeCoinII(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("changeCoinII() = %v, want %v", got, tt.want)
			}
		})
	}
}

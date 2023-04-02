package main

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				haystack: "abc",
				needle:   "abcd",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

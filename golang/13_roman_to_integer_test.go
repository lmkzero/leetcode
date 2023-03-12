package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "4",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "5",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

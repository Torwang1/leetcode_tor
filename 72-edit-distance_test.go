package leetcode

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"base 1",
			args{
				word1: "",
				word2: "i",
			},
			1,
		},
		{
			"example 1",
			args{
				word1: "horse",
				word2: "ros",
			},
			3,
		},
		{
			"example 2",
			args{
				word1: "intention",
				word2: "execution",
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

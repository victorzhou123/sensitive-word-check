package main

import "testing"

var c Checker

func init() {
	var err error
	if c, err = newChecker(); err != nil {
		return
	}
}

func Test_checker_Check(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		c    Checker
		args args
		want bool
	}{
		{
			"test1",
			c,
			args{
				"招聘",
			},
			true,
		},
		{
			"test2",
			c,
			args{
				"正常",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Check(tt.args.word); got != tt.want {
				t.Errorf("checker.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

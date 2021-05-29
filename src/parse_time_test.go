package main

import "testing"

func TestParseTime(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{a: "20210529-000449"},
			want: "2021-05-29T09:04:49+09:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTime(tt.args.a)
			if err != nil {
				t.Errorf("parseTime() = err: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("parseTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

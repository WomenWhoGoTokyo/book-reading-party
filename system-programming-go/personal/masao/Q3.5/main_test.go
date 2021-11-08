package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_myCopyN(t *testing.T) {
	type args struct {
		src    io.Reader
		length int
	}
	tests := []struct {
		name     string
		args     args
		wantDest string
	}{
		{
			name: "切り出す長さが正の値",
			args: args{
				src:    strings.NewReader("abcdefg"),
				length: 3,
			},
			wantDest: "abc",
		},
		{
			name: "切り出す長さが0",
			args: args{
				src:    strings.NewReader("abcdefg"),
				length: 0,
			},
			wantDest: "",
		},
		{
			name: "切り出す長さが元の入力よりも大きい値",
			args: args{
				src:    strings.NewReader("abcdefg"),
				length: 8,
			},
			wantDest: "abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dest := &bytes.Buffer{}
			myCopyN(dest, tt.args.src, tt.args.length)
			if gotDest := dest.String(); gotDest != tt.wantDest {
				t.Errorf("myCopyN() = %v, want %v", gotDest, tt.wantDest)
			}
		})
	}
}

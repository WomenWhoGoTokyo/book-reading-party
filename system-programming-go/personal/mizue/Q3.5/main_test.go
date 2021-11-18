package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCopyN(t *testing.T) {

	type input struct {
		dst io.Writer
		src io.Reader
		len int64
	}
	type expected struct {
		dst string
		len int64
		err error
	}
	tests := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name: "正常",
			input: input{
				src: strings.NewReader("abcdefghijklmn"),
				len: 3,
			},
			expected: expected{
				"abc",
				3,
				nil,
			},
		},
		{
			name: "srcがlengthより短い場合",
			input: input{
				src: strings.NewReader("abcdefg"),
				len: 100,
			},
			expected: expected{
				"abcdefg",
				7,
				io.EOF,
			},
		},
		{
			name: "lengthが 0 の場合",
			input: input{
				src: strings.NewReader("abcdefg"),
				len: 0,
			},
			expected: expected{
				"",
				0,
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := new(bytes.Buffer)
			got, err := CopyN(stdout, tt.input.src, tt.input.len)
			if got != tt.expected.len || stdout.String() != tt.expected.dst || err != tt.expected.err {
				t.Errorf("CopyN() = %v %v %v, expected %v %v %v", got, err, stdout.String(), tt.expected.len, tt.expected, tt.expected.err)
			}
		})
	}
}

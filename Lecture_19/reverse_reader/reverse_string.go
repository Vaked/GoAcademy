package reverseReader

import "io"

type ReverseStringReader struct {
	Text string
	done bool
}

func reverseString(str string) string{
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
	   byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
 }

func NewReverseStringReader(input string) *ReverseStringReader {
	input = reverseString(input)
	return &ReverseStringReader{input, false}
}

func (r *ReverseStringReader) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
  	}

  	for i, b := range []byte(r.Text) {
		p[i] = b
  	}

	r.done = true
	return len(r.Text), nil
}
package main

import "io"

type MyReader struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
		data: buffer,
		pos:  0,
	}
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	i := 0
	for r.pos < len(r.data) && i < len(p) {
		p[i] = r.data[r.pos]
		i++
		r.pos++
	}

	if i > 0 {
		return i, nil
	}

	return 0, io.EOF
}

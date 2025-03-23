package main

import "io"

type MyReader struct {
	data []byte
	pos  int
}

type MyWriter struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
		data: buffer,
		pos:  0,
	}
}

func NewWriterToBuffer(buffer []byte) *MyWriter {
	return &MyWriter{
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

func (w *MyWriter) Write(p []byte) (n int, err error) {
	startPos := w.pos
	for w.pos < len(w.data) {
		w.data[w.pos] = p[w.pos]
		w.pos++
	}

	if startPos == w.pos {
		return 0, io.ErrShortWrite
	} else if len(p) > len(w.data)-w.pos {
		return w.pos - startPos, io.ErrShortBuffer
	}

	return w.pos - startPos, nil
}

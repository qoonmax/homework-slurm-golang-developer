Напишите программу, которая имплементирует io.Reader и io.Writer.

Для локальной проверки см. io.Copy

```go
package homework

type MyReader struct {
	data []byte
	pos  int
}

type MyWriter struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	// ваш код
}

func NewWriterToBuffer(buffer []byte) *MyWriter {
	// ваш код
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	// ваш код
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	// ваш код
}
```
Реализуйте маршалинг json в структуру

```go
package homework

type Example struct {
	A int    `json:"a"`
	B string `json:"B,omitempty"`
}

func jsonToStruct(s []byte) (*Example, error) {
	// Ваш код
}
```
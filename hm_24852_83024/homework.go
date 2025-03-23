package homework

import "encoding/json"

type Example struct {
	A int    `json:"a"`
	B string `json:"B,omitempty"`
}

func jsonToStruct(s []byte) (*Example, error) {
	var example Example

	err := json.Unmarshal(s, &example)
	if err != nil {
		return nil, err
	}

	return &example, nil
}

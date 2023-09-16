package presentation

import "github.com/bytedance/sonic"

type JsonSerializer struct{}

// Marshal object data to JSON
func (j *JsonSerializer) Marshal(data interface{}) ([]byte, error) {
	return sonic.Marshal(data)
}

// Unmarshal from JSON to object
func (j *JsonSerializer) Unmarshal(data []byte, model *interface{}) error {
	return sonic.Unmarshal(data, model)
}

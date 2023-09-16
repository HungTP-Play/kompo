package presentation

import "gopkg.in/yaml.v3"

type YamlSerializer struct{}

// Marshal from data object to YAML
func (y *YamlSerializer) Marshal(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

// Unmarshal from YAML to data object
func (y *YamlSerializer) Unmarshal(buffer []byte, data interface{}) error {
	return yaml.Unmarshal(buffer, data)
}

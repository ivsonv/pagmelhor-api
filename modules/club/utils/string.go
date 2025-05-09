package utils

import (
	"encoding/json"
)

func Serialize(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Deserialize(body []byte, v any) error {
	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	return nil
}

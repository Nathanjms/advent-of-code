package sharedcode

import "encoding/json"

func ParseJSON(jsonString string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	return data, err
}

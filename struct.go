package utils

import "encoding/json"

func CopyStructUsingJSON(from interface{}, to interface{}) error {
	fromJson, err := json.Marshal(from)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fromJson, to)
	if err != nil {
		return err
	}
	return nil
}

package utils

import (
	"io/ioutil"
	"os"
)

import "encoding/json"

// ParseJSON parse every json file.
func ParseJSON(e interface{}, filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &e)
}

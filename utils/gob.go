package utils

import (
	"bytes"
	"encoding/gob"
)

func GobEncode(res interface{}) ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(res)
	return result.Bytes(), err
}

func GobDeCode(res []byte, result interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(res))
	err := decoder.Decode(result)
	return err
}

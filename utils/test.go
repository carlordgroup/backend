package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJson(v interface{}) io.Reader {
	if v == nil {
		return nil
	}
	data, _ := json.Marshal(v)
	return bytes.NewReader(data)
}

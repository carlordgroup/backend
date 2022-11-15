package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToJson(v interface{}) io.Reader {
	data, _ := json.Marshal(v)
	return bytes.NewReader(data)
}

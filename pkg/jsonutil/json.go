package jsonutil

import (
  "encoding/json"
)

func DataTransfer(data interface{}, v interface{}) {
  b, _ := json.Marshal(data)
  json.Unmarshal(b, v)
}

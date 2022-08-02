package json

import "encoding/json"

func ToString(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

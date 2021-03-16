package utils

import "encoding/json"

//ToJSON
func ToJSON(v interface{}, indents ...bool) string {
	var dt []byte
	if len(indents) > 0 && indents[0] {
		dt, _ = json.MarshalIndent(v, "", "  ")
	} else {
		dt, _ = json.Marshal(v)
	}
	return string(dt)
}

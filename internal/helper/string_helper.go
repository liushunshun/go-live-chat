package helper

import "encoding/json"

func ToJsonString(object any) []byte {
	jsonbytes, _ := json.Marshal(object)
	return jsonbytes
}

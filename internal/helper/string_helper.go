package helper

import (
	"encoding/json"
	"strconv"
)

func ToJsonString(object any) []byte {
	jsonbytes, _ := json.Marshal(object)
	return jsonbytes
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return i
}

package helper

import "os"

func ReadBytes(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return []byte(content)
}

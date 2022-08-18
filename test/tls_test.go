package test

import (
	"fmt"
	"os"
	"testing"
)

func TestReadPem(t *testing.T) {
	content, err := os.ReadFile("../config/cer.pem")
	if err != nil {
		t.Error(err)
	}
	data := []byte(content)
	fmt.Println(string(data))

}

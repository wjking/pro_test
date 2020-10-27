package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	input := []byte("playerName:万金")
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Print(encodeString)

	decodeString, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		return
	}
	fmt.Print(string(decodeString))
}

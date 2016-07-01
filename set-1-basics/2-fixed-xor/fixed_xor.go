package main

import (
	"fmt"
	"encoding/hex"
	"bytes"
)

func main() {
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	hex1, _ := hex.DecodeString(str1)
	hex2, _ := hex.DecodeString(str2)
	length := len(hex1)
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = hex1[i] ^ hex2[i]
	}

	expectedBytes, _ := hex.DecodeString(expected)

	if bytes.Compare(result, expectedBytes) == 0 {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure.")
	}

	fmt.Println("Expected:", hex.EncodeToString(expectedBytes))
	fmt.Println("Got:", hex.EncodeToString(result))
}

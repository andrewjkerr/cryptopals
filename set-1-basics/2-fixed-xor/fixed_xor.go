package main

import (
	"fmt"
	"encoding/hex"
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

	expectedHex, _ := hex.DecodeString(expected)
	expectedStr := string(expectedHex)
	resultStr := string(result)

	if resultStr == expectedStr {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure.")
	}

	expectedHexBytes, _ := hex.DecodeString(expectedStr)

	fmt.Println("Expected:", expectedHexBytes)
	fmt.Println("Got:", result)
}

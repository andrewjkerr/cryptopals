package main

import (
	"encoding/base64"
	"encoding/hex"
	"os"
)

func main() {
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	
	// Hex to ASCII values
	asciiString, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	// Create new base64 encoder and write encoded string to Stdout
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(asciiString)
	encoder.Close()
}

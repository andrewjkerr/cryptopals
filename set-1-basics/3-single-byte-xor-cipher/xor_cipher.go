package main

import (
    "fmt"
    "encoding/hex"
)

func main() {
    hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    str, _ := hex.DecodeString(hexStr)
    characters := []byte(str)
    numCharacters := len(characters)

    scores := make(map[byte]int)
    for i := 0; i < numCharacters; i++ {
        scores[characters[i]] = scores[characters[i]] + 1
    }

    var max int
    var maxByte byte
    for character, score := range scores {
        if score > max {
            max = score
            maxByte = character
        }
    }

    commonCharactersStr := "aehinorst "
    commonCharacters := []byte(commonCharactersStr)

    englishCharactersStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    englishCharacters := []byte(englishCharactersStr)

    for _, englishCharacter := range englishCharacters {
        xorResult := englishCharacter ^ maxByte

        // Check if we have a common character!
        for _, a := range commonCharacters {
            if a == xorResult {
                fmt.Println("~ potential result ~")

                potentialResult := make([]byte, numCharacters)
                for j := 0; j < numCharacters; j++ {
                    potentialResult[j] = characters[j] ^ englishCharacter
                }
                potentialResultStr := string(potentialResult[:numCharacters])
                fmt.Println(string(englishCharacter), ":", potentialResultStr)

                fmt.Println("~ end potential result ~")
            }
        }
    }
}

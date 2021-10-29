package symmetry

import (
	"fmt"
	"strings"
)

// Encrypt a short string
func Encrypt(m, key string) string {
	var cypher strings.Builder

	for i, v := range m {
		// convert the character to an in between 0 and 126
		charValue := numVal(v)

		// get a int value from the corresponding key character
		keyNumberValue := numValFrmStr(i, key)

		// sum the int values of the message and the key value
		sum := keyNumberValue + charValue

		// Get an int between 0 and 126, then return its character value
		c := cypherMod(sum)
		fmt.Fprintf(&cypher, "%c", c)
	}
	return cypher.String()
}

func numVal(r rune) int {
	return int(r)
}

func numValFrmStr(index int, key string) int {
	r := key[index]
	return numVal(rune(r))
}

func cypherMod(i int) rune {
	if i <= 0 {
		return rune(i + 125)
	}
	r := i % 125
	return rune(r)
}

// Decrypt a short string
func Decrypt(cypher, key string) string {
	var msg strings.Builder
	for i, v := range cypher {
		// calculate the number value of the character
		cypherValue := numVal(v)

		// get the number value of the corresponding character in the key
		keyNumValue := numValFrmStr(i, key)

		// subtract the key number value from the cypher number value
		sub := cypherValue - keyNumValue

		// get the ascii character from the value
		c := cypherMod(sub)
		fmt.Fprintf(&msg, "%c", c)
	}
	return msg.String()
}

package symmetry

import (
	"fmt"
	"strings"
)

func EncrptLongMsg(msg, key string) (cypherMsg string) {
	var cypherBuild strings.Builder
	for j := 0; j < len(msg); j++ {
		if j != 0 && j%5 == 0 {
			chunkMsg := msg[j-5 : j]
			chunkCypher := Encrypt(chunkMsg, key)
			fmt.Fprintf(&cypherBuild, "%s", chunkCypher)
		}
	}

	mod := len(msg) % 5
	if mod != 0 {
		rem := msg[len(msg)-mod:]
		shortCypher := Encrypt(rem, key)
		fmt.Fprintf(&cypherBuild, "%s", shortCypher)
	}

	cypherMsg = cypherBuild.String()
	return
}

// Decrypt a long string
func DecryptLongMsg(cypher, key string) (msg string) {
	var msgBuilder strings.Builder
	for j := 0; j < len(cypher); j++ {
		if j != 0 && j%5 == 0 {
			chunkCypher := cypher[j-5 : j]
			chunkMsg := Decrypt(chunkCypher, key)
			fmt.Fprintf(&msgBuilder, "%s", chunkMsg)
		}
	}

	mod := len(cypher) % 5
	if mod != 0 {
		rem := cypher[len(cypher)-mod:]
		shortMsg := Decrypt(rem, key)
		fmt.Fprintf(&msgBuilder, "%s", shortMsg)
	}

	msg = msgBuilder.String()
	return
}

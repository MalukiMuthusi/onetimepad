package cmd

import (
	"fmt"
	"os"

	"github.com/malukimuthusi/onetimepad/symmetry"
	"github.com/spf13/cobra"
)

var (
	msg     string
	key     string
	rootCmd = &cobra.Command{
		Use:   "crypt",
		Short: "Crypt is a cryptographic tool that uses symmetric algorithm",
		Long:  `A free and open source encryption algorithm build with open source technologies`,
		Run:   crypt,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&msg, "msg", "m", "Hello World", "message to be encrypted")
	rootCmd.MarkFlagRequired("msg")

	rootCmd.Flags().StringVarP(&key, "key", "k", "secrt", "secrete key to encrypt msg with it")
	rootCmd.MarkFlagRequired("key")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func crypt(cmd *cobra.Command, args []string) {
	fmt.Printf("\t\tEncrypting message\t\t\n")

	cypher := symmetry.EncrptLongMsg(msg, key)

	fmt.Printf("Cypher produced is\n\t%s\n", cypher)

	fmt.Printf("\t\tDecrypting message\t\t\n")
	decryptedMsg := symmetry.DecryptLongMsg(cypher, key)
	fmt.Printf("Decrypted message is:\n%s\n", decryptedMsg)
}

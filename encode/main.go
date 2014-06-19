package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"encoding/base64"
)

func main() {
	var root = &cobra.Command{
		Use: "encode",
	}

	base64Command := &cobra.Command{
		Use: "b64",
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range(args) {
				fmt.Println(toBase64(arg))
			}
		},
	}

	root.AddCommand(base64Command)
	root.Execute()
}

func toBase64(str string) string{
	return base64.StdEncoding.EncodeToString([]byte(str))
}

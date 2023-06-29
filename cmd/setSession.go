/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alfreddobradi/pwdm/store"
	"github.com/spf13/cobra"
)

var setSessionCmd = &cobra.Command{
	Use:   "set-session <session-key>",
	Short: "Sets the session key used for encryption and decryption",
	Long: `This command sets the session key used for encrypting and decrypting values during the session.
The application has no knowledge of the validity of the session key,
decrypting values will only work if the same session key was used for encryption`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := store.SetSessionKey([]byte(args[0]))
		cobra.CheckErr(err)

		fmt.Printf("Session key was set to \"%s\"\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(setSessionCmd)
}

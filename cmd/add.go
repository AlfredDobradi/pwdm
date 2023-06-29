/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/alfreddobradi/pwdm/store"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <key> <value>",
	Short: "Stores a value with the given key",
	Long: `If the session key is set, the application encrypts and stores
		the given value in the database`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := store.Set([]byte(args[0]), []byte(args[1]))
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/alfreddobradi/pwdm/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Returns a stored value by key",
	Long:  `The value will only be returned if the session key is set and encryption is successful`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		val, err := store.Get([]byte(args[0]))
		cobra.CheckErr(err)

		if viper.GetBool("clipboard") {
			err = clipboard.Init()
			cobra.CheckErr(err)

			clipboard.Write(clipboard.FmtText, val)
			return
		}

		tpl := "%s"
		if viper.GetBool("newline") {
			tpl = "%s\n"
		}
		fmt.Printf(tpl, string(val))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().BoolP("clipboard", "c", false, "Write value to clipboard")
	getCmd.Flags().BoolP("newline", "n", false, "Add newline character after the value (only when writing to stdout)")
	viper.BindPFlag("clipboard", getCmd.Flags().Lookup("clipboard")) // nolint
	viper.BindPFlag("newline", getCmd.Flags().Lookup("newline"))     // nolint
}

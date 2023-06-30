/*
Copyright © 2023 Alfred Dobradi <alfreddobradi@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/alfreddobradi/pwdm/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.design/x/clipboard"
)

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

	getCmd.Flags().BoolP("clipboard", "c", false, "Write value to clipboard")
	getCmd.Flags().BoolP("newline", "n", false, "Add newline character after the value (only when writing to stdout)")
	viper.BindPFlag("clipboard", getCmd.Flags().Lookup("clipboard")) // nolint
	viper.BindPFlag("newline", getCmd.Flags().Lookup("newline"))     // nolint
}

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

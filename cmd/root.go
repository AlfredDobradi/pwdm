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
	"os"

	"github.com/alfreddobradi/pwdm/store"
	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	databasePath string
)

var (
	db *badger.DB
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwdm",
	Short: "A simple password manager in your terminal",
	Long: `The application stores and recalls encrypted strings from a key-value database.
		It has no concept of users or authentication, the get operation will only succeed if the
		correct key was used for decryption.`,
}

func Execute() {
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	initDatabase()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwdm/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&databasePath, "db-path", "", "path to the database (default is $HOME/.pwdm/store)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))   // nolint
	viper.BindPFlag("db-path", rootCmd.PersistentFlags().Lookup("db-path")) // nolint
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(fmt.Sprintf("%s/.pwdm", home))
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initDatabase() {
	path := viper.GetString("db-path")
	if path == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		path = fmt.Sprintf("%s/.pwdm/store", home)
	}

	cobra.CheckErr(store.Init(path))
}

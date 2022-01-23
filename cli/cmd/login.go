/*
Copyright © 2022 Teresa Romero <hello@teresaromero.dev>

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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save credentials to access the api",

	Run: func(cmd *cobra.Command, args []string) {

		token := cmd.Flag("token").Value.String()
		user := cmd.Flag("user").Value.String()

		if token == "" || user == "" {
			log.Fatal("no config flags, please login and use the flags")
		}

		if token != "" {
			viper.Set("IG_ACCESS_TOKEN", token)
		}
		if user != "" {
			viper.Set("IG_USER_ID", user)
		}

		configFile := viper.ConfigFileUsed()
		if err := viper.WriteConfigAs(configFile); err != nil {
			log.Fatal(err)
		}

		log.Printf("login data saved at %s", configFile)

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("token", "t", "", "Access Token provided by Instagram")
	loginCmd.Flags().StringP("user", "u", "", "User ID for your instagram user")

}

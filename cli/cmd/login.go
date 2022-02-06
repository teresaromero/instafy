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
	"fmt"

	"github.com/teresaromero/instafy/client"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save credentials to access the api",

	Run: func(cmd *cobra.Command, args []string) {

		api, err := client.NewIgBasicAPI(false)
		cobra.CheckErr(err)

		fmt.Println("A browser window will be opened now, go there and complete the auth process with your ig account")

		clientID, err := api.LaunchLogin()
		cobra.CheckErr(err)

		fmt.Println("Waiting for auth to be completed....")

		if err := api.ListenSSE(clientID); err != nil {
			cobra.CheckErr(err)
		}

		fmt.Println("SUCCESS: Login completed!")

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

}

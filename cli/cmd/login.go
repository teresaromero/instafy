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
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/r3labs/sse/v2"
	"github.com/teresaromero/instafy/constants"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type authInfo struct {
	AccessToken string `json:"access_token"`
	UserID      int    `json:"user_id"`
	ClientID    string `json:"client_id"`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save credentials to access the api",

	Run: func(cmd *cobra.Command, args []string) {

		clientID := uuid.New()

		baseURL := viper.GetString("INSTAFY_API_BASE_URL")
		url, err := url.Parse(baseURL + fmt.Sprintf("/login?client_id=%s", clientID.String()))
		cobra.CheckErr(err)

		if err := browser.OpenURL(url.String()); err != nil {
			cobra.CheckErr(err)
		}

		sseURL, err := url.Parse(baseURL + fmt.Sprintf("/stream-login?client_id=%s", clientID.String()))
		cobra.CheckErr(err)

		sseClient := sse.NewClient(sseURL.String())
		sseClient.SubscribeRaw(func(msg *sse.Event) {
			var authData authInfo
			if err := json.Unmarshal(msg.Data, &authData); err != nil {
				cobra.CheckErr(err)
			}
			viper.Set(constants.AccessTokenEnv, authData.AccessToken)
			viper.Set(constants.ClientIDEnv, authData.ClientID)
			viper.Set(constants.UserIDEnv, authData.UserID)

			if err := viper.WriteConfig(); err != nil {
				cobra.CheckErr(err)
			}
		})

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

}

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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teresaromero/instafy/models"
)

// mediaCmd represents the media command
var mediaCmd = &cobra.Command{
	Use:   "media",
	Short: "Returns your latest 10 media objects, for the last 60 days",

	Run: func(cmd *cobra.Command, args []string) {

		accessToken := viper.GetString("IG_ACCESS_TOKEN")
		userID := viper.GetString("IG_USER_ID")

		if accessToken == "" || userID == "" {
			log.Fatal("no config, you have to login")
		}

		client := &http.Client{}

		baseURL := os.Getenv("INSTAFY_API_BASE_URL")
		url, err := url.Parse(fmt.Sprintf("%s/media", baseURL))
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("x-access-token", accessToken)
		req.Header.Set("x-user-id", userID)

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		data := models.MediaResponse{}
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatal(err)
		}

		log.Printf("Got %v media objects", len(data))

	},
}

func init() {
	rootCmd.AddCommand(mediaCmd)
}

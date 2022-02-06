package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/r3labs/sse/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teresaromero/instafy/constants"
)

type authResponse struct {
	AccessToken string `json:"access_token"`
	UserID      int    `json:"user_id"`
	ClientID    string `json:"client_id"`
}

func (i *IgBasicAPI) ListenSSE(clientID string) error {
	sseURL, err := url.Parse(i.baseURL + fmt.Sprintf("/stream-login?client_id=%s", clientID))
	if err != nil {
		return err
	}

	sseClient := sse.NewClient(sseURL.String())
	return sseClient.SubscribeRaw(func(msg *sse.Event) {
		var auth authResponse
		if err := json.Unmarshal(msg.Data, &auth); err != nil {
			cobra.CheckErr(err)
		}
		viper.Set(constants.AccessTokenEnv, auth.AccessToken)
		viper.Set(constants.ClientIDEnv, auth.ClientID)
		viper.Set(constants.UserIDEnv, auth.UserID)

		if err := viper.WriteConfig(); err != nil {
			cobra.CheckErr(err)
		}
		log.Println("auth set at config file")
	})

}

// LaunchLogin starts login auth
func (i *IgBasicAPI) LaunchLogin() (string, error) {
	clientID := uuid.New()

	url, err := url.Parse(i.baseURL + fmt.Sprintf("/login?client_id=%s", clientID.String()))
	if err != nil {
		return "", err
	}

	return clientID.String(), browser.OpenURL(url.String())

}

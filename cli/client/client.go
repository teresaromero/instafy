package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/teresaromero/instafy/constants"
)

// IgBasicAPI represents the Instagram Basic Visualization API
type IgBasicAPI struct {
	baseURL     string
	client      *http.Client
	accessToken string
	userID      string
}

// NewIgBasicAPI returns a new IgBasicAPI
func NewIgBasicAPI() (*IgBasicAPI, error) {

	url := viper.GetString("INSTAFY_API_BASE_URL")
	if url == "" {
		return nil, fmt.Errorf("can't setup client, invalid url: %s", url)
	}

	token := viper.GetString(constants.AccessTokenEnv)
	user := viper.GetString(constants.UserIDEnv)

	if token == "" || user == "" {
		return nil, errors.New("can't setup client, invalid url, token or user")
	}

	return &IgBasicAPI{
		baseURL:     url,
		client:      &http.Client{Timeout: time.Duration(10 * time.Second)},
		accessToken: token,
		userID:      user,
	}, nil
}

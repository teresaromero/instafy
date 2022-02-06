package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Media represents a single unit of media content return from API
type Media struct {
	Caption   string `json:"caption"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
}

// MediaResponse represents an array of Media
type MediaResponse = []Media

// GetMedia gets the media from IG API
func (i *IgBasicAPI) GetMedia() (MediaResponse, error) {
	url, err := url.Parse(fmt.Sprintf("%s/media", i.baseURL))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-access-token", i.accessToken)
	req.Header.Set("x-user-id", i.userID)

	res, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data MediaResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil

}

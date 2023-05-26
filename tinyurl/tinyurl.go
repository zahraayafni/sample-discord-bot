package tinyurl

import (
	"errors"
	"fmt"
	"os"

	"github.com/zahraayafni/sample-discord-bot/pkg/http"
)

type (
	CreateShortLinkRequest struct {
		URL string `json:"url"`
	}

	CreateShortLinkResponse struct {
		Data struct {
			TinyURL string `json:"tiny_url"`
		} `json:"data"`
		Code   int      `json:"code"`
		Errors []string `json:"errors"`
	}
)

func CreateShortLink(url string) (string, error) {
	req := CreateShortLinkRequest{
		URL: url,
	}

	resp := CreateShortLinkResponse{}

	client := generateTinyUrlClientParams(req, resp)
	clientResp, err := client.Post()
	if err != nil {
		return "", err
	}

	if clientResp.Code != 0 {
		return "", errors.New(fmt.Sprintf("%+v", clientResp.Errors))
	}

	return clientResp.Data.TinyURL, nil
}

func generateTinyUrlClientParams[payloadStruct any, responseStruct any](payload payloadStruct, response responseStruct) http.ClientHttpParams[payloadStruct, responseStruct] {
	tinyurlToken := "Bearer " + os.Getenv("DISCORD_TOKEN")
	return http.ClientHttpParams[payloadStruct, responseStruct]{
		BaseUrl:  "https://api.tinyurl.com",
		Path:     "/create",
		Token:    tinyurlToken,
		Header:   map[string]string{},
		Payload:  payload,
		Response: response,
	}
}

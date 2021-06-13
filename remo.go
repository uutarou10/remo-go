package remo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	path_ "path"
)

const (
	defaultApiEndpoint = "https://api.nature.global"
)

type Client struct {
	endpoint   string
	token      string
	httpClient *http.Client
}

func New(token string) *Client {
	return &Client{
		endpoint:   defaultApiEndpoint,
		token:      token,
		httpClient: nil,
	}
}

func (c *Client) getApi(path string, dest interface{}) error {
	url_, err := url.Parse(c.endpoint)
	if err != nil {
		log.Panicf("failed to parse api endpoint: %s\n", err.Error())
	}
	url_.Path = path_.Join(url_.Path, path)

	request, err := http.NewRequest(http.MethodGet, url_.String(), nil)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(dest); err != nil {
		return err
	}

	return nil
}

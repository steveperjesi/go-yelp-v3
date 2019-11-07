package yelp

import (
	"errors"
)

type Client struct {
	AuthOptions AuthOptions
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is missing")
	}

	newClient := new(Client)
	newClient.AuthOptions.ApiKey = apiKey

	return newClient, nil
}

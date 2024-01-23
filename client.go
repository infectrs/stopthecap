package stopthecap

import "errors"

func NewClient(apiKey string) (*CapsolverClient, error) {

	if apiKey == "" {
		return nil, errors.New("stopthecap: api key is missing")
	}

	client := CapsolverClient{
		ClientKey: apiKey,
	}

	return &client, nil
}

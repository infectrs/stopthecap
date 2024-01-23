package stopthecap


func NewClient(apiKey string) (*CapsolverClient, error) {

	if apiKey == "" {
		return nil, apiKeyError
	}

	client := CapsolverClient{
		ClientKey: apiKey,
	}

	return &client, nil
}

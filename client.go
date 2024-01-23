package stopthecap

// NewClient creates a new CapSolver client
func NewClient(apiKey string) (*CapsolverClient, error) {

	if apiKey == "" {
		return nil, apiKeyError
	}

	client := CapsolverClient{
		ClientKey: apiKey,
	}

	return &client, nil
}

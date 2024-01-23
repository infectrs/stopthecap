package stopthecap

const (
	getBalanceEndpoint = "/getBalance"
)

func (client CapsolverClient) GetBalance() (*CapsolverResponse, error) {

	reqBody := CapsolverRequest{
		ClientKey: client.ClientKey,
	}

	resp, err := newReq(getBalanceEndpoint, reqBody)

	if err != nil {
		return nil, balanceError
	}

	return resp, nil
}

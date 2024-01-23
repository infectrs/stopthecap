package stopthecap

import "errors"

const (
	getBalanceEndpoint = "/getBalance"
)

func (client CapsolverClient) GetBalance() (*CapsolverResponse, error) {

	reqBody := CapsolverRequest{
		ClientKey: client.ClientKey,
	}

	resp, err := newReq(getBalanceEndpoint, reqBody)

	if err != nil {
		return nil, errors.New("stopthecap: couldn't get balance")
	}

	return resp, nil
}

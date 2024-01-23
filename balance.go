package stopthecap

const (
	getBalanceEndpoint = "/getBalance"
)

// GetBalance retrieves the balance of a CapSolver account
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

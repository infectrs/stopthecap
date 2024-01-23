package stopthecap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseUrl = "https://api.capsolver.com"
)

func newReq(apiEndpoint string, reqBody CapsolverRequest) (*CapsolverResponse, error) {
	url := fmt.Sprintf("%s%s", baseUrl, apiEndpoint)

	jsonReqBody, err := json.Marshal(reqBody)

	if err != nil {
		return nil, err
	}

	jsonReqBodyBytes := bytes.NewReader(jsonReqBody)

	resp, err := http.Post(url, "application/json", jsonReqBodyBytes)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var respBodyJson CapsolverResponse

	if err := json.Unmarshal(respBodyBytes, &respBodyJson); err != nil {
		return nil, err
	}

	return &respBodyJson, nil

}

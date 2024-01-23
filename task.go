package stopthecap

const (
	createTaskEndpoint    = "/createTask"
	getTaskResultEndpoint = "/getTaskResult"
)

func (client CapsolverClient) createTask(captchaTask map[string]any) (*CapsolverResponse, error) {

	reqBody := CapsolverRequest{
		ClientKey: client.ClientKey,
		Task:      captchaTask,
	}

	resp, err := newReq(createTaskEndpoint, reqBody)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (client CapsolverClient) getTaskResult(taskId string) (*CapsolverResponse, error) {

	reqBody := CapsolverRequest{
		ClientKey: client.ClientKey,
		TaskId:    taskId,
	}

	resp, err := newReq(getTaskResultEndpoint, reqBody)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

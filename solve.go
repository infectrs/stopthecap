package stopthecap

import (
	"errors"
	"strings"
	"time"
)

const (
	readyStatus = "ready"
)

var (
	supportedModes = []string{"HCaptchaTask", "HCaptchaTaskProxyLess", "FunCaptchaTaskProxyLess", "GeeTestTask", "GeeTestTaskProxyLess", "ReCaptchaV2Task", "ReCaptchaV2EnterpriseTask", "ReCaptchaV2TaskProxyLess", "ReCaptchaV2EnterpriseTaskProxyLess", "ReCaptchaV3Task", "ReCaptchaV3EnterpriseTask", "ReCaptchaV3TaskProxyLess", "ReCaptchaV3EnterpriseTaskProxyLess", "ReCaptchaV3M1TaskProxyLess", "MTCaptcha"}
)

func (client CapsolverClient) Solve(captchaTask map[string]any, retry int, timeout time.Duration) (*CapsolverResponse, error) {
	taskType, exists := captchaTask["type"].(string)

	if !exists {
		return nil, errors.New("stopthecap: captcha type is missing")
	}

	normalizedTaskType := strings.ToLower(taskType)

	if !contains(supportedModes, normalizedTaskType) {
		return nil, errors.New("stopthecap: captcha type is not supported")
	}

	task, err := client.createTask(captchaTask)

	if err != nil {
		return nil, errors.New("stopthecap: couldn't create task")
	}

	if task.ErrorID == 1 {
		return nil, errors.New("stopthecap: " + task.ErrorDescription)
	}

	var taskResult *CapsolverResponse
	var taskResultErr error

	for i := 0; i < retry; i++ {
		taskResult, taskResultErr = client.getTaskResult(task.TaskId)

		if taskResultErr != nil {
			return nil, errors.New("stopthecap: couldn't get task result")
		}

		time.Sleep(timeout)

		if taskResult.ErrorID == 1 {
			return nil, errors.New("stopthecap: " + task.ErrorDescription)
		}

		if taskResult.Status == readyStatus {
			break
		}
	}

	return taskResult, nil
}

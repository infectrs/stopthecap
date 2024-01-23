package stopthecap

import (
	"strings"
	"time"
)

const (
	readyStatus = "ready"
)

var (
	supportedModes = []string{"HCaptchaTask", "HCaptchaTaskProxyLess", "FunCaptchaTaskProxyLess", "GeeTestTask", "GeeTestTaskProxyLess", "ReCaptchaV2Task", "ReCaptchaV2EnterpriseTask", "ReCaptchaV2TaskProxyLess", "ReCaptchaV2EnterpriseTaskProxyLess", "ReCaptchaV3Task", "ReCaptchaV3EnterpriseTask", "ReCaptchaV3TaskProxyLess", "ReCaptchaV3EnterpriseTaskProxyLess", "ReCaptchaV3M1TaskProxyLess", "MTCaptcha"}
)

// Solve solves a captcha using CapSolver
func (client CapsolverClient) Solve(captchaTask map[string]any, retry int, timeout time.Duration) (*CapsolverResponse, error) {
	taskType, exists := captchaTask["type"].(string)

	if !exists {
		return nil, missingCaptchaTypeError
	}

	normalizedTaskType := strings.ToLower(taskType)

	if !contains(supportedModes, normalizedTaskType) {
		return nil, notSupportedCaptchaTypeError
	}

	task, err := client.createTask(captchaTask)

	if err != nil {
		return nil, createTaskError
	}

	if task.ErrorID == 1 {
		return nil, errorIdError
	}

	var taskResult *CapsolverResponse
	var taskResultErr error

	for i := 0; i < retry; i++ {
		taskResult, taskResultErr = client.getTaskResult(task.TaskId)

		if taskResultErr != nil {
			return nil, taskResultError
		}

		time.Sleep(timeout)

		if taskResult.ErrorID == 1 {
			return nil, errorIdError
		}

		if taskResult.Status == readyStatus {
			break
		}
	}

	return taskResult, nil
}

package stopthecap

import "errors"

// List of stopthecap errors
var (
	apiKeyError                  = errors.New("stopthecap: api key is missing")
	balanceError                 = errors.New("stopthecap: couldn't get balance")
	missingCaptchaTypeError      = errors.New("stopthecap: captcha type is missing")
	notSupportedCaptchaTypeError = errors.New("stopthecap: captcha type is not supported")
	createTaskError              = errors.New("stopthecap: couldn't create task")
	taskResultError              = errors.New("stopthecap: couldn't get task result")
	errorIdError                 = errors.New("stopthecap: got an error while doing task")
)

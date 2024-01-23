package stopthecap

// CapsolverClient holds the CapSolver client information
type CapsolverClient struct {
	ClientKey string
}

// CapsolverRequest holds the CapSolver request body data
type CapsolverRequest struct {
	ClientKey string         `json:"clientKey"`
	Task      map[string]any `json:"task,omitempty"`
	TaskId    string         `json:"taskId,omitempty"`
}

// CapsolverSolution holds the CapSolver solution data
type CapsolverSolution struct {
	UserAgent        string `json:"userAgent,omitempty"`
	ExpireTime       string `json:"expireTime,omitempty"`
	Timestamp        string `json:"timestamp,omitempty"`
	CaptchaKey       string `json:"captchaKey,omitempty"`
	CaptchaSolution  string `json:"gRecaptchaRespons,omitempty"`
	CaptchaId        string `json:"captcha_id,omitempty"`
	CaptchaOutput    string `json:"captcha_output,omitempty"`
	CaptchaChallenge string `json:"challenge,omitempty"`
	CaptchaValidate  string `json:"validate,omitempty"`
	GenTime          string `json:"gen_time,omitempty"`
	LotNumber        string `json:"lot_number,omitempty"`
	PassToken        string `json:"pass_token,omitempty"`
	RiskType         string `json:"risk_type,omitempty"`
}

// CapsolverResponse holds the CapSolver response body data
type CapsolverResponse struct {
	ErrorID          int64              `json:"errorId"`
	ErrorCode        string             `json:"errorCode,omitempty"`
	ErrorDescription string             `json:"errorDescription,omitempty"`
	TaskId           string             `json:"taskId,omitempty"`
	Solution         *CapsolverSolution `json:"solution,omitempty"`
	Balance          float64            `json:"balance,omitempty"`
	Packages         []string           `json:"packages,omitempty"`
	Status           string             `json:"status,omitempty"`
}

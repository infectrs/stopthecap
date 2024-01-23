# stopthecap
Go Package to easily solve captchas using CapSolver

## Supported Captcha Modes

- `HCaptchaTask`
- `HCaptchaTaskProxyLess`
- `FunCaptchaTaskProxyLess`
- `GeeTestTask`
- `GeeTestTaskProxyLess`
- `ReCaptchaV2Task`
- `ReCaptchaV2EnterpriseTask`
- `ReCaptchaV2TaskProxyLess`
- `ReCaptchaV2EnterpriseTaskProxyLess`
- `ReCaptchaV3Task`
- `ReCaptchaV3EnterpriseTask`
- `ReCaptchaV3TaskProxyLess`
- `ReCaptchaV3EnterpriseTaskProxyLess`
- `ReCaptchaV3M1TaskProxyLess`
- `MTCaptcha`
## Installation

`go get github.com/infectrs/stopthecap`
## Usage

### Get Balance

```golang
package main

import (
	"fmt"

	"github.com/infectrs/stopthecap"
)

func main() {
	clientKey := "CAP-..."
	captchaClient, err := stopthecap.NewClient(clientKey)

	if err != nil {
		fmt.Println(err)
	}

	balance, err := captchaClient.GetBalance()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Balance: %.2f", balance.Balance)
}
```

### Solve Captcha

```golang
package main

import (
	"fmt"
	"time"

	"github.com/infectrs/stopthecap"
)

func main() {
	clientKey := "CAP-..."
	captchaClient, err := stopthecap.NewClient(clientKey)

	if err != nil {
		fmt.Println(err)
	}

	captchaRetry := 45
	captchaTimeout := time.Second * 1

	captchaTask := map[string]any{
		"type":       "GeeTestTaskProxyLess",
		"websiteURL": "https://mywebsite.com/geetest/test.php",
		"captchaId":  "",
	}

	solvedCaptcha, err := captchaClient.Solve(captchaTask, captchaRetry, captchaTimeout)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Solved Captcha: %s", solvedCaptcha.Solution.CaptchaOutput) // Note: Response varies based on the type of captcha that is being used
}
```

package invitation

import (
	"errors"
	"log"
	"net/http"
	"os"
)

const (
	envSlackSubDomain     = "SLACK_SUBDOMAIN"
	envSlackAPIToken      = "SLACK_API_TOKEN"
	envReCaptchaSecretKey = "GOOGLE_CAPTCHA_SECRET"
)

func New() http.Handler {
	if err := envValidate(); err != nil {
		log.Fatal(err)
	}

	slackAPIToken := os.Getenv(envSlackAPIToken)
	workspace := os.Getenv(envSlackSubDomain)
	reCaptchaSecretKey := os.Getenv(envReCaptchaSecretKey)
	return slackInviteHandler(workspace, slackAPIToken, reCaptchaSecretKey)
}

func envValidate() error {
	var failed bool
	for _, env := range []string{
		envSlackSubDomain,
		envSlackAPIToken,
		envReCaptchaSecretKey,
	} {
		if os.Getenv(env) == "" {
			failed = true
			log.Printf("unset environment variables: %s\n", env)
		}
	}
	if failed {
		return errors.New("unset environment variables")
	}
	return nil
}

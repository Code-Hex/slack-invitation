package invitation

import (
	"log"
	"net/http"

	"github.com/Code-Hex/slack-invitaion/internal/recaptcha"
	"github.com/Code-Hex/slack-invitaion/internal/slack"
)

func slackInviteHandler(workspace, token, reCaptureSecret string) http.HandlerFunc {
	slackAPIClient := slack.NewInviteClient(workspace, token)
	reCaptchaClient := &recaptcha.Client{
		SecretKey: reCaptureSecret,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		response := r.FormValue("g-recaptcha-response")
		ok, err := reCaptchaClient.Verify(ctx, response)
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if !ok {
			writeJSONResponse(w, http.StatusBadRequest, "reCapture challenge is failed")
			return
		}

		email := r.FormValue("email")
		resp, err := slackAPIClient.Invite(ctx, email)
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if !resp.OK {
			log.Printf("failed to send invite mail: %#v\n", resp)

			return
		}

	}
}

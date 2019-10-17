package invitation

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Code-Hex/slack-invitaion/internal/ip"
	"github.com/Code-Hex/slack-invitaion/internal/recaptcha"
	"github.com/Code-Hex/slack-invitaion/internal/slack"
)

// This regex is got from slack sign-in page
var emailValidation = regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")

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
		email := r.PostFormValue("email")
		if !emailValidation.MatchString(email) {
			writeJSONResponse(w, http.StatusBadRequest, &Response{
				Op:  "validation",
				Msg: "email validation is failed",
			})
			return
		}

		addr, err := ip.Normalize(r.RemoteAddr)
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, &Response{
				Op:  "validation",
				Msg: "invalid remote address: " + err.Error(),
			})
			return
		}

		ctx := r.Context()

		response := r.PostFormValue("g-recaptcha-response")
		ok, err := reCaptchaClient.Verify(ctx, response, addr)
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, &Response{
				Op:  "reCAPTCHA",
				Msg: "request has been failed: " + err.Error(),
			})
			return
		}
		if !ok {
			writeJSONResponse(w, http.StatusBadRequest, &Response{
				Op:  "reCAPTCHA",
				Msg: "challenge is failed (one more try please!!)",
			})
			return
		}

		resp, err := slackAPIClient.Invite(ctx, email)
		if err != nil {
			writeJSONResponse(w, http.StatusBadRequest, &Response{
				Op:  "slack",
				Msg: "request has been failed",
			})
			return
		}
		if !resp.OK {
			// already invited
			if resp.Provided == "already_in_team" {
				http.Redirect(w, r, "https://"+workspace+".slack.com", http.StatusSeeOther)
				return
			}
			writeJSONResponse(w, http.StatusBadRequest, &Response{
				Op:  "slack",
				Msg: slackErrorHandle(resp),
			})
			return
		}
		writeJSONResponse(w, http.StatusOK, &Response{
			Op:  "slack",
			Msg: "invite successful :D",
		})
	}
}

const (
	msgErrMissingAdminToken = "Missing admin scope: The token you provided is for an account that is not an admin. You must provide a token from an admin account in order to invite users through the Slack API."
	msgErrAlreadyInvited    = "You have already been invited to Slack. Check for an email from feedback@slack.com."
)

func slackErrorHandle(resp *slack.Response) string {
	if resp.Provided == "missing_scope" && resp.Needed == "admin" {
		return msgErrMissingAdminToken
	}
	if resp.Provided == "already_invited" {
		return msgErrAlreadyInvited
	}
	return fmt.Sprintf("provided: %s, message: %s", resp.Provided, resp.Error)
}

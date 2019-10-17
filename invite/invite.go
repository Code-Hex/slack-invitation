package invite

import (
	"net/http"

	"github.com/Code-Hex/slack-invitaion/internal/invitation"
)

var h = invitation.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	h.ServeHTTP(w, r)
}

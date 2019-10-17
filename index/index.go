package index

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://codehex.dev/slack_invitation.html", http.StatusSeeOther)
}

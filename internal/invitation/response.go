package invitation

import (
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, code int, msg string) {
	w.Write([]byte(`{"msg":"` + msg + `"}`))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

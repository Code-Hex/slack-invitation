package invitation

import (
	"encoding/json"
	"net/http"
)

// Response indicates response format (json)
type Response struct {
	Op  string `json:"op"`
	Msg string `json:"message"`
}

func writeJSONResponse(w http.ResponseWriter, code int, msg *Response) {
	w.WriteHeader(code)
	b, _ := json.Marshal(msg)
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
}

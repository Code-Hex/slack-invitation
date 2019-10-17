package invitation

import (
	"encoding/json"
	"net/http"
)

// OpError indicates response error (json)
type OpError struct {
	Op  string `json:"op"`
	Msg string `json:"message"`
}

func writeJSONResponse(w http.ResponseWriter, code int, err *OpError) {
	w.WriteHeader(code)
	b, _ := json.Marshal(err)
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
}

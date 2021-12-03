package render

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

func JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	_ = enc.Encode(v)
}

func Text(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(t))
}

func Error(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(err.Error()))
}

func Html(w http.ResponseWriter, t string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(t))
}

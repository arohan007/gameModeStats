package Controller

import (
	"net/http"
)

func DummyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("foo"))
	if err != nil {
		return
	}
}

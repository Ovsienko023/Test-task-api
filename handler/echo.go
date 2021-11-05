package handler

import (
	"net/http"
	"time"
)

func EchoData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}

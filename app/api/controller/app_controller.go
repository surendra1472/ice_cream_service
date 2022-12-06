package controller

import (
	"net/http"
)

func GetHeartBeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

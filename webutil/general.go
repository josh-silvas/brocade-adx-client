package webutil

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func LogRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"remote_address": r.RemoteAddr,
		"method":         r.Method,
		"url":            r.URL,
	}).Info("request")
}

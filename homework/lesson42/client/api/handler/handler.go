package handler

import (
	"net/http"
	"strings"
)

type handler struct {
	client *http.Client
}

func New() *handler {
	return &handler{
		client: &http.Client{},
	}
}

func GetIDFromURL(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		return parts[2]
	}
	return ""
}

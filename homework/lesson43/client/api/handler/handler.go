package handler

import (
	"net/http"
	"strings"
)

type handler struct {
	client *http.Client
}

type HandlerConfig struct {
}

func New(c *HandlerConfig) *handler {
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

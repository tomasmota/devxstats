package util

import (
	"net/http"
	"time"

	"github.com/drone/go-scm/scm/transport"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func NewBearerHttpClient(token string) *http.Client {
	c := NewHttpClient()
	c.Transport = &transport.BearerToken{
		Token: token,
	}
	return c
}

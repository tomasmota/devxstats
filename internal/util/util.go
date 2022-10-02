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

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

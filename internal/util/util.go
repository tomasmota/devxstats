package util

import (
	"net/http"
	"time"
)

type bearerRoundTripper struct {
	token string // Bearer token
}

func (t *bearerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+t.token)
	return http.DefaultTransport.RoundTrip(r)
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func NewBearerHttpClient(token string) *http.Client {
	c := NewHttpClient()
	c.Transport = &bearerRoundTripper{
		token: token,
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

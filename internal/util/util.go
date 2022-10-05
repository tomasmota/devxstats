package util

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

type BearerToken struct {
	Token string // Bearer token
}

func (t *BearerToken) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+t.Token)
	return http.DefaultTransport.RoundTrip(r)
}

func NewBearerHttpClient(token string) *http.Client {
	c := NewHttpClient()
	c.Transport = &BearerToken{
		Token: token,
	}
	c.Timeout = time.Second * 10
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

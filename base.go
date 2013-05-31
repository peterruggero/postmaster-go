package postmaster

import (
	"github.com/jmcvetta/restclient"
	"net/http"
	"net/url"
)

type PostmasterError struct {
	Message string
}

func (e *PostmasterError) Error() string {
	return e.Message
}

type Postmaster struct {
	ApiKey   string
	BaseUrl  string
	Client   *restclient.Client
	Userinfo *url.Userinfo
	Headers  *http.Header
}

func New(key string) *Postmaster {
	client := restclient.New()
	client.UnsafeBasicAuth = true
	userinfo := url.UserPassword(key, "")
	header := http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent":   []string{"Postmaster/1.0 Go"},
	}
	return &Postmaster{
		ApiKey:   key,
		Client:   client,
		Userinfo: userinfo,
		Headers:  &header,
	}
}

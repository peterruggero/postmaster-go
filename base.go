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

// Postmaster is base library structure. Don't use it, invoke New() instead.
type Postmaster struct {
	ApiKey   string
	BaseUrl  string
	Client   *restclient.Client
	Userinfo *url.Userinfo
	Headers  *http.Header
}

// New returns freshly squeezed Postmaster object with all dependants initialized.
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

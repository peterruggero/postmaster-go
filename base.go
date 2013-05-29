package postmaster

import (
	"github.com/jmcvetta/restclient"
	"net/url"
)

type Postmaster struct {
	ApiKey   string
	BaseUrl  string
	Client   *restclient.Client
	Userinfo *url.Userinfo
}

func New(key string) *Postmaster {
	client := restclient.New()
	client.UnsafeBasicAuth = true
	userinfo := url.UserPassword(key, "")
	return &Postmaster{
		ApiKey:   key,
		Client:   client,
		Userinfo: userinfo,
	}
}

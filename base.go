/*
postmaster-go is the Postmaster.io's API library for Go language.

Postmaster takes the pain out of sending shipments via UPS, Fedex, and USPS. Save money
before you ship, while you ship, and after you ship.

For more information, see https://www.postmaster.io/.

For more information about structures' fields, see https://www.postmaster.io/docs.
*/
package postmaster

import (
	"fmt"
	"github.com/jmcvetta/restclient"
	"net/http"
	"net/url"
	"strings"
)

// PostmasterError is returned as error by every function, and is not nil when
// something bad happens.
type PostmasterError struct {
	Message string
	Code  int
}

// Error returns nice error message.
func (e *PostmasterError) Error() string {
	if e.Code != 0 {
		return fmt.Sprintf("%d: %s", e.Code, e.Message)
	} else {
		return e.Message
	}
}

// Postmaster is base library structure. Don't use it, invoke New() instead.
// In case you need to change API base URL, SetBaseUrl() is there for you.
type Postmaster struct {
	apiKey   string
	baseUrl  string
	client   *restclient.Client
	userinfo *url.Userinfo
	headers  *http.Header
}

// New returns freshly squeezed Postmaster object with all dependants initialized.
func New(key string) *Postmaster {
	client := restclient.New()
	userinfo := url.UserPassword(key, "")
	header := http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{fmt.Sprintf("Postmaster/%.1f Go", VERSION)},
	}
	return &Postmaster{
		apiKey:   key,
		client:   client,
		userinfo: userinfo,
		headers:  &header,
	}
}

// SetBaseUrl sets API base URL.
func (p *Postmaster) SetBaseUrl(url string) {
	p.baseUrl = url
	if strings.HasPrefix(url, "https://") {
		p.client.UnsafeBasicAuth = false
	} else {
		p.client.UnsafeBasicAuth = true
	}
}
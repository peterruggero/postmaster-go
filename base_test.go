package postmaster

import (
	"testing"
)

func TestNew(t *testing.T) {
	pm := New("someapikey")
	if pm.userinfo.String() != "someapikey:" {
		t.Error("wrong userinfo value")
	}
}

func TestSetBaseUrl(t *testing.T) {
	pm := New("someapikey")
	pm.SetBaseUrl("http://not-ssl-addr")
	if !pm.client.UnsafeBasicAuth {
		t.Error("UnsafeBasicAuth should be true")
	}
	pm.SetBaseUrl("https://ssl-addr")
	if pm.client.UnsafeBasicAuth {
		t.Error("UnsafeBasicAuth should be false")
	}
}

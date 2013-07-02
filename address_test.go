package postmaster

import (
	"testing"
)

func TestValidate(t *testing.T) {
	// Mock
	c := make(chan *restMockObj, 1)
	post = restMock(c, nil, 100, nil)

	pm := New("apikey")
	addr := new(Address)
	addr.City = "Austin"
	addr.State = "TX"
	pm.Validate(addr)
	ret := <-c
	if ret.endpoint != "validate" {
		t.Error("wrong endpoint")
	}
	if ret.version != "v1" {
		t.Error("wrong version")
	}
	params := ret.params.(*Address)
	if params.City != "Austin" {
		t.Error("wrong param (city)")
	}
	if params.State != "TX" {
		t.Error("wrong param (state)")
	}
}

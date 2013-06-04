package postmaster

import (
	"testing"
)

func TestMakeUrl(t *testing.T) {
	pm := New("key")
	if pm.makeUrl("v", "endp") != "http://api.postmaster.io/v/endp" {
		t.Error("wrong url for default BaseUrl")
	}
	pm.BaseUrl = "something"
	if pm.makeUrl("v", "endp") != "something/v/endp" {
		t.Error("wrong url for custom BaseUrl")
	}
}

func TestUrlencode(t *testing.T) {
	m := make(map[string]string)
	m["some"] = "thing"
	if urlencode(m) != "&some=thing&" {
		t.Error("failed for 1 param")
	}
	m["any"] = "one"
	if urlencode(m) != "&some=thing&any=one&" {
		t.Error("failed for 2 params")
	}
	m["blank"] = ""
	if urlencode(m) != "&some=thing&any=one&" {
		t.Error("failed for blank param")
	}
}

type N struct {
	A string
	B int
	C float32
}

type S struct {
	A string
	B int
	C float32
	D *N
}

func TestMapStruct(t *testing.T) {
	var m map[string]string
	s := new(S)
	m = mapStructNested(s, "")
	if len(m) != 0 {
		t.Error("non-empty map for empty struct")
	}
	s.A = "test"
	s.B = 5
	s.C = 3.2
	m = mapStructNested(s, "")
	if len(m) != 3 {
		t.Error("map should contain exactly 3 items")
	}
	if m["a"] != "test" {
		t.Error("wrong value for A")
	}
	if m["b"] != "5" {
		t.Error("wrong value for B")
	}
	if m["c"] != "3.2" {
		t.Error("wrong value for C")
	}
	s.D = new(N)
	s.D.B = 5
	m = mapStructNested(s, "")
	if len(m) != 4 {
		t.Error("map should contain exactly 4 items")
	}
	if m["d[b]"] != "5" {
		t.Error("wrong value for D.B")
	}
}

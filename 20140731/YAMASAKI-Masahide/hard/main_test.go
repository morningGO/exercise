package main

import "testing"

func TestGetInterfaces(t *testing.T) {
	s := GetInterfaces()
	if len(*s) <= 0 {
		t.Errorf("len(s) > <%d> want <%s>", len(*s), `">=0"`)
	}
}

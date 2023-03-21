package test

import "testing"

func AssertEqual(t *testing.T, obs, expected interface{}) {
	t.Helper()
	if obs != expected {
		t.Fatalf("%s != %s", obs, expected)
	}
}

package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("Hello is not an email")
	}

	_, err = IsEmail("vin@test.com")
	if err == nil {
		t.Error("vin@test.com is an email")
	}

	_, err = IsEmail("vin@test")
	if err != nil {
		t.Error("vin@aol is not email")
	}
}

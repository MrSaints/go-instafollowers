package util

import (
	"os"
	"testing"
)

const (
	testFile = "config_test.json"
	fakeFile = "nonexistent.json"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig(testFile)
	if err != nil {
		t.Fatalf("LoadConfig returned unexpected error: %v", err)
	}
	if got, want := c.ClientID, "successfulID"; got != want {
		t.Errorf("LoadConfig ClientID is %v, want %v", got, want)
	}
	if got, want := c.ClientSecret, "successfulSecret"; got != want {
		t.Errorf("LoadConfig ClientSecret is %v, want %v", got, want)
	}
	if got, want := c.AccessToken, "successfulToken"; got != want {
		t.Errorf("LoadConfig AccessToken is %v, want %v", got, want)
	}
}

func TestLoadConfig_noFile(t *testing.T) {
	_, err := LoadConfig(fakeFile)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*os.PathError); !ok {
		t.Errorf("Expected error to be of type *PathError, got %+v", err)
	}
}

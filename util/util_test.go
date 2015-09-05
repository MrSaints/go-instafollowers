package util

import (
	"os"
	"testing"
)

const (
	TestFile = "config_test.json"
	FakeFile = "nonexistent.json"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig(TestFile)
	if err != nil {
		t.Fatalf("LoadConfig returned unexpected error: %+v", err)
	}
	if got, want := c.ClientID, "successfulID"; got != want {
		t.Errorf("Client id is %+v, want %+v", got, want)
	}
	if got, want := c.ClientSecret, "successfulSecret"; got != want {
		t.Errorf("Client secret is %+v, want %+v", got, want)
	}
	if got, want := c.AccessToken, "successfulToken"; got != want {
		t.Errorf("Access token is %+v, want %+v", got, want)
	}
}

func TestLoadConfig_noFile(t *testing.T) {
	_, err := LoadConfig(FakeFile)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*os.PathError); !ok {
		t.Errorf("Expected an OS path error, got %+v", err)
	}
}

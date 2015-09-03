package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("config_test.json")
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
	_, err := LoadConfig("nonexistent.json")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*os.PathError); !ok {
		t.Errorf("Expected error to be of type *PathError, got %+v", err)
	}
}

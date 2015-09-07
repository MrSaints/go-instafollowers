package util

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}

func LoadConfig(fn string) (Configuration, error) {
	c := Configuration{}

	f, err := os.Open(fn)
	if err != nil {
		return c, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&c)

	if err != nil {
		return c, err
	}

	return c, nil
}

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

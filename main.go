package main

import (
	"encoding/json"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

const (
	configFile = "config.json"
)

var (
	client *instagram.Client
)

type Configuration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}

func LoadConfig(fn string) (Configuration, error) {
	c := Configuration{}
	f, err := os.Open(fn)
	defer f.Close()

	if err != nil {
		return c, err
	}

	err = json.NewDecoder(f).Decode(&c)

	if err != nil {
		return c, err
	}

	return c, nil
}

func main() {
	// Load config
	config, err := LoadConfig(configFile)
	failOnError(err)

	// Create an Instagram client
	client = instagram.NewClient(nil)
	client.ClientID = config.ClientID
	client.ClientSecret = config.ClientSecret
	client.AccessToken = config.AccessToken

	// Commands
	commands := []cli.Command{
		{
			Name:    "FollowsBack",
			Aliases: []string{"fb"},
			Usage:   "Returns a list of Instagram users who does not follow you back",
			Action:  FollowsBack,
		},
	}

	// Create CLI
	app := cli.NewApp()
	app.Name = "go-instafollowers"
	app.Usage = "Manage your Instagram followers"
	app.Version = "1.0.0"
	app.Commands = commands
	app.Run(os.Args)
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

const (
	configFile    = "config.json"
	followersFile = "followers.json"
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

	if _, err := os.Stat(fn); os.IsNotExist(err) {
		err = errors.New(fmt.Sprintf("The configuration file (\"%s\") cannot be found.", fn))
		return c, err
	}

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

	if config.ClientID == "" || config.ClientSecret == "" {
		err = errors.New(fmt.Sprintf("Please set your `client_id` and `client_secret` config in \"%v\".\n", configFile))
		failOnError(err)
	}

	// Create an Instagram client
	client = instagram.NewClient(nil)
	client.ClientID = config.ClientID
	client.ClientSecret = config.ClientSecret
	client.AccessToken = config.AccessToken

	// Commands
	commands := []cli.Command{
		{
			Name:    "Followers",
			Aliases: []string{"fl"},
			Usage:   "Returns a list of users who are following you",
			Action:  Followers,
		},
		{
			Name:    "Following",
			Aliases: []string{"fw"},
			Usage:   "Returns a list of users who you are following",
			Action:  Following,
		},
		{
			Name:    "FollowsBack",
			Aliases: []string{"fb"},
			Usage:   "Returns a list of users who are not following you back",
			Action:  FollowsBack,
		},
		{
			Name:    "Unfollowed",
			Aliases: []string{"un"},
			Usage:   "Returns a list of users who unfollowed you (since you last ran the command)",
			Action:  Unfollowed,
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

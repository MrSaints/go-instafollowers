package main

import (
	"errors"
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
	"github.com/mrsaints/go-instafollowers/util"
	"os"
)

const (
	configFile    = "config.json"
	followersFile = "followers.json"
)

var (
	client *instagram.Client
)

func main() {
	// Load config
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err = errors.New(fmt.Sprintf("The configuration file (\"%s\") cannot be found.\n", configFile))
		util.FailOnError(err)
	}

	config, err := util.LoadConfig(configFile)
	util.FailOnError(err)

	if config.AccessToken == "" {
		err = errors.New(fmt.Sprintf("This app requires an authenticated user. Please set your `access_token` config in \"%v\".\n", configFile))
		util.FailOnError(err)
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
	app.Name = "instafollowers"
	app.Authors = []cli.Author{cli.Author{"Ian Lai", "os@fyianlai.com"}}
	app.Usage = "Manage your Instagram followers"
	app.Version = "1.4.3"
	app.Commands = commands
	app.Run(os.Args)
}

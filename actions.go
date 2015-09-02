package main

import (
	"encoding/json"
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
)

func userIsFollowing(u instagram.User, followers []instagram.User) bool {
	for _, follower := range followers {
		if u == follower {
			return true
		}
	}
	return false
}

func FollowsBack(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	failOnError(err)

	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	var total int
	fmt.Println("Users who are not following you back:")

	for _, u := range following {
		if !userIsFollowing(u, followers) {
			total++
			fmt.Printf("%s (%s)\n", u.Username, u.FullName)
		}
	}

	fmt.Printf("Total: %v", total)
}

func Followers(c *cli.Context) {
	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	var total int
	fmt.Println("Users who are following you:")

	for _, u := range followers {
		total++
		fmt.Printf("%s (%s)\n", u.Username, u.FullName)
	}

	fmt.Printf("Total: %v", total)
}

func Following(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	failOnError(err)

	var total int
	fmt.Println("Users who you are following:")

	for _, u := range following {
		total++
		fmt.Printf("%s (%s)\n", u.Username, u.FullName)
	}

	fmt.Printf("Total: %v", total)
}

func Unfollowed(c *cli.Context) {
	// Load history
	f, err := ioutil.ReadFile(followersFile)
	var history []instagram.User

	// Not running for the first time
	if err == nil {
		err = json.Unmarshal(f, &history)
		failOnError(err)
	} else {
		log.Println(err)
		log.Println("Running the command for the first time. You should have received a \"no such file or directory\" error.")
	}

	// Load current followers
	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	var total int
	fmt.Println("Users who have unfollowed you (since you last ran the command):")

	// Match history with current followers
	for _, u := range history {
		if !userIsFollowing(u, followers) {
			total++
			fmt.Printf("%s (%s)\n", u.Username, u.FullName)
		}
	}

	fmt.Printf("Total: %v", total)

	// Save current followers for future reference
	if total != 0 {
		d, err := json.Marshal(followers)
		failOnError(err)
		err = ioutil.WriteFile(followersFile, d, 0644)
		failOnError(err)
	}
}

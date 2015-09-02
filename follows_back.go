package main

import (
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
)

func FollowsBack(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	failOnError(err)

	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	var total int
	for _, u := range following {
		if !userIsFollowing(u, followers) {
			total++
			fmt.Printf("%s (%s)\n", u.Username, u.FullName)
		}
	}

	fmt.Printf("Total: %v", total)
}

func userIsFollowing(u instagram.User, followers []instagram.User) bool {
	for _, follower := range followers {
		if u == follower {
			return true
		}
	}
	return false
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func userIsFollowing(u instagram.User, followers []instagram.User) bool {
	for _, follower := range followers {
		if u == follower {
			return true
		}
	}
	return false
}

func getNonFollowers(ref, f []instagram.User) []instagram.User {
	// Reuse slice
	notFollowing := ref[:0]
	for _, u := range ref {
		if !userIsFollowing(u, f) {
			notFollowing = append(notFollowing, u)
		}
	}
	return notFollowing
}

func printUsers(users []instagram.User) int {
	var total int
	for _, u := range users {
		total++
		fmt.Printf("%s (%s)\n", u.Username, u.FullName)
	}
	fmt.Printf("Total: %v\n", total)
	return total
}

func FollowsBack(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	failOnError(err)
	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	fmt.Println("Users who are not following you back:")
	u := getNonFollowers(following, followers)
	_ = printUsers(u)
}

func Followers(c *cli.Context) {
	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	fmt.Println("Users who are following you:")
	_ = printUsers(followers)
}

func Following(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	failOnError(err)

	fmt.Println("Users who you are following:")
	_ = printUsers(following)
}

func Unfollowed(c *cli.Context) {
	// Load history
	f, err := os.Open(followersFile)
	defer f.Close()

	var history []instagram.User

	// Not running for the first time
	if err == nil {
		err = json.NewDecoder(f).Decode(&history)
		failOnError(err)
	} else {
		log.Println(err)
		log.Println("Running the command for the first time. You should have received a \"no such file or directory\" error.")
	}

	// Load current followers
	followers, _, err := client.Relationships.FollowedBy("")
	failOnError(err)

	fs, err := f.Stat()
	failOnError(err)

	fmt.Printf("Last modified: %v\n", fs.ModTime())
	fmt.Println("Users who have unfollowed you (since you last ran the command):")

	// Match history with current followers
	u := getNonFollowers(history, followers)
	total := printUsers(u)

	// Save current followers for future reference
	if total != 0 {
		d, err := json.Marshal(followers)
		failOnError(err)
		_, err = f.Write(d)
		failOnError(err)
		f.Sync()
	}
}

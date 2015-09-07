package main

import (
	"encoding/json"
	"fmt"
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/codegangsta/cli"
	"github.com/mrsaints/go-instafollowers/util"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// ByUsername implements sort.Interface for []instagram.User based on the Username field.
type ByUsername []instagram.User

func (m ByUsername) Len() int           { return len(m) }
func (m ByUsername) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByUsername) Less(i, j int) bool { return m[i].Username < m[j].Username }

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
	sort.Sort(ByUsername(users))
	for _, u := range users {
		total++
		fmt.Printf("%s (%s)\n", u.Username, u.FullName)
	}
	fmt.Printf("Total: %v\n", total)
	return total
}

func FollowsBack(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	util.FailOnError(err)
	followers, _, err := client.Relationships.FollowedBy("")
	util.FailOnError(err)

	fmt.Println("Users who are not following you back:")
	u := getNonFollowers(following, followers)
	_ = printUsers(u)
}

func Followers(c *cli.Context) {
	followers, _, err := client.Relationships.FollowedBy("")
	util.FailOnError(err)

	fmt.Println("Users who are following you:")
	_ = printUsers(followers)
}

func Following(c *cli.Context) {
	following, _, err := client.Relationships.Follows("")
	util.FailOnError(err)

	fmt.Println("Users who you are following:")
	_ = printUsers(following)
}

func Unfollowed(c *cli.Context) {
	// Load history
	f, err := os.Open(followersFile)

	var history []instagram.User
	var fs os.FileInfo

	// Not running for the first time
	if err == nil {
		err = json.NewDecoder(f).Decode(&history)
		util.FailOnError(err)

		fs, err = f.Stat()
		util.FailOnError(err)
		fmt.Printf("Last modified: %v\n", fs.ModTime())
	} else {
		log.Println(err)
		log.Println("Running the command for the first time. You should have received a 'no such file or directory' error.")
	}

	defer f.Close()

	// Load current followers
	followers, _, err := client.Relationships.FollowedBy("")
	util.FailOnError(err)

	fmt.Println("Users who have unfollowed you (since you last ran the command):")

	// Match history with current followers
	u := getNonFollowers(history, followers)
	total := printUsers(u)

	// Save current followers for future reference
	if fs == nil || total != 0 {
		d, err := json.Marshal(followers)
		util.FailOnError(err)

		err = ioutil.WriteFile(followersFile, d, 0644)
		util.FailOnError(err)
	}
}

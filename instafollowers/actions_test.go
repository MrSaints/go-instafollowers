package main

import (
	"github.com/Invoiced/go-instagram/instagram"
	"github.com/manveru/faker"
	"math/rand"
	"testing"
)

func genFakeUser(f *faker.Faker) instagram.User {
	if f == nil {
		f, _ = faker.New("en")
		f.Rand = rand.New(rand.NewSource(42))
	}

	return instagram.User{
		Username: f.Name(),
		FullName: f.UserName(),
	}
}

func genFakeUsers(total int) []instagram.User {
	f, _ := faker.New("en")
	f.Rand = rand.New(rand.NewSource(42))
	var fakeUsers []instagram.User
	for i := 0; i < total; i++ {
		fakeUsers = append(fakeUsers, genFakeUser(f))
	}
	return fakeUsers
}

func TestDo_userIsFollowing(t *testing.T) {
	f := genFakeUsers(5)
	if got, want := userIsFollowing(f[0], f), true; got != want {
		t.Errorf("userIsFollowing return is %v, want %v", got, want)
	}
	if got, want := userIsFollowing(f[0], f[1:]), false; got != want {
		t.Errorf("userIsFollowing return is %v, want %v", got, want)
	}
}

func TestDo_getNonFollowers(t *testing.T) {
	r := genFakeUsers(5)
	nonFollowers := getNonFollowers(r, r[1:])
	if len(nonFollowers) != 1 {
		t.Errorf("Expected return length to be 1")
	}
	if got, want := nonFollowers[0], r[0]; got != want {
		t.Errorf("getNonFollowers return is %v, want %v", got, want)
	}

	nonFollowers = getNonFollowers(r, r)
	if len(nonFollowers) != 0 {
		t.Errorf("Expected return length to be 0")
	}
}

func TestDo_printUsers(t *testing.T) {
	u := genFakeUsers(5)
	if got, want := printUsers(u), 5; got != want {
		t.Errorf("printUsers returned count is %v, want %v", got, want)
	}
}

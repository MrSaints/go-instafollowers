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
	fakeUsers := make([]instagram.User, total)
	for i := 0; i < total; i++ {
		fakeUsers[i] = genFakeUser(f)
	}
	return fakeUsers
}

func TestDo_userIsFollowing(t *testing.T) {
	f := genFakeUsers(5)
	if got, want := userIsFollowing(f[0], f), true; got != want {
		t.Errorf("userIsFollowing returned %+v, want %+v", got, want)
	}
	if got, want := userIsFollowing(f[0], f[1:]), false; got != want {
		t.Errorf("userIsFollowing returned %+v, want %+v", got, want)
	}
}

func TestDo_getNonFollowers(t *testing.T) {
	r := genFakeUsers(5)
	nonFollowers := getNonFollowers(r, r[1:])
	if got := len(nonFollowers); got != 1 {
		t.Errorf("getNonFollowers returned %+v results, want 1", got)
	}
	if got, want := nonFollowers[0], r[0]; got != want {
		t.Errorf("getNonFollowers returned %+v, want %+v", got, want)
	}

	nonFollowers = getNonFollowers(r, r)
	if got := len(nonFollowers); got != 0 {
		t.Errorf("getNonFollowers returned %+v results, want 0", got)
	}
}

func TestDo_printUsers(t *testing.T) {
	u := genFakeUsers(5)
	if got, want := printUsers(u), 5; got != want {
		t.Errorf("printUsers return is %+v, want %+v", got, want)
	}
}

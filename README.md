# go-instafollowers

A simple Go CLI app for managing your Instagram followers.


## Features

**To avoid any confusion:** It CANNOT help you get more Instagram followers.

It CAN however:

- [x] List users you are following / being followed by;
- [x] List users who are not following you back;
- [x] Help you keep track of who followed / unfollowed you;
- [ ] Help you follow / unfollow user(s) _(WIP: Still trying to figure out the best way to handle this)_.

You can do all that in the comfort of your computer, and hence, you will not have to deal with any pesky ads, loading times or having to give away your personal information. This app is entirely open source too.


## Installation

```shell
go get github.com/mrsaints/go-instafollowers
```

Alternatively, you can download the [source][].


## Usage

You can either run the app through the Go command or by executing the pre-built [amd64][] binary. Ensure you are in the project / app directory before proceeding with the commands below.

### Go command

- Via `go run`:

    ```shell
    go run main.go actions.go
    ```

- Via `go build`:

    ```shell
    go build && ./go-instafollowers
    ```

### Pre-built binary

```
./bin/go-instafollowers
```

Run the app without any arguments to see the usage information.


[source]: https://github.com/MrSaints/go-steamwebapi/archive/master.zip
[amd64]: https://github.com/MrSaints/go-instafollowers/tree/master/bin
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


## Configuration

The [`instafollowers/`][main] app relies on an `access_token` which you must set in `config.json`. A sample `config.new.json` is provided. Rename it, and fill in your client information accordingly.

If you do not have an access token, refer to the [section below](#access-token) on obtaining one.

The [`instatoken/`][instatoken] web app relies on a `client_id`, and a `client_secret` being set. You can obtain these credentials by registering an [Instagram API client][register].

The `config.json` should be present in the directory you choose to run the apps. For example, if you are running the binary from the [root][] directory `./bin/instafollowers`, you will need to ensure that `config.json` is copied into the [root][] directory.


## Usage

You can either run the app through the Go command or by executing the pre-built [amd64][] binary.

### Go command

Ensure you are in the **main project / app directory** (i.e. [`instafollowers/`][main]) before proceeding with the commands below. Now, compile, and run:

- Either via `go run`:

    ```shell
    go run main.go actions.go
    ```

- Or via `go build`:

    ```shell
    go build && ./instafollowers
    ```

### Pre-built binary

In the [root][] directory (i.e. where the `LICENSE` is), run:

```shell
./bin/instafollowers
```

Run the app without any arguments to see the usage information.


## Access Token

To obtain your access token (for individual / private use), you can generate it using the [`instatoken/`][instatoken] web app included in this package.

### Go command

First, ensure you are in the [`instatoken/`][instatoken] directory. Now, compile, and run:

    ```
    go build && ./instatoken
    ```

### Pre-built binary

Alternatively, you can execute the pre-built [amd64][] binary:

    ```shell
    ./bin/instatoken
    ```

The above command assumes that you are in the [root][] directory.

Now, navigate to [http://localhost:8080](http://localhost:8080), and follow the generated link.


[source]: https://github.com/MrSaints/go-steamwebapi/archive/master.zip
[amd64]: https://github.com/MrSaints/go-instafollowers/tree/master/bin
[root]: https://github.com/MrSaints/go-instafollowers
[main]: https://github.com/MrSaints/go-instafollowers/tree/master/instafollowers
[instatoken]: https://github.com/MrSaints/go-instafollowers/tree/master/instatoken
[register]: https://instagram.com/developer/clients/manage/
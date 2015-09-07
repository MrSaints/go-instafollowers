package main

import (
	"fmt"
	"github.com/mrsaints/go-instafollowers/util"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	serverPort        = ":8080"
	redirectURL       = "http://localhost:8080/handshake"
	instagramAuthURL  = "https://api.instagram.com/oauth/authorize"
	instagramTokenURL = "https://api.instagram.com/oauth/access_token"
)

const (
	configFile = "config.json"
)

var (
	igConf *oauth2.Config
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	url := igConf.AuthCodeURL("")
	fmt.Fprintf(w, "<b>Visit the access code URL:</b> <a href=\"%v\">%v</a>", url, url)
}

func Handshake(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	code := r.FormValue("code")
	t, err := igConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v\n", t)
	fmt.Fprintf(w, "<b>Your access token is:</b> %v", t.AccessToken)
}

func main() {
	config, err := util.LoadConfig(configFile)
	util.FailOnError(err)

	if config.ClientID == "" || config.ClientSecret == "" {
		err = fmt.Errorf("This app requires a registered API client. Please set your `client_id`, and `client_secret` config in '%v'.\n", configFile)
		util.FailOnError(err)
	}

	igConf = &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  instagramAuthURL,
			TokenURL: instagramTokenURL,
		},
	}

	http.HandleFunc("/", Home)
	http.HandleFunc("/handshake", Handshake)
	fmt.Printf("Listening and serving HTTP on %s\n", serverPort)
	http.ListenAndServe(serverPort, nil)
}

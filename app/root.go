package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}

	return 0
}

type appEnv struct {
	hc     http.Client
	domain string
}

func (app *appEnv) fromArgs(args []string) error {
	app.hc = *http.DefaultClient
	app.domain = ""
	if len(args) > 0 {
		app.domain = args[0]
	}

	flag.Usage = func() {
		fmt.Printf(`dontbore-cli - 0.0.1

Find and share logins provided by the community

Usage:

	dontbore-cli [domain]
`)
	}
	flag.Parse()

	return nil
}

func (app *appEnv) run() error {
	u := BuildURL(app.domain)

	var logins []Login
	if err := app.fetchJSON(u, &logins); err != nil {
		return err
	}

	return prettyPrint(logins)
}

func (app *appEnv) fetchJSON(url string, data interface{}) error {
	resp, err := app.hc.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func prettyPrint(logins []Login) error {
	var err error
	for i, login := range logins {
		_, err = fmt.Printf(
			"Index: %d\nDate: %s\nUsername: %s\nPassword: %s\nVote: %d\n",
			i,
			login.Date,
			login.User,
			login.Pass,
			login.Vote,
		)
	}

	return err
}

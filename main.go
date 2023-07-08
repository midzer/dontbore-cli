package main

import (
	"os"

	"github.com/midzer/dontbore-cli/app"
)

func main() {
	os.Exit(app.CLI(os.Args[1:]))
}

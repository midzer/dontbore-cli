package app

import (
	"fmt"
)

func BuildURL(domain string) string {
	return fmt.Sprintf("https://api.dontbo.re/%s", domain)
}

type Login struct {
	Date       string `json:"date"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	Vote       int    `json:"vote"`
}

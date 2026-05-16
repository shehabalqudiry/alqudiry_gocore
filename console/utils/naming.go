package utils

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var plural = pluralize.NewClient()

func Normalize(name string) string {

	name = strings.ToLower(name)

	name = strings.ReplaceAll(
		name,
		"-",
		"_",
	)

	return name
}

func Singular(name string) string {
	return plural.Singular(name)
}

func ModuleName(name string) string {

	return cases.Title(
		language.English,
	).String(
		Singular(name),
	)
}

func ToPublicName(name string) string {

	return cases.Title(
		language.English,
	).String(name)
}
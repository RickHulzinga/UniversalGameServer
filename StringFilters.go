package UniversalGameServer

import (
	"regexp"
)

var alphanumeric *regexp.Regexp

func initalizeRegex() {
	sm, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		Log.Fatal(err)
	}

	alphanumeric = sm
}

func getAlphanumericString(st string) string {

	return alphanumeric.ReplaceAllString(st, "")

}

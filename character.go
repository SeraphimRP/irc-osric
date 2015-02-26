// characters and moar by vypr
package main

import (
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"strconv"
)

var charmap map[string]*jason.Object

func importChar(nick string) bool {
	file, err := ioutil.ReadFile("json/" + nick + ".json")

	if err != nil {
		return false
	}

	char, _ := jason.NewObjectFromBytes(file)

	charmap = map[string]*jason.Object{
		nick: char,
	}

	return true
}

func exportChar(nick string) {
	// TODO
}

func printChar(nick string, cat string, scat string, item string) string {
	// TODO: Checking if nick exists, to prevent crashing the bot.

	if scat == "nil" {
		test, _ := charmap[nick].GetString(cat, item)

		if len(test) == 0 {
			val, _ := charmap[nick].GetInt64(cat, item)
			return strconv.FormatInt(val, 10)
		} else {
			val, _ := charmap[nick].GetString(cat, item)
			return val
		}

		return "does not exist"
	} else {
		test, _ := charmap[nick].GetString(cat, scat, item)

		if len(test) == 0 {
			val, _ := charmap[nick].GetInt64(cat, scat, item)
			return strconv.FormatInt(val, 10)
		} else {
			val, _ := charmap[nick].GetString(cat, scat, item)
			return val
		}

		return "does not exist"
	}
}

// characters and moar by vypr
package main

import (
	"github.com/antonholmquist/jason"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func printChar(nick string, cat string, scat string, item string) string {
	if _, err := os.Stat("json/" + nick + ".json"); os.IsNotExist(err) {
		return "does not exist"
	}

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

func setChar(nick string, cat string, scat string, item string, value string) bool {
	file, err := ioutil.ReadFile("json/" + nick + ".json")

	if err != nil {
		return false
	}

	lines := strings.Split(string(file), "\n")

	for i, l := range lines {
		if strings.Contains(l, item) {
			if scat == "nil" {
				test, _ := charmap[nick].GetString(cat, item)

				if len(test) == 0 {
					lines[i] = "\"" + item + "\": " + value + ","
				} else {
					lines[i] = "\"" + item + "\": \"" + value + "\","
				}
			} else {
				test, _ := charmap[nick].GetString(cat, scat, item)

				if len(test) == 0 {
					lines[i] = "\"" + item + "\": " + value + ","
				} else {
					lines[i] = "\"" + item + "\": \"" + value + "\","
				}
			}
		}
	}

	oput := strings.Join(lines, "\n")
	err = ioutil.WriteFile("json/"+nick+".json", []byte(oput), 0644)

	if err != nil {
		return false
	}

	if !importChar(nick) {
		return false
	}

	return true
}

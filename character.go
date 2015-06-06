// characters and moar by vypr
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Character struct {
	irc		string
	personal	PersonalData
	equipment	EquipmentData
	stats		StatData
	wealth		WealthData
}

type PersonalData struct {
	name		string
	lvl		int
	race		string
	xp		int
	height		int
	alignment	string
	classes		[]string
	weight		int
	sex		string
	hp		int
	ac		int
	age		int
}

type EquipmentData struct {
	armour		[]string
	weapons		[]string
	items		[]string
	missiles	[]string
}

type StatData struct {
	con	int
	str	int
	intl	int
	cha	int
	wis	int
	dex	int
}

type WealthData struct {
	coins	int
	other	[]string
	gems	[]string
}

var charmap map[string]*Character

func importChar(nick string) bool {
	file, err := ioutil.ReadFile("json/" + nick + ".json")

	if err != nil {
		return false
	}

	var char Character

	json.Unmarshal(file, char)
	
	charmap = map[string]*jason.Object{
		char.irc: char,
	}

	return true
}

func printChar(nick string, cat string, scat string, item string) string {
	if _, err := os.Stat("json/" + nick + ".json"); os.IsNotExist(err) {
		return "does not exist"
	}
	
	var fixedCat string
	
	
	
	switch (cat) {
		case "personal":
			fixedCat = "PersonalData"
			break
		case "equipment":
			fixedCat = "EquipmentData"
			break
		case "stats":
			fixedCat = "StatData"
			break
		case "wealth":
			fixedCat = "WealthData"
			break
		case "nil":
			fixedCat = "nil"
			break
		default:
			return "invalid category"
	}

	if scat == "nil" {
		return charmap[nick].fixedCat.
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

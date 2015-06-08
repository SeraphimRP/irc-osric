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

func accessChar(set bool, nick string, cat string, item string, val string) string {
	var pdata, edata, sdata, wdata bool
	
	switch (cat) {
		case "personal":
			pdata = true
			break
		case "equipment":
			edata = true
			break
		case "stats":
			sdata = true
			break
		case "wealth":
			wdata = true
			break
		case "nil":
			return charmap[nick].irc // can't be anything else but
			break
		default:
			return "invalid category"
	}
	
	if personal {
		switch (item) {
		case "name":
			return charmap[nick].personal.name
			break
		case "lvl":
			return strconv.Itoa(charmap[nick].personal.lvl)
			break
		case "race":
			return charmap[nick].personal.race
			break
		case "xp":
			return strconv.Itoa(charmap[nick].personal.xp)
			break
		case "height":
			return strconv.Itoa(charmap[nick].personal.height)
			break
		case "alignment":
			return charmap[nick].personal.alignment
			break
		case "classes":
			var classes string
			for i in range charmap[nick].personal.classes {
				if len(charmap[nick].personal.classes > 0 {
					if len(classes) > 0 {
						classes = classes + ", " + i	
					} else {
						classes = i
					}
				}
			}
			break
		case "weight":
			return strconv.Itoa(charmap[nick].personal.weight)
			break
		case "sex":
			return charmap[nick].personal.sex
			break
		case "hp":
			return strconv.Itoa(charmap[nick].personal.hp)
			break
		case "ac":
			return strconv.Itoa(charmap[nick].personal.ac)
			break
		case "age":
			return strconv.Itoa(charmap[nick].personal.age)
			break
		default:
			return "invalid item"
			break
		}
	} else if equipment {
		switch (item) {
		case "armour":
			var armour string
			for i in range charmap[nick].equipment.armour {
				if len(charmap[nick].equipment.armour) > 0 {
					if len(armour) > 0 {
						armour = armour + ", " + i	
					} else {
						armour = i
					}
				}
			}
			break
		case "weapons":
			var weapons string
			for i in range charmap[nick].equipment.weapons {
				if len(charmap[nick].equipment.weapons) > 0 {
					if len(weapons) > 0 {
						weapons = weapons + ", " + i	
					} else {
						weapons = i
					}
				}
			}
			break
		case "items":
			var items string
			for i in range charmap[nick].equipment.items {
				if len(charmap[nick].equipment.items) > 0 {
					if len(items) > 0 {
						items = items + ", " + i	
					} else {
						items = i
					}
				}
			}
			break
		case "missiles":
			var missiles string
			for i in range charmap[nick].equipment.missiles {
				if len(charmap[nick].equipment.missiles) > 0 {
					if len(missiles) > 0 {
						missiles = missiles + ", " + i	
					} else {
						missiles = i
					}
				}
			}
			break
		default:
			return "invalid item"
			break
		}
	} else if stats {
		switch (item) {
		case "con":
			return strconv.Itoa(charmap[nick].stats.con)
			break
		case "str":
			return strconv.Itoa(charmap[nick].stats.str)
			break
		case "intl":
			return strconv.Itoa(charmap[nick].stats.intl)
			break
		case "cha":
			return strconv.Itoa(charmap[nick].stats.cha)
			break
		case "wis":
			return strconv.Itoa(charmap[nick].stats.wis)
			break
		case "dex":
			return strconv.Itoa(charmap[nick].stats.dex)
			break
		default:
			return "invalid item"
			break
		}
	} else if wealth {
		switch (item) {
		case "coins":
			return strconv.Itoa(charmap[nick].wealth.coins)
			break
		case "other":
			var other string
			for i in range charmap[nick].wealth.other {
				if len(charmap[nick].wealth.other) > 0 {
					if len(other) > 0 {
						other = other + ", " + i	
					} else {
						other = i
					}
				}
			}
			break
		case "gems":
			var gems string
			for i in range charmap[nick].wealth.gems {
				if len(charmap[nick].wealth.gems) > 0 {
					if len(gems) > 0 {
						gems = gems + ", " + i	
					} else {
						gems = i
					}
				}
			}
			break
		}
	}
}

func importChar(nick string) bool {
	file, err := ioutil.ReadFile("json/" + nick + ".json")

	if err != nil {
		return false
	}

	var char Character

	json.Unmarshal(file, char)
	
	charmap = map[string]*Character{
		char.irc: char,
	}

	return true
}

func setChar(nick string, cat string, scat string, item string, value string) bool {
	// With the changes I'm currently making, this will have to be changed as well.
}

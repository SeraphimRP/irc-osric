// characters and moar by vypr
package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Character struct {
	Equipment struct {
		Armour   []interface{} `json:"armour"`
		Items    []interface{} `json:"items"`
		Missiles []interface{} `json:"missiles"`
		Weapons  []interface{} `json:"weapons"`
	} `json:"equipment"`
	Personal struct {
		Ac        int      `json:"ac"`
		Age       int      `json:"age"`
		Alignment string   `json:"alignment"`
		Classes   []string `json:"classes"`
		Height    int      `json:"height"`
		Hp        int      `json:"hp"`
		Lvl       int      `json:"lvl"`
		Name      string   `json:"name"`
		Race      string   `json:"race"`
		Sex       string   `json:"sex"`
		Weight    int      `json:"weight"`
		Xp        int      `json:"xp"`
	} `json:"personal"`
	Stats struct {
		Cha  int `json:"cha"`
		Con  int `json:"con"`
		Dex  int `json:"dex"`
		Intl int `json:"intl"`
		Str  int `json:"str"`
		Wis  int `json:"wis"`
	} `json:"stats"`
	Wealth struct {
		Coins int           `json:"coins"`
		Gems  []interface{} `json:"gems"`
		Other []interface{} `json:"other"`
	} `json:"wealth"`
}

var charmap map[string]*Character

func accessChar(nick string, cat string, item string) string {
	var pdata, edata, sdata, wdata bool

	switch cat {
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
	default:
		return "invalid category"
	}

	if pdata {
		switch item {
		case "name":
			return charmap[nick].Personal.Name
			break
		case "lvl":
			return strconv.Itoa(charmap[nick].Personal.Lvl)
			break
		case "race":
			return charmap[nick].Personal.Race
			break
		case "xp":
			return strconv.Itoa(charmap[nick].Personal.Xp)
			break
		case "height":
			return strconv.Itoa(charmap[nick].Personal.Height)
			break
		case "alignment":
			return charmap[nick].Personal.Alignment
			break
		case "classes":
			var classes string
			for _, j := range charmap[nick].Personal.Classes {
				if len(charmap[nick].Personal.Classes) > 0 {
					if len(classes) > 0 {
						classes = classes + ", " + j
					} else {
						classes = j
					}
				}
			}
			if len(classes) > 0 {
				return classes
			}
			return "you've got no classes, m80"
			break
		case "weight":
			return strconv.Itoa(charmap[nick].Personal.Weight)
			break
		case "sex":
			return charmap[nick].Personal.Sex
			break
		case "hp":
			return strconv.Itoa(charmap[nick].Personal.Hp)
			break
		case "ac":
			return strconv.Itoa(charmap[nick].Personal.Ac)
			break
		case "age":
			return strconv.Itoa(charmap[nick].Personal.Age)
			break
		default:
			return "invalid item"
			break
		}
	} else if edata {
		switch item {
		case "armour":
			var armour string
			for _, j := range charmap[nick].Equipment.Armour {
				value, _ := j.(string)
				if len(charmap[nick].Equipment.Armour) > 0 {
					if len(armour) > 0 {
						armour = armour + ", " + value
					} else {
						armour = value
					}
				}
			}
			if len(armour) > 0 {
				return armour
			}
			return "you've got no armour, m80"
			break
		case "weapons":
			var weapons string
			for _, j := range charmap[nick].Equipment.Weapons {
				value, _ := j.(string)
				if len(charmap[nick].Equipment.Weapons) > 0 {
					if len(weapons) > 0 {
						weapons = weapons + ", " + value
					} else {
						weapons = value
					}
				}
			}
			if len(weapons) > 0 {
				return weapons
			}
			return "you've got no weapons, m80"
			break
		case "items":
			var items string
			for _, j := range charmap[nick].Equipment.Items {
				value, _ := j.(string)
				if len(charmap[nick].Equipment.Items) > 0 {
					if len(items) > 0 {
						items = items + ", " + value
					} else {
						items = value
					}
				}
			}
			if len(items) > 0 {
				return items
			}
			return "you've got no items, m80"
			break
		case "missiles":
			var missiles string
			for _, j := range charmap[nick].Equipment.Missiles {
				value, _ := j.(string)
				if len(charmap[nick].Equipment.Missiles) > 0 {
					if len(missiles) > 0 {
						missiles = missiles + ", " + value
					} else {
						missiles = value
					}
				}
			}
			if len(missiles) > 0 {
				return missiles
			}
			return "you've got no missiles, m80"
			break
		default:
			return "invalid item"
			break
		}
	} else if sdata {
		switch item {
		case "con":
			return strconv.Itoa(charmap[nick].Stats.Con)
			break
		case "str":
			return strconv.Itoa(charmap[nick].Stats.Str)
			break
		case "intl":
			return strconv.Itoa(charmap[nick].Stats.Intl)
			break
		case "cha":
			return strconv.Itoa(charmap[nick].Stats.Cha)
			break
		case "wis":
			return strconv.Itoa(charmap[nick].Stats.Wis)
			break
		case "dex":
			return strconv.Itoa(charmap[nick].Stats.Dex)
			break
		default:
			return "invalid item"
			break
		}
	} else if wdata {
		switch item {
		case "coins":
			return strconv.Itoa(charmap[nick].Wealth.Coins)
			break
		case "other":
			var other string
			for _, j := range charmap[nick].Wealth.Other {
				value, _ := j.(string)
				if len(charmap[nick].Wealth.Other) > 0 {
					if len(other) > 0 {
						other = other + ", " + value
					} else {
						other = value
					}
				}
			}
			if len(other) > 0 {
				return other
			}
			return "you've got no other items, m80"
			break
		case "gems":
			var gems string
			for _, j := range charmap[nick].Wealth.Gems {
				value, _ := j.(string)
				if len(charmap[nick].Wealth.Gems) > 0 {
					if len(gems) > 0 {
						gems = gems + ", " + value
					} else {
						gems = value
					}
				}
			}
			if len(gems) > 0 {
				return gems
			}
			return "you've got no gems, m80"
			break
		}
	}

	return "*shrug*"
}

func importChar(nick string) bool {
	file, err := ioutil.ReadFile("json/" + nick + ".json")

	if err != nil {
		panic(err)
	}

	var char Character

	err = json.Unmarshal([]byte(file), &char)

	if err != nil {
		panic(err)
	}

	charmap = map[string]*Character{
		nick: &char,
	}

	return true
}

func setChar(nick string, cat string, item string, value string) bool {
	var pdata, edata, sdata, wdata bool

	switch cat {
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
	default:
		return false
	}

	if pdata {
		switch item {
		case "name":
			charmap[nick].Personal.Name = value
			break
		case "lvl":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Lvl = value
			break
		case "race":
			charmap[nick].Personal.Race = value
			break
		case "xp":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Xp = value
			break
		case "height":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Height = value
			break
		case "alignment":
			charmap[nick].Personal.Alignment = value
			break
		case "classes":
			charmap[nick].Personal.Classes = append(charmap[nick].Personal.Classes, value)
			break
		case "weight":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Weight = value
			break
		case "sex":
			charmap[nick].Personal.Sex = value
			break
		case "hp":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Hp = value
			break
		case "ac":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Ac = value
			break
		case "age":
			value, _ := strconv.Atoi(value)
			charmap[nick].Personal.Age = value
			break
		default:
			return false
			break
		}
	} else if edata {
		switch item {
		case "armour":
			var armour string
			armour = value
			charmap[nick].Equipment.Armour = append(charmap[nick].Equipment.Armour, armour)
			break
		case "weapons":
			var weapons string
			weapons = value
			charmap[nick].Equipment.Weapons = append(charmap[nick].Equipment.Weapons, weapons)
			break
		case "items":
			var items string
			items = value
			charmap[nick].Equipment.Items = append(charmap[nick].Equipment.Items, items)
			break
		case "missiles":
			var missiles string
			missiles = value
			charmap[nick].Equipment.Missiles = append(charmap[nick].Equipment.Missiles, missiles)
			break
		default:
			return false
			break
		}
	} else if sdata {
		switch item {
		case "con":
			value, _ := strconv.Atoi(value)
			charmap[nick].Stats.Con = value
			break
		case "str":
			value, _ := strconv.Atoi(value)
			charmap[nick].Stats.Str = value
			break
		case "intl":
			value, _ := strconv.Atoi(value)
			charmap[nick].Stats.Intl = value
			break
		case "cha":
			value, _ := strconv.Atoi(value)
			charmap[nick].Stats.Cha = value
			break
		case "wis":
			value, _ := strconv.Atoi(value)
			charmap[nick].Stats.Wis = value
			break
		case "dex":
			value, _ := strconv.Atoi(value)
			charmap[nick].Stats.Dex = value
			break
		default:
			return false
			break
		}
	} else if wdata {
		switch item {
		case "coins":
			value, _ := strconv.Atoi(value)
			charmap[nick].Wealth.Coins = value
			break
		case "other":
			var other string
			other = value
			charmap[nick].Wealth.Other = append(charmap[nick].Wealth.Other, other)
			break
		case "gems":
			var gems string
			gems = value
			charmap[nick].Wealth.Gems = append(charmap[nick].Wealth.Gems, gems)
			break
		default:
			return false
			break
		}
	}

	data, _ := json.MarshalIndent(charmap[nick], "", "    ")

	err := ioutil.WriteFile("json/"+nick+".json", data, 0644)

	if err != nil {
		return false
	}

	return true
}

package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bot struct {
	Conn    *irc.Connection
	Nick    string
	Server  string
	Channel string
}

var (
	admins = []string{"vypr", "bajr", "kirby"}
	dieopt = []string{"4", "6", "8", "10", "12", "20"}
	dunmas = ""

	rulemod = make([]string, len(modeopt))
	modeopt = []string{"adminoverride", "saves", "logging"}
	initLog = true

	charmap = make(map[string]map[string]map[string]string)
	monsmap = make(map[string]string)

	filename = "log/log_" + time.Now().Local().Format("20060102") + "_" + time.Now().Local().Format("150405") + ".txt"
)

var dict = map[string]string{
	"hp":   "health points",
	"ap":   "armour points",
	"algn": "alignment",
	"xp":   "experience points",
	"str":  "strength",
	"con":  "constitution",
	"dex":  "dexterity",
	"wis":  "wisdom",
	"cha":  "charisma",
	"lvl":  "level",
	"hgt":  "height",
	"wgt":  "weight",
	"cls":  "class",
}

var argmap = map[string]int{
	".set":     4,
	".print":   3,
	".d":       1,
	".mode":    1,
	".rmmode":  1,
	".dm":      1,
	".resetdm": 1,
	".quit":    0,
}

func fillCharmap(nick string, cat string, item string, val string) {
	charmap = map[string]map[string]map[string]string{
		nick: map[string]map[string]string{
			cat: map[string]string{
				item: val,
			},
		},
	}
}

func roll(amount int, side int) int {
	var numbers = make([]int, amount)
	var finaln = 0

	for i := 0; i < amount; i++ {
		number := rand.Intn(side) + 1
		numbers[i] = number
	}

	for i := 0; i < len(numbers); i++ {
		if i == len(numbers) {
			return finaln
		}

		finaln = finaln + numbers[i]
	}

	return finaln
}

// TODO: Create functions related to character import/export.

func (b *Bot) Command(nick string, msg string) {
	var args = make([]string, len(strings.Split(msg, " "))-1)
	var command = ""

	if stringInSlice(modeopt[2], rulemod) {
		b.Log(nick+": "+msg, initLog)
		initLog = false
	}

	for i, j := range strings.Split(msg, " ") {
		if j != " " && i != 0 {
			args[i-1] = strings.Split(msg, " ")[i]
		}
	}

	for i := range argmap {
		if i == strings.Split(msg, " ")[0] {
			command = strings.Split(msg, " ")[0]
			break
		}
	}

	if argmap[strings.Split(msg, " ")[0]] != len(args) {
		return
	}

	switch command {
	case ".set":
		if nick == dunmas {
			fillCharmap(args[0], args[1], args[2], args[3])
			fmt.Println("[cmd] set - " + args[0] + "'s " + args[2] + "in " + args[1] + " is set to " + args[3])
		} else if stringInSlice(nick, admins) && !stringInSlice(modeopt[0], rulemod) {
			fillCharmap(args[0], args[1], args[2], args[3])
			fmt.Println("[cmd] set - " + args[0] + "'s " + args[2] + " in " + args[1] + " is set to " + args[3])
			b.Say(nick + " used override, it's super effective!")
		}
		break

	case ".print":
		if len(charmap[args[0]][args[1]][args[2]]) == 0 {
			b.Say("there is no setting for " + args[0] + "'s " + args[2])
		} else {
			fmt.Println("[cmd] print - " + args[0] + "'s " + args[2] + " in " + args[1])
			b.Say(args[0] + "'s " + args[2] + " is set to " + charmap[args[0]][args[1]][args[2]])
		}
		break

	case ".d":
		amount, _ := strconv.Atoi(strings.Split(args[0], "d")[0])
		side, _ := strconv.Atoi(strings.Split(args[0], "d")[1])

		if stringInSlice(strconv.Itoa(side), dieopt) && len(strings.Split(args[0], "d")) < 3 && amount > 0 && amount <= 20 {
			fmt.Println("[cmd] rolling " + args[0])
			b.Say(nick + " rolled a " + strconv.Itoa(roll(amount, side)))
		}
		break

	case ".mode":
		if stringInSlice(args[0], rulemod) {
			b.Say(args[0] + " is already set to true")
		} else if stringInSlice(args[0], modeopt) {
			fmt.Println("[cmd] mode - " + args[0])
			b.Say(args[0] + " is now enabled")
			rulemod = append(rulemod, args[0])
		}
		break

	case ".rmmode":
		if removeItemInSlice(args[0], rulemod) {
			fmt.Println("[cmd] rmmode - " + args[0])
			b.Say(args[0] + " has been removed from the list of modes")
		} else {
			b.Say(args[0] + " isn't in the list of modes")
		}
		break

	case ".dm":
		if len(dunmas) == 0 {
			dunmas = args[0]
			fmt.Println("[cmd] dm - " + dunmas)
			b.Say("dm is now set to " + dunmas)
		} else {
			b.Say("dm has already been set, the current DM is " + dunmas)
		}
		break

	case ".resetdm":
		if nick == dunmas || stringInSlice(nick, admins) {
			dunmas = ""
			fmt.Println("[cmd] resetdm")
			b.Say("dm has been reset")
		}
		break

	case ".quit":
		if stringInSlice(nick, admins) {
			fmt.Println("[cmd] shutdown from " + nick)
			os.Exit(1)
		}
		break
	}
}

func (b *Bot) Log(line string, initLog bool) {
	if initLog {
		os.Create(filename)

		file, fileerr := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)

		if fileerr != nil {
			panic(fileerr)
		}

		file.WriteString("IRC-OSRIC by Elliott Pardee (vypr)\n")
		file.WriteString("----------------------------------\n\n")
		file.WriteString(line + "\n")
	} else {
		file, fileerr := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)

		if fileerr != nil {
			panic(fileerr)
		}

		file.WriteString(line + "\n")
	}
}

func (b *Bot) Say(msg string) {
	if stringInSlice(modeopt[2], rulemod) {
		b.Log("bot: "+msg, initLog)
	}

	b.Conn.Privmsg(b.Channel, msg)
}

func (b *Bot) Listen() {
	err := b.Conn.Connect(b.Server)

	if err != nil {
		panic(err)
	}

	b.Conn.AddCallback("001", func(e *irc.Event) {
		b.Conn.Join(b.Channel)
	})

	b.Conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		b.Command(e.Nick, e.Message())
	})

	b.Conn.Loop()
}

func NewBot(server string, channel string, nick string) *Bot {
	return &Bot{Conn: irc.IRC(nick, nick), Server: server, Channel: channel, Nick: nick}
}

func main() {
	b := NewBot("irc.iotek.org:6667", "#d20", "bot")
	b.Listen()
}

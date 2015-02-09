package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"os"
	"strings"
    "time"
    "math/rand"
    "strconv"
)

type Bot struct {
	Conn    *irc.Connection
	Nick    string
	Server  string
	Channel string
}

var (
	admins = []string{"vypr", "bajr"}
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
    ".die":     2,
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

func die(side int, amount int) {
    // TODO: Works, just not correctly.
    var numbers = make([]int, amount + 1)
    var tmpnum = make([]int, amount)

    for i := 0; i < amount; i++ {
        r1 := rand.NewSource(0)
        r2 := rand.New(r1)
        number := r2.Intn(side)
        numbers[i] = number

        if amount > i {
            tmpnum[i] = numbers[i + 1] + number
        }

        fmt.Println(tmpnum[i])
    }
}

// TODO: Create functions related to character import/export.

func (b *Bot) Command(nick string, msg string) {
	var args = make([]string, len(strings.Split(msg, " "))-1)
	var command = ""

	if stringInSlice(modeopt[2], rulemod) {
		b.Log(nick + ": " + msg, initLog)
	    initLog = false
    }

	for i, j := range strings.Split(msg, " ") {
		if j != " " && i != 0 {
			args[i - 1] = strings.Split(msg, " ")[i]
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

	case ".die":
        arg1, _ := strconv.Atoi(args[0])
        arg2, _ := strconv.Atoi(args[1])
        die(arg1, arg2)
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
		b.Log("bot: " + msg, initLog)
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

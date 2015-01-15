package main

import (
    "github.com/thoj/go-ircevent"
    "strings"
    "fmt"
    "os"
)

type Bot struct {
    Conn    *irc.Connection
    Nick    string
    Server  string
    Channel string
}

var (
    admins = []string{"vypr"}
    dunmas = ""

    rulemod = make([]string, len(modeopt))
    modeopt = []string{"adminoverride", "saves", "logging"}

    charmap = make(map[string]map[string]map[string]string)
    monsmap = make(map[string]string)
)

var dict = map[string]string{
    "hp": "health points",
    "ap": "armour points",
    "algn": "alignment",
    "xp": "experience points",
    "str": "strength",
    "dex": "dexterity",
    "wis": "wisdom",
    "cha": "charisma",
    "lvl": "level",
    "hgt": "height",
    "wgt": "weight",
    "cls": "class",
}

var argmap = map[string]int{
    ".set": 4,
    ".print": 3,
    ".mode": 1,
    ".rmmode": 1,
    ".dm": 1,
    ".resetdm": 1,
    ".quit": 0,
}

func fillCharmap(nick string, cat string, item string, val string) {
    charmap = map[string]map[string]map[string]string { nick: map[string]map[string]string{ cat: map[string]string{ item: val, }, }, }
}

func (b *Bot) Command(nick string, msg string) {
    var args = make([]string, len(strings.Split(msg, " ")) - 1)
    var command = ""

    for i, j := range strings.Split(msg, " ") {
        if j != " " && i !=0 {
            args[i - 1] = strings.Split(msg, " ")[i]
        }
    }

    for i := range argmap {
        if i == strings.Split(msg, " ")[0] {
            command = strings.Split(msg, " ")[0]
            break
        }
    }

    if argmap[strings.Split(msg, " ")[0]] != len(args) { return }

    // TODO: Check if mode is enabled and if command can be applied.

    switch command {
    case ".set":
        if nick == dunmas {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2]  + "in " + args[1] + " is set to " + args[3])
        } else if stringInSlice(nick, admins) && !stringInSlice("adminoverride", rulemod) {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2] + " in " + args[1] + " is set to " + args[3])
            b.Say(nick + " used override, it's super effective!")
        }
        break

    case ".print":
        fmt.Println("[cmd] print - " + args[0] + "'s " + args[2] + " in " + args[1])
        b.Say(args[0] + "'s " + args[2] + " is set to " + charmap[args[0]][args[1]][args[2]])
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

func (b *Bot) Say(msg string) {
    b.Conn.Privmsg(b.Channel, msg)
}

func (b *Bot) Listen() {
    err := b.Conn.Connect(b.Server)

    if err != nil { panic(err) }

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

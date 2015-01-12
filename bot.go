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

func fillCharmap(nick string, cat string, item string, val string) {
    charmap = map[string]map[string]map[string]string { nick: map[string]map[string]string{ cat: map[string]string{ item: val, }, }, }
}

func (b *Bot) Command(nick string, msg string) {
    var args = make([]string, len(strings.Split(msg, " ")) - 1)

    for i, j := range strings.Split(msg, " ") {
        if j != " " && i !=0 {
            args[i - 1] = strings.Split(msg, " ")[i]
        }
    }

    // TODO: Check if mode is enabled and if command can be applied.

    if strings.HasPrefix(msg, ".set") && len(args) == 4 {
        if nick == dunmas {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2]  + "in " + args[1] + " is set to " + args[3] + ".")
        } else if stringInSlice(nick, admins) && !stringInSlice("nocharoverride", rulemod) {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2] + " in " + args[1] + " is set to " + args[3] + ".")
            b.Say(nick + " used override, it's super effective!")
        }
    } else if strings.HasPrefix(msg, ".print") && len(args) == 3 {
        fmt.Println("[cmd] print - " + args[0] + "'s " + args[2] + " in " + args[1] + ".")
        b.Say(args[0] + "'s " + args[2] + " is set to " + charmap[args[0]][args[1]][args[2]] + ".")
    } else if strings.HasPrefix(msg, ".mode") && len(args) == 1 {
        if stringInSlice(args[0], rulemod) {
            fmt.Println("[cmd] mode - change to " + args[0] + " failed, already set to true")
            b.Say(args[0] + " is already set to true.")
        } else {
            fmt.Println("[cmd] mode - " + args[0])
            b.Say(args[0] + " is now enabled.")
        }
    } else if strings.HasPrefix(msg, ".rmmode") && len(args) == 1 {
        if removeItemInSlice(args[0], rulemod) {
            fmt.Println("[cmd] rmmode - " + args[0])
            b.Say(args[0] + " has been removed from the list of modes.")
        } else {
            b.Say(args[0] + " isn't in the list of modes.")
        }
    } else if strings.HasPrefix(msg, ".dm") && len(args) == 1 {
        if len(dunmas) == 0 {
            dunmas = args[0]
            fmt.Println("[cmd] dm - " + dunmas)
            b.Say("dm is now set to " + dunmas)
        } else {
            b.Say("dm has already been set, the current DM is " + dunmas)
        }
    } else if msg == ".resetdm" && (nick == dunmas || stringInSlice(nick, admins)) {
        dunmas = ""
        fmt.Println("[cmd] resetdm")
        b.Say("dm has been reset")
    } else if msg == ".quit" && stringInSlice(nick, admins) {
        fmt.Println("[cmd] shutdown from " + nick)
        os.Exit(1)
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

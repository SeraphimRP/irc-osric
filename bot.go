package main

import (

    "github.com/thoj/go-ircevent"
    "strings"
    "strconv"
    "fmt"
    "os"
)

var (
    server = "irc.iotek.org"
    port = 6667
    channel = "#d20"
    nickname = "bot"

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

func stringInSlice(a string, list []string) bool {
    // thanks stackoverflow
    for _, b := range list {
        if b == a { return true }
    }
    return false
}

func removeItemInSlice(a string, list []string) bool {
    for c, d := range list {
        if d == a {
            list[c] = ""
            return true
        }
    }
    return false
}

func findArguments(msg string) []string {
    var args = make([]string, len(strings.Split(msg, " ")) - 1) // thanks jmbi
    var nmsg = strings.Split(msg, " ")

    for i, j := range nmsg {
        if j != " " && i != 0 {
            args[i - 1] = nmsg[i]
        }
    }

    return args
}

func loadSheet(file string) {
    // TODO: JSON character sheets.
    // http://github.com/kirbyman62/osric-character-sheet-to-json
}

func exportSheet(nick string) {
    // TODO
}

func save() {
    // TODO
}

func fillCharmap(nick string, cat string, item string, val string) {
    // thanks jmbi
    charmap = map[string]map[string]map[string]string { nick: map[string]map[string]string{ cat: map[string]string{ item: val, }, }, }
}

func handleMessage(nick string, msg string, conn *irc.Connection) {
    var args = findArguments(msg)

    // TODO: Check if mode is enabled and if command can be applied.

    if strings.HasPrefix(msg, ".set") && len(args) == 4 {
        if nick == dunmas {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2]  + "in " + args[1] + " is set to " + args[3] + ".")
        } else if stringInSlice(nick, admins) && !stringInSlice("nocharoverride", rulemod) {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2] + " in " + args[1] + " is set to " + args[3] + ".")
            conn.Privmsg(channel, nick + " used override, it's super effective!")
        }
    } else if strings.HasPrefix(msg, ".print") && len(args) == 3 {
        fmt.Println("[cmd] print - " + args[0] + "'s " + args[1] + " in " + args[3] + ".")
        conn.Privmsg(channel, args[0] + "'s " + args[0] + " is set to " + charmap[args[0]][args[1]][args[2]] + ".")
    } else if strings.HasPrefix(msg, ".mode") && len(args) == 1 {
        if stringInSlice(args[0], rulemod) {
            fmt.Println("[cmd] mode - change to " + args[0] + " failed, already set to true")
            conn.Privmsg(channel, args[0] + " is already set to true.")
        } else {
            fmt.Println("[cmd] mode - " + args[0])
            conn.Privmsg(channel, args[0] + " is now enabled.")
        }
    } else if strings.HasPrefix(msg, ".rmmode") && len(args) == 1 {
        if removeItemInSlice(args[0], rulemod) {
            fmt.Println("[cmd] rmmode - " + args[0])
            conn.Privmsg(channel, args[0] + " has been removed from the list of modes.")
        } else {
            conn.Privmsg(channel, args[0] + " isn't in the list of modes.")
        }
    } else if strings.HasPrefix(msg, ".dm") && len(args) == 1 {
        if len(dunmas) == 0 {
            dunmas = args[0]
            fmt.Println("[cmd] dm - " + dunmas)
            conn.Privmsg(channel, "dm is now set to " + dunmas)
        } else {
            conn.Privmsg(channel, "dm has already been set, the current DM is " + dunmas)
        }
    } else if msg == ".resetdm" && (nick == dunmas || stringInSlice(nick, admins)) {
        dunmas = ""
        fmt.Println("[cmd] resetdm")
        conn.Privmsg(channel, "dm has been reset")
    } else if msg == ".quit" && stringInSlice(nick, admins) {
        fmt.Println("[cmd] shutdown from " + nick)
        os.Exit(1)
    }
}

func main() {
    conn := irc.IRC(nickname, nickname)

    err := conn.Connect(server + ":" + strconv.Itoa(port))
    if err != nil {
        fmt.Print("[err] connection failed - ")
        fmt.Println(err)
    }

    conn.AddCallback("001", func(e *irc.Event) { conn.Join(channel) })

    conn.AddCallback("PRIVMSG", func(e *irc.Event) {
        handleMessage(e.Nick, e.Message(), conn)
    })

    conn.Loop()
}

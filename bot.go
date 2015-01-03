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

    rulemod = []string{"nocharoverride"}

    charmap = make(map[string]map[string]map[string]string)
    monsmap = make(map[string]string)
)

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
    var args []string
    var nmsg = strings.Split(msg, " ")

    fmt.Println(nmsg)

    for i, j := range nmsg {
        if j != " " || i != 0 {
            args[i - 1] = nmsg[i]
        }
    }

    return args
}

func fillCharmap(nick string, cat string, item string, val string) {
    charmap = map[string]map[string]map[string]string { nick: map[string]map[string]string{ cat: map[string]string{ item: val, }, }, }

func handleMessage(nick string, msg string, conn *irc.Connection) {
    var args = findArguments(msg)
    var argc = len(args)

    if strings.HasPrefix(msg, ".set") && argc == 5 {
        // thanks to jmbi (github.com/karlmcg) for helping me think
        if nick == dunmas {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2]  + "in " + args[1] + " is set to " + args[3] + ".")
        } else if stringInSlice(nick, admins) && !stringInSlice("nocharoverride", rulemod) {
            fillCharmap(args[0], args[1], args[2], args[3])
            fmt.Println("[cmd] set - " + args[0] + "'s " + args[2] + " in " + args[1] + " is set to " + args[3] + ".")
            conn.Privmsg(channel, nick + " used override, it's super effective!")
        }
    } else if strings.HasPrefix(msg, ".print") && argc == 4 {
        fmt.Println("[cmd] print - " + args[0] + "'s " + args[1] + " in " + args[3] + ".")
        conn.Privmsg(channel, args[0] + "'s " + args[0] + " is set to " + charmap[args[0]][args[1]][args[2]] + ".")
    } else if strings.HasPrefix(msg, ".mode") && argc == 2 {
        if stringInSlice(args[0], rulemod) {
            fmt.Println("[cmd] mode - change failed, already set to true")
            conn.Privmsg(channel, args[0] + " is already set to true.")
        } else {
            fmt.Println("[cmd] mode " + args[0])
            conn.Privmsg(channel, "set " + args[0] + " to true.")
        }
    } else if strings.HasPrefix(msg, ".rmmode") && argc == 2 {
        if removeItemInSlice(args[0], rulemod) {
            fmt.Println("[cmd] rmmode - " + args[0])
            conn.Privmsg(channel, args[0] + " has been removed from the list of modes.")
        } else {
            conn.Privmsg(channel, args[0] + " isn't in the list of modes.")
        }
    } else if strings.HasPrefix(msg, ".dm") && argc == 2 {
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

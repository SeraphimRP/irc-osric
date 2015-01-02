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

func processMsg(nick string, msg string, conn *irc.Connection) {
    // At some point, this will probably need refactoring.
    if strings.HasPrefix(msg, ".set") && len(strings.Split(msg, " ")) == 5 {
        var arg1 = strings.Split(msg, " ")[1] // nick
        var arg2 = strings.Split(msg, " ")[2] // category
        var arg3 = strings.Split(msg, " ")[3] // stat/info
        var arg4 = strings.Split(msg, " ")[4] // value

        if nick == dunmas {
            var charmap { arg1 { arg2 { arg3: arg4, }, }, }
            fmt.Println("[cmd] set " + arg1 + "'s " + arg3  + "in " + arg2 + " is set to " + arg4 + ".")
        } else if stringInSlice(nick, admins) {
            var charmap = { arg1 = { arg2 = { arg3: arg4, }, }, }
            fmt.Println("[cmd] set " + arg1 + "'s " + arg3 + " in " + arg2 + " is set to " + arg4 + ".")
            conn.Privmsg(channel, nick + " used override, it's super effective!")
        }
    } else if msg == ".print" && len(strings.Split(msg, " ")) == 4 {
        var arg1 = strings.Split(msg, " ")[1] // nick
        var arg2 = strings.Split(msg, " ")[2] // category
        var arg3 = strings.Split(msg, " ")[3] // stat/info

        fmt.Println("[cmd] print " + arg1 + "'s " + arg3 + "in " + arg2 + ".")
        conn.Privmsg(channel, arg1 + "'s " + arg3 + " is set to " + charmap[arg1][arg2][arg3] + ".")
    } else if strings.HasPrefix(msg, ".dm") && len(strings.Split(msg, " ")) == 2 {
        if len(dunmas) == 0 && len(strings.Split(msg, " ")[1]) > 0 {
            dunmas = strings.Split(msg, " ")[1]
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
        processMsg(e.Nick, e.Message(), conn)
    })

    conn.Loop()
}

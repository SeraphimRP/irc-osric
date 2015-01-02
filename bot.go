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
    for _, b := range list {
        if b == a {
            b = " "
            return true
        }
    }
    return false
}

func processMsg(nick string, msg string, conn *irc.Connection) {
    // At some point, this will probably need refactoring.
    if strings.HasPrefix(msg, ".set") && len(strings.Split(msg, " ")) == 5 {

        // thanks to jmbi (github.com/karlmcg) for helping me think
        var arg1 = strings.Split(msg, " ")[1] // nick
        var arg2 = strings.Split(msg, " ")[2] // category
        var arg3 = strings.Split(msg, " ")[3] // stat/info
        var arg4 = strings.Split(msg, " ")[4] // value

        if arg1 == " " || arg2 == " " || arg3 == " " || arg4 == " " {
            conn.Privmsg(channel, "invalid setting")
            return
        }

        if nick == dunmas {
            charmap = map[string]map[string]map[string]string{ arg1: map[string]map[string]string{ arg2: map[string]string{ arg3: arg4, }, }, }
            fmt.Println("[cmd] set - " + arg1 + "'s " + arg3  + "in " + arg2 + " is set to " + arg4 + ".")
        } else if stringInSlice(nick, admins) && !stringInSlice("nocharoverride", rulemod) {
            charmap = map[string]map[string]map[string]string{ arg1: map[string]map[string]string{ arg2: map[string]string{ arg3: arg4, }, }, }
            fmt.Println("[cmd] set - " + arg1 + "'s " + arg3 + " in " + arg2 + " is set to " + arg4 + ".")
            conn.Privmsg(channel, nick + " used override, it's super effective!")
        }
    } else if strings.HasPrefix(msg, ".print") && len(strings.Split(msg, " ")) == 4 {
        var arg1 = strings.Split(msg, " ")[1] // nick
        var arg2 = strings.Split(msg, " ")[2] // category
        var arg3 = strings.Split(msg, " ")[3] // stat/info

        if arg1 == " " || arg2 == " " || arg3 == " " {
            conn.Privmsg(channel, "invalid setting")
            return
        }

        fmt.Println("[cmd] print - " + arg1 + "'s " + arg3 + " in " + arg2 + ".")
        conn.Privmsg(channel, arg1 + "'s " + arg3 + " is set to " + charmap[arg1][arg2][arg3] + ".")
    } else if strings.HasPrefix(msg, ".mode") && len(strings.Split(msg, " ")) == 2 {
       if len(strings.Split(msg, " ")[1]) > 0 {
            if stringInSlice(strings.Split(msg, " ")[1], rulemod) {
                conn.Privmsg(channel, strings.Split(msg, " ")[1] + " is already set to true.")
            } else {
                fmt.Println("[cmd] mode " + strings.Split(msg, " ")[1])
                conn.Privmsg(channel, "set " + strings.Split(msg, " ")[1] + " to true.")
            }
       }
    } else if strings.HasPrefix(msg, ".rmmode") && len(strings.Split(msg, " ")) == 2 {
        if len(strings.Split(msg, " ")[1]) > 0 {
            if stringInSlice(strings.Split(msg, " ")[1], rulemod) {
                if removeItemInSlice(strings.Split(msg, " ")[1], rulemod) {
                    fmt.Println("[cmd] rmmode - " + strings.Split(msg, " ")[1])
                    conn.Privmsg(channel, strings.Split(msg, " ")[1] + " has been removed from the list of modes.")
                }
            } else {
                conn.Privmsg(channel, strings.Split(msg, " ")[1] + " isn't in the list of modes.")
            }
        }
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

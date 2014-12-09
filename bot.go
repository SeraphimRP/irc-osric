package main

import (

    "github.com/thoj/go-ircevent"
    "strings"
    "strconv"
    "fmt"
    "os"

)

var (
    server = "localhost"
    port = 6667
    channel = "#bot"
    nickname = "bot"

    gabens = []string{"vypr"}
    dunmas = ""

    charmap = make(map[string]string)
    monsmap = make(map[string]string)
)

func stringInSlice(a string, list []string) bool {
    // thanks stackoverflow
    for _, b := range list {
        if b == a { return true }
    }
    return false
}

func fillCharmap(nick string, sect string, val string) {
    // TODO: Command that creates a map of maps for a nick.
    // vypr["stat"]["armor"] = 9001
    // vypr["stat"]["health"] = 100
    // vypr["info"]["name"] = "Peapod"
    // make(map[string]map[string]string)
    var isDungeonMaster = false
    if nick == dunmas { isDungeonMaster = true }
}

func processMsg(nick string, msg string, conn *irc.Connection) {
    if strings.HasPrefix(msg, ".set") && len(strings.Split(msg, " ")) == 4 {
        var arg1 = strings.Split(msg, " ")[1]
        var arg2 = strings.Split(msg, " ")[2]
        var arg3 = strings.Split(msg, " ")[3]
        fillCharmap(arg1, arg2, arg3)
    } else if strings.HasPrefix(msg, ".dm") && len(strings.Split(msg, " ")) == 2 {
        if len(dunmas) == 0 && len(strings.Split(msg, " ")[1]) > 0 {
            dunmas = strings.Split(msg, " ")[1]
            fmt.Println("[cmd] dm - " + dunmas)
            conn.Privmsg(channel, "dm is now set to " + dunmas)
        } else {
            conn.Privmsg(channel, "dm has already been set, the current DM is " + dunmas)
        }
    } else if msg == ".resetdm" && (nick == dunmas || stringInSlice(nick, gabens)) {
        dunmas = ""
        fmt.Println("[cmd] resetdm")
        conn.Privmsg(channel, "dm has been reset")
    } else if msg == ".quit" && stringInSlice(nick, gabens) {
        fmt.Println("[cmd] lord " + nick + " has requested a shutdown")
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

    conn.AddCallback("001", func(e *irc.Event) { conn.Join("#bot") })

    conn.AddCallback("PRIVMSG", func(e *irc.Event) {
        processMsg(e.Nick, e.Message(), conn)
    })

    conn.Loop()
}

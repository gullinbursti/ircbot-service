package main

import (
	"fmt"
	"go-ircevent"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
//	"time"
)

var userName = "faroutrob"
var roomName = "#" + userName

func main() {
	fmt.Println("yo")
	con := irc.IRC("mbtester", "mbtester") //nick, user
	con.Password = "oauth:4xin8bvtejcs15hyddk279dduicc68"

	var result []string

	response, err := http.Get("http://beta.modd.live/api/stream_bot2.php")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		result = strings.Split(string(contents), "\n")
		for i := range result {
			fmt.Println(result[i])
		}
	}
	///////////////////////////////////////
	fmt.Println("hi")

	con.AddCallback("001", func(e *irc.Event) {
		fmt.Println("001")
		for i := range result {
			con.Join("#" + result[i])
		}
	})
	con.AddCallback("JOIN", func(e *irc.Event) {
		room := e.Arguments[0]
		nick := room[1:len(room)]
//		fmt.Println("Join")
		con.Privmsg(room, "/w " + nick + " This is the new ModdBot bitches!!!")
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		room := e.Arguments[0]
		fmt.Println("roomname = " + room)
		fmt.Println("privmsg")
		subscribe := regexp.MustCompile(`(?i)!subscribe`)
		gamecard := regexp.MustCompile(`(?i)!gamecard`)
                support := regexp.MustCompile(`(?i)!support`)
		help := regexp.MustCompile(`(?i)!help`)
                highlight := regexp.MustCompile(`(?i)!highlight`)
                notify := regexp.MustCompile(`(?i)!notify`)
		var msg = e.Message()

		  switch {
                  case subscribe.MatchString(msg):
		          con.Privmsg(room, "Subscribe to my Stream Moments and get updates throughout my broadcast.")
		  case support.MatchString(msg):
		        response, err := http.Get("http://beta.modd.live/api/bot_url.php?type=game&name=" + userName)
                        if err != nil {
                                fmt.Printf("%s", err)
                                os.Exit(1)
                        } else {

                                defer response.Body.Close()
                                contents2, err := ioutil.ReadAll(response.Body)
                                if err != nil {
                                        fmt.Printf("%s", err)
                                        os.Exit(1)
                                } else {
                                        con.Privmsg(room, string(contents2))
                                        fmt.Println("yo" + "http://beta.modd.live/api/bot_url.php?type=game&name=" + userName)
                                }
                        }

                  case gamecard.MatchString(msg):
                        response, err := http.Get("http://beta.modd.live/api/bot_url.php?type=promo&name=" + userName)
                        if err != nil {
                                fmt.Printf("%s", err)
                                os.Exit(1)
                        } else {

                                defer response.Body.Close()
                                contents2, err := ioutil.ReadAll(response.Body)
                                if err != nil {
                                        fmt.Printf("%s", err)
                                        os.Exit(1)
                                } else {
                                        con.Privmsg(room, string(contents2))
                                }
                        }

		   case help.MatchString(msg):
			fmt.Println("matched help")
			con.Privmsg(room, "Visit 104.131.141.147/help.html for a list of commands")
                   case highlight.MatchString(msg):
                        con.Privmsg(room, "highlight")
                   case notify.MatchString(msg):
                        con.Privmsg(room, "To be notified each time this streamer goes live click here s.00m.co/"  + userName)
		   default:
		  fmt.Println("is not recognized")
		  }

	})

	ircerr := con.Connect("irc.chat.twitch.tv:6667")
	if ircerr != nil {
		fmt.Println("Failed connecting")
		return
	}


 /*messageTicker := time.NewTicker(time.Millisecond * 3600000)
    go func() {
        for t := range messageTicker.C {
            //loop through message timeTable for users and if no  messages within the hour send out subscribe, gamecard and/or cross promote round robin  
        }
    }()


 modTicker := time.NewTicker(time.Millisecond * 1800000)
    go func() {
        for t := range modTicker.C {
		//loop through modTable 
		//if  modTable[user] == 0 and moddbot shows as moderator then set modTable[user] to 1 and message "moddbot has been activate"
		//if modTble[user == 1 and  moddbot does NOT show up as moderator then exit the channel and  
        }
    }()

 milestoneTicker := time.NewTicker(time.Millisecond * 1800000)
    go func() {
        for t := range milestoneTicker.C {
                //loop through users and check if user has reached milestones 
        }
    }()

*/
	con.Loop()
}

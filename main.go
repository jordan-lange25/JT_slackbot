package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net/http"
	"github.com/jordan-lange25/JT_slackbot/pkg/chunky"
	"github.com/gorilla/mux"
	"github.com/slack-go/slack/slackevents"
)

func main() {

	token := os.Getenv("CHUNKY_TOKEN")
	if token == "" {
		fmt.Println("CHUNKY_TOKEN not set")
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/slack/post", eventsHandler)

	chunky := chunky.Chunky{}
	chunky.InitalizeClient(token)

	/*err := chunky.PostMessage("#botspam", "That's a chunky")
	if err != nil{
		fmt.Println(err)
	}*/

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func eventsHandler(w http.ResponseWriter, req *http.Request) {

	token := os.Getenv("CHUNKY_TOKEN")
	if token == "" {
		fmt.Println("CHUNKY_TOKEN not set")
		os.Exit(1)
	}

	chunky := chunky.Chunky{}

	chunky.InitalizeClient(token)


	// Reads the requetst body and returns it in a []byte
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// In the case of some error, return a bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse any events in the body into an EventsAPIEvent struct
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		// In the case of some error, return a 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If the event is a URL verification, return the expected challenge 
	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		// Get the challenge field out of the slack event
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Return the challenge field back to slack to verify our app
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			fmt.Println(ev)
			chunky.PostMessage("#botspam", ev.User + " " + ev.Text)

		}
	}

}

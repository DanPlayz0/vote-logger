package botlists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/config"
	"github.com/Would-You-Bot/vote-logger/helpers"
)

type TopWebhookData struct {
	Bot  string `json:"bot"`
	User string `json:"user"`
}

func HandleTopgg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received vote from top.gg")

	if !helpers.Validate(r, w, config.Conf.BotList.Topgg.Auth) {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)

	var v TopWebhookData

	err := dec.Decode(&v)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := helpers.GetUserData(v.User)

	message := fmt.Sprintf("https://top.gg/bot/%s/vote", v.Bot)

	helpers.SendVoteWebhook(response, message)

	w.WriteHeader(http.StatusOK)

	fmt.Println("Vote received from " + v.User + " for " + v.Bot)
}

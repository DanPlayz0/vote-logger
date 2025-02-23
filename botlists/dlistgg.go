package botlists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/helpers"
	"github.com/Would-You-Bot/vote-logger/types"
	"github.com/Would-You-Bot/vote-logger/config"
)

func HandleDlistgg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received vote from dlist.gg")

	if !helpers.Validate(r, w, config.Conf.BotList.Dlist.Auth) {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)

	var v types.Vote

	err := dec.Decode(&v)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := helpers.GetUserData(v)

	message := fmt.Sprintf("https://discordlist.gg/bot/%s/vote", v.Bot)

	helpers.SendVoteWebhook(response, message)

	w.WriteHeader(http.StatusOK)

	fmt.Println("Vote received from " + v.User + " for " + v.Bot)
}

package botlists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/config"
	"github.com/Would-You-Bot/vote-logger/helpers"
)

type DscWebhookData struct {
	ListingId string `json:"listing_id"`
	BotId     string `json:"bot_id"`
	UserId    string `json:"user_id"`
}

func HandleDscbot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received vote from dsc.bot")

	if !helpers.Validate(r, w, config.Conf.BotList.Dscbot.Auth) {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)

	var v DscWebhookData

	err := dec.Decode(&v)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := helpers.GetUserData(v.UserId)

	message := fmt.Sprintf("https://dsc.bot/%s/vote", v.ListingId)

	helpers.SendVoteWebhook(response, message)

	w.WriteHeader(http.StatusOK)

	fmt.Println("Vote received from " + v.UserId + " for " + v.BotId)
}

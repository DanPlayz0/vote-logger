package botlists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/helpers"
	"github.com/Would-You-Bot/vote-logger/types"
)

func HandleTopgg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received vote from top.gg")

	if !helpers.Validate(r, w) {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	fmt.Println("line 24")

	var v types.TopggVote

	err := dec.Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("line 34")

	response := helpers.GetUserData(v)

	message := fmt.Sprintf("https://top.gg/bot/%s/vote", v.Bot)
	fmt.Println("line 39")
	helpers.SendVoteWebhook(response, message)
	fmt.Println("line 41")

	w.WriteHeader(http.StatusOK)

	fmt.Println("Vote received from " + v.User + " for " + v.Bot + " with type " + v.Type)
}

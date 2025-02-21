package botlists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/helpers"
	"github.com/Would-You-Bot/vote-logger/types"
)

func HandleTopgg(w http.ResponseWriter, r *http.Request) {

	if !helpers.Validate(r, w) {
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var v types.TopggVote

	err := dec.Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := helpers.GetUserData(v)

	helpers.SendVoteWebhook(response, v)

	w.WriteHeader(http.StatusOK)

	fmt.Println("Vote received from " + v.User + " for " + v.Bot + " with type " + v.Type)
}

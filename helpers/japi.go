package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/types"
)

func GetUserData(v types.Vote) types.Response {
	res, err := http.Get("https://japi.rest/discord/v1/user/" + v.User)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var response types.Response

	resp := json.NewDecoder(res.Body)
	resp.DisallowUnknownFields()

	resp.Decode(&response)

	return response
}

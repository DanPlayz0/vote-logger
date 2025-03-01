package botlists

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/config"
	"github.com/Would-You-Bot/vote-logger/helpers"
	"github.com/golang-jwt/jwt/v5"
)

type DlistWebhookData struct {
	BotId  string `json:"bot_id"`
	UserId string `json:"user_id"`
	IsTest bool   `json:"is_test"`
	jwt.RegisteredClaims
}

func HandleDlist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received vote from dlist.gg")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jwtString := string(body)
	data := &DlistWebhookData{}

	token, err := jwt.ParseWithClaims(jwtString, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.BotList.Dlist.Auth), nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Invalid authorization")
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	fmt.Printf("Bot ID: %s, User ID: %s\n", data.BotId, data.UserId)

	response := helpers.GetUserData(data.UserId)

	message := fmt.Sprintf("https://dlist.gg/bot/%s/vote", data.BotId)

	helpers.SendVoteWebhook(response, message)

	w.WriteHeader(http.StatusOK)

	fmt.Println("Vote received from " + data.UserId + " for " + data.BotId)
}

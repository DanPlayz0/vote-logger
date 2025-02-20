package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Vote struct {
	Bot       string `json:"bot"`
	User      string `json:"user"`
	Type      string `json:"type"`
	IsWeekend bool   `json:"isWeekend"`
	Query     string `json:"query,omitempty"`
}

type Response struct {
	Data User `json:"data"`
}

type User struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatarURL"`
}

func cleanUsername(username string) string {
	for _, word := range []string{"Discord", "discord", "Everyone", "everyone"} {
		username = strings.ReplaceAll(username, word, "")
	}
	return strings.TrimSpace(username)
}

func handleTopgg(w http.ResponseWriter, r *http.Request) {

	auth := ""

	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	if r.Header.Get("Authorization") != auth {
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var v Vote
	fmt.Println(r.Header.Get("Authorization"))

	err := dec.Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	resp, err := http.Get("https://japi.rest/discord/v1/user/" + v.User)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var resonse Response

	res := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()

	res.Decode(&resonse)

	webhookURL := "https://discord.com/api/webhooks/yourwebhookurl"

	emojis := [7]string{"<a:jammiesyou:1009965703484424282>",
		"<a:nyancatyou:1009965705808056350>",
		"<a:partyparrotyou:1009965704621080678>",
		"<a:shootyou:1009965706978267136>",
		"<a:catjamyou:1009965950101110806>",
		"<a:patyou:1009964589678612581>",
		"<a:patyoufast:1009964759216574586>"}

	i := rand.Intn(len(emojis))

	payload := map[string]interface{}{
		"content":     fmt.Sprintf("%s Voted for me on https://top.gg/bot/%s/vote", emojis[i], v.Bot),
		"embeds":      nil,
		"attachments": []interface{}{},
		"username":    cleanUsername(resonse.Data.Username),
		"avatar_url":  resonse.Data.AvatarURL,
		"tts":         false,
		"components": []map[string]interface{}{
			{
				"type": 1,
				"components": []map[string]interface{}{
					{
						"type":  2,
						"label": "Vote",
						"emoji": map[string]interface{}{
							"id":       nil,
							"name":     "💻",
							"animated": false,
						},
						"style": 5,
						"url":   fmt.Sprintf("https://top.gg/bot/%s/vote", v.Bot),
					},
				},
			},
		},
	}

	// Convert the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp2.Body.Close()

	// Check response
	if resp2.StatusCode == http.StatusOK || resp2.StatusCode == http.StatusNoContent {
		fmt.Println("Message sent successfully!")
	} else {
		fmt.Printf("Failed to send message. Status: %d\n", resp2.StatusCode)
	}

	fmt.Println("Vote received from " + v.User + " for " + v.Bot + " with type " + v.Type)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/topgg", handleTopgg).Methods("POST")

	err := http.ListenAndServe(":8000", router)
	log.Fatal(err)
}

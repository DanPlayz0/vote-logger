package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Would-You-Bot/vote-logger/config"
	"github.com/Would-You-Bot/vote-logger/emojis"
	"github.com/Would-You-Bot/vote-logger/types"
)

func SendVoteWebhook(response types.Response, message string) {
	payload := map[string]interface{}{
		"content":     fmt.Sprintf("%s Voted for me on %s", emojis.GetRandomEmoji(), message),
		"embeds":      nil,
		"attachments": []interface{}{},
		"username":    CleanUsername(response.Data.Username),
		"avatar_url":  response.Data.AvatarURL,
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
						"url":   message,
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
	req, err := http.NewRequest("POST", config.Conf.WebhookURL, bytes.NewBuffer(jsonData))
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
}

package helpers

import "strings"

func CleanUsername(username string) string {
	for _, word := range []string{"Discord", "discord", "Everyone", "everyone"} {
		username = strings.ReplaceAll(username, word, "")
	}
	return strings.TrimSpace(username)
}
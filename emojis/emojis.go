package emojis

import "math/rand"

var Emojis = [7]string{"<a:jammiesyou:1009965703484424282>",
"<a:nyancatyou:1009965705808056350>",
"<a:partyparrotyou:1009965704621080678>",
"<a:shootyou:1009965706978267136>",
"<a:catjamyou:1009965950101110806>",
"<a:patyou:1009964589678612581>",
"<a:patyoufast:1009964759216574586>"}

func GetRandomEmoji() string {
	return Emojis[rand.Intn(len(Emojis))]
}
package types

type TopggVote struct {
	Bot       string `json:"bot"`
	Entity    string `json:"entity"`
	User      string `json:"user"`
	Type      string `json:"type"`
	IsWeekend bool   `json:"isWeekend"`
	Query     string `json:"query,omitempty"`
}

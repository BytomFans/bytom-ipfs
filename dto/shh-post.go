
package dto

type SHHPostParameters struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Topics   []string `json:"topics"`
	Payload  string   `json:"payload"`
	Priority string   `json:"priority"`
	TTL      string   `json:"ttl"`
}

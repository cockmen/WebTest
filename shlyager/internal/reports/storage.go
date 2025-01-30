package reports

import "time"

type Report struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"descriptions"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}

package service

import "time"

type Word struct {
	Title       string `json:"title"`
	Translation string `jsong:"translation"`
}

type Report struct {
	Title       string    `json:"title"`
	Description string    `json:"descriptions"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}

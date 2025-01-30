package service

type Word struct {
	Title       string `json:"title"`
	Translation string `jsong:"translation"`
}

type Report struct {
	Title       string `json:"title"`
	Description string `json:"descriptions"`
}

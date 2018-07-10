package server

type Code struct {
	HTML string `json:"html,omitempty"`
	CSS  string `json:"css,omitempty"`
	JS   string `json:"js,omitempty"`
}

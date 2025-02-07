package model

// Note struct
type Note struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

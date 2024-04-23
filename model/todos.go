package model

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"boolean"`
}

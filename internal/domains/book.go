package domains

type Books struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Status      *bool  `json:"status"`
	Pages       int    `json:"pages"`
	Description string `json:"description"`
	Published   string `json:"published"`
	Image       string `json:"image"`
}

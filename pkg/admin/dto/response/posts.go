package response

type PostResponse struct {
	Title     string `json:"title"`
	Category  string `json:"category"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
}

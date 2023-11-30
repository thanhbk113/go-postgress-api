package response

type PostResponse struct {
	Title     string `json:"title"`
	Category  string `json:"category"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
}

type LikeResponse struct {
	Like int `json:"like"`
}

type PostResposeAll struct {
	PostResponse []PostResponse `json:"post_response"`
	Total        int            `json:"total"`
	Limit        int32          `json:"limit"`
}

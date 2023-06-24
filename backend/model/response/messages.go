package response

type GetMessages struct {
	Messages []*Message `json:"messages"`
}

type Message struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Type     string `jsn:"type"`
	Content  string `json:"content"`
	PostedAt string `json:"timestamp"`
}

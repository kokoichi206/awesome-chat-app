package request

type PostMessage struct {
	UserID   string `json:"user_id"`
	Type     string `jsn:"type"`
	Content  string `json:"content"`
	PostedAt string `json:"posted_at"`
}

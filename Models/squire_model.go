package Models

type Squire struct {
	ID       string `json:"id"`
	UserId   string `json:"user_id"`
	Url      string `json:"url"`
	Password string `json:"password"`
	Username string `json:"username"`
}

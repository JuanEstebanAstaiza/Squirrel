package Models

type profile struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type credentials struct {
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

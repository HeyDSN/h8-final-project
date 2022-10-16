package models

type Response struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Error  *Error      `json:"error,omitempty"`
}

type Error struct {
	Fields  []string    `json:"fields"`
	Message string      `json:"message"`
	Extends interface{} `json:"extends,omitempty"`
}

type PhotoResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User      struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
}

type CommentResponse struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	PhotoID   uint   `json:"photo_id"`
	UserID    uint   `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	} `json:"photo"`
}

type SocialMediaResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	User           struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

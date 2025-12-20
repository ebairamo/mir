package domain

import "time"

// Post представляет пост/тред на доске
type Post struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Text            string     `json:"text"`
	UserID          int        `json:"user_id"`
	UserName        string     `json:"user_name"`
	UserAvatarURL   string     `json:"user_avatar_url"`
	ImageURL        string     `json:"image_url"`
	CreatedAt       time.Time  `json:"created_at"`
	LastCommentTime time.Time  `json:"last_comment_time"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

// Comment представляет комментарий к посту или другому комментарию
type Comment struct {
	ID            int       `json:"id"`
	Text          string    `json:"text"`
	PostID        int       `json:"post_id"`
	ReplyToID     *int      `json:"reply_to_id"` // nil если ответ на Post
	UserID        int       `json:"user_id"`
	UserName      string    `json:"user_name"	`
	UserAvatarURL string    `json:"user_avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
}

// Session представляет сессию пользователя
type Session struct {
	SessionID      string    `json:"session_id"`
	UserID         int       `json:"user_id"`
	AvatarID       int       `json:"avatar_id"`
	CustomUserName string    `json:"custom_user_name"`
	ExpiresAt      time.Time `json:"expires_at"`
}

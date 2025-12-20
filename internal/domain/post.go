package domain

import "context"

type PostRepository interface {
	CreatePost(ctx context.Context, post *Post) error
	GetPost(ctx context.Context, postID int) (*Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	GetArchive(ctx context.Context) ([]Post, error)
	UpdatePost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, postID int) error
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *Comment) error
	GetComment(ctx context.Context, commentID int) (*Comment, error)
	GetCommentByPostID(ctx context.Context, postID int) ([]Comment, error)
	UpdateComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, commentID int) error
}

type SessionRepository interface {
	CreateSession(ctx context.Context, session *Session) error
	GetSessionByID(ctx context.Context, sessionID string) (*Session, error)
	UpdateSessionByID(ctx context.Context, session *Session) error
	DeleteSession(ctx context.Context, sessionID string) error
}

type FileStorage interface {
	UploadFile(ctx context.Context, bucket string, key string, data []byte) (string, error)
	GetFile(ctx context.Context, bucket string, key string) ([]byte, error)
	DeleteFile(ctx context.Context, bucket string, key string) error
}

type AvatarService interface {
	GetAvatarURL(ctx context.Context, characterID int) (string, error)
	GetRandomCharacter(ctx context.Context) (int, string, error)
	GetTotalCharacters(ctx context.Context) (int, error)
}

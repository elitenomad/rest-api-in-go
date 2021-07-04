package comment

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}
type CommentService interface {
	// CRUD
	GetComment(ID uint) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

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
	GetCommentsBySlug(slug string) ([]Comment, error)
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

func (service *Service) GetComment(ID uint64) (Comment, error) {
	var comment Comment
	if result := service.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (service *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := service.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) UpdateComment(ID uint64, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *Service) GetAllComments() []Comment {
	var comments []Comment
	_ = service.DB.Find(&comments)
	return comments
}

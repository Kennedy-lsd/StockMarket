package services

import (
	"github.com/Kennedy-lsd/StockMarket/data"
	"github.com/Kennedy-lsd/StockMarket/internal/repos"
)

type CommentService struct {
	CommentRepository *repos.CommentRepository
}

func NewCommentService(r *repos.CommentRepository) *CommentService {
	return &CommentService{
		CommentRepository: r,
	}
}

func (s *CommentService) GetAll() ([]data.Comment, error) {
	return s.CommentRepository.FindAll()
}

func (s *CommentService) Create(comment *data.CommentCreate) error {
	return s.CommentRepository.Post(comment)
}

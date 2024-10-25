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

func (s *CommentService) GetById(id int64) (*data.Comment, error) {
	return s.CommentRepository.FindById(id)
}

func (s *CommentService) DeleteById(id int64) error {
	return s.CommentRepository.DeleteById(id)
}

func (s *CommentService) UpdateById(id int64, comment *data.CommentUpdate) error {
	return s.CommentRepository.UpdateById(id, comment)
}

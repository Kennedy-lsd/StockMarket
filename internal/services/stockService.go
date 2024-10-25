package services

import (
	"github.com/Kennedy-lsd/StockMarket/data"
	"github.com/Kennedy-lsd/StockMarket/internal/repos"
)

type StockService struct {
	StockRepo *repos.StockRepository
}

func NewStockService(r *repos.StockRepository) *StockService {
	return &StockService{
		StockRepo: r,
	}
}

func (s *StockService) GetAll() ([]data.Stock, error) {
	return s.StockRepo.FindAll()
}

func (s *StockService) Create(stock *data.CreatedStock) error {
	return s.StockRepo.Post(stock)
}

func (s *StockService) GetById(id int64) (*data.Stock, error) {
	return s.StockRepo.FindById(id)
}

func (s *StockService) DeleteById(id int64) error {
	return s.StockRepo.DeleteById(id)
}

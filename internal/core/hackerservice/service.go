package hackerservice

import (
	"test_hackers/internal/core/domain"
	"test_hackers/internal/core/ports"
)

type Service struct {
	storage ports.HackerRepository
}

func NewService(storage ports.HackerRepository) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetAll() ([]domain.Hacker, error) {
	return s.storage.GetAll()
}

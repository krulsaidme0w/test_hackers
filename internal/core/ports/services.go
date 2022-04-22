package ports

import (
	"test_hackers/internal/core/domain"
)

type HackerService interface {
	GetAll() ([]domain.Hacker, error)
}

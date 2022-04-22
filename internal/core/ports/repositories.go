package ports

import "test_hackers/internal/core/domain"

type HackerRepository interface {
	GetAll() ([]domain.Hacker, error)
}

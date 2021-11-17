package repository

import (
	"backend-task/pgk/logger"
)

type Repository struct {
	ttl int64
}

func New(ttl int64, logger *logger.Logger) *Repository {
	return &Repository{ttl: ttl}
}

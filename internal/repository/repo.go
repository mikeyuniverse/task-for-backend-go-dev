package repository

import (
	"backend-task/pgk/config"
	"backend-task/pgk/logger"
)

type Repository struct{}

func New(config *config.Config, logger *logger.Logger) *Repository {
	return &Repository{}
}

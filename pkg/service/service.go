package service

import "github.com/bitwyre/template-go/pkg/repository"

type Service struct {
	HealthCheckService IHealthCheckService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		HealthCheckService: NewHealthCheckService(repository),
	}
}

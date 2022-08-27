package service

import "github.com/mohammedaugi13/simple-rest-go/pkg/repository"

type Service struct {
	HealthCheckService IHealthCheckService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		HealthCheckService: NewHealthCheckService(repository),
	}
}

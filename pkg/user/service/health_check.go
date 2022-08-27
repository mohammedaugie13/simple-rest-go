package service

import "github.com/mohammedaugi13/simple-rest-go/pkg/repository"

type IHealthCheckService interface {
	HealthCheck() *responseMsg
}

type healthCheckService struct {
	repository *repository.Repository
}

func NewHealthCheckService(repository *repository.Repository) *healthCheckService {
	return &healthCheckService{repository: repository}
}

type responseMsg struct {
	Message string `json:"message"`
}

func (s healthCheckService) HealthCheck() *responseMsg {
	resp := responseMsg{
		Message: "Everything Fine Bro",
	}

	return &resp
}

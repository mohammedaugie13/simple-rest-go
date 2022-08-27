package service

import "github.com/bitwyre/template-go/pkg/repository"

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
		Message: "User Module OK",
	}

	return &resp
}

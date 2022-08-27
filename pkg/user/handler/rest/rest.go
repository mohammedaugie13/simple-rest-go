package rest

import "github.com/mohammedaugi13/simple-rest-go/pkg/user/service"

type Rest struct {
	service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		Service: *service,
	}
}

package service

import (
	"github.com/mohammedaugi13/simple-rest-go/pkg/handler/graph/model"
	"github.com/mohammedaugi13/simple-rest-go/pkg/helper"
	"github.com/mohammedaugi13/simple-rest-go/pkg/repository"
)

type IUserAuthService interface {
	CheckUserStatus(input *model.InfoQueryInput) (*model.InfoQuery, error)
}

type userAuthService struct {
	repo *repository.Repository
}

func NewUserAuthService(repo *repository.Repository) *userAuthService {
	return &userAuthService{
		repo: repo,
	}
}

func (s *userAuthService) CheckUserStatus(input *model.InfoQueryInput) (*model.InfoQuery, error) {
	var isRegistered = true
	var isEmailVerified = true

	_, err := s.repo.UserAuthRepository.FindByEMail(input.Email)
	token := helper.GenSHA256(input.Email, 10)

	if err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}

		isRegistered = false
		isEmailVerified = false
	}

	return &model.InfoQuery{
		Email:           input.Email,
		IsRegistered:    isRegistered,
		IsEmailVerified: isEmailVerified,
		ActionToken:     token,
	}, nil
}

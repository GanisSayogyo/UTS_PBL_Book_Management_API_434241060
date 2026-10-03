package service

import (
	"context"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/model"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/repository"
)

type UserService interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}
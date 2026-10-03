package service

import (
	"context"
	"errors"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/model"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/repository"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/helper"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.User, string, error)
}

type authService struct {
	userRepo repository.UserRepository
	cfg      config.Config
}

func NewAuthService(userRepo repository.UserRepository, cfg config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *authService) Register(ctx context.Context, req model.RegisterRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (s *authService) Login(ctx context.Context, req model.LoginRequest) (*model.User, string, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, "", ErrInvalidCredentials
		}

		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, "", ErrInvalidCredentials
	}

	token, err := helper.GenerateToken(user.ID, user.Role, s.cfg)
	if err != nil {
		return nil, "", err
	}

	user.Password = ""

	return user, token, nil
}
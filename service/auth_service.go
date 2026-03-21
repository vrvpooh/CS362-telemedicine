package service

import "cs362-telemedicine/model"

type AuthService interface {
	Register(req model.RegisterRequest) error
	Login(req model.LoginRequest) (string, error)
	GetProfile(userID string) (*model.User, error)
}
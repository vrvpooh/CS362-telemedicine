package repository

import "cs362-telemedicine/model"

type UserRepository interface {
	SaveUser(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(id string) (*model.User, error)
}
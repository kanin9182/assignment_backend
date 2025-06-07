package services

import "assignment/internals/core/domain"

type UserServices interface {
	GetUser() (*domain.UserResponse, error)
	GetUserById(userId string) (*domain.GetUserByIdRequest, error)
	GetUserProfile(userId string) (*domain.GetUserMain, error)
}

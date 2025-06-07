package repositories

import "assignment/internals/repositories/models"

type UserRepositories interface {
	GetUser() (*models.User, error)
	GetUserById(userId string) (*models.User, error)
	GetGreetingAndBanner(userId string) (*models.GreetingAndBanner, error)
	GetAccountInfo(userId string) (*[]models.AccountInfo, error)
	GetDebitCardInfo(userId string) (*[]models.DebitCardInfo, error)
}

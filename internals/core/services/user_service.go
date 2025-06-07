package services

import (
	"assignment/internals/core/domain"
	"assignment/internals/core/ports/repositories"
	"assignment/internals/repositories/models"
)

type UserServices struct {
	userRepo repositories.UserRepositories
}

func NewUserService(userRepo repositories.UserRepositories) *UserServices {
	return &UserServices{
		userRepo: userRepo,
	}
}

func (s *UserServices) GetUser() (*domain.UserResponse, error) {
	res, err := s.userRepo.GetUser()
	if err != nil {
		return nil, err
	}

	return &domain.UserResponse{
		UserID: res.UserID,
		Name:   res.Name,
	}, nil
}

func (s *UserServices) GetUserById(userId string) (*domain.GetUserByIdRequest, error) {
	res, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return &domain.GetUserByIdRequest{
		UserId:  res.UserID,
		Name:    res.Name,
		PinHash: res.PinHash,
	}, nil
}

func (s *UserServices) GetUserProfile(userId string) (*domain.GetUserProfileResponse, error) {
	resGreetingAndBanner, err := s.userRepo.GetGreetingAndBanner(userId)
	if err != nil {
		return nil, err
	}

	resGetAccountInfo, err := s.userRepo.GetAccountInfo(userId)
	if err != nil {
		return nil, err
	}

	resGetDebitCardInfo, err := s.userRepo.GetDebitCardInfo(userId)
	if err != nil {
		return nil, err
	}

	return &domain.GetUserProfileResponse{
		GreetingAndBanner: domain.GreetingAndBanner{
			BannerId:    resGreetingAndBanner.BannerId,
			Name:        resGreetingAndBanner.Name,
			Greeting:    resGreetingAndBanner.Greeting,
			Title:       resGreetingAndBanner.Title,
			Description: resGreetingAndBanner.Description,
			ImageUrl:    resGreetingAndBanner.Image,
		},
		AccountInfos:   mapAccountInfoModelsToDomain(resGetAccountInfo),
		DebitCardInfos: mapDebitCardInfoModelsToDomain(resGetDebitCardInfo),
	}, nil
}

func mapAccountInfoModelsToDomain(models *[]models.AccountInfo) *[]domain.AccountInfo {
	if models == nil {
		return nil
	}

	accountMap := make(map[string]*domain.AccountInfo)
	for _, m := range *models {
		acc := accountMap[m.AccountID]

		if acc == nil {
			accountMap[m.AccountID] = &domain.AccountInfo{
				AccountID:     m.AccountID,
				Type:          m.Type,
				Currency:      m.Currency,
				AccountNumber: m.AccountNumber,
				Color:         m.Color,
				IsMainAccount: m.IsMainAccount,
				FlagValue:     append([]string{}, m.FlagValue),
				Amount:        m.Amount,
			}
			continue
		}
		acc.Amount += m.Amount
		acc.FlagValue = append(acc.FlagValue, m.FlagValue)
	}

	result := make([]domain.AccountInfo, 0, len(accountMap))
	for _, acc := range accountMap {
		result = append(result, *acc)
	}

	return &result
}

func mapDebitCardInfoModelsToDomain(models *[]models.DebitCardInfo) *[]domain.DebitCardIno {
	if models == nil {
		return nil
	}
	var result []domain.DebitCardIno
	for _, m := range *models {
		result = append(result, domain.DebitCardIno{
			CardID:      m.CardID,
			Name:        m.Name,
			Issuer:      m.Issuer,
			Number:      m.Number,
			Status:      m.Status,
			Color:       m.Color,
			BorderColor: m.BorderColor,
		})
	}
	return &result
}

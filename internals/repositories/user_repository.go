package repositories

import (
	"assignment/internals/core/ports/repositories"
	"assignment/internals/repositories/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepositories {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUser() (*models.User, error) {
	var user models.User
	if err := r.db.Order("RAND()").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserById(userId string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetGreetingAndBanner(userId string) (*models.GreetingAndBanner, error) {
	var result models.GreetingAndBanner
	query := `
        select u.name, ug.greeting, b.banner_id, b.title, b.description, b.image from user_greetings ug 
		inner join banners b 
		on ug.user_id = b.user_id
		inner join users u 
		on u.user_id = b.user_id
		where ug.user_id = ?
    `
	if err := r.db.Raw(query, userId).Scan(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *userRepository) GetAccountInfo(userId string) (*[]models.AccountInfo, error) {
	var accounts []models.AccountInfo
	query := `
		select a.account_id, a.type, a.currency, a.account_number, 
		ad.color, ad.is_main_account,
		af.flag_value,
		ab.amount
		from accounts a
		inner join account_details ad
		on a.account_id  = ad.account_id 
		inner join account_flags af 
		on a.account_id = af.account_id
		inner join account_balances ab 
		on a.account_id = ab.account_id
		where a.user_id = ?
	`
	if err := r.db.Raw(query, userId).Scan(&accounts).Error; err != nil {
		return nil, err
	}
	return &accounts, nil
}

func (r *userRepository) GetDebitCardInfo(userId string) (*[]models.DebitCardInfo, error) {
	var debitCards []models.DebitCardInfo
	query := `
		select dc.card_id, dc.name,
		dcd.issuer, dcd.number,
		dcs.status,
		dcd2.color, dcd2.border_color
		from debit_cards dc 
		inner join debit_card_details dcd 
		on dc.card_id = dcd.card_id
		inner join debit_card_status dcs 
		on dc.card_id = dcs.card_id
		inner join debit_card_design dcd2 
		on dc.card_id = dcd2.card_id
		where dc.user_id = ?
	`
	if err := r.db.Raw(query, userId).Scan(&debitCards).Error; err != nil {
		return nil, err
	}
	return &debitCards, nil
}

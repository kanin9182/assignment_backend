package domain

type UserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type GetUserByIdRequest struct {
	UserId  string `json:"user_id"`
	Name    string `json:"name"`
	PinHash string `json:"pin_hash"`
}

type LoginRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	ExpiresIn   int           `json:"expires_in"`
	User        *UserResponse `json:"user"`
}

type GetUserMain struct {
	UserID            string            `json:"user_id"`
	Name              string            `json:"name"`
	GreetingAndBanner GreetingAndBanner `json:"greeting_and_banner"`
	AccountInfos      *[]AccountInfo    `json:"account_info"`
	DebitCardInfos    *[]DebitCardIno   `json:"debit_card_info"`
}

type GreetingAndBanner struct {
	BannerId    string `json:"banner_id"`
	Name        string `json:"name"`
	Greeting    string `json:"greeting"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image"`
}

type AccountInfo struct {
	AccountID     string   `json:"account_id"`
	Type          string   `json:"type"`
	Currency      string   `json:"currency"`
	AccountNumber string   `json:"account_number"`
	Color         string   `json:"color"`
	IsMainAccount bool     `json:"is_main_account"`
	FlagValue     []string `json:"flag_value"`
	Amount        float64  `json:"amount"`
}

type DebitCardIno struct {
	CardID      string `json:"card_id"`
	Name        string `json:"name"`
	Issuer      string `json:"issuer"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	BorderColor string `json:"border_color"`
}

package models

type AccountInfo struct {
	AccountID     string  `json:"account_id"`
	Type          string  `json:"type"`
	Currency      string  `json:"currency"`
	AccountNumber string  `json:"account_number"`
	Color         string  `json:"color"`
	IsMainAccount bool    `json:"is_main_account"`
	FlagValue     string  `json:"flag_value"`
	Amount        float64 `json:"amount"`
}

package models

type DebitCardInfo struct {
	CardID      string `json:"card_id"`
	Name        string `json:"name"`
	Issuer      string `json:"issuer"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	Color       string `json:"color"`
	BorderColor string `json:"border_color"`
}

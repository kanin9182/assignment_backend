package models

type GreetingAndBanner struct {
	BannerId    string `json:"banner_id"`
	Name        string `json:"name"`
	Greeting    string `json:"greeting"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

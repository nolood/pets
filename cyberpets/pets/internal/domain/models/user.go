package models

type User struct {
	ID           uint64 `json:"id"`
	TgID         int64  `json:"tg_id"`
	Username     string `json:"username"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
	PhotoUrl     string `json:"photo_url"`
}

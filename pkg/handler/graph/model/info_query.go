package model

type InfoQuery struct {
	Email           string `json:"email"`
	IsRegistered    bool   `json:"isRegistered"`
	IsEmailVerified bool   `json:"isEmailVerified"`
	ActionToken     string `json:"actionToken"`
}

type InfoQueryInput struct {
	Email string `json:"email" validate:"required, email"`
}

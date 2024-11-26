package models

type SignInResponse struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}

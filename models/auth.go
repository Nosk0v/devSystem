package models

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignInOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenOutput struct {
	AccessToken string `json:"access_token"`
}

type SignUpInput struct {
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=8,max=32"`
	Name           *string `json:"name"`
	Role           int
	OrganizationID int `json:"organization"`
}

type SignUpOutput struct {
	Message string `json:"message"`
}

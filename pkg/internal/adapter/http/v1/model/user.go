package model

type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	PicUrl    string `json:"picUrl"`
	District  string `json:"district"`
	City      string `json:"city"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      string `json:"gender"`
	PicUrl      string `json:"picUrl"`
	District    string `json:"district"`
	City        string `json:"city"`
	AccessToken string `json:"accessToken"`
}

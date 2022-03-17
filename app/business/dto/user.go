package dto

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponseDto struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}

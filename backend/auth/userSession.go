package auth

type UserSession struct {
	Id       int    `json:"Id" validate:"required"`
	UserId   string `json:"UserId" validate:"required"`
	Username string `json:"Username" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

func NewUserSession(Id int, UserId string, Username string, Password string) *UserSession {
	return &UserSession{
		UserId:   UserId,
		Username: Username,
		Password: Password,
	}
}

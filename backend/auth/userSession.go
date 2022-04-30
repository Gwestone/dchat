package auth

type UserSession struct {
	Id       int    `json:"Id"`
	UserId   string `json:"UserId"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func NewUserSession(Id int, UserId string, Username string, Password string) *UserSession {
	return &UserSession{
		UserId:   UserId,
		Username: Username,
		Password: Password,
	}
}

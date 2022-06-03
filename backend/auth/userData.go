package auth

//local variable for JSON
type userDataJSON struct {
	Username string `json:"Username" validate:"required,min=2,max=100"`
	Password string `json:"Password" validate:"required,min=2,max=100"`
}

type PublicUserData struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
}

func NewPublicUserData() *PublicUserData {
	return &PublicUserData{}
}

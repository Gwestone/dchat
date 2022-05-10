package auth

//local variable for JSON
type userDataJSON struct {
	Username string `json:"username" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=2,max=100"`
}

type PublicUserData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func NewPublicUserData() *PublicUserData {
	return &PublicUserData{}
}

package auth

type userDataJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PublicUserData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func NewPublicUserData() *PublicUserData {
	return &PublicUserData{}
}

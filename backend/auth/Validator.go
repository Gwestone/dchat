package auth

func EmptyValidate(session UserSession) bool {
	if (session == UserSession{}) {
		return true
	} else {
		return false
	}
}

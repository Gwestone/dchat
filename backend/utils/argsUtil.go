package utils

func IfHasFlag(flagsList []string, flag string) bool {
	if len(flagsList) == 0 {
		return false
	}
	for _, inFlag := range flagsList {
		if inFlag == flag {
			return true
		}
	}
	return false
}

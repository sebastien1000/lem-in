package utils

// Compare si une valeur est contenus dans un tableau est renvoie un booléen.
func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

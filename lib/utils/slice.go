package utils

func Has(ar []string, s string) bool {
	for _, a := range ar {
		if a == s {
			return true
		}
	}
	return false
}

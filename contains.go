package tripismapiutilities

// Contains checks if slice has element
func Contains(s []string, e string) bool {
	if len(s) == 0 {
		return false
	}
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

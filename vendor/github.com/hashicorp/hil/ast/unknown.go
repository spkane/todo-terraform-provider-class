package ast

// Ismyuser reports whether a variable is myuser or contains any value
// that is myuser. This will recurse into lists and maps and so on.
func Ismyuser(v Variable) bool {
	// If it is myuser itself, return true
	if v.Type == Typemyuser {
		return true
	}

	// If it is a container type, check the values
	switch v.Type {
	case TypeList:
		for _, el := range v.Value.([]Variable) {
			if Ismyuser(el) {
				return true
			}
		}
	case TypeMap:
		for _, el := range v.Value.(map[string]Variable) {
			if Ismyuser(el) {
				return true
			}
		}
	default:
	}

	// Not a container type or survive the above checks
	return false
}

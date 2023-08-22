package util

// IntOrDef returns the value of i if it is not nil, otherwise it returns def.
func IntOrDef(i *int, def int) int {
	if i != nil {
		return *i
	}
	return def
}

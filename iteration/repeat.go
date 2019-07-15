package iteration

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
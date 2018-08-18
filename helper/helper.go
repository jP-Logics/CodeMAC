//Package helper is to define all common helper functions in the application
package helper

// Split just splits
func Split(r rune) bool {
	return r == '\v' || r == '\t' || r == '\n' || r == '\f' || r == '\r' || r == ' '
}

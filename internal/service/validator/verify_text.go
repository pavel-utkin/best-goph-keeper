package validator

/*
 * Text rules:
 * at least 7 letters
 */

func VerifyText(s string) bool {
	var (
		hasMinLen = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	return hasMinLen
}

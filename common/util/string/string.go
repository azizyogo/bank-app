package string

import (
	"regexp"
)

func IsValidString(str string) bool {
	// Regex pattern for emojis and other unwanted characters
	emojiPattern := regexp.MustCompile(`[\p{So}\p{C}\p{Zs}]`)

	// Check if the string contains any emojis
	if emojiPattern.MatchString(str) {
		return false
	}

	// Define allowed characters (alphanumeric and specific symbols)
	var allowedPattern = `^[a-zA-Z0-9@._!#&\s-]+$`
	matched, _ := regexp.MatchString(allowedPattern, str)
	return matched

}

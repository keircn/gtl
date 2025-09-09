package titlecase

import (
	"strings"
	"unicode"
)

var SmallWords = map[string]bool{
	"a": true, "an": true, "and": true, "as": true, "at": true, "but": true,
	"by": true, "for": true, "if": true, "in": true, "nor": true, "of": true,
	"on": true, "or": true, "the": true, "to": true, "up": true, "yet": true,
	"so": true, "with": true,
}

func ToTitleCase(text string) string {
	if text == "" {
		return text
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	result := make([]string, len(words))

	for i, word := range words {
		result[i] = titleWord(word, i == 0 || i == len(words)-1)
	}

	return strings.Join(result, " ")
}

func titleWord(word string, isFirstOrLast bool) string {
	if word == "" {
		return word
	}

	lowerWord := strings.ToLower(word)

	if isFirstOrLast {
		return capitalizeFirst(lowerWord)
	}

	if SmallWords[lowerWord] {
		return lowerWord
	}

	return capitalizeFirst(lowerWord)
}

func capitalizeFirst(word string) string {
	if word == "" {
		return word
	}

	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

package titlecase

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	MaxInputLength = 10000
)

var (
	ErrInputTooLong   = errors.New("input text exceeds maximum length")
	ErrInvalidUnicode = errors.New("input contains invalid unicode")
	ErrEmptyInput     = errors.New("input cannot be empty")
)

var SmallWords = map[string]bool{
	"a": true, "an": true, "and": true, "as": true, "at": true, "but": true,
	"by": true, "for": true, "if": true, "in": true, "nor": true, "of": true,
	"on": true, "or": true, "the": true, "to": true, "up": true, "yet": true,
	"so": true, "with": true,
}

func ToTitleCase(text string) (string, error) {
	if text == "" {
		return "", ErrEmptyInput
	}

	if len(text) > MaxInputLength {
		return "", ErrInputTooLong
	}

	if !utf8.ValidString(text) {
		return "", ErrInvalidUnicode
	}

	tokens := tokenize(text)
	if len(tokens) == 0 {
		return "", ErrEmptyInput
	}

	return processTokens(tokens)
}

func titleWord(word string, isFirstOrLast bool) (string, error) {
	if word == "" {
		return word, nil
	}

	if !utf8.ValidString(word) {
		return "", ErrInvalidUnicode
	}

	if strings.Contains(word, "-") {
		return titleHyphenatedWord(word, isFirstOrLast)
	}

	return titleSingleWord(word, isFirstOrLast)
}

func titleHyphenatedWord(word string, isFirstOrLast bool) (string, error) {
	parts := strings.Split(word, "-")
	titleParts := make([]string, len(parts))

	for i, part := range parts {
		if part == "" {
			titleParts[i] = part
			continue
		}

		isPartFirstOrLast := isFirstOrLast && (i == 0 || i == len(parts)-1)
		titlePart, err := titleSingleWord(part, isPartFirstOrLast)
		if err != nil {
			return "", err
		}
		titleParts[i] = titlePart
	}

	return strings.Join(titleParts, "-"), nil
}

func titleSingleWord(word string, isFirstOrLast bool) (string, error) {
	if word == "" {
		return word, nil
	}

	if !utf8.ValidString(word) {
		return "", ErrInvalidUnicode
	}

	return preserveOrCapitalize(word, isFirstOrLast)
}

func capitalizeFirst(word string) (string, error) {
	if word == "" {
		return word, nil
	}

	if !utf8.ValidString(word) {
		return "", ErrInvalidUnicode
	}

	runes := []rune(word)
	if len(runes) == 0 {
		return word, nil
	}

	runes[0] = unicode.ToUpper(runes[0])
	return string(runes), nil
}

func shouldPreserveOriginalCasing(word string) bool {
	if len(word) < 2 {
		return false
	}

	runes := []rune(word)
	allUpper := true
	hasLetter := false
	letterCount := 0

	for _, r := range runes {
		if unicode.IsLetter(r) {
			hasLetter = true
			letterCount++
			if !unicode.IsUpper(r) {
				allUpper = false
			}
		}
	}

	if !hasLetter {
		return false
	}

	if allUpper && letterCount >= 2 && letterCount <= 6 {
		return true
	}

	return false
}

func preserveOrCapitalize(word string, isFirstOrLast bool) (string, error) {
	if shouldPreserveOriginalCasing(word) {
		return word, nil
	}

	lowerWord := strings.ToLower(word)

	if isFirstOrLast {
		return capitalizeFirst(lowerWord)
	}

	if SmallWords[lowerWord] {
		return lowerWord, nil
	}

	return capitalizeFirst(lowerWord)
}

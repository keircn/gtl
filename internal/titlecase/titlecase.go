package titlecase

import (
	"errors"
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

type Token struct {
	Text          string
	IsWord        bool
	IsPunctuation bool
}

func tokenize(text string) []Token {
	var tokens []Token
	runes := []rune(text)
	var currentToken strings.Builder
	var isInWord bool

	for i, r := range runes {
		if unicode.IsSpace(r) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, Token{
					Text:          currentToken.String(),
					IsWord:        isInWord,
					IsPunctuation: !isInWord,
				})
				currentToken.Reset()
			}
			tokens = append(tokens, Token{
				Text:          string(r),
				IsWord:        false,
				IsPunctuation: false,
			})
			isInWord = false
		} else if unicode.IsLetter(r) || isValidWordCharacter(r, runes, i) {
			if !isInWord && currentToken.Len() > 0 {
				tokens = append(tokens, Token{
					Text:          currentToken.String(),
					IsWord:        false,
					IsPunctuation: true,
				})
				currentToken.Reset()
			}
			currentToken.WriteRune(r)
			isInWord = true
		} else {
			if isInWord && currentToken.Len() > 0 {
				tokens = append(tokens, Token{
					Text:          currentToken.String(),
					IsWord:        true,
					IsPunctuation: false,
				})
				currentToken.Reset()
			}
			currentToken.WriteRune(r)
			isInWord = false
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, Token{
			Text:          currentToken.String(),
			IsWord:        isInWord,
			IsPunctuation: !isInWord,
		})
	}

	return tokens
}

func isValidWordCharacter(r rune, runes []rune, index int) bool {
	if r == '-' {
		return true
	}

	if r == '\'' {
		return index > 0 && index < len(runes)-1 &&
			unicode.IsLetter(runes[index-1]) &&
			unicode.IsLetter(runes[index+1])
	}

	return false
}

func processTokens(tokens []Token) (string, error) {
	var result strings.Builder
	wordCount := 0
	var wordIndices []int

	for i, token := range tokens {
		if token.IsWord {
			wordIndices = append(wordIndices, i)
			wordCount++
		}
	}

	for i, token := range tokens {
		if token.IsWord {
			wordIndex := -1
			for j, idx := range wordIndices {
				if idx == i {
					wordIndex = j
					break
				}
			}

			isFirstOrLast := wordIndex == 0 || wordIndex == wordCount-1
			shouldCapitalizeAfterPunctuation := shouldCapitalizeAfterPunctuation(tokens, i)

			titleWord, err := titleWord(token.Text, isFirstOrLast || shouldCapitalizeAfterPunctuation)
			if err != nil {
				return "", err
			}
			result.WriteString(titleWord)
		} else {
			result.WriteString(token.Text)
		}
	}

	return result.String(), nil
}

func shouldCapitalizeAfterPunctuation(tokens []Token, currentIndex int) bool {
	lastPunctuation := ""
	for i := currentIndex - 1; i >= 0; i-- {
		token := tokens[i]
		if token.IsWord {
			return false
		}
		if token.IsPunctuation {
			lastPunctuation = strings.TrimSpace(token.Text)
		}
	}

	return lastPunctuation == "(" ||
		lastPunctuation == ":" ||
		lastPunctuation == "\"" ||
		lastPunctuation == "'"
}

func isOpeningQuote(tokens []Token, quoteIndex int) bool {
	quoteText := strings.TrimSpace(tokens[quoteIndex].Text)
	if !strings.ContainsAny(quoteText, "\"'") {
		return false
	}

	for i := quoteIndex - 1; i >= 0; i-- {
		if tokens[i].IsWord {
			return true
		}
		if tokens[i].IsPunctuation && strings.TrimSpace(tokens[i].Text) != "" {
			return true
		}
	}

	return true
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

	hasWords := false
	for _, token := range tokens {
		if token.IsWord {
			hasWords = true
			break
		}
	}

	if !hasWords {
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

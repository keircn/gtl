package titlecase

import "strings"

func ToTitleCase(text string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "and": true, "as": true, "at": true, "but": true,
		"by": true, "for": true, "if": true, "in": true, "nor": true, "of": true,
		"on": true, "or": true, "the": true, "to": true, "up": true, "yet": true,
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	result := make([]string, len(words))

	for i, word := range words {
		lowerWord := strings.ToLower(word)

		if i == 0 || i == len(words)-1 {
			result[i] = strings.Title(lowerWord)
		} else if smallWords[lowerWord] {
			result[i] = lowerWord
		} else {
			result[i] = strings.Title(lowerWord)
		}
	}

	return strings.Join(result, " ")
}

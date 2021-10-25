package refactor

import "strings"

func findFirstString(str string) string {
	if len(str) == 0 {
		return ""
	}

	wordsAfterFirstBracket := findWordsAfterBracket(str)
	if wordsAfterFirstBracket == "" {
		return ""
	}

	words := findWordsBeforeCloseBracket(wordsAfterFirstBracket)

	return words
}

func findWordsAfterBracket(str string) string {
	runes := []rune(str)
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	return string(runes[indexFirstBracketFound:len(str)])
}

func findWordsBeforeCloseBracket(str string) string {
	runes := []rune(str)
	indexClosingBracketFound := strings.Index(str, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	return string(runes[1 : indexClosingBracketFound-1])
}

package main

import "strings"

func findFirstStringInBracket(str string) string {
	lenstr := len(str)

	if lenstr > 0 && strindex(str, "(") >= 0 {
		indexFirstBracketFound := strindex(str, "(")
		runes := []rune(str)
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:lenstr])
		indexClosingBracketFound := strindex(wordsAfterFirstBracket, ")")
		if indexClosingBracketFound >= 0 {
			runes := []rune(wordsAfterFirstBracket)
			return string(runes[1 : indexClosingBracketFound-1])
		}

	} else {
		return ""
	}
	return ""
}

func strindex(value string, index string) int {
	return strings.Index(value, index)
}

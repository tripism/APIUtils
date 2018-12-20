package tripismapiutilities

import "strings"

// TitleCase returns title cased string :) all letters that begin words mapped to their title case.
// Certain small words skipped though.
func TitleCase(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to of in and "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = strings.ToLower(word)
		} else {
			words[index] = strings.Title(strings.ToLower(word))
		}
	}
	return strings.Join(words, " ")
}
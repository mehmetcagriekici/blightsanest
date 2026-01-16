package search

import(
	"strings"
	"regexp"
	"slices"

	"github.com/caneroj1/stemmer"
)

// text processing + tokenize + stem
func Tokenize(text string) []string {
	// convert the text to lowercase
	lower := strings.ToLower(text)

	// remove punbtuation except float points
	noPunc := removePunc(lower)

	// tokenization
	tokens := slices.DeleteFunc(strings.Split(noPunc, " "), func(word string) bool {
		if word == "" {
			return true
		}

		if _, isStop := stopwords[word]; isStop {
			return true
		}
		
		return false
	})

	// stem tokens
	stemmed := stemmer.StemMultiple(tokens)
	return stemmed
}

// Source - https://stackoverflow.com/a/62343672
// Posted by segFault, modified by community. See post 'Timeline' for change history
// Retrieved 2026-01-16, License - CC BY-SA 4.0
func removePunc(text string) string {
	// Regexp that finds all puncuation characters grouping the characters that wrap it
	re := regexp.MustCompile(`(.{0,1})([^\w\s])(.{0,1})`)

	// Regexp that determines if a given string is wrapped by digit characters
	isFloat := regexp.MustCompile(`\d([^\w\s])\d`)

	// Get the parts using the punctuation regexp... e.g. "t. "
	parts := re.FindAllString(text, -1)
	for _, part := range parts {
		// determine if the part is a float
		if !isFloat.MatchString(part) {
			newPart := re.ReplaceAllString(part, "$1$3")
			text = strings.Replace(text, part, newPart, 1)
		}
	}
	
	return text
}

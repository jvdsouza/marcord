package markovapi

import "strings"

type MarkovAPI struct{}

var words []string = make([]string, 0)

func (m MarkovAPI) ProcessTextMessage(s string) string {
	if strings.HasPrefix(s, "!read") || strings.HasPrefix(s, "!r") {
		return strings.Join(words, ", ")
	}
	words = append(words, s)
	return ""
}

package markovapi

type MarkovAPI struct{}

func (m MarkovAPI) ProcessTextMessage(s string) string {
	if s == "learn" {
		return "entering learning mode"
	}
	if s == "talk" {
		return "entering talking mode"
	}

	return s
}

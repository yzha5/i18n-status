package status

type Status struct {
	currentLang Lang
	messages    map[string]string
	code        Code
	msg         string
}

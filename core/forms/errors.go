package forms

type errors map[string][]string

// Add a new error to the errors list
func (e errors) Add(filed, message string) {
	e[filed] = append(e[filed], message)
}

// Get return the first error message
func (e errors) Get(filed string) string {
	es := e[filed]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

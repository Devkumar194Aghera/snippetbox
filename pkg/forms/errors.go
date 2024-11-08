package forms

type errors map[string][]string

func (e errors) Add(fields string, message string) {

	e[fields] = append(e[fields], message)
}

func (e errors) Get(fields string) string {
	errorString := e[fields]

	if len(errorString) == 0 {
		return ""
	}

	return errorString[0]

}

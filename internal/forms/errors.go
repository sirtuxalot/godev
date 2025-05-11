/* 99-My_Experimenting/internal/forms/errors.go */

package forms

type errors map[string][]string

// adds an error message for a given form field
func (fieldError errors) Add(field, message string) {
	fieldError[field] = append(fieldError[field], message)
}

// returns the first error message
func (fieldError errors) Get(field string) string {
	fieldMessage := fieldError[field]
	if len(fieldMessage) == 0 {
		return ""
	}
	return fieldMessage[0]
}

/* 99-My_Experimenting/internal/forms/forms.go */

package forms

import (
	//  "fmt"
	//  "github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

// creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// check for required fields
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field may not be blank")
		}
	}
}

// checks if form field in the post that is required is not empty
//func (form *Form) Has(field string) bool {
//  requiredField := form.Get(field)
//  if requiredField == "" {
//    return false
//  }
//  return true
//}

// returns true if form is valid, otherwise it returns false
func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

// checkts for minimum length of field responce
//func (form *Form) MinLength(field string, length int) bool {
//  if len(field) < length {
//    form.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
//    return false
//  }
//  return true
//}

// validates email provided in form
//func (form *Form) IsEmail(field string) {
//  if !govalidator.IsEmail(form.Get(field)) {
//    form.Errors.Add(field, "Email address is not valid")
//  }
//}

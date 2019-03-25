package validators

import (
	"net/http"

	"gopkg.in/thedevsaddam/govalidator.v1"
)

func Handler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	rules := govalidator.MapData{
		"slug":  []string{"required"},
		"title": []string{"required"},
	}
	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		RequiredDefault: true,  // all the field to be pass the rules
	}

	validator := govalidator.New(opts)
	validationErrors := validator.Validate()
	err := map[string]interface{}{"validationError": validationErrors}
	if len(validationErrors) != 0 {
		return err
	}
	return nil
}

func Category(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	return nil
}

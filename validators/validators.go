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

	messages := govalidator.MapData{
		"slug":  []string{"required:"},
		"title": []string{"required:"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	if len(e) == 0 {
		return nil
	}
	err := map[string]interface{}{"validationError": e}
	if e != nil {
		return err
	}
	return nil
}

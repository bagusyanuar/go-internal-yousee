package common

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translate_id "github.com/go-playground/validator/v10/translations/id"
)

func Validate(v *validator.Validate, request any) (err error, objErrors map[string]string) {
	err = v.Struct(request)
	if err != nil {
		localeID := id.New()
		uni := ut.New(localeID, localeID)
		trans, _ := uni.GetTranslator("id")
		translate_id.RegisterDefaultTranslations(v, trans)
		var tmpMap = make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			f, _ := reflect.TypeOf(request).Elem().FieldByName(field)
			jsonName, _ := f.Tag.Lookup("json")
			tmpMap[jsonName] = strings.ToLower(e.Translate(trans))
		}
		objErrors = tmpMap
	}
	return err, objErrors
}

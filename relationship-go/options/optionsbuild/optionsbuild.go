package optionsbuild

import (
	"reflect"

	"github.com/young2j/relationship/relationship-go/helpers"
	"github.com/young2j/relationship/relationship-go/options"
	"github.com/young2j/relationship/relationship-go/rule"
)

func FromMapping(m map[string]any) *options.ParameterOptions {
	opts := options.Default()

	optsVal := reflect.ValueOf(opts).Elem()
	for k, v := range m {
		k = helpers.Capitalize(k)
		field := optsVal.FieldByName(k)
		if !field.IsValid() {
			continue
		}

		if field.CanSet() {
			val := reflect.ValueOf(v)
			if field.Type() != reflect.TypeOf(v) {
				if reflect.TypeOf(v).ConvertibleTo(field.Type()) {
					val = reflect.ValueOf(v).Convert(field.Type())
				}
			}
			field.Set(val)
		}
	}

	return opts
}

func FromString(text string) *options.ParameterOptions {
	opts := options.Default()
	for _, item := range rule.Expression {
		match := item.Exp.FindStringSubmatch(text)
		if match != nil {
			opts = item.Opt(match)
			return opts
		}
	}

	opts.Text = text

	return opts
}

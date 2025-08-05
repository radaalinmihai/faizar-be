package auth

import (
	"fmt"
	"reflect"
)

type AccessToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
}

func (a *AccessToken) FromMap(data map[string]interface{}) error {
	v := reflect.ValueOf(a).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := t.Field(i)

		mapKey := structField.Tag.Get("json")
		if mapKey == "" {
			mapKey = structField.Name
		}

		if val, ok := data[mapKey]; ok {
			valValue := reflect.ValueOf(val)

			if field.Kind() == reflect.Uint && valValue.Kind() == reflect.Float64 {
				field.SetUint(uint64(val.(float64)))
				continue
			}

			if valValue.Type().AssignableTo(field.Type()) {
				field.Set(valValue)
			} else {
				return fmt.Errorf("cannot assign %v to field %s", valValue.Type(), structField.Name)
			}
		}
	}

	return nil
}

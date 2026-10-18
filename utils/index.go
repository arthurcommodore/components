package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

type validationError struct {
	Field   string
	Message string
}

func Validator(model any) ([]validationError, error) {
	val := reflect.ValueOf(model)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("the validator expects a struct, received %s", val.Kind())
	}

	typ := val.Type()
	var validationErrors []validationError

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		fieldName := fieldType.Name
		tag := fieldType.Tag.Get("validate")

		if strings.Contains(tag, "required") {
			if IsZeroValue(fieldVal) {
				validationErrors = append(validationErrors, validationError{
					Field:   fieldName,
					Message: fmt.Sprintf("o campo '%s' é obrigatório", fieldName),
				})
			}
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors, fmt.Errorf("validação falhou com %d erro(s)", len(validationErrors))
	}

	return nil, nil
}

func IsZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Pointer, reflect.Interface, reflect.Slice, reflect.Map:
		return !v.IsValid() || v.IsNil()
	default:
		return false
	}
}

package app

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validate *validator.Validate = validator.New()	

func ExistsColumnInTable(field validator.FieldLevel) bool {
	if field.Param() == "" {
		return false
	}
	params := strings.Split(field.Param(), ".")
	tableName := params[0]
	columName := "id"
	if len(params) > 1 {
		columName = params[1]
	}
	value := field.Field().Interface()
	if reflect.TypeOf(value) == reflect.TypeOf(uuid.UUID{}) {
		value = value.(uuid.UUID).String()
	} else {
		value = value.(string)
	}
	var count int64
	DB.Table(tableName).Where(columName +" = ?", value).Count(&count)
	return count > 0
}
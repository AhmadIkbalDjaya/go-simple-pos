package app

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validate *validator.Validate = validator.New()	

// format "exists=tableName.columName"
// tableName[required]
// columnName[optional,default="id"]
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

// format "uniqueRow=tableName.columnName.exceptColum.exceptValue"
// tableName[required]
// columnName[optional,default="id"]
// exceptColumn[optional]
// exceptValue[optional] required if exceptColumn Nut Null
func CheckUniqColumnInTable(field validator.FieldLevel) bool {
	if field.Param() == "" {
		return false
	}
	value := field.Field().Interface()
	if reflect.TypeOf(value) == reflect.TypeOf(uuid.UUID{}) {
		value = value.(uuid.UUID).String()
	} else {
		value = value.(string)
	}

	params := strings.Split(field.Param(), ".")
	tableName := params[0]
	columnName := "id"
	if len(params) > 1 {
		columnName = params[1]
	}
	var count int64 
	query := DB.Table(tableName).Where(columnName+" = ?", value)
	if len(params) == 4 {
		exceptColumn := params[2]
		exceptValue := params[3]
		query.Where(exceptColumn+" != ?", exceptValue).Count(&count)
	} else if len(params) == 2 || len(params) == 1 {
		query.Count(&count)
	} else {
		return false
	}
	return count <= 0
}
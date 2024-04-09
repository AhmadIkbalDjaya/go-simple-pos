package model

import (
	"reflect"

	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/exception"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Models interface {
	Category | Product
}

func GetModelById[T Models](ctx *fiber.Ctx, model *T, param string) (error) {
	id := ctx.Params(param)
	err := app.DB.First(model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewNotFoundError(err.Error(), reflect.TypeOf(model).Elem().Name())
		}
		return err
	}
	return nil
}
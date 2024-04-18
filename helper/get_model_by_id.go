package helper

import (
	"reflect"

	"github.com/AhmadIkbalDjaya/go-simple-pos/exception"
	"github.com/AhmadIkbalDjaya/go-simple-pos/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetModelById[T models.Models](ctx *fiber.Ctx, model *T, param string, db *gorm.DB) (error) {
	id := ctx.Params(param)
	err := db.First(model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewNotFoundError(err.Error(), reflect.TypeOf(model).Elem().Name())
		}
		return err
	}
	return nil
}
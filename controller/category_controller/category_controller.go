package categorycontroller

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/exception"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/api"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/category"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(ctx *fiber.Ctx) error {
	categories := []model.Category{}
	app.DB.Find(&categories)
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Category",
		Data: category.ToCategoryResponses(categories),
	})
}

func Show(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")
	var findCategory model.Category
	err := app.DB.First(&findCategory, categoryId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewNotFoundError(err.Error(), "Category")
		}
		return err
		// if err == gorm.ErrRecordNotFound {
		// 	return ctx.Status(fiber.StatusNotFound).JSON(api.BaseError{
		// 		Code: fiber.StatusNotFound,
		// 		Status: "Not Found",
		// 		Message: "Category Not Found",
		// 		Errors: err.Error(),
		// 	})
		// }
		// return ctx.Status(fiber.StatusInternalServerError).JSON(api.BaseError{
		// 	Code: fiber.StatusInternalServerError,
		// 	Status: "Internal Server Error",
		// 	Message: "Internal Server Error",
		// 	Errors: err.Error(),
		// })
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get Category",
		Data: category.ToCategoryResponse(findCategory),
	})
}

func Create(ctx *fiber.Ctx) error {
	categoryRequest := category.CategoryRequest{}
	ctx.BodyParser(&categoryRequest)
	err := app.Validate.Struct(categoryRequest)
	if err != nil {
		return err
		// errorMessages := helper.GetFieldErrors(err.(validator.ValidationErrors))
		// return ctx.Status(fiber.StatusBadRequest).JSON(api.BaseError{
		// 	Code: fiber.ErrBadRequest.Code,
		// 	Status: "Bad Request",
		// 	Message: "Data Is Not Valid",
		// 	Errors: errorMessages,
		// })
	}
	
	Newcategory := category.CategoryRequestToCategory(categoryRequest)
	err = app.DB.Create(&Newcategory).Error
	if err != nil {
		return nil
		// return ctx.Status(fiber.StatusInternalServerError).JSON(api.BaseError{
		// 	Code: fiber.StatusInternalServerError,
		// 	Status: "Internal Server Error",
		// 	Message: "Internal Server Error",
		// 	Errors: err.Error(),
		// })
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Category",
		Data: category.ToCategoryResponse(Newcategory),
	})
}

func Update(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")
	var findCategory model.Category
	err := app.DB.First(&findCategory, categoryId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewNotFoundError(err.Error(), "Category")
		}
		return err
	}
	
	categoryRequest := category.CategoryRequest{}
	ctx.BodyParser(&categoryRequest)
	err = app.Validate.Struct(categoryRequest)
	if err != nil {
		return err
	}

	findCategory.Name = categoryRequest.Name

	err = app.DB.Save(&findCategory).Error
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Update Category",
		Data: category.ToCategoryResponse(findCategory),
	})
}

func Delete(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")
	var findCategory model.Category
	err := app.DB.First(&findCategory, categoryId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewNotFoundError(err.Error(), "Category")
		}
		return err
	}
	err = app.DB.Delete(&findCategory).Error
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Delete Category",
	})
}
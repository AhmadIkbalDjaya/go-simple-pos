package categorycontroller

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/api"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/category"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	query := app.DB.Model(&model.Category{})
	categories := []model.Category{}
	var metaPaginate api.MetaPaginate

	model.PaginateModel(ctx, query, &categories, &metaPaginate)

	return ctx.Status(fiber.StatusOK).JSON(api.PaginateResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Category",
		Data: category.ToCategoryResponses(categories),
		Meta: metaPaginate,
	})
}

func Show(ctx *fiber.Ctx) error {
	var findCategory model.Category
	err := model.GetModelById(ctx, &findCategory, "categoryId")
	if err != nil {
		return err
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
	}
	
	newcategory := category.CategoryRequestToCategory(categoryRequest)
	err = app.DB.Create(&newcategory).Error
	if err != nil {
		return nil
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Create Category",
		Data: category.ToCategoryResponse(newcategory),
	})
}

func Update(ctx *fiber.Ctx) error {
	var findCategory model.Category
	err := model.GetModelById(ctx, &findCategory, "categoryId")
	if err != nil {
		return err
	}

	categoryRequest := category.CategoryRequest{}
	ctx.BodyParser(&categoryRequest)
	err = app.Validate.Struct(categoryRequest)
	if err != nil {
		return err
	}

	// findCategory.Name = categoryRequest.Name

	updateCategory := category.CategoryRequestToCategory(categoryRequest)
	// updateCategory.ID = findCategory.ID

	// err = app.DB.Save(&updateCategory).Error
	err = app.DB.Model(&findCategory).Updates(updateCategory).Error
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
	var findCategory model.Category
	err := model.GetModelById(ctx, &findCategory, "categoryId")
	if err != nil {
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
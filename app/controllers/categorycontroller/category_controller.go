package categorycontroller

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/helper"
	"github.com/AhmadIkbalDjaya/go-simple-pos/models"
	"github.com/AhmadIkbalDjaya/go-simple-pos/models/api"
	"github.com/AhmadIkbalDjaya/go-simple-pos/models/requests"
	"github.com/AhmadIkbalDjaya/go-simple-pos/models/responses"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	query := app.DB.Model(&models.Category{})
	categories := []models.Category{}
	var metaPaginate api.MetaPaginate

	helper.PaginateModel(ctx, query, &categories, &metaPaginate)

	return ctx.Status(fiber.StatusOK).JSON(api.PaginateResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Category",
		Data: responses.ToCategoryResponses(categories),
		Meta: metaPaginate,
	})
}

func Show(ctx *fiber.Ctx) error {
	var findCategory models.Category
	err := helper.GetModelById(ctx, &findCategory, "categoryId", app.DB)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get Category",
		Data: responses.ToCategoryResponse(findCategory),
	})
}

func Create(ctx *fiber.Ctx) error {
	categoryRequest := requests.CategoryRequest{}
	ctx.BodyParser(&categoryRequest)
	err := app.Validate.Struct(categoryRequest)
	if err != nil {
		return err
	}
	
	newcategory := requests.CategoryRequestToCategory(categoryRequest)
	err = app.DB.Create(&newcategory).Error
	if err != nil {
		return nil
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Create Category",
		Data: responses.ToCategoryResponse(newcategory),
	})
}

func Update(ctx *fiber.Ctx) error {
	var findCategory models.Category
	err := helper.GetModelById(ctx, &findCategory, "categoryId", app.DB)
	if err != nil {
		return err
	}

	categoryRequest := requests.CategoryRequest{}
	ctx.BodyParser(&categoryRequest)
	err = app.Validate.Struct(categoryRequest)
	if err != nil {
		return err
	}

	// findCategory.Name = categoryRequest.Name

	updateCategory := requests.CategoryRequestToCategory(categoryRequest)
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
		Data: responses.ToCategoryResponse(findCategory),
	})
}

func Delete(ctx *fiber.Ctx) error {
	var findCategory models.Category
	err := helper.GetModelById(ctx, &findCategory, "categoryId", app.DB)
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
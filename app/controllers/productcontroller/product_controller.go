package productcontroller

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
	query := app.DB.Model(&models.Product{}).Preload("Category")
	products := []models.Product{}
	metaPaginate := api.MetaPaginate{}

	helper.PaginateModel(ctx, query, &products, &metaPaginate)

	return ctx.Status(fiber.StatusOK).JSON(api.PaginateResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Product",
		Data: responses.ToProductResponses(products),
		Meta: metaPaginate,
	})
}

func Show(ctx *fiber.Ctx) error {
	var findProduct models.Product
	query := app.DB.Model(&models.Product{}).Preload("Category")
	err := helper.GetModelById(ctx, &findProduct, "productId", query)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get Product",
		Data: responses.ToProductResponse(findProduct),
	})
}

func Create(ctx *fiber.Ctx) error {
	productRequest := requests.ProductRequest{}
	ctx.BodyParser(&productRequest)
	err := app.Validate.Struct(productRequest)
	if err != nil {
		return err
	}

	newProduct := requests.ProductRequestToProduct(productRequest)
	err = app.DB.Create(&newProduct).Error
	if err != nil {
		return err
	}
	app.DB.Preload("Category").Take(&newProduct)
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Create Product",
		Data: responses.ToProductResponse(newProduct),
	})
}

func Update(ctx *fiber.Ctx) error {
	var findProduct models.Product
	query := app.DB.Model(&models.Product{}).Preload("Category")
	err := helper.GetModelById(ctx, &findProduct, "productId", query)
	if err != nil {
		return err
	}

	productRequest := requests.ProductRequest{}
	ctx.BodyParser(&productRequest)
	err = app.Validate.Struct(productRequest)
	if err != nil {
		return err
	}

	updateProduct := requests.ProductRequestToProduct(productRequest)

	err = app.DB.Model(&findProduct).Updates(updateProduct).Error
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Update Product",
		Data: responses.ToProductResponse(findProduct),
	})
}

func Delete(ctx *fiber.Ctx) error {
	var findProduct models.Product
	query := app.DB.Model(&models.Product{}).Preload("Category")
	err := helper.GetModelById(ctx, &findProduct, "productId", query)
	if err != nil {
		return err
	}
	err = app.DB.Delete(&findProduct).Error
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Delete Product",
	})
}
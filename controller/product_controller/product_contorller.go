package productcontroller

import (
	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/api"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/product"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	query := app.DB.Model(&model.Product{})
	products := []model.Product{}
	metaPaginate := api.MetaPaginate{}

	model.PaginateModel(ctx, query, &products, &metaPaginate)

	return ctx.Status(fiber.StatusOK).JSON(api.PaginateResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Product",
		Data: product.ToProductResponses(products),
		Meta: metaPaginate,
	})
}

func Show(ctx *fiber.Ctx) error {
	var findProduct model.Product
	err := model.GetModelById(ctx, &findProduct, "Product")
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get Product",
		Data: product.ToProductResponse(findProduct),
	})
}

func Create(ctx *fiber.Ctx) error {
	productRequest := product.ProductRequest{}
	ctx.BodyParser(&productRequest)
	err := app.Validate.Struct(productRequest)
	if err != nil {
		return err
	}

	newProduct := product.ProductRequestToProduct(productRequest)
	err = app.DB.Create(&newProduct).Error
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Create Product",
		Data: product.ToProductResponse(newProduct),
	})
}

func Update(ctx *fiber.Ctx) error {
	var findProduct model.Product
	err := model.GetModelById(ctx, &findProduct, "Product")
	if err != nil {
		return err
	}

	productRequest := product.ProductRequest{}
	ctx.BodyParser(&productRequest)
	err = app.Validate.Struct(productRequest)
	if err != nil {
		return err
	}

	updateProduct := product.ProductRequestToProduct(productRequest)

	err = app.DB.Model(&findProduct).Updates(updateProduct).Error
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Update Product",
		Data: product.ToProductResponse(findProduct),
	})
}

func Delete(ctx *fiber.Ctx) error {
	var findProduct model.Product
	err := model.GetModelById(ctx, &findProduct, "Product")
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
package categorycontroller

import (
	"math"
	"strconv"

	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/api"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/category"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	perpage, _ := strconv.Atoi(ctx.Query("perpage", "10"))
	search := ctx.Query("search", "")

	var totalItems int64
	query := app.DB.Model(&model.Category{})
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	query.Count(&totalItems)
	totalPage := int(math.Ceil(float64(totalItems)/float64(perpage)))

	offset := (page-1) * perpage
	categories := []model.Category{}
	query.Offset(offset).Limit(perpage).Find(&categories)
	return ctx.Status(fiber.StatusOK).JSON(api.PaginateResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Category",
		Data: category.ToCategoryResponses(categories),
		Meta: api.MetaPaginate{
			Page: page,
			Perpage: perpage,
			TotalPage: totalPage,
			TotalItem: int(totalItems),
		},
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
	
	Newcategory := category.CategoryRequestToCategory(categoryRequest)
	err = app.DB.Create(&Newcategory).Error
	if err != nil {
		return nil
	}
	return ctx.Status(fiber.StatusOK).JSON(api.BaseResponse{
		Code: fiber.StatusOK,
		Status: "OK",
		Message: "Success Get All Category",
		Data: category.ToCategoryResponse(Newcategory),
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
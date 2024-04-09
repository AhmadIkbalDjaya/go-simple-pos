package model

import (
	"math"
	"strconv"

	"github.com/AhmadIkbalDjaya/go-simple-pos/model/api"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PaginateModel[T Models](ctx *fiber.Ctx, tx *gorm.DB, models *[]T, metaPaginate *api.MetaPaginate) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	perpage, _ := strconv.Atoi(ctx.Query("perpage", "10"))
	search := ctx.Query("search", "")

	var totalItems int64
	query := tx
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	query.Count(&totalItems)
	totalPage := int(math.Ceil(float64(totalItems)/float64(perpage)))

	offset := (page-1) * perpage
	err := query.Offset(offset).Limit(perpage).Find(&models).Error
	if err != nil {
		return err
	}
	metaPaginate.Page = page
	metaPaginate.Perpage = perpage
	metaPaginate.TotalPage = totalPage
	metaPaginate.TotalItem = int(totalItems)
	return nil
}
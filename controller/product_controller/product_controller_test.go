package productcontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model/product"
	"github.com/AhmadIkbalDjaya/go-simple-pos/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var category1 = model.Category{
	Name: "Makanan",
}
var category2 = model.Category{
	Name: "Minuman",
}

func TestMain(m *testing.M) {
	app.DB = app.SetUpTestDatabase()
	app.DB.Migrator().DropTable(&model.Category{}, &model.Product{})
	app.DB.AutoMigrate(&model.Category{}, &model.Product{})
	
	app.DB.Create(&category1)
	app.DB.Create(&category2)

	m.Run()

	app.DB.Migrator().DropTable(&model.Product{}, &model.Category{})
}

func TestIndex(t *testing.T) {
	product1 := model.Product{
		Code: "P001",
		Name: "Oreo",
		Unit: "pcs",
		CategoryId: category1.ID,
		Stock: 100,
		PurchasePrice: int64(1000),
		SellingPrice: int64(1500),
	}
	product2 := model.Product{
		Code: "P002",
		Name: "Fanta",
		Unit: "pcs",
		CategoryId: category2.ID,
		Stock: 200,
		PurchasePrice: int64(3000),
		SellingPrice: int64(5000),
	}
	app.DB.Create(&product1)
	app.DB.Create(&product2)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	request := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	response, err := fiberApp.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)

	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Get All Product", responseJSON["message"])

	products := responseJSON["data"].([]interface{})
	// fmt.Println(products[0].(map[string]interface{})["category"].(map[string]interface{})["name"/])
	assert.Equal(t, 2, len(products))
}

func TestShow(t *testing.T) {
	product3 := model.Product{
		Code: "P003",
		Name: "Fresh Tea",
		Unit: "pcs",
		CategoryId: category2.ID,
		Stock: 300,
		PurchasePrice: int64(4000),
		SellingPrice: int64(5000),
	}
	app.DB.Create(&product3)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	request := httptest.NewRequest(http.MethodGet, "/api/products/" + product3.ID.String(), nil)
	response, err := fiberApp.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)

	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Get Product", responseJSON["message"])

	responseProduct := responseJSON["data"].(map[string]interface{})
	assert.Equal(t, product3.ID.String(), responseProduct["id"])
	assert.Equal(t, product3.Code, responseProduct["code"])
	assert.Equal(t, product3.Name, responseProduct["name"])
	assert.Equal(t, product3.CategoryId.String(), responseProduct["category"].(map[string]interface{})["id"])
	assert.Equal(t, category2.Name, responseProduct["category"].(map[string]interface{})["name"])
}

func TestCreate(t *testing.T) {
	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	var productsCountBefore int64
	app.DB.Model(&model.Product{}).Count(&productsCountBefore)

	newProduct := product.ProductRequest{
		Code: "P004",
		Name: "Otside",
		Unit: "pcs",
		CategoryId: category2.ID.String(),
		Stock: 320,
		PurchasePrice: 6500,
		SellingPrice: 8000,
	}
	reqBody, err := json.Marshal(newProduct)
	assert.Nil(t, err)
	request := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewBuffer(reqBody))
	request.Header.Add("Content-type", "application/json")
	response, err := fiberApp.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	
	var productsCountAfter int64
	app.DB.Model(&model.Product{}).Count(&productsCountAfter)
	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)
	
	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Create Product", responseJSON["message"])
	
	responseProduct := responseJSON["data"].(map[string]interface{})
	assert.Equal(t, newProduct.Code, responseProduct["code"])
	assert.Equal(t, newProduct.Name, responseProduct["name"])
	assert.Equal(t, newProduct.CategoryId, responseProduct["category"].(map[string]interface{})["id"])
	assert.Equal(t, category2.Name, responseProduct["category"].(map[string]interface{})["name"])
	assert.Equal(t, productsCountBefore + 1, productsCountAfter)
}

func TestUpdate(t *testing.T) {
	product5 := model.Product{
		Code: "P005",
		Name: "Ship",
		Unit: "pcs",
		CategoryId: category1.ID,
		Stock: 500,
		PurchasePrice: int64(1000),
		SellingPrice: int64(15000),
	}
	err := app.DB.Create(&product5).Error
	assert.Nil(t, err)
	updateProduct := product.ProductRequest{
		Code: "P005",
		Name: "Shippp",
		Unit: "box",
		CategoryId: category1.ID.String(),
		Stock: 500,
		PurchasePrice: int64(10000),
		SellingPrice: int64(15000),
	}
	reqBody, err := json.Marshal(updateProduct)
	assert.Nil(t, err)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	// requestBody := strings.NewReader(`{"name":"Shipp"}`)
	request := httptest.NewRequest(http.MethodPut, "/api/products/"+ product5.ID.String(), bytes.NewBuffer(reqBody))
	request.Header.Add("Content-type", "application/json")

	response, err := fiberApp.Test(request)
	fmt.Println(response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)
	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Update Product", responseJSON["message"])
	
	productResponse := responseJSON["data"].(map[string]interface{})
	assert.NotEqual(t, product5.Name, productResponse["name"])
	assert.NotEqual(t, product5.Unit, productResponse["unit"])
}

func TestDelete(t *testing.T) {
	product6 := model.Product{
		Code: "P006",
		Name: "Good Day",
		Unit: "pcs",
		CategoryId: category2.ID,
		Stock: 500,
		PurchasePrice: int64(500),
		SellingPrice: int64(1000),
	}
	app.DB.Create(&product6)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	var productsCountBefore int64
	app.DB.Model(&model.Product{}).Count(&productsCountBefore)

	request := httptest.NewRequest(http.MethodDelete, "/api/products/" + product6.ID.String(), nil)
	response, err := fiberApp.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)
	
	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Delete Product", responseJSON["message"])

	var productsCountAfter int64
	app.DB.Model(&model.Product{}).Count(&productsCountAfter)
	assert.Equal(t, productsCountBefore - 1 , productsCountAfter)
	
	
	
}
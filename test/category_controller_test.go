package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/model"
	"github.com/AhmadIkbalDjaya/go-simple-pos/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)


func TestMain(m *testing.M) {
	app.DB = app.SetUpTestDatabase()
	app.DB.Migrator().DropTable(&model.Category{})
	app.DB.AutoMigrate(&model.Category{})

	m.Run()

	app.DB.Migrator().DropTable(&model.Category{})
}

func TestIndex(t *testing.T) {
	app.DB.Create(&model.Category{Name: "Makanan"})
	app.DB.Create(&model.Category{Name: "Minuman"})

	app := fiber.New()
	routes.SetUpRoutes(app)

	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)

	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Get All Category", responseJSON["message"])

	categories := responseJSON["data"].([]interface{})
	assert.Equal(t, 2, len(categories))	
}

func TestShow(t *testing.T) {
	category1 := model.Category{
		Name: "Makanan",
	}
	category2 := model.Category{
		Name: "Minuman",
	}
	app.DB.Create(&category1)
	app.DB.Create(&category2)

	app := fiber.New()
	routes.SetUpRoutes(app)

	request := httptest.NewRequest(http.MethodGet, "/api/categories/" + category1.ID.String(), nil)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)

	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Get Category", responseJSON["message"])

	category := responseJSON["data"].(map[string]interface{})
	assert.Equal(t, "Makanan", category["name"])
	assert.Equal(t, category1.ID.String(), category["id"])
}

func TestCreate(t *testing.T) {
	app.DB.Create(&model.Category{Name: "Perkakas"})
	app := fiber.New()
	routes.SetUpRoutes(app)
	
	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("Content-type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	
	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)
	
	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Create Category", responseJSON["message"])
	
	data := responseJSON["data"].(map[string]interface{})
	assert.Equal(t, "Gadget", data["name"].(string))
}

func TestUpdate(t *testing.T) {
	category1 := model.Category{
		Name: "Perkakas",
	}
	app.DB.Create(&category1)
	app := fiber.New()
	routes.SetUpRoutes(app)
	
	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+ category1.ID.String(), requestBody)
	request.Header.Add("Content-type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	
	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)
	
	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Update Category", responseJSON["message"])
	
	data := responseJSON["data"].(map[string]interface{})
	assert.Equal(t, "Gadget", data["name"].(string))
}

func TestDelete(t *testing.T) {
	category1 := model.Category{
		Name: "Makanan",
	}
	category2 := model.Category{
		Name: "Minuman",
	}
	app.DB.Create(&category1)
	app.DB.Create(&category2)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	categoriesBefore := []model.Category{}
	app.DB.Find(&categoriesBefore)
	request := httptest.NewRequest(http.MethodDelete, "/api/categories/" + category1.ID.String(), nil)
	response, err := fiberApp.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)

	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Delete Category", responseJSON["message"])

	categories := []model.Category{}
	app.DB.Find(&categories)
	assert.Equal(t, len(categoriesBefore)-1, len(categories))
	
}
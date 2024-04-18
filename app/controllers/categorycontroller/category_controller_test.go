package categorycontroller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AhmadIkbalDjaya/go-simple-pos/app"
	"github.com/AhmadIkbalDjaya/go-simple-pos/models"
	"github.com/AhmadIkbalDjaya/go-simple-pos/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)


func TestMain(m *testing.M) {
	app.DB = app.SetUpTestDatabase()
	app.DB.Migrator().DropTable(&models.Category{})
	app.DB.AutoMigrate(&models.Category{})

	m.Run()

	app.DB.Migrator().DropTable(&models.Category{})
}

func TestIndex(t *testing.T) {
	app.DB.Create(&models.Category{Name: "Makanan"})
	app.DB.Create(&models.Category{Name: "Minuman"})

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	response, err := fiberApp.Test(request)

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
	category1 := models.Category{
		Name: "Makanan",
	}
	category2 := models.Category{
		Name: "Minuman",
	}
	app.DB.Create(&category1)
	app.DB.Create(&category2)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	request := httptest.NewRequest(http.MethodGet, "/api/categories/" + category1.ID.String(), nil)
	response, err := fiberApp.Test(request)

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
	app.DB.Create(&models.Category{Name: "Perkakas"})
	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)
	
	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("Content-type", "application/json")

	response, err := fiberApp.Test(request)
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
	category1 := models.Category{
		Name: "Perkakas",
	}
	app.DB.Create(&category1)
	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)
	
	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "/api/categories/"+ category1.ID.String(), requestBody)
	request.Header.Add("Content-type", "application/json")

	response, err := fiberApp.Test(request)
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
	category1 := models.Category{
		Name: "Makanan",
	}
	category2 := models.Category{
		Name: "Minuman",
	}
	app.DB.Create(&category1)
	app.DB.Create(&category2)

	fiberApp := fiber.New()
	routes.SetUpRoutes(fiberApp)

	var categoryCountBefore int64
	app.DB.Model(&models.Category{}).Count(&categoryCountBefore)
	request := httptest.NewRequest(http.MethodDelete, "/api/categories/" + category1.ID.String(), nil)
	response, err := fiberApp.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseJSON map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	assert.Nil(t, err)

	assert.Equal(t, "OK", responseJSON["status"])
	assert.Equal(t, "Success Delete Category", responseJSON["message"])

	var categoryCountAfter int64
	app.DB.Model(&models.Category{}).Count(&categoryCountAfter)
	assert.Equal(t, categoryCountBefore - 1, categoryCountAfter)
	
}
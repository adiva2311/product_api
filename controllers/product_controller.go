package controllers

import (
	"fmt"
	"net/http"

	"github.com/adiva2311/product_api.git/helpers"
	"github.com/adiva2311/product_api.git/models"
	"github.com/adiva2311/product_api.git/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController interface {
	Create(c echo.Context) error
}

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func (p *ProductControllerImpl) Create(c echo.Context) error {
	productPayload := new(helpers.ProductRequest)
	err := c.Bind(productPayload)
	if err != nil {
		return err
	}

	// Ambil user_id dari context
	userIdInterface := c.Get("user_id")
	if userIdInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	// Ambil Role dari context
	roleInterface := c.Get("role")
	if roleInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Role Undetected"})
	}
	fmt.Println(roleInterface)

	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}

	productPayload.UserID = uint(userIdFloat)

	if roleInterface == "admin" {
		result := p.ProductService.Create(models.Product{
			ProductName: productPayload.ProductName,
			Total:       productPayload.Total,
			Price:       productPayload.Price,
			UserID:      productPayload.UserID,
		})

		apiResponse := helpers.ApiResponse{
			Status:  http.StatusOK,
			Message: "Berhasil Tambah Data",
			Data:    result,
		}

		return c.JSON(http.StatusOK, apiResponse)
	}
	return c.JSON(http.StatusUnauthorized, echo.Map{"message": "You Are Not Admin"})

}

func NewProductController(db *gorm.DB) ProductControllerImpl {
	service := services.NewProductService(db)
	controller := ProductControllerImpl{
		ProductService: service,
	}
	return controller
}

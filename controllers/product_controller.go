package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adiva2311/product_api.git/helpers"
	"github.com/adiva2311/product_api.git/models"
	"github.com/adiva2311/product_api.git/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductController interface {
	Create(c echo.Context) error
	GetByUserId(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
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

	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}

	productPayload.UserID = uint(userIdFloat)

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

func (p *ProductControllerImpl) GetByUserId(c echo.Context) error {
	// Ambil user_id dari context
	userIdInterface := c.Get("user_id")
	if userIdInterface == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user_id type"})
	}

	userID := int(userIdFloat)

	result := p.ProductService.GetByUserId(userID)

	apiResponse := helpers.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Ambil Data Berdasarkan UserID",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (p *ProductControllerImpl) Update(c echo.Context) error {
	productPayload := new(helpers.UpdateProductRequest)
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

	// Ambil product id dari param
	productID := c.Param("product_id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "ID tidak valid",
		})
	}
	productPayload.ID = uint(id)

	if roleInterface == "admin" {
		result := p.ProductService.Update(int(productPayload.ID), int(productPayload.UserID), models.Product{
			ID:          productPayload.ID,
			ProductName: productPayload.ProductName,
			Total:       productPayload.Total,
			Price:       productPayload.Price,
			UserID:      productPayload.UserID,
		})

		apiResponse := helpers.ApiResponse{
			Status:  http.StatusOK,
			Message: "Berhasil Ubah Data",
			Data:    result,
		}

		return c.JSON(http.StatusOK, apiResponse)
	}
	return c.JSON(http.StatusInternalServerError, helpers.ApiResponse{
		Status:  http.StatusNonAuthoritativeInfo,
		Message: "Cannot Update! You are not admin",
	})
}

func (p *ProductControllerImpl) Delete(c echo.Context) error {
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
	User_id := uint(userIdFloat)

	// Ambil product id dari param
	productID := c.Param("product_id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "productID tidak valid",
		})
	}
	product_id := uint(id)

	if roleInterface == "admin" {
		err = p.ProductService.Delete(int(product_id), int(User_id))
		if err != nil {
			fmt.Printf("Delete error: %v\n", err)
			return c.JSON(http.StatusInternalServerError, helpers.ApiResponse{
				Status:  http.StatusBadRequest,
				Message: "Failed to Delete Data",
			})
		}

		apiResponse := helpers.ApiResponse{
			Status:  http.StatusOK,
			Message: "Berhasil Hapus Data",
		}

		return c.JSON(http.StatusOK, apiResponse)
	}
	return c.JSON(http.StatusInternalServerError, helpers.ApiResponse{
		Status:  http.StatusNonAuthoritativeInfo,
		Message: "Cannot Delete! You are not admin",
	})
}

func NewProductController(db *gorm.DB) ProductControllerImpl {
	service := services.NewProductService(db)
	controller := ProductControllerImpl{
		ProductService: service,
	}
	return controller
}

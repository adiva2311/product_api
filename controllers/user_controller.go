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

type UserController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type UserControllerImpl struct {
	UserService services.UserService
}

// Login implements UserController.
func (u *UserControllerImpl) Login(c echo.Context) error {
	panic("unimplemented")
}

// Register implements UserController.
func (u *UserControllerImpl) Register(c echo.Context) error {
	userPayload := new(helpers.RegisterRequest)

	err := c.Bind(userPayload)
	if err != nil {
		return err
	}

	result, err := u.UserService.Register(models.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Password: userPayload.Password,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal register", "error": err.Error()})
	}

	fmt.Println(result)

	apiResponse := helpers.ApiResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Register",
		Data:    result,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func NewUserController(db *gorm.DB) UserControllerImpl {
	service := services.NewUserService(db)
	controllers := UserControllerImpl{
		UserService: service,
	}
	return controllers
}

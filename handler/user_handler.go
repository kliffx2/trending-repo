package handler

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/kliffx2/trending-repo/model"
	req "github.com/kliffx2/trending-repo/model/req"
	"github.com/kliffx2/trending-repo/repository"
	"github.com/kliffx2/trending-repo/security"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil{
		//log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil{
		//log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		//log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message: err.Error(),
			Data: nil,
		})
	}

	user := model.User{
		UserId:    userId.String(),
		FullName:  req.Email,
		Email:     req.Email,
		Password:  hash,
		Role:      role,
		
		Token:     "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil{
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	
	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       user,
	})
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "Dat",
		"email": "dat@gmail.com",
	})
}
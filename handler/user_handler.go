package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"github.com/kliffx2/trending-repo/fault"
	"github.com/kliffx2/trending-repo/model"
	"github.com/kliffx2/trending-repo/model/req"
	"github.com/kliffx2/trending-repo/repository"
	"github.com/kliffx2/trending-repo/security"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	UserRepo repository.UserRepo
}

// SignUp godoc
// @Summary Create new account
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.ReqSignUp true "user"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/sign-up [post]
func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil{
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	if err := c.Validate(req); err != nil{
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
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message: err.Error(),
			Data: nil,
		})
	}

	user := model.User{
		UserId:    userId.String(),
		FullName:  req.FullName,
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

	// generate token
	token, err := security.GenToken(user) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}
	user.Token = token

	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Signup successfully",
		Data:       user,
	})
}

// SignIn godoc
// @Summary Sign in to access your account
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.ReqSignIn true "user"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/sign-in [post]
func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req := req.ReqSignIn{}
	if err := c.Bind(&req); err != nil{
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	if err := c.Validate(req); err != nil{
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message: err.Error(),
			Data: nil,
		})
	}

	//check password
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame{
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message: "Login failed",
			Data: nil,
		})
	}
	
	// generate token
	token, err := security.GenToken(user) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}
	user.Token = token

	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Login successfully",
		Data: user,
	})
}

// Profile godoc
// @Summary get user profile
// @Tags user-service
// @Accept  json
// @Produce  json
// @Security jwt
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/profile [get]
func (u *UserHandler) Profile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil{
		if err == fault.UserNotFound{
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message: err.Error(),
				Data: nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Success",
		Data: user,
	})
}

// UpdateProfile godoc
// @Summary get user profile
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.ReqUpdateUser true "user"
// @Security jwt
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/profile/update [put]
func (u UserHandler) UpdateProfile(c echo.Context) error {
	req := req.ReqUpdateUser{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate sent info
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	user := model.User{
		UserId:   claims.UserId,
		FullName: req.FullName,
		Email:    req.Email,
	}

	user, err = u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Success",
		Data:       user,
	})
}

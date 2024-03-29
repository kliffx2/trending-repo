package handler

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kliffx2/trending-repo/model"
	"github.com/kliffx2/trending-repo/model/req"
	"github.com/kliffx2/trending-repo/repository"
	"github.com/labstack/echo/v4"
)

type RepoHandler struct {
	GithubRepo repository.GithubRepo
}

// GetTrendingRepo godoc
// @Summary get trending repos
// @Tags repo-service
// @Accept  json
// @Produce  json
// @Security jwt
// @Success 200 {object} model.Response
// @Router /github/trending [get]
func (r RepoHandler) RepoTrending(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	repos, _ := r.GithubRepo.SelectRepos(c.Request().Context(), claims.UserId, 25)
	for i, repo := range repos {
		repos[i].Contributors = strings.Split(repo.BuildBy, ",")
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       repos,
	})
}

// SelectBookmark godoc
// @Summary get bookmark list
// @Tags repo-service
// @Accept  json
// @Produce  json
// @Security jwt
// @Success 200 {object} model.Response
// @Router /bookmark/list [get]
func (r RepoHandler) SelectBookmarks(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	repos, _ := r.GithubRepo.SelectAllBookmarks(
		c.Request().Context(),
		claims.UserId)

	for i, repo := range repos {
		repos[i].Contributors = strings.Split(repo.BuildBy, ",")
	}	

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       repos,
	})
}

// Bookmark godoc
// @Summary add bookmark 
// @Tags repo-service
// @Accept  json
// @Produce  json
// @Security jwt
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /bookmark/add [post]
func (r RepoHandler) Bookmark(c echo.Context) error {
	req := req.ReqBookmark{}
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

	bId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err = r.GithubRepo.Bookmark(
		c.Request().Context(),
		bId.String(),
		req.RepoName,
		claims.UserId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Bookmark successfully",
		Data:       nil,
	})
}

// DelBookmark godoc
// @Summary delete bookmark 
// @Tags repo-service
// @Accept  json
// @Produce  json
// @Security jwt
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /bookmark/delete [delete]
func (r RepoHandler) DelBookmark(c echo.Context) error {
	req := req.ReqBookmark{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate thông tin gửi lên
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	err = r.GithubRepo.DelBookmark(
		c.Request().Context(),
		req.RepoName, claims.UserId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete bookmark successfully",
		Data:       nil,
	})
}
package repository

import (
	"context"

	"github.com/kliffx2/trending-repo/model"
	"github.com/kliffx2/trending-repo/model/req"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)	
	SelectUserById(context context.Context, userId string) (model.User, error)	
	UpdateUser(context context.Context, user model.User) (model.User, error)
}
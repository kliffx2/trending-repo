package repository

import (
	"context"

	"github.com/kliffx2/trending-repo/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
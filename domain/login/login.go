package login

import "github.com/rcgc/go-ecommerce/model"

type UseCase interface {
	Login(email, password, jwtSecretKey string) (model.User, string, error)
}

type UseCaseUser interface {
	Login(email, password string) (model.User, error)
}

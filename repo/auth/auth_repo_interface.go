package auth

import "rest/entities"

type AuthRepoInterface interface {
	Login(user entities.User) (entities.User, error)
}

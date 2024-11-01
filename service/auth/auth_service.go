package auth

import (
	"rest/constant"
	"rest/entities"
	"rest/middleware"
	"rest/repo/auth"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(ar auth.AuthRepoInterface, jt middleware.JwtInterface) *AuthService {
	return &AuthService{
		authRepoInterface: ar,
		jwtInterface:      jt,
	}
}

type AuthService struct {
	authRepoInterface auth.AuthRepoInterface
	jwtInterface      middleware.JwtInterface
}

func (authService AuthService) Login(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, constant.EMAIL_IS_EMPTY
	} else if user.Password == "" {
		return entities.User{}, constant.PASSWORD_IS_EMPTY
	}

	var err error
	user, err = authService.authRepoInterface.Login(user)
	if err != nil {
		return entities.User{}, err
	}
	// disini logika naruh JWT
	token, err := authService.jwtInterface.GenerateJWT(user.ID, user.Nama)
	if err != nil {
		return entities.User{}, err
	}
	user.Token = token

	return user, nil
}
func (authService AuthService) Register(user entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Add(a, b int) int {
	result := a + b
	if result < 0 {
		return 0
	}

	return result
}

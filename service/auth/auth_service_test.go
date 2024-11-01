package auth

// import (
// 	"rest/entities"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var authService AuthService

// type AuthRepoDummy struct{}

// func (ard AuthRepoDummy) Login(user entities.User) (entities.User, error) {
// 	return entities.User{
// 		ID:       1,
// 		Nama:     "Fajar",
// 		Email:    "test@mail.com",
// 		Password: "123",
// 		Token:    "321",
// 	}, nil
// }

// func setup() {
// 	// authRepoDummy := AuthRepoDummy{}
// 	// authService = *NewAuthService(authRepoDummy)
// }

// func TestAuthService_Login(t *testing.T) {
// 	setup()
// 	t.Run("sukses login", func(t *testing.T) {
// 		// login, respond ni
// 		user, err := authService.Login(entities.User{Email: "alta@gmail.com", Password: "123"})

// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, user.ID)
// 	})

// 	t.Run("gagal login email kosong", func(t *testing.T) {
// 		user, err := authService.Login(entities.User{Password: "123"})

// 		assert.NotNil(t, err)
// 		assert.Equal(t, "email empty", err.Error())
// 		assert.Equal(t, 0, user.ID)
// 	})
// }

// func TestAdd(t *testing.T) {
// 	t.Run("keduanya positif", func(t *testing.T) {
// 		result := Add(5, 5)
// 		if result != 10 {
// 			t.Error("5 + 5 hasilnya bukan 10")
// 		}
// 	})

// 	t.Run("keduanya negatif", func(t *testing.T) {
// 		result := Add(-5, -5)
// 		if result != 0 {
// 			t.Error("-5 + -5 hasilnya harus 0")
// 		}
// 	})
// }

package service

import (
	"errors"
	"fmt"
	"majoo/repo"
)

type AuthService interface {
	VerifyCredential(username string, password string) error
	// CreateUser(user dto.) entity.User
}

type authService struct {
	userRepo repo.UserRepository
}

func NewAuthService(userRepo repo.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (c *authService) VerifyCredential(username string, password string) error {
	user, err := c.userRepo.FindByUserName(username)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	isValidPassword := comparePassword(user.Password, password)
	if !isValidPassword {
		return errors.New("invalid username or password")
	}

	return nil

}

func comparePassword(hashedPwd string, plainPassword string) bool {
	md5Hash := repo.HashAndSalt(plainPassword)
	return md5Hash == hashedPwd
}


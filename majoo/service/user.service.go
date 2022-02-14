package service

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"majoo/dto"
	"majoo/repo"
	"majoo/response"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) (*response.UserResponse, error)
	FindUserByUserName(username string) (*response.UserResponse, error)
	FindUserByID(userID string) (*response.UserResponse, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}


func (c *userService) CreateUser(registerRequest dto.RegisterRequest) (*response.UserResponse, error) {
	user, err := c.userRepo.FindByUserName(registerRequest.UserName)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	user, err = c.userRepo.InsertUser(user)
	if err == nil {
		return nil, errors.New("failed saved to db")
	}
	res := response.NewUserResponse(user)
	return &res, nil

}

func (c *userService) FindUserByUserName(username string) (*response.UserResponse, error) {
	user, err := c.userRepo.FindByUserName(username)

	if err != nil {
		return nil, err
	}

	userResponse := response.NewUserResponse(user)
	return &userResponse, nil
}

func (c *userService) FindUserByID(userID string) (*response.UserResponse, error) {
	user, err := c.userRepo.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

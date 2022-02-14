package repo

import (
	"io"
	"fmt"
	"majoo/entity"
	"crypto/md5"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	FindByUserName(username string) (entity.User, error)
	FindByUserID(userID string) (entity.User, error)
}

type userRepo struct {
	connection *gorm.DB
}

func NewUserRepo(connection *gorm.DB) UserRepository {
	return &userRepo{
		connection: connection,
	}
}

func (c *userRepo) InsertUser(user entity.User) (entity.User, error) {
	user.Password = HashAndSalt(user.Password)
	c.connection.Save(&user)
	return user, nil
}



func (c *userRepo) FindByUserName(username string) (entity.User, error) {
	var user entity.User
	res := c.connection.Where("user_name = ?", username).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (c *userRepo) FindByUserID(userID string) (entity.User, error) {
	var user entity.User
	res := c.connection.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func HashAndSalt(pwd string ) string {
	h := md5.New()
	io.WriteString(h, pwd)
	return fmt.Sprintf("%x", h.Sum(nil))
	
	
}

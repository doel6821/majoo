package response

import "majoo/entity"

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	UserName string `json:"user_name"`
	Token string `json:"token"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		UserName: user.UserName,
		Name:  user.Name,
	}
}

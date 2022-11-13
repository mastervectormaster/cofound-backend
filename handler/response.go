package handler

import (
	_ "github.com/gofiber/fiber/v2"
	"github.com/mastervectormaster/cofound-backend/model"
	"github.com/mastervectormaster/cofound-backend/user"
	"github.com/mastervectormaster/cofound-backend/utils"
)

type userResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

type profileResponse struct {
	Profile struct {
		Username  string  `json:"username"`
		Bio       *string `json:"bio"`
		Image     *string `json:"image"`
		Following bool    `json:"following"`
	} `json:"profile"`
}

func newProfileResponse(us user.Store, userID uint, u *model.User) *profileResponse {
	r := new(profileResponse)
	r.Profile.Username = u.Username
	return r
}

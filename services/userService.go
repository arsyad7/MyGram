package services

import (
	"mygram/helpers"
	"mygram/models"
	"mygram/params"
	"mygram/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	userService repositories.UserRepo
}

var db = gorm.DB{}
var repo = repositories.NewUserRepo(&db)

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{
		userService: *repo,
	}
}

func (u *UserService) Register(req *params.CreateUser) (*params.UserResponse, *params.Response) {
	model := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}

	err := u.userService.Register(&model)
	if err != nil {
		errResp := params.Response{
			Status:         400,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return &params.UserResponse{
		Age:      model.Age,
		Email:    model.Email,
		ID:       int(model.ID),
		Username: model.Username,
	}, nil
}

func (u *UserService) Login(req *params.CreateUser) (*params.UserResponse, *params.Response) {
	user, err := u.userService.FindUserByEmail(req.Email)
	if err != nil {
		errResp := params.Response{
			Status:  404,
			Message: err.Error(),
		}
		return nil, &errResp
	}

	checkPass := helpers.ComparePass(user.Password, req.Password)
	if !checkPass {
		return nil, &params.Response{
			Status:  401,
			Message: "Wrong email / password",
		}
	}

	accessToken := helpers.GenerateToken(user.ID, user.Email)
	return &params.UserResponse{
		Token: accessToken,
	}, nil
}

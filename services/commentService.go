package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
	"net/http"
)

type CommentService struct {
	commentService repositories.CommentRepo
}

func NewCommentService(repo *repositories.CommentRepo) *CommentService {
	return &CommentService{commentService: *repo}
}

func (c *CommentService) CreateComment(req *params.CreateComment, id int) (*params.CommentResponse, *params.Response) {
	comment := models.Comment{
		Message: req.Message,
		PhotoID: uint(req.PhotoID),
		UserID:  uint(id),
	}

	err := c.commentService.CreateComment(&comment)
	if err != nil {
		return nil, &params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.CommentResponse{
		ID:        int(comment.ID),
		Message:   comment.Message,
		PhotoID:   int(comment.PhotoID),
		UserID:    int(comment.UserID),
		CreatedAt: &comment.CreatedAt,
	}, nil
}

func (c *CommentService) GetComments() (*[]params.CommentResponse, *params.Response) {
	comments, err := c.commentService.GetComments()
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	var list []params.CommentResponse
	for _, v := range *comments {
		list = append(list, params.CommentResponse{
			ID:        int(v.ID),
			Message:   v.Message,
			PhotoID:   int(v.PhotoID),
			UserID:    int(v.UserID),
			CreatedAt: &v.CreatedAt,
			UpdatedAt: &v.UpdatedAt,
			User: &params.UserComment{
				ID:       int(v.User.ID),
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: &params.PhotoResponse{
				ID:       int(v.Photo.ID),
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoUrl: v.Photo.PhotoUrl,
				UserID:   int(v.Photo.UserID),
			},
		})
	}
	return &list, nil
}

func (c CommentService) UpdateComment(req *params.CreateComment, id int) (*params.CommentResponse, *params.Response) {
	comment := models.Comment{
		Message: req.Message,
	}

	res, err := c.commentService.UpdateComment(&comment, id)
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return &params.CommentResponse{
		ID:        int(res.ID),
		Message:   res.Message,
		PhotoID:   int(res.PhotoID),
		UpdatedAt: &res.UpdatedAt,
		UserID:    int(res.UserID),
	}, nil
}

func (c *CommentService) DeleteComment(id int) *params.Response {
	err := c.commentService.DeleteComment(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Your comment has been successfully deleted",
	}
}

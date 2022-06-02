package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
	"net/http"
)

type SocialMediaService struct {
	socialMediaService repositories.SocialMediaRepo
}

func NewSocialMediaService(repo *repositories.SocialMediaRepo) *SocialMediaService {
	return &SocialMediaService{socialMediaService: *repo}
}

func (s *SocialMediaService) CreateSocialMedia(req *params.CreateSocialMedia, id int) (*params.SocialMediaResponse, *params.Response) {
	socmed := models.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         uint(id),
	}

	err := s.socialMediaService.CreateSocialMedia(&socmed)
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return &params.SocialMediaResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaUrl: socmed.SocialMediaUrl,
		UserID:         int(socmed.UserID),
		CreatedAt:      &socmed.CreatedAt,
	}, nil
}

func (s *SocialMediaService) GetSocialMedias() (*params.Socmed, *params.Response) {
	socmeds, err := s.socialMediaService.GetSocialMedias()
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	var list params.Socmed

	for _, v := range *socmeds {
		list.SocialMedias = append(list.SocialMedias, params.SocialMediaResponse{
			ID:             int(v.ID),
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserID:         int(v.UserID),
			CreatedAt:      &v.CreatedAt,
			UpdatedAt:      &v.UpdatedAt,
			User: &params.UserResponse{
				ID:       int(v.User.ID),
				Username: v.User.Username,
				Email:    v.User.Email,
			},
		})
	}

	return &list, nil
}

func (s *SocialMediaService) UpdateSocialMedia(req *params.CreateSocialMedia, id int) (*params.SocialMediaResponse, *params.Response) {
	socmed := models.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}

	res, err := s.socialMediaService.UpdateSocialMedia(&socmed, id)
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return &params.SocialMediaResponse{
		ID:             int(res.ID),
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         int(res.UserID),
		UpdatedAt:      &res.UpdatedAt,
	}, nil
}

func (s *SocialMediaService) DeleteSocialMedia(id int) *params.Response {
	err := s.socialMediaService.DeleteSocialMedia(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Your social media has been successfully deleted",
	}
}

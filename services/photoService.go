package services

import (
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
	"net/http"
)

type PhotoService struct {
	photoService repositories.PhotoRepo
}

func NewPhotoService(repo *repositories.PhotoRepo) *PhotoService {
	return &PhotoService{photoService: *repo}
}

func (p *PhotoService) CreatePhoto(req *params.CreatePhoto, id int) (*params.PhotoResponse, *params.Response) {
	photo := models.Photo{
		Title:    req.Title,
		PhotoUrl: req.PhotoUrl,
		Caption:  req.Caption,
		UserID:   uint(id),
	}

	err := p.photoService.CreatePhoto(&photo)
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return &params.PhotoResponse{
		ID:        int(photo.ID),
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    int(photo.UserID),
		CreatedAt: &photo.CreatedAt,
	}, nil
}

func (p *PhotoService) GetPhotos() (*[]params.PhotoResponse, *params.Response) {
	photos, err := p.photoService.GetPhotos()
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	var list []params.PhotoResponse

	for _, v := range *photos {
		list = append(list, params.PhotoResponse{
			ID:        int(v.ID),
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserID:    int(v.UserID),
			CreatedAt: &v.CreatedAt,
			UpdatedAt: &v.UpdatedAt,
			User: &params.UserPhoto{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return &list, nil
}

func (p *PhotoService) UpdatePhoto(req *params.CreatePhoto, id int) (*params.PhotoResponse, *params.Response) {
	photo := models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}

	res, err := p.photoService.UpdatePhoto(&photo, id)
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return &params.PhotoResponse{
		ID:        int(res.ID),
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UserID:    int(res.UserID),
		UpdatedAt: &res.UpdatedAt,
	}, nil
}

func (p *PhotoService) GetPhotoById(id int) (*models.Photo, *params.Response) {
	photo, err := p.photoService.FindById(id)
	if err != nil {
		errResp := params.Response{
			Status:         http.StatusNotFound,
			Message:        "NotFound",
			AdditionalInfo: err.Error(),
		}
		return nil, &errResp
	}

	return photo, nil
}

func (p *PhotoService) DeletePhoto(id int) *params.Response {
	err := p.photoService.DeletePhoto(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Message:        "Bad Request",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Your photo has been successfully deleted",
	}
}

package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepo interface {
	CreatePhoto(photo *models.Photo) error
	GetPhotos() (*[]models.Photo, error)
	UpdatePhoto(payload *models.Photo, id int) (*models.Photo, error)
	FindById(id int) (*models.Photo, error)
	DeletePhoto(id int) error
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &photoRepo{db}
}

func (p *photoRepo) CreatePhoto(photo *models.Photo) error {
	return p.db.Create(photo).Error
}

func (p *photoRepo) GetPhotos() (*[]models.Photo, error) {
	var photos []models.Photo
	err := p.db.Preload(clause.Associations).Find(&photos).Error
	return &photos, err
}

func (p *photoRepo) UpdatePhoto(payload *models.Photo, id int) (*models.Photo, error) {
	var photo models.Photo

	err := p.db.Model(&photo).Table("photos").Clauses(clause.Returning{}).Where("id = ?", id).Updates(models.Photo{Title: payload.Title, Caption: payload.Caption, PhotoUrl: payload.PhotoUrl}).Order("created_at ASC").Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (p *photoRepo) FindById(id int) (*models.Photo, error) {
	var photo models.Photo
	err := p.db.First(&photo).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (p *photoRepo) DeletePhoto(id int) error {
	var photo models.Photo
	err := p.db.Where("id = ?", id).First(&photo).Error
	if err == nil {
		err = p.db.Where("id = ?", id).Delete(&photo).Error
	}
	return err
}

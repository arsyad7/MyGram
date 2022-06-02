package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepo interface {
	CreateSocialMedia(socmed *models.SocialMedia) error
	GetSocialMedias() (*[]models.SocialMedia, error)
	UpdateSocialMedia(p *models.SocialMedia, id int) (*models.SocialMedia, error)
	DeleteSocialMedia(id int) error
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) SocialMediaRepo {
	return &socialMediaRepo{db}
}

func (s *socialMediaRepo) CreateSocialMedia(socmed *models.SocialMedia) error {
	return s.db.Create(socmed).Error
}

func (s *socialMediaRepo) GetSocialMedias() (*[]models.SocialMedia, error) {
	var socialmedias []models.SocialMedia
	err := s.db.Preload(clause.Associations).Find(&socialmedias).Error
	return &socialmedias, err
}

func (s *socialMediaRepo) UpdateSocialMedia(p *models.SocialMedia, id int) (*models.SocialMedia, error) {
	var socialmedia models.SocialMedia

	err := s.db.Model(&socialmedia).Clauses(clause.Returning{}).Where("id = ?", id).Updates(models.SocialMedia{Name: p.Name, SocialMediaUrl: p.SocialMediaUrl}).Order("created_at ASC").Error
	if err != nil {
		return nil, err
	}
	return &socialmedia, nil
}

func (s *socialMediaRepo) DeleteSocialMedia(id int) error {
	var socmed models.SocialMedia
	err := s.db.Where("id = ?", id).First(&socmed).Error
	if err == nil {
		err = s.db.Where("id = ?", id).Delete(&socmed).Error
	}
	return err
}

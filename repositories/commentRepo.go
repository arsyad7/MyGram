package repositories

import (
	"mygram/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepo interface {
	CreateComment(comment *models.Comment) error
	GetComments() (*[]models.Comment, error)
	UpdateComment(p *models.Comment, id int) (*models.Comment, error)
	DeleteComment(id int) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db}
}

func (c *commentRepo) CreateComment(comment *models.Comment) error {
	return c.db.Create(comment).Error
}

func (c *commentRepo) GetComments() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Preload(clause.Associations).Find(&comments).Error
	return &comments, err
}

func (c *commentRepo) UpdateComment(p *models.Comment, id int) (*models.Comment, error) {
	var comment models.Comment

	err := c.db.Model(&comment).Table("comments").Clauses(clause.Returning{}).Where("id = ?", id).Updates(models.Comment{Message: p.Message}).Order("created_at ASC").Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentRepo) DeleteComment(id int) error {
	var comment models.Comment
	err := c.db.Where("id = ?", id).First(&comment).Error
	if err == nil {
		err = c.db.Where("id = ?", id).Delete(&comment).Error
	}
	return err
}

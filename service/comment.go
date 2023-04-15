package service

import (
	"errors"
	"finalproject_mygram/models"
	"finalproject_mygram/repository"
)

type CService interface {
	GetOneComment(id uint) (*models.Comment, error)
	GetAllComment() (*[]models.Comment, error)
}

type CommentService struct {
	Repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *CommentService {
	return &CommentService{repository}
}

func (service *CommentService) GetOneComment(id uint) (*models.Comment, error) {
	comment := service.Repository.FindById(id)
	if comment == nil {
		return comment, errors.New("comment not found")
	}
	return comment, nil
}

func (service *CommentService) GetAllComment() (*[]models.Comment, error) {
	comment := service.Repository.FindAll()
	if comment == nil || len(*comment)==0{
		return comment, errors.New("data comment not available")
	}
	return comment, nil
}

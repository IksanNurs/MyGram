package service

import (
	"finalproject_mygram/models"
	"finalproject_mygram/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var commentRepository = &repository.CommentRepositoryMock{Mock: mock.Mock{}}
var commentService = CommentService{Repository: commentRepository}

func TestCommentServiceGetOneCommentNotFound(t *testing.T) {
	commentRepository.Mock.On("FindById", uint(1)).Return(nil)

	comment, err := commentService.GetOneComment(uint(1))

	assert.Nil(t, comment)
	assert.NotNil(t, err)
	assert.Equal(t, "comment not found", err.Error(), "Error response has to be 'comment not found'")
}

func TestCommentServiceGetOneComment(t *testing.T) {
	comment := models.Comment{
		GormModel: models.GormModel{
			ID: 2,
		},
		UserID: 1,
	}

	commentRepository.Mock.On("FindById", uint(2)).Return(comment)

	result, err := commentService.GetOneComment(uint(2))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, comment.GormModel.ID, result.GormModel.ID, "Result has to be '2'")
	//assert.Equal(t, comment.Title, result.Title, "Result has to be 'Kaca Mata'")
	assert.Equal(t, &comment, result, "Result has to be comment data with id '2'")
}

func TestCommentServiceGetAllCommentNotFound(t *testing.T) {
	commentRepository.Mock.On("FindAll").Return(nil)

	comment, err := commentService.GetAllComment()

	assert.Nil(t, comment)
	assert.NotNil(t, err)
	assert.Equal(t, "comment not found", err.Error(), "Error response has to be 'comment not found'")
}

func TestCommentServiceGetAllComment(t *testing.T) {
	comment := []models.Comment{
		{
			GormModel: models.GormModel{
				ID: 1,
			},
			UserID: 2,
		},

		{
			GormModel: models.GormModel{
				ID: 2,
			},
			UserID: 2,
		},
	}
	commentRepository.Mock.On("FindAll").Return(comment)
	result, err := commentService.GetAllComment()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(comment), len(*result), "Result lenght")
	assert.Equal(t, comment, *result, "Result value")
}

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
		InputComment: models.InputComment{
			Message: "final project dts",
			PhotoID: 12,
		},
	}

	commentRepository.Mock.On("FindById", uint(2)).Return(comment)

	result, err := commentService.GetOneComment(uint(2))

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, comment.GormModel.ID, result.GormModel.ID, "Result has to be '2'")
	assert.Equal(t, comment.InputComment.Message, result.InputComment.Message, "Result has to be 'final project dts'")
	assert.Equal(t, comment.InputComment.PhotoID, result.InputComment.PhotoID, "Result has to be 12")
	assert.Equal(t, &comment, result, "Result has to be comment data with id '2'")
}

func TestCommentServiceGetAllCommentNotAvailable(t *testing.T) {
	commentRepository.Mock.On("FindAll").Return(nil)

	comment, err := commentService.GetAllComment()

	assert.Nil(t, comment)
	assert.NotNil(t, err)
	assert.Equal(t, "data comment not available", err.Error(), "Error response has to be 'data comment not available'")
}

func TestCommentServiceGetAllComment(t *testing.T) {
	comment := []models.Comment{
		{
			GormModel: models.GormModel{
				ID: 1,
			},
			UserID: 2,
			InputComment: models.InputComment{
				Message: "final project dts",
				PhotoID: 12,
			},
		},

		{
			GormModel: models.GormModel{
				ID: 2,
			},
			UserID: 2,
			InputComment: models.InputComment{
				Message: "showcase",
				PhotoID: 123,
			},
		},
	}
	commentRepository.Mock.On("FindAll").Return(comment)
	result, err := commentService.GetAllComment()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(comment), len(*result), "Result lenght")
	assert.Equal(t, comment, *result, "Result value")
}

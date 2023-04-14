package controllers

import (
	"finalproject_mygram/database"
	"finalproject_mygram/helpers"
	"finalproject_mygram/models"
	"finalproject_mygram/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	Service service.CService
}

func NewCommentController(Service service.CService) *CommentController {
	return &CommentController{Service}
}

// CreateComment godoc
// @Summary Post details for a given Id
// @Description Post details of comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param models.InputComment body models.InputComment true "create comment"
// @Success 200 {object} models.Comment
// @Router /comments [post]
// @Security Bearer 
func (pc *CommentController) CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	inputComment := models.InputComment{}
	userID := uint(userData["id"].(float64))
	User := models.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&inputComment)
	} else {
		c.ShouldBind(&inputComment)
	}
	Comment := models.Comment{
		InputComment: inputComment,
	}
	errA := db.First(&User, "id=?", userID).Error
	if errA != nil {
		response := helpers.APIResponse(errA.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	Comment.UserID = userID
	Comment.User = &User
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil menambahkan data comment", http.StatusOK, Comment)
	c.JSON(http.StatusCreated, response)
}

// UpdateComment godoc
// @Summary Update comment identified by the given Id
// @Description Update the comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment to be updated"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [put]
// @Security Bearer 
func (pc *CommentController) UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	inputComment := models.InputComment{}
	User := models.User{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&inputComment)
	} else {
		c.ShouldBind(&inputComment)
	}

	errA := db.First(&User, "id=?", userID).Error
	if errA != nil {
		response := helpers.APIResponse(errA.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	Comment := models.Comment{}
	Comment.InputComment=inputComment
	Comment.UserID = userID
	Comment.User = &User
	Comment.ID = uint(commentId)
	err := db.Model(&Comment).Where("id=?", commentId).Updates(Comment).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil update data comment "+c.Param("commentId"), http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

// GetComment godoc
// @Summary Get details
// @Description Get details of all comment
// @Tags comments
// @Accept json
// @Success 200 {object} models.Comment
// @Router /comments [get]
// @Security Bearer 
func (pc *CommentController) GetAllComment(ctx *gin.Context) {
	comment, err := pc.Service.GetAllComment()
	if err != nil {
		response := helpers.APIResponse(fmt.Sprint("data comment not available:", err), http.StatusBadRequest, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil menampilkan semua data comment", http.StatusOK, comment)
	ctx.JSON(http.StatusOK, response)
}

// GetCommentId godoc
// @Summary Get details for a given id
// @Description Get details of comment corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param  commentId path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [get]
// @Security Bearer 
func (pc *CommentController) GetOneComment(ctx *gin.Context) {
	commentID, _ := strconv.Atoi(ctx.Param("commentId"))
	comment, err := pc.Service.GetOneComment(uint(commentID))
	if err != nil {
		response := helpers.APIResponse(fmt.Sprint("Error finding comment:", err), http.StatusUnprocessableEntity, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("berhasil menampilkan data comment dengan id "+ctx.Param("commentId"), http.StatusOK, comment)
	ctx.JSON(http.StatusOK, response)

}

// DeleteComment godoc
// @Summary Delete comment identified by the given Id
// @Description Delete the book corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "ID of the comment to be deleted"
// @Success 200 "Ok"
// @Router /comments/{commentId} [delete]
// @Security Bearer 
func (pc *CommentController) DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))
	Comment.UserID = userID
	Comment.ID = uint(commentId)
	err := db.Where("id= ?", Comment.ID).Delete(&Comment).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helpers.APIResponse("berhasil menghapus data id "+c.Param("commentId"), http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

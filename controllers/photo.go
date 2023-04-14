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

type PhotoController struct {
	Service service.PService
}

func NewPhotoController(Service service.PService) *PhotoController {
	return &PhotoController{Service}
}

// CreatePhoto godoc
// @Summary Post details for a given Id
// @Description Post details of photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param models.InputPhoto body models.InputPhoto true "create photo"
// @Success 200 {object} models.Photo
// @Router /photos [post]
// @Security Bearer 
func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	inputPhoto := models.InputPhoto{}
	userID := uint(userData["id"].(float64))
	User := models.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&inputPhoto)
	} else {
		c.ShouldBind(&inputPhoto)
	}
	Photo := models.Photo{
		InputPhoto: inputPhoto,
	}
	errA := db.First(&User, "id=?", userID).Error
	if errA != nil {
		response := helpers.APIResponse(errA.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	Photo.UserID = userID
	Photo.User = &User
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil menambahkan data foto", http.StatusOK, Photo)
	c.JSON(http.StatusCreated, response)
}

// UpdatePhoto godoc
// @Summary Update photo identified by the given Id
// @Description Update the photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo to be updated"
// @Param models.InputPhoto body models.InputPhoto true "update photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [put]
// @Security Bearer 
func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	inputPhoto := models.InputPhoto{}
	User := models.User{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&inputPhoto)
	} else {
		c.ShouldBind(&inputPhoto)
	}

	errA := db.First(&User, "id=?", userID).Error
	if errA != nil {
		response := helpers.APIResponse(errA.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	Photo := models.Photo{}
	Photo.InputPhoto=inputPhoto
	Photo.UserID = userID
	Photo.User = &User
	Photo.ID = uint(photoId)
	err := db.Model(&Photo).Where("id=?", photoId).Updates(Photo).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil update data foto "+c.Param("photoId"), http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

// GetPhoto godoc
// @Summary Get details
// @Description Get details of all photo
// @Tags photos
// @Accept json
// @Success 200 {object} models.Photo
// @Router /photos [get]
// @Security Bearer 
func (pc *PhotoController) GetAllPhoto(ctx *gin.Context) {
	photo, err := pc.Service.GetAllPhoto()
	if err != nil {
		response := helpers.APIResponse(fmt.Sprint("data photo not available:", err), http.StatusBadRequest, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)

		return
	}
	response := helpers.APIResponse("berhasil menampilkan semua data foto", http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)
}

// GetPhotoId godoc
// @Summary Get details for a given id
// @Description Get details of photo corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param  photoId path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [get]
// @Security Bearer 
func (pc *PhotoController) GetOnePhoto(ctx *gin.Context) {
	photoID, _ := strconv.Atoi(ctx.Param("photoId"))
	photo, err := pc.Service.GetOnePhoto(uint(photoID))
	if err != nil {
		response := helpers.APIResponse(fmt.Sprint("Error finding photo:", err), http.StatusUnprocessableEntity, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helpers.APIResponse("berhasil menampilkan data foto dengan id "+ctx.Param("photoId"), http.StatusOK, photo)
	ctx.JSON(http.StatusOK, response)

}

// DeletePhoto godoc
// @Summary Delete photo identified by the given Id
// @Description Delete the book corresponding to the input Id
// @Tags photos
// @Accept json
// @Produce json
// @Param photoId path int true "ID of the photo to be deleted"
// @Success 200 "Ok"
// @Router /photos/{photoId} [delete]
// @Security Bearer 
func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))
	Photo.UserID = userID
	Photo.ID = uint(photoId)
	err := db.Where("id= ?", Photo.ID).Delete(&Photo).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helpers.APIResponse("berhasil menghapus data id "+c.Param("photoId"), http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

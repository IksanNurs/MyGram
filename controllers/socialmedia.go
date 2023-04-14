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

type SocialMediaController struct {
	Service service.SService
}

func NewSocialMediaController(Service service.SService) *SocialMediaController {
	return &SocialMediaController{Service}
}

// CreateSocialMedia godoc
// @Summary Post details for a given Id
// @Description Post details of socialmedia corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param models.InputSocialMedia body models.InputSocialMedia true "create socialmedia"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias [post]
// @Security Bearer 
func (pc *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	inputSocialMedia := models.InputSocialMedia{}
	userID := uint(userData["id"].(float64))
	User := models.User{}
	if contentType == appJSON {
		c.ShouldBindJSON(&inputSocialMedia)
	} else {
		c.ShouldBind(&inputSocialMedia)
	}
	SocialMedia := models.SocialMedia{
		InputSocialMedia: inputSocialMedia,
	}
	errA := db.First(&User, "id=?", userID).Error
	if errA != nil {
		response := helpers.APIResponse(errA.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	SocialMedia.UserID = userID
	SocialMedia.User = &User
	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil menambahkan data socialmedia", http.StatusOK, SocialMedia)
	c.JSON(http.StatusCreated, response)
}

// UpdateSocialMedia godoc
// @Summary Update socialmedia identified by the given Id
// @Description Update the socialmedia corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param socialmediaId path int true "ID of the socialmedia to be updated"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias/{socialmediaId} [put]
// @Security Bearer 
func (pc *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	inputSocialMedia := models.InputSocialMedia{}
	User := models.User{}
	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		c.ShouldBindJSON(&inputSocialMedia)
	} else {
		c.ShouldBind(&inputSocialMedia)
	}

	errA := db.First(&User, "id=?", userID).Error
	if errA != nil {
		response := helpers.APIResponse(errA.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	SocialMedia := models.SocialMedia{}
	SocialMedia.InputSocialMedia=inputSocialMedia
	SocialMedia.UserID = userID
	SocialMedia.User = &User
	SocialMedia.ID = uint(socialmediaId)
	err := db.Model(&SocialMedia).Where("id=?", socialmediaId).Updates(SocialMedia).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("berhasil update data socialmedia "+c.Param("socialmediaId"), http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

// GetSocialMedia godoc
// @Summary Get details
// @Description Get details of all socialmedia
// @Tags socialmedias
// @Accept json
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias [get]
// @Security Bearer 
func (pc *SocialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	socialmedia, err := pc.Service.GetAllSocialMedia()
	if err != nil {
		response := helpers.APIResponse(fmt.Sprint("data socialmedia not available:", err), http.StatusBadRequest, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)

		return
	}
	response := helpers.APIResponse("berhasil menampilkan semua data socialmedia", http.StatusOK, socialmedia)
	ctx.JSON(http.StatusOK, response)
}

// GetSocialMediaId godoc
// @Summary Get details for a given id
// @Description Get details of socialmedia corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param  socialmediaId path int true "ID of the socialmedia"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias/{socialmediaId} [get]
// @Security Bearer 
func (pc *SocialMediaController) GetOneSocialMedia(ctx *gin.Context) {
	socialmediaID, _ := strconv.Atoi(ctx.Param("socialmediaId"))
	socialmedia, err := pc.Service.GetOneSocialMedia(uint(socialmediaID))
	if err != nil {
		response := helpers.APIResponse(fmt.Sprint("Error finding socialmedia:", err), http.StatusUnprocessableEntity, nil)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("berhasil menampilkan data socialmedia dengan id "+ctx.Param("socialmediaId"), http.StatusOK, socialmedia)
	ctx.JSON(http.StatusOK, response)

}

// DeleteSocialMedia godoc
// @Summary Delete socialmedia identified by the given Id
// @Description Delete the book corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param socialmediaId path int true "ID of the socialmedia to be deleted"
// @Success 200 "Ok"
// @Router /socialmedias/{socialmediaId} [delete]
// @Security Bearer 
func (pc *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}
	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	userID := uint(userData["id"].(float64))
	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialmediaId)
	err := db.Where("id= ?", SocialMedia.ID).Delete(&SocialMedia).Error

	if err != nil {
	response := helpers.APIResponse(err.Error(), http.StatusNotFound, nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := helpers.APIResponse("berhasil menghapus data id "+c.Param("socialmediaId"), http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

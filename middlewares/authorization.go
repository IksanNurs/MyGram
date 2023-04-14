package middlewares

import (
	"finalproject_mygram/database"
	"finalproject_mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const COMMENT = 0
const PHOTO = 1
const SOCIALMEDIA = 2

func Authorization(groupRequest *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		if *groupRequest == COMMENT {
			Comment := models.Comment{}
			productId, err := strconv.Atoi(c.Param("commentId"))

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request",
					"message": "invalid paramter",
				})
				return
			}
			err = db.Select("user_id").First(&Comment, uint(productId)).Error

			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doest exist",
				})
				return
			}
			if Comment.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}
		if *groupRequest == PHOTO {
			Photo := models.Photo{}
			photoId, err := strconv.Atoi(c.Param("photoId"))

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request",
					"message": "invalid paramter",
				})
				return
			}
			err = db.Select("user_id").First(&Photo, uint(photoId)).Error

			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doest exist",
				})
				return
			}
			if Photo.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}
		if *groupRequest == SOCIALMEDIA {
			SocialMedia := models.SocialMedia{}
			productId, err := strconv.Atoi(c.Param("socialmediaId"))

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "bad request",
					"message": "invalid paramter",
				})
				return
			}
			err = db.Select("user_id").First(&SocialMedia, uint(productId)).Error

			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data Not Found",
					"message": "data doest exist",
				})
				return
			}
			if SocialMedia.UserID != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}
		c.Next()
	}
}

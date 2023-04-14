package controllers

import (
	"finalproject_mygram/database"
	"finalproject_mygram/helpers"
	"finalproject_mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// CreateUser godoc
// @Summary Post details for a given Id
// @Description Post details of user corresponding to the input Id
// @Tags users
// @Accept json
// @Produce json
// @Param models.InputUser body models.InputUser true "create user"
// @Success 200 {object} models.User
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	if contentType != appJSON {
		response := helpers.APIResponse("content type must be application/json", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUser models.InputUser
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user := models.User{
		Email:    inputUser.Email,
		UserName: inputUser.UserName,
		Password: inputUser.Password,
		Age:      inputUser.Age,
	}

	if err := db.Debug().Create(&user).Error; err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("berhasil menambahkan data user", http.StatusOK, user)
	c.JSON(http.StatusCreated, response)
}

// CreateUser godoc
// @Summary Post details for a given Id
// @Description Post details of user corresponding to the input Id
// @Tags users
// @Accept json
// @Produce json
// @Param models.LoginUser body models.LoginUser true "create user"
// @Success 200 {object} models.User
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	User1 := models.LoginUser{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User1)
	}
	password = User1.Password

	err := db.Debug().Where("email=?", User1.Email).Take(&User).Error

	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusUnauthorized, nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		response := helpers.APIResponse("unauthoried2", http.StatusUnauthorized, nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

package user_routes

import (
	"encoding/json"
	"net/http"

	internal_security "github.com/Mateus-MS/Xeubiart.git/backend/internal/security"
	user_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/model"
	user_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/user/service"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name     string `json:"name"     bson:"Name"`
	Password string `json:"password" bson:"Password"`
	Email    string `json:"email"    bson:"Email"`
	Phone    string `json:"phone"    bson:"Phone"`
}

func UserRegisterRoute(userService user_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestBody

		// Bind the JSON from the body into RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 1 - Create the pass hash
		passHash, err := internal_security.HashPassword(reqBody.Password)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 2 - Create the Email
		email, err := user_model.NewEmail(reqBody.Email)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 3 - If not empty, create the Phone
		var phone user_model.Phone
		if reqBody.Phone != "" {
			phone, err = user_model.NewPhone(reqBody.Phone)
		}

		// 4 - Create the UserEntity
		userEntity, err := user_model.NewUserEntity(reqBody.Name, passHash, email, &phone)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 5 - Register the user
		err = userService.Register(c.Request.Context(), userEntity)
		if err != nil {
			if err == user_service.ErrEmailTaken {
				c.AbortWithError(http.StatusConflict, err)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		data, err := json.Marshal(reqBody)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.String(http.StatusOK, string(data))
	}
}

package appointment_routes

import (
	"errors"
	"net/http"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	appointment_model "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/model"
	appointment_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/appointment/service"
	routes_utils "github.com/Mateus-MS/Xeubiart.git/backend/utils/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestBody struct {
	UserID primitive.ObjectID `json:"userID" bson:"UserID"`
	Date   string             `json:"date"   bson:"Date"`
}

func AppointmentRegisterRoute(appointmentService appointment_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestBody

		// Bind the JSON from the body into RequestBody
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Load the timezone from cookie
		loc, err := routes_utils.LoadLocationFromCookie(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Convert the received string into a LocalTime
		lt, err := internal_datetime.NewLocalFromString(reqBody.Date, loc)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Create the appointment object
		appointment, err := appointment_model.NewEntity(reqBody.UserID, lt)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Try to register the new appointment
		err = appointmentService.Create(c.Request.Context(), appointment)
		if err != nil {
			if errors.Is(err, appointment_service.ErrInvalidAppointmentDate) {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.String(http.StatusOK, appointment.Date.String())
	}
}

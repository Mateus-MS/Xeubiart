package schedule_routes

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	internal_datetime "github.com/Mateus-MS/Xeubiart.git/backend/internal/datetime"
	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
	routes_utils "github.com/Mateus-MS/Xeubiart.git/backend/utils/routes"
	"github.com/gin-gonic/gin"
)

func ScheduleReadMonthlyRoute(scheduleService schedule_service.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the Year param
		yearStr := c.Param("year")
		if yearStr == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid year"))
			return
		}
		yearInt, err := strconv.Atoi(yearStr)

		// Get the path param as a time.Month
		monthStr := c.Param("month")
		monthInt, err := strconv.Atoi(monthStr)
		if err != nil || monthInt < 1 || monthInt > 12 {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid month"))
			return
		}
		month := time.Month(monthInt)

		// Get the timezone cookie
		loc, err := routes_utils.LoadLocationFromCookie(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Create the local time with received data
		localTime, err := internal_datetime.NewLocalFromTime(time.Date(yearInt, month, 1, 0, 0, 0, 0, loc))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Try to read from DB converting the local time to UTC
		schedule, err := scheduleService.ReadByMonth(c.Request.Context(), localTime.ToUTCTime())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, schedule)
	}
}

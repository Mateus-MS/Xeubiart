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
		// Get the path param as a time.Month
		monthOffsetStr := c.Param("month")
		monthOffsetInt, err := strconv.Atoi(monthOffsetStr)
		if err != nil || monthOffsetInt < 0 || monthOffsetInt > 12 {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid month"))
			return
		}

		// Get the timezone cookie
		loc, err := routes_utils.LoadLocationFromCookie(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Create the local time with received data
		localTime, err := internal_datetime.NewLocalFromTime(time.Now().In(loc).AddDate(0, monthOffsetInt, 0))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Converts localtime to UTC
		utcTime, err := internal_datetime.NewUTCTimeFromTime(localTime.Time.UTC())
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// Try to read from DB
		schedule, err := scheduleService.ReadByOffsetMonth(c.Request.Context(), int(utcTime.Month()))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, schedule)
	}
}

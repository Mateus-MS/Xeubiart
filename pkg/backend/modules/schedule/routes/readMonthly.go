package schedule_routes

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	schedule_service "github.com/Mateus-MS/Xeubiart.git/backend/modules/schedule/service"
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
		// loc, err := routes_utils.LoadLocationFromCookie(c)
		// if err != nil {
		// 	c.AbortWithError(http.StatusBadRequest, err)
		// 	return
		// }

		schedule, err := scheduleService.ReadByMonth(c.Request.Context(), yearInt, month)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, schedule)
	}
}

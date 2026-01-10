package routes_utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoadLocationFromCookie(c *gin.Context) (*time.Location, error) {
	loc := &time.Location{}

	tz, err := c.Cookie("timezone")
	if err != nil {
		return loc, fmt.Errorf("timezone cookie not provided: %w", err)
	}

	loc, err = time.LoadLocation(tz)
	if err != nil {
		return loc, fmt.Errorf("the given timezone is invalid: %w", err)
	}

	return loc, nil
}

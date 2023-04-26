package gapi

import (
	"net/http"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (g *gapi) postJobs(c echo.Context) error {

	var def bo.TaskDefinition
	if err := c.Bind(&def); err != nil {
		log.Errorf("extract task %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"err": err.Error(),
		})
	}
}

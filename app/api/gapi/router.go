package gapi

import (
	"github.com/labstack/echo/v4"
)

func (g *gapi) addRoutes() {
	grp := g.e.Group("/api")

	g.postProjectRouter(grp)
	g.postScriptsRouter(grp)
	g.postRunRouter(grp)

}

func (g *gapi) postProjectRouter(grp *echo.Group) {
	grp.POST("/project", g.postProject)
}

func (g *gapi) postScriptsRouter(grp *echo.Group) {
	grp.POST("/scripts", g.postScripts)
}

func (g *gapi) postRunRouter(grp *echo.Group) {
	grp.POST("/run", g.postRun)
}

package gapi

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (g *gapi) addRoutes() {
	grp := g.e.Group("/api")
	// grp.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{
	// 		echo.HeaderAccessControlAllowMethods,
	// 		echo.HeaderAccessControlAllowOrigin,
	// 		echo.HeaderContentType,
	// 		echo.HeaderAccept,
	// 		echo.HeaderOrigin,
	// 	},
	// }))
	grp.Use(middleware.CORS())

	g.postProjectRouter(grp)
	g.postScriptsRouter(grp)
	g.postRunRouter(grp)
	grp.POST("/test", func(c echo.Context) error {
		time.Sleep(time.Second * 60)
		return c.JSON(http.StatusOK, "hello")
	})

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

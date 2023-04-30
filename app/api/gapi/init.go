package gapi

import (
	"github.com/gabrieldiaziv/wghidra/app/bo/defs"
	"github.com/labstack/echo/v4"
)

type gapi struct {
	wghidra defs.WGhidra
	e       *echo.Echo
	port    string
}

func NewGAPI(port string, wghidra defs.WGhidra) *gapi {
	e := echo.New()
	return &gapi{
		wghidra: wghidra,
		e:       e,
		port:    port,
	}
}

func (g *gapi) Start() {
	g.addRoutes()
	g.e.Logger.Fatal(g.e.Start(g.port))
}

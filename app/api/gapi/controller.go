package gapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gabrieldiaziv/wghidra/app/repo/mock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func skip(c echo.Context) error {
	var p project_out
	if err := json.Unmarshal([]byte(mock.Project), &p); err != nil {
		log.Errorf("could marshal skip: %v", err)
		return gError(c, "could not skip", http.StatusInternalServerError)
	}

	return gSuccess(c, p)
}

func (g *gapi) postProject(c echo.Context) error {

	var p_in project_in
	if err := c.Bind(&p_in); err != nil {
		return gError(c, err.Error(), http.StatusBadRequest)
	}

	if p_in.Skip != 0 {
		return skip(c)
	}

	exe, err := c.FormFile("project")
	if err != nil {
		log.Errorf("missing project file: %v", err)
		return gError(c, "missing project file", http.StatusBadRequest)
	}

	project, err := exe.Open()
	if err != nil {
		log.Errorf("cannot open file: %v", err)
		return gError(c, "cannot open file", http.StatusBadRequest)
	}

	// c.Request().Context()
	projectId, functions, asm, err :=
		g.wghidra.ParseProject(context.Background(), project)

	if err != nil {
		log.Errorf("could not parse: %v", err)
		return gError(c, "cannot parse file", http.StatusBadRequest)
	}

	return gSuccess(c, project_out{
		Functions: functions,
		Assembly:  asm,
		ProjectID: projectId,
	})
}

func (g *gapi) postScripts(c echo.Context) error {

	var jobs scripts_in
	if err := c.Bind(&jobs); err != nil {
		log.Errorf("bad input: %v", err)
		return gError(c, "bad input", http.StatusBadRequest)
	}

	tasksRes, err := g.wghidra.RunScripts(
		c.Request().Context(),
		jobs.ProjectID,
		Scripts2Def(jobs.ProjectID, jobs.Scripts),
	)

	if err != nil {
		log.Errorf("could not run scripts: %v", err)
		return gError(c, "could not run sciprts", http.StatusInternalServerError)
	}

	return gSuccess(c, scripts_out{
		Results: tasksRes,
	})
}

func (g *gapi) postRun(c echo.Context) error {
	var run run_in

	if err := c.Bind(&run); err != nil {
		log.Errorf("bad input: %v", err)
		return gError(c, "bad input", http.StatusBadRequest)
	}

	result, err := g.wghidra.PyRun(c.Request().Context(), run.ExecuteFunction, run.Parameters, run.Functions)
	if err != nil {
		log.Errorf("could not run : %v", err)
		return gError(c, "could not run", http.StatusBadRequest)
	}

	return gSuccess(c, scripts_out{
		Results: result,
	})
}

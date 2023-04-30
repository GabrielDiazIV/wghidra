package gapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (g *gapi) postProject(c echo.Context) error {

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

	defer project.Close()

	projectId, functions, err :=
		g.wghidra.ParseProject(c.Request().Context(), project)

	if err != nil {
		log.Errorf("could not parse: %v", err)
		return gError(c, "cannot open file", http.StatusBadRequest)
	}

	return gSuccess(c, project_out{
		Functions: functions,
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

	result, err := g.wghidra.PyRun(c.Request().Context(), run.ExecuteFunction, run.Functions)
	if err != nil {
		log.Errorf("could not run : %v", err)
		return gError(c, "could not run", http.StatusBadRequest)
	}

	return gSuccess(c, result)
}

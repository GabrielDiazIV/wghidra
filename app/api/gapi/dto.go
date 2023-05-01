package gapi

import (
	"net/http"
	"strings"

	"github.com/gabrieldiaziv/wghidra/app/bo"
	"github.com/labstack/echo/v4"
)

type gerror struct {
	Msg  string `json:"msg,omitempty"`
	Code int    `json:"code,omitempty"`
}

type gresponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error gerror      `json:"error,omitempty"`
}

type project_in struct {
}

type project_out struct {
	Functions []bo.Function `json:"functions,omitempty"`
	ProjectID string        `json:"project_id,omitempty"`
	Assembly  string        `json:"assembly,omitempty"`
}

type scripts_in struct {
	ProjectID string `json:"project_id,omitempty"`
	Scripts   []bo.ScriptTask
}
type scripts_out struct {
	Results []bo.TaskResult
}

type run_in struct {
	Functions       []bo.Function `json:"functions,omitempty"`
	ExecuteFunction string        `json:"execute_function,omitempty"`
	Parameters      []string      `json:"parameters,omitempty"`
}

func gSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, gresponse{
		Data: data,
	})
}

func gError(c echo.Context, msg string, code int) error {
	return c.JSON(code, gresponse{
		Error: gerror{msg, code},
	})
}

func gResponse(c echo.Context, data interface{}, msg string, code int) error {
	return c.JSON(code, gresponse{
		Data:  data,
		Error: gerror{msg, code},
	})
}

func Scripts2Def(projectID string, scripts []bo.ScriptTask) bo.TaskDefinition {

	uts := make([]bo.UnitTask, len(scripts))
	for i, script := range scripts {

		ut := bo.NewScriptTask(name)
		ut.Name = strings.Join([]string{projectID, script.ScriptName}, "-")
		ut.Task = script
		uts[i] = ut
	}

	return bo.TaskDefinition{Tasks: uts}
}

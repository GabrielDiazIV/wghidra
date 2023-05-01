package bo

import (
	"io"
)

type TaskDefinition struct {
	Version string     `yaml:"version,omitempty" json:"version,omitempty"`
	Exe     io.Reader  `yaml:"-"                 json:"-"`
	Tasks   []UnitTask `yaml:"tasks,omitempty"   json:"tasks,omitempty"`
}

type UnitTask struct {
	Name    string `yaml:"name,omitempty"    json:"name,omitempty"`
	Task    ScriptTask
	Cleanup bool `yaml:"cleanup,omitempty" json:"cleanup,omitempty"`
}

type ScriptTask struct {
	ScriptName string   `json:"scriptName,omitempty"`
	Parameters []string `json:"parameters,omitempty"`
}

type TaskResult struct {
	Name   string                 `yaml:"name,omitempty"    json:"name,omitempty"`
	Output map[string]interface{} `yaml:"output,omitempty"      json:"output,omitempty"`
	Error  *TaskError             `json:"error,omitempty"`
}

type TaskError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type Function struct {
	Name       string   `json:"name,omitempty"`
	Parameters []string `json:"parameters,omitempty"`
	Body       string   `json:"body,omitempty"`
}

func TaskFailed(ut UnitTask, code int, msg string) TaskResult {
	return TaskResult{
		Name: ut.Name,
		Error: &TaskError{
			Code: code,
			Msg:  msg,
		},
	}
}

type ImagePullStatus struct {
	Status         string `json:"status,omitempty"`
	Error          string `json:"error,omitempty"`
	Progress       string `json:"progress,omitempty"`
	ProgressDetail struct {
		Current int `json:"current,omitempty"`
		Total   int `json:"total,omitempty"`
	} `json:"progressDetail,omitempty"`
}

const (
	DecompileTaskName  = "decompiler"
	DissasmbleTaskName = "asmer"
	RunTaskName        = "runner"
)

const (
	PythonGhidraScript = "gpython"
	JavaGhidraScript   = "gjava"
	PyRunScript        = "pyrun"
)

func NewScriptTask(name string, task ScriptTask) UnitTask {
	return UnitTask{
		Name:    name,
		Cleanup: true,
		Task:    task,
	}
}

func NewDecompileTask() UnitTask {
	return UnitTask{
		Name:    DecompileTaskName,
		Cleanup: true,
		Task: ScriptTask{
			ScriptName: "extract.py",
		},
	}
}
func NewDissasemblyTask() UnitTask {
	return UnitTask{
		Name:    DissasmbleTaskName,
		Cleanup: true,
		Task: ScriptTask{
			ScriptName: "ExportAssembly.java",
		},
	}
}

func NewRunTask(paramters []string) UnitTask {
	return UnitTask{
		Name:    RunTaskName,
		Cleanup: true,
		Task: ScriptTask{
			ScriptName: PyRunScript,
			Parameters: paramters,
		},
	}
}

func (st ScriptTask) Cmd() []string {

	isTypePython := st.ScriptName[len(st.ScriptName)-2:] == "py"
	isTypeJava := st.ScriptName[len(st.ScriptName)-4:] == "java"

	var argv []string

	if isTypePython {
		argv = append(argv, PythonGhidraScript)
	} else if isTypeJava {
		argv = append(argv, JavaGhidraScript)
	}

	argv = append(argv, st.ScriptName)
	argv = append(argv, st.Parameters...)

	return argv

}

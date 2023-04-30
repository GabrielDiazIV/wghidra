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
	Cleanup bool      `yaml:"cleanup,omitempty" json:"cleanup,omitempty"`
	Exe     io.Reader `yaml:"-"                 json:"-"`
}

type ScriptTask struct {
	ScriptName string   `json:"scriptName,omitempty"`
	Parameters []string `json:"parameters,omitempty"`
}

type TaskResult struct {
	Name   string     `yaml:"name,omitempty"    json:"name,omitempty"`
	Output string     `yaml:"output,omitempty"      json:"output,omitempty"`
	Error  *TaskError `json:"error,omitempty"`
}

type TaskError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type Function struct {
	Name       string
	Parameters []string
	Body       string
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
	DecompileTaskName  = "Decompile"
	DissasmbleTaskName = "Dissasmble"
	RunTaskName        = "Run"
)

func NewScriptTask(name string, fstream io.Reader, task ScriptTask) UnitTask {
	return UnitTask{
		Name:    name,
		Cleanup: true,
		Exe:     fstream,
		Task:    task,
	}
}

func NewDecompileTask(fstream io.Reader) UnitTask {
	return UnitTask{
		Name:    DecompileTaskName,
		Cleanup: true,
		Exe:     fstream,
		Task: ScriptTask{
			ScriptName: "extract.py",
		},
	}
}
func NewDissasemblyTask(fstream io.Reader) UnitTask {
	return UnitTask{
		Name:    DissasmbleTaskName,
		Cleanup: true,
		Exe:     fstream,
		Task: ScriptTask{
			ScriptName: "dissasmbly",
		},
	}
}

func NewRunTask(fstream io.Reader, paramters []string) UnitTask {
	return UnitTask{
		Name:    RunTaskName,
		Cleanup: true,
		Exe:     fstream,
		Task: ScriptTask{
			ScriptName: "run.py",
			Parameters: paramters,
		},
	}
}

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
			ScriptName: "dissasmbly",
		},
	}
}

func NewRunTask(paramters []string) UnitTask {
	return UnitTask{
		Name:    RunTaskName,
		Cleanup: true,
		Task: ScriptTask{
			ScriptName: "run.py",
			Parameters: paramters,
		},
	}
}

func (st ScriptTask) Cmd() []string {

	argv := make([]string, len(st.Parameters)+1)
	argv[0] = st.ScriptName
	for i, arg := range st.Parameters {
		argv[i+1] = arg
	}

	return argv
}

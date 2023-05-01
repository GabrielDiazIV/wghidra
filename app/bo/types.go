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
	Output any        `yaml:"output,omitempty"      json:"output,omitempty"`
	Error  *TaskError `json:"error,omitempty"`
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
	DecompileTaskName  = "Decompile"
	DissasmbleTaskName = "Dissasmble"
	RunTaskName        = "Run"
)

const (
	PythonGhidraScript = "pghidra"
	JavaGhidraScript   = "jghidra"
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
			ScriptName: "dissasmbly",
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

	argv_len := 1
	offset := 0

	if isTypePython || isTypeJava {
		offset = 2
	}

	if st.Parameters != nil {
		argv_len += len(st.Parameters)
	}

	argv := make([]string, argv_len+offset)

	if isTypePython {
		argv[0] = PythonGhidraScript
		argv[1] = st.ScriptName
	} else if isTypeJava {
		argv[0] = JavaGhidraScript
		argv[1] = st.ScriptName
	} else {
		argv[0] = st.ScriptName
	}

	for i, arg := range st.Parameters {
		argv[i+offset] = arg
	}

	return argv

}

package bo

import (
	"io"
)

type TaskDefinition struct {
	Version string     `yaml:"version,omitempty" json:"version,omitempty"`
	Tasks   []UnitTask `yaml:"tasks,omitempty"   json:"tasks,omitempty"`
}

type UnitTask struct {
	Name    string        `yaml:"name,omitempty"    json:"name,omitempty"`
	ID      string        `yaml:"id,omitempty"      json:"id,omitempty"`
	Runner  string        `yaml:"runner,omitempty"  json:"runner,omitempty"`
	Command []string      `yaml:"command,omitempty" json:"command,omitempty"`
	Cleanup bool          `yaml:"cleanup,omitempty" json:"cleanup,omitempty"`
	Exe     io.ReadCloser `yaml:"-"                 json:"-"`
}

type TaskResult struct {
	Name      string        `yaml:"name,omitempty"    json:"name,omitempty"`
	ID        string        `yaml:"id,omitempty"      json:"id,omitempty"`
	Link      string        `yaml:"link,omitempty"      json:"link,omitempty"`
	TarStream io.ReadCloser `yaml:"-"      json:"-"`
	Error     *TaskError    `json:"error,omitempty"`
}

type TaskError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func TaskFailed(ut UnitTask, code int, msg string) TaskResult {
	return TaskResult{
		Name: ut.Name,
		ID:   ut.ID,
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

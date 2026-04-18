package taskig

import "encoding/json"

type (
	State   string
	JobHook string
	Hook    string

	StatePair struct {
		Start State `json:"start"`
		End   State `json:"end"`
	}

	JobType struct {
		Namespace string `json:"namespace"`
		Kind      string `json:"kind"`
	}

	JobDefinition struct {
		Type   *JobType     `json:"type"`
		States []*StatePair `json:"states"`
		Errors []State      `json:"errors"`
	}

	ParentDefinition struct {
		Ref  JobRef  `json:"ref"`            // the parent
		Hook JobHook `json:"hook"`           // before|after|error
		Prio int     `json:"prio,omitempty"` // prio of this hook compared to other similar hooks
	}

	// Job defines a standard job.
	Job struct {
		Type   *JobType          `json:"type"`
		Meta   *Meta             `json:"meta,omitempty"`
		Task   json.RawMessage   `json:"task"`
		Parent *ParentDefinition `json:"parent,omitempty"`
		State  State             `json:"state"`
	}

	// Meta contains data that could be meaninful to the executor.
	Meta struct {
		Name   string   `json:"name,omitempty"`
		Labels []string `json:"labels,omitempty"`
	}

	JobRef string

	Executor interface {
		Schedule(*Job) (JobRef, error)
		Update(JobRef, State, error) error
	}

	// SyncExecutorSupport allows to execute a task on the spot, requires executor with InstantSupport.
	SyncExecutorSupport interface {
		Executor
		Execute(*Job) (any, error)
	}

	ExecutorHookData struct {
		Ref     JobRef   `json:"ref"`
		JobType *JobType `json:"jobType"`
		Hook    Hook     `json:"hook"` // scheduled|updated|error|success
		Before  State    `json:"before,omitempty"`
		State   State    `json:"state"`
		Error   string   `json:"error,omitempty"`
	}

	ExecutorHook interface {
		Hook(*ExecutorHookData) ([]JobRef, error)
	}

	Worker interface {
		// Execute allows scheduler to notify executor of work.
		Execute(JobRef) error
	}

	// SyncWorkerSupport allows to execute a task on the spot.
	SyncWorkerSupport interface {
		Worker
		// Work, execute the task on the spot.
		Work(*Job) (any, error)
	}
)

var (
	BeforeJobHook = JobHook("before")
	JobAfterHook  = JobHook("after")
	JobErrorHook  = JobHook("error")

	ScheduledHook = Hook("scheduled")
	UpdatedHook   = Hook("updated")
	ErrorHook     = Hook("error")
	SuccessHook   = Hook("success")
)

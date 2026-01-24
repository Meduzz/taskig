package xcute

import "encoding/json"

type (
	State string

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

	// Job defines a standard job.
	Job struct {
		Type   *JobType        `json:"type"`
		Meta   *Meta           `json:"meta,omitempty"`
		Task   json.RawMessage `json:"task"`
		Parent *JobRef         `json:"parent,omitempty"`
		Start  State           `json:"state"`
	}

	// Meta contains data that could be meaninful to the executor.
	Meta struct {
		Name   string   `json:"name,omitempty"`
		Labels []string `json:"labels,omitempty"`
	}

	JobRef string

	Executor interface {
		Schedule(*Job) (JobRef, error)
		RegisterWorker(*JobDefinition, Worker)
	}

	// SyncExecutorSupport allows to execute a task on the spot, requires executor with InstantSupport.
	SyncExecutorSupport interface {
		Executor
		Execute(*Job) (any, error)
	}

	// A scheduler extension not tied to the scheduler...?
	JobApi interface {
		Load(JobRef) (*Job, error)
		Status(JobRef) (State, error)
		Update(JobRef, State) error
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

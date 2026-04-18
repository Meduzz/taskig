package taskig_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/Meduzz/taskig"
)

type (
	Task struct {
		Format  string `json:"format"`
		Message string `json:"message"`
	}
)

func TestModel(t *testing.T) {
	var executor Executor
	var Created, Pending, Error, Done State
	task := &Task{
		Format:  "Hello %s!",
		Message: "World",
	}

	definition := DefineJob(func(b JobDefinitionBuilder) {
		b.Type("test", "test")
		b.Transition(Created, Pending) // from created to pending
		b.Transition(Pending, Error)   // from pending to error
		b.Transition(Pending, Done)    // or from pending to done.
		b.Error(Error)                 // mark error as error state
	})

	job, err := CreateJob(func(b JobBuilder) error {
		b.Type("test", "test")
		b.Meta(func(m MetaBuilder) {
			m.Name("test")
		})
		b.StartState(Created)
		return b.Task(task)
	})

	if err != nil {
		t.Errorf("creating job threw error: %v", err)
	}

	t.Run("Fake schedule job", func(t *testing.T) {
		ref, err := executor.Schedule(job)

		if err != nil {
			t.Error(err)
		}

		if ref.Namespace() != definition.Type.Namespace || ref.Kind() != definition.Type.Kind {
			t.Error("namespace or kind did not match")
		}
	})

	t.Run("Fake executor running a job", func(t *testing.T) {
		var executor Executor
		var jobRef JobRef

		err = executor.Update(jobRef, Pending, nil)

		if err != nil {
			t.Error(err)
		}

		args := make(map[string]string)
		err = json.Unmarshal(job.Task, args)

		if err != nil {
			err2 := executor.Update(jobRef, Error, err)

			if err2 != nil {
				t.Error(err, err2)
			} else {
				t.Error(err)
			}
		}

		fmt.Printf(args["format"], args["message"])

		err = executor.Update(jobRef, Done, nil)

		if err != nil {
			t.Error(err)
		}
	})
}

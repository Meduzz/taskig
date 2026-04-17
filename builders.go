package taskig

import "encoding/json"

type (
	JobDefinitionBuilder interface {
		Type(namespace, kind string)
		Transition(from, to State)
		Error(State)
	}

	definitionBuilder struct {
		def *JobDefinition
	}

	JobBuilder interface {
		Type(namespace, kind string)
		Meta(func(MetaBuilder))
		StartState(State)
		Task(any) error
	}

	jobBuilder struct {
		job *Job
	}

	MetaBuilder interface {
		Name(string)
		Labels(...string)
	}

	metaBuilder struct {
		job *Job
	}
)

func newDefinitionBuiler(def *JobDefinition) JobDefinitionBuilder {
	return &definitionBuilder{def}
}

func (j *definitionBuilder) Type(namespace, kind string) {
	j.def.Type = &JobType{
		Namespace: namespace,
		Kind:      kind,
	}
}

func (j *definitionBuilder) Transition(from, to State) {
	j.def.States = append(j.def.States, &StatePair{
		Start: from,
		End:   to,
	})
}

func (j *definitionBuilder) Error(it State) {
	j.def.Errors = append(j.def.Errors, it)
}

func newJobBuilder(job *Job) JobBuilder {
	return &jobBuilder{job}
}

func (j *jobBuilder) Type(namespace, kind string) {
	j.job.Type = &JobType{
		Namespace: namespace,
		Kind:      kind,
	}
}

func (j *jobBuilder) Meta(cb func(MetaBuilder)) {
	builder := newMetaBuilder(j.job)
	cb(builder)
}

func (j *jobBuilder) StartState(start State) {
	j.job.State = start
}

func (j *jobBuilder) Task(task any) error {
	bs, err := json.Marshal(task)

	if err != nil {
		return err
	}

	j.job.Task = bs

	return nil
}

func newMetaBuilder(job *Job) MetaBuilder {
	if job.Meta == nil {
		job.Meta = &Meta{}
	}

	return &metaBuilder{job}
}

func (m *metaBuilder) Name(name string) {
	m.job.Meta.Name = name
}

func (m *metaBuilder) Labels(labels ...string) {
	m.job.Meta.Labels = append(m.job.Meta.Labels, labels...)
}

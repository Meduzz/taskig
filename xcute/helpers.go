package xcute

func DefineJob(cb func(JobDefinitionBuilder)) *JobDefinition {
	def := &JobDefinition{}

	builder := newDefinitionBuiler(def)
	cb(builder)

	return def
}

func CreateJob(cb func(JobBuilder) error) (*Job, error) {
	job := &Job{}

	builder := newJobBuilder(job)
	err := cb(builder)

	return job, err
}

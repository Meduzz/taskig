package xcute

import "fmt"

func (j *JobType) JobRef(id string) JobRef {
	return JobRef(fmt.Sprintf("%s/%s#%s", j.Namespace, j.Kind, id))
}

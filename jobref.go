package taskig

import "strings"

func (it JobRef) Namespace() string {
	idx1 := strings.LastIndex(string(it), "/")
	return string(it)[:idx1]
}

func (it JobRef) Kind() string {
	split1 := strings.Split(string(it), "#")

	idx2 := strings.LastIndex(split1[0], "/") + 1
	return string(split1[0][idx2:])
}

func (it JobRef) ID() string {
	split1 := strings.Split(string(it), "#")
	return split1[1]
}

func (it JobRef) JobType() *JobType {
	return &JobType{
		Namespace: it.Namespace(),
		Kind:      it.Kind(),
	}
}

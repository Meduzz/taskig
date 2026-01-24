package xcute_test

import (
	"fmt"
	"testing"

	. "github.com/Meduzz/taskig/xcute"
)

func TestJobRef(t *testing.T) {
	simple := JobRef("namespace/kind#id")
	complex := JobRef("namespace/v1/kind#id")

	const (
		namespace = "namespace"
		kind      = "kind"
		id        = "id"
	)

	t.Run("with simple namespace", func(t *testing.T) {
		if simple.Namespace() != namespace {
			t.Error("namespace does not match", simple.Namespace())
		}

		if simple.Kind() != kind {
			t.Error("kind does not match", simple.Kind())
		}

		if simple.ID() != id {
			t.Error("id does not match", simple.ID())
		}
	})

	t.Run("with complex namespace", func(t *testing.T) {
		complexNS := fmt.Sprintf("%s/v1", namespace)

		if complex.Namespace() != complexNS {
			t.Error("namespace does not match", complex.Namespace())
		}

		if complex.Kind() != kind {
			t.Error("kind does not match", complex.Kind())
		}

		if complex.ID() != id {
			t.Error("id does not match", complex.ID())
		}
	})
}

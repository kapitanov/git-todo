package idgen

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate_FirstItem(t *testing.T) {
	const title = "Foo Bar"

	id := iterFirst(Generate(title))
	t.Logf("Generate(%q): %v", title, id)

	assert.NotEmpty(t, id)
	assert.Len(t, id, IDLength)
}

func TestGenerate_NextItems(t *testing.T) {
	const title = "Foo Bar"
	const count = 10

	ids := iterTake(Generate(title), count)
	t.Logf("Generate(%q): %v", title, ids)

	assert.Len(t, ids, count)
	uniq := make(map[string]struct{})
	for _, id := range ids {
		assert.NotContains(t, uniq, id)
		uniq[id] = struct{}{}
	}
}

func iterTake[V any](it iter.Seq[V], n int) []V {
	var vs []V
	it(func(v V) bool {
		if n > 0 {
			n--
			vs = append(vs, v)
			return true
		}

		return false
	})
	return vs
}

func iterFirst[V any](it iter.Seq[V]) V {
	vs := iterTake(it, 1)
	if len(vs) == 0 {
		panic("no items")
	}

	return vs[0]
}

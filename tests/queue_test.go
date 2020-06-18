package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Workiva/go-datastructures/queue"
)

func TestPut(t *testing.T) {
	q := queue.New(10)

	q.Put(`test`)
	assert.Equal(t, int64(1), q.Len())

	results, err := q.Get(1)
	assert.Nil(t, err)

	result := results[0]
	assert.Equal(t, `test`, result)
	assert.True(t, q.Empty())

	q.Put(`test2`)
	assert.Equal(t, int64(1), q.Len())

	results, err = q.Get(1)
	assert.Nil(t, err)

	result = results[0]
	assert.Equal(t, `test2`, result)
	assert.True(t, q.Empty())
}

func TestGet(t *testing.T) {
	q := queue.New(10)

	q.Put(uint64(1))
	q.Put(uint64(2))
	q.Put(uint64(3))
	assert.Equal(t, int64(3), q.Len())
	results, err := q.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, uint64(1), results[0])
	assert.Equal(t, int64(2), q.Len())
	results, err = q.Get(1)
	assert.Equal(t, uint64(2), results[0])
	results, err = q.Get(1)
	assert.Equal(t, uint64(3), results[0])
	assert.Equal(t, true, q.Empty())

	q.Put(uint64(4))
	results, err = q.Get(1)
	assert.Equal(t, uint64(4), results[0])
}

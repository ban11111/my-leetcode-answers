package skiplist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	var length = 500
	sl := NewSkipList(16)

	var data []int
	for i := 0; i < length; i++ {
		data = append(data, i)
	}

	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	for _, d := range data {
		assert.True(t, sl.insert(d))
	}

	node := sl.find(250)
	assert.NotNil(t, node)
	assert.EqualValues(t, 250, node.Value())

	assert.EqualValues(t, length, sl.len())

	assert.True(t, sl.delete(250))
	node = sl.find(250)
	assert.Nil(t, node)
	assert.EqualValues(t, length-1, sl.len())
}

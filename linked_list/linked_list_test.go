package linked_list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3))
}
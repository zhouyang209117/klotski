package tests

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	for k, v := range a {
		fmt.Println(fmt.Sprintf("k=%d,v=%d", k, v))
	}
	assert.Equal(t, 3, len(a))
}

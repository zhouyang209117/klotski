package tests

import (
	"testing"
	"github.com/Workiva/go-datastructures/set"
)

func TestAddItems(t *testing.T) {
	set := set.New()
	set.Add(`test`)
	set.Add(`test1`)

	firstSeen := false
	secondSeen := false
	// order is not guaranteed
	for _, item := range set.Flatten() {
		if item.(string) == `test` {
			firstSeen = true
		} else if item.(string) == `test1` {
			secondSeen = true
		}
	}

	if !firstSeen || !secondSeen {
		t.Errorf(`Not all items seen in set.`)
	}
}

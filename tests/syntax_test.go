package tests

import (
	"testing"
	"klotski/model"
	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	s := model.State{}
	assert.Nil(t, s.Pre)
}

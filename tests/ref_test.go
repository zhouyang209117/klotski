package tests

import (
	"testing"
	"klotski/util"
)

func SetInt(a *uint32, value uint32) {
	*a = value
}

func TestRef(t *testing.T) {
	a := uint32(3)
	SetInt(&a, 5)
	util.AssertUint32(t, a, uint32(5))
}

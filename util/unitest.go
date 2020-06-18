package util

import (
	"testing"
	"reflect"
)

func AssertUint32(t *testing.T, got, want uint32)  {
	if want != got {
		t.Errorf("got: %d, want: %d.", got, want)
		panic("unit test error")
	}
}

func AssertUint8(t *testing.T, got, want uint8)  {
	if want != got {
		t.Errorf("got: %d, want: %d.", got, want)
		panic("unit test error")
	}
}


func AssertUint64(t *testing.T, got, want uint64)  {
	if want != got {
		t.Errorf("got: %d, want: %d.", got, want)
		panic("unit test error")
	}
}

func AssertDeep(t *testing.T, got, want interface{})  {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d, want: %d.", got, want)
		panic("unit test error")
	}
}

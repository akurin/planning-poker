package assertions

import (
	"reflect"
	"testing"
)

func DeepEqual(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

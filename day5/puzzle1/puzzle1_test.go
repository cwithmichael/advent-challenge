package puzzle1

import (
	"reflect"
	"testing"
)

func TestParseOpcodes(t *testing.T) {
	var testSeq = []int{1002, 4, 3, 4, 33}
	var expectedSeq = []int{1002, 4, 3, 4, 99}
	actualSeq := ParseOpcodes(testSeq, 1)
	if !reflect.DeepEqual(actualSeq, expectedSeq) {
		t.Errorf("Incorrect result, got: %v, want: %v.", actualSeq, expectedSeq)
	}

	testSeq = []int{1101, 100, -1, 4, 0}
	expectedSeq = []int{1101, 100, -1, 4, 99}
	actualSeq = ParseOpcodes(testSeq, 1)
	if !reflect.DeepEqual(actualSeq, expectedSeq) {
		t.Errorf("Incorrect result, got: %v, want: %v.", actualSeq, expectedSeq)
	}
}

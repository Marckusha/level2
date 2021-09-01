package task2

import (
	"level2/pkq/task2"
	"testing"
)

func TestUnpuck1(t *testing.T) {
	var data = "a4bc2d5e"
	var testOk = "aaaabccddddde"

	result, err := task2.Unpack(data)

	if err != nil {
		t.Errorf("test for OK Failed - error %v", err)
	}
	if result != testOk {
		t.Errorf("test for OK Failed - result not match\n %v %v", result, testOk)
	}
}

func TestUnpuck2(t *testing.T) {
	var data = "abcd"
	var testOk = "abcd"

	result, err := task2.Unpack(data)

	if err != nil {
		t.Errorf("test for OK Failed - error %v", err)
	}
	if result != testOk {
		t.Errorf("test for OK Failed - result not match\n %v %v", result, testOk)
	}
}

func TestUnpuck3(t *testing.T) {
	var data = ""
	var testOk = ""

	result, err := task2.Unpack(data)

	if err != nil {
		t.Errorf("test for OK Failed - error %v", err)
	}
	if result != testOk {
		t.Errorf("test for OK Failed - result not match\n %v %v", result, testOk)
	}
}

func TestUnpuckForError(t *testing.T) {
	var data = "45"
	_, err := task2.Unpack(data)

	if err == nil {
		t.Errorf("test for OK Failed - error"
	}
}

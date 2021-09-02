package task2

import (
	"level2/pkq/task2"
	"testing"
)

func TestUnpuck(t *testing.T) {

	testCases := []struct {
		str    string
		expStr string
	}{
		{
			str:    "a4bc2d5e",
			expStr: "aaaabccddddde",
		},
		{
			str:    "abcd",
			expStr: "abcd",
		},
		{
			str:    "",
			expStr: "",
		},
	}

	for _, test := range testCases {
		result, err := task2.Unpack(test.str)

		if err != nil {
			t.Errorf("test for OK Failed - error %v", err)
		}
		if result != test.expStr {
			t.Errorf("test for OK Failed - result not match\n %v %v", result, test.expStr)
		}

	}
}

func TestUnpuckForError(t *testing.T) {
	var data = "45"
	_, err := task2.Unpack(data)

	if err == nil {
		t.Errorf("test for OK Failed - error")
	}
}

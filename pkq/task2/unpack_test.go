package task2

import (
	"level2/pkq/task2"
	"testing"

	"github.com/stretchr/testify/assert"
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
		result, _ := task2.Unpack(test.str)
		assert.Equal(t, result, test.expStr, "error")
	}
}

func TestUnpuckForError(t *testing.T) {
	var data = "45"
	_, err := task2.Unpack(data)

	if err == nil {
		t.Errorf("test for OK Failed - error")
	}
}

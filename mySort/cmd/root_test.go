package cmd

import (
	"embed"
	"mySort/cmd"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed tests/*
var data embed.FS

func TestOk(t *testing.T) {

	args := make([]string, 0)

	countTests := 7
	for i := 1; i <= countTests; i++ {

		//fmt.Println("TEST N ", i)
		sNum := strconv.Itoa(i)
		pathTestFile := "tests/testFile" + sNum + ".txt"
		pathExpectFile := "tests/expectFile" + sNum + ".txt"
		pathArgs := "tests/args" + sNum + ".txt"

		expectFile, _ := data.ReadFile(pathExpectFile)

		argsData, _ := data.ReadFile(pathArgs)
		argsFile := strings.Split(string(argsData), "\n")

		args = []string{}
		copy(args, os.Args)
		args = append(args, pathTestFile)
		args = append(args, argsFile...)

		c := cmd.NewCommand()
		cmd.SetFlags(c)
		c.SetArgs(args)
		c.Execute()

		//fmt.Println(cmd.TestString == string(expectFile))

		assert.Equal(t, cmd.TestString, string(expectFile), "uncorrect result. Test № "+sNum)
	}
}

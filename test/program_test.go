package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/lrm25/programWrapper"
	"github.com/stretchr/testify/require"
)

var executable string
var expectedOutput string

func init() {
	flag.StringVar(&executable, "exe", "", "Executable to run")
	flag.StringVar(&expectedOutput, "output", "", "Output to check against")
}

func Test_program_cmdError(t *testing.T) {

	cmdErr := errors.New("test error")

	program := programWrapper.NewProgram("")
	program.GetCombinedOutputFunc = func(cmd *exec.Cmd) ([]byte, error) {
		return nil, cmdErr
	}
	defer func() { program.GetCombinedOutputFunc = program.GetCombinedOutput }()
	_, err := program.Run()
	require.Equal(t, cmdErr, err)
}

var nothing int

func Test_program(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	path := filepath.Join(wd, "../testProgram.exe")

	program := programWrapper.NewProgram(path)
	output, err := program.Run()
	require.NoError(t, err)
	require.Contains(t, output, "default message")

	newMessage := "different message"
	program = program.WithParam("message", newMessage)
	output, err = program.Run()
	require.NoError(t, err)
	require.Contains(t, output, newMessage)
}

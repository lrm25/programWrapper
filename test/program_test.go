package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lrm25/wrapper"
	"github.com/stretchr/testify/require"
)

var executable string
var expectedOutput string

func init() {
	flag.StringVar(&executable, "exe", "", "Executable to run")
	flag.StringVar(&expectedOutput, "output", "", "Output to check against")
}

func Test_program(t *testing.T) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	path := filepath.Join(wd, "../testProgram.exe")

	program := wrapper.NewProgram(path)
	err = program.Compile("../testProgram/main.go")
	require.NoError(t, err)

	output, err := program.Run()
	require.NoError(t, err)
	require.Contains(t, output, "default message")

	newMessage := "different message"
	program = program.WithParam("message", newMessage)
	output, err = program.Run()
	require.NoError(t, err)
	require.Contains(t, output, newMessage)
}

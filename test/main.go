package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lrm25/outputWrapper"
)

var executable string
var expectedOutput string

func init() {
	flag.StringVar(&executable, "exe", "", "Executable to run")
	flag.StringVar(&expectedOutput, "output", "", "Output to check against")
}

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	path := filepath.Join(wd, "../testprogram.exe")

	program := outputWrapper.NewProgram(path)
	output, err := program.Run()
	fmt.Println(output)
}

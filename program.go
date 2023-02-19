package wrapper

import (
	"errors"
	"fmt"
	"os/exec"
)

type Program struct {
	executable string
	params     map[string]interface{}
}

func NewProgram(executable string) *Program {
	program := &Program{
		executable: executable,
		params:     make(map[string]interface{}),
	}
	return program
}

func (p *Program) WithParam(name string, value interface{}) *Program {
	p.params[name] = value
	return p
}

var getCombinedOutputFunc = getCombinedOutput

func getCombinedOutput(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

func (p *Program) Run() (string, error) {

	var params []string
	for name, value := range p.params {
		params = append(params, "-"+name, fmt.Sprintf("%v", value))
	}

	cmd := exec.Command(p.executable, params...)
	output, err := getCombinedOutputFunc(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (p *Program) Compile(mainPath string) error {
	cmd := exec.Command("go", "build", "-o", p.executable, mainPath)
	output, err := getCombinedOutputFunc(cmd)
	if err != nil {
		return err
	}
	outputStr := string(output)
	if outputStr != "" {
		return errors.New("Error compiling: " + outputStr)
	}
	return nil
}

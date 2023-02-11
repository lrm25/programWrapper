package programWrapper

import (
	"fmt"
	"os/exec"
)

type Program struct {
	executable            string
	params                map[string]interface{}
	getCombinedOutputFunc func(cmd *exec.Cmd) ([]byte, error)
}

func NewProgram(executable string) *Program {
	program := &Program{
		executable: executable,
		params:     make(map[string]interface{}),
	}
	program.getCombinedOutputFunc = program.getCombinedOutput
	return program
}

func (p *Program) WithParam(name string, value interface{}) *Program {
	p.params[name] = value
	return p
}

func (p *Program) getCombinedOutput(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

func (p *Program) Run() (string, error) {

	var params []string
	for name, value := range p.params {
		params = append(params, "-"+name, fmt.Sprintf("%v", value))
	}

	cmd := exec.Command(p.executable, params...)
	output, err := p.getCombinedOutputFunc(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

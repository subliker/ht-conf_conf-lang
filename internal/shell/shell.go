package shell

import (
	"fmt"

	"github.com/subliker/ht-conf_conf-lang/internal/converter"
)

type Shell struct {
	input     string
	output    string
	converter *converter.Converter
}

// New creates new instance of Shell
func New() *Shell {
	c := converter.New()
	return &Shell{
		input:     _inputPath,
		output:    _outputPath,
		converter: c,
	}
}

// Run starts shell
func (sh *Shell) Run() {
	splitLineChan := make(chan []string)

	go sh.ParseInput(splitLineChan)

	for splitLine := range splitLineChan {
		sh.converter.ParseLine(splitLine)
	}
	fmt.Print("ready")
}

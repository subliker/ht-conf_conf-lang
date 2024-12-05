package shell

import "flag"

var _inputPath string
var _outputPath string

func init() {
	flag.StringVar(&_inputPath, "input", "./input.nya", "set input config file")
	flag.StringVar(&_outputPath, "output", "./output.toml", "set output config file")
}

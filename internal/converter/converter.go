package converter

import (
	"fmt"
	"os"
)

type Converter struct {
	vars     map[string]interface{}
	handlers []func(string) bool

	f *os.File
}

func New(outputPath string) *Converter {
	var c Converter

	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	c.f = f

	c.vars = make(map[string]interface{})

	c.handlers = append(c.handlers, c.HandlerRem, c.HandleValue)
	return &c
}

func (c *Converter) Vars() map[string]interface{} {
	return c.vars
}

func (c *Converter) ParseLine(line string) {
	lenLine := len(line)
	if lenLine == 0 {
		return
	}

	for _, handler := range c.handlers {
		if ok := handler(line); ok {
			return
		}
	}
	panic(fmt.Errorf("error in line: %s", line))
}

func (c *Converter) Close() {
	c.f.Close()
}

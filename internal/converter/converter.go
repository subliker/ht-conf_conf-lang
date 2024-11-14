package converter

import "fmt"

type Converter struct {
	vars     map[string]interface{}
	handlers []func(string) bool
}

func New() *Converter {
	var c Converter

	c.vars = make(map[string]interface{})

	c.handlers = append(c.handlers, c.HandlerRem)
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

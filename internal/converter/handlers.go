package converter

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	varDeclRegex = regexp.MustCompile("^var ([a-zA-Z_][a-zA-Z0-9_]*) = (.+?)$")
	listRegex    = regexp.MustCompile(`^list\((?:\s?\d\,?)+\)$`)
)

func (c *Converter) HandlerRem(line string) bool {
	if strings.HasPrefix(line, "REM") || line == "" {
		return true
	}

	return false
}

func (c *Converter) HandleValue(line string) bool {
	if ss := varDeclRegex.FindStringSubmatch(line); ss != nil {
		name := ss[1]
		value := ss[2]
		parsedValue, err := c.ParseValue(value)
		if err != nil {
			return false
		}
		c.vars[name] = parsedValue
		return true
	}
	return false
}

func (c *Converter) ParseValue(value string) (interface{}, error) {
	if ss := listRegex.FindStringSubmatch(value); ss != nil {
		itemsStr := strings.TrimPrefix(strings.TrimSuffix(value, ")"), "list(")
		items := strings.Split(itemsStr, ",")
		var itemsList []interface{}
		for _, item := range items {
			item = strings.TrimSpace(item)
			parsedItem, err := c.ParseSingleValue(item)
			if err != nil {
				return nil, err
			}
			itemsList = append(itemsList, parsedItem)
		}
		return itemsList, nil
	}

	return c.ParseSingleValue(value)
}

func (c *Converter) ParseSingleValue(value string) (interface{}, error) {
	d, err := strconv.Atoi(value)

	return d, err
}

package converter

import (
	"errors"
	"fmt"
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
		fmt.Fprintf(c.f, "# %s\n", strings.TrimSpace(strings.TrimPrefix(line, "REM")))
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
			fmt.Print(err)
			return false
		}
		c.vars[name] = parsedValue

		strVal, err := valToStr(parsedValue)
		if err != nil {
			panic(1)
		}

		fmt.Fprintf(c.f, "%s=%s\n", name, strVal)
		return true
	}
	return false
}

func valToStr(val interface{}) (string, error) {
	switch v := val.(type) {
	case int:
		return strconv.Itoa(v), nil
	case []interface{}:
		strVal := "["
		for i, av := range v {
			s, err := valToStr(av)
			if err != nil {
				return "", err
			}
			strVal += s
			if i != len(v)-1 {
				strVal += ", "
			}
		}
		strVal += "]"
		return strVal, nil
	}
	return "", errors.New("type")
}

func (c *Converter) ParseValue(value string) (interface{}, error) {
	if ss := listRegex.FindStringSubmatch(value); ss != nil {
		itemsStr := strings.TrimPrefix(strings.TrimSuffix(value, ")"), "list(")
		items := strings.Split(itemsStr, ",")
		var itemsList []interface{}
		for _, item := range items {
			item = strings.TrimSpace(item)
			parsedItem, err := c.ParseValue(item)
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

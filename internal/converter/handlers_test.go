package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerRem(t *testing.T) {
	assert := assert.New(t)
	c := New(".nya")

	assert.Equal(c.HandlerRem("REM this is comment!"), true)
	assert.Equal(c.HandlerRem(""), true)
	assert.Equal(c.HandlerRem("var a = 1"), false)
}

func TestHandleValue(t *testing.T) {
	assert := assert.New(t)
	c := New(".nya")

	assert.Equal(true, c.HandleValue("var a = 1"))
	assert.Equal(1, c.Vars()["a"])

	assert.Equal(true, c.HandleValue("var l = list(1,2,3)"))
	assert.Equal([]interface{}([]interface{}{1, 2, 3}), c.Vars()["l"])

	assert.Equal(true, c.HandleValue("var l = list(1,2,3)"))
	assert.Equal([]interface{}([]interface{}{1, 2, 3}), c.Vars()["l"])

	assert.Equal(false, c.HandleValue("var b = asdasd"))
	assert.Equal(nil, c.Vars()["b"])
}

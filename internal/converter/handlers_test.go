package converter

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerRem(t *testing.T) {
	assert := assert.New(t)
	c := New(".nya")
	defer os.Remove(".nya")

	assert.Equal(c.HandlerRem("REM this is comment!"), true)
	assert.Equal(c.HandlerRem("var a = 1"), false)
	assert.Equal(c.HandlerRem("nothing"), false)

	c.Close()

	f, err := os.Open(".nya")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	assert.Equal(`# this is comment!
`, string(d))
}

func TestHandleValue(t *testing.T) {
	assert := assert.New(t)
	c := New(".nya")
	defer os.Remove(".nya")

	assert.Equal(c.HandleValue("REM this is comment!"), false)
	assert.Equal(c.HandleValue("var a = 1"), true)
	assert.Equal(c.HandleValue("var arr = [1, 2, 3]"), false)
	assert.Equal(c.HandleValue("var arr = list(1, 2, 3)"), true)
	assert.Equal(c.HandleValue("var arr2 = list(3,4,5)"), true)
	assert.Equal(c.HandleValue("var arrLen = .[arr len() +]."), false)
	assert.Equal(c.HandleValue("var arrLen = .[arr len()]."), true)
	assert.Equal(c.HandleValue("var sumAAndArrLen = .[a arrLen +]."), true)
	assert.Equal(c.HandleValue("var c = .[sumAAndArrLen arr *]."), false)
	assert.Equal(c.HandleValue("var c = .[sumAAndArrLen sumAAndArrLen *]."), true)
	assert.Equal(c.HandleValue("nothing"), false)

	c.Close()

	f, err := os.Open(".nya")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	assert.Equal(`a=1
arr=[1, 2, 3]
arr2=[3, 4, 5]
arrLen=3
sumAAndArrLen=4
c=16
`, string(d))
}

package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	Name string
	Age  int
}

func TestCBOR(t *testing.T) {
	assert := assert.New(t)

	app := &App{}
	app.UseCBOR()

	me := &testStruct{
		Name: "John Doe",
		Age:  32323,
	}

	enc, err := app.MarshalCBOR(me)
	if err != nil {
		t.Fatal(err)
		return
	}

	you := &testStruct{}
	err = app.UnmarshalCBOR(enc, you)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(me, you)
}

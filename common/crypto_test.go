package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryption(t *testing.T) {
	assert := assert.New(t)

	app := &App{}
	app.UseCBOR()

	me := &testStruct{
		Name: "John Doe",
		Age:  32323,
	}

	key := app.SHA256([]byte("hello"))

	enc, err := app.MarshalCBOR(me)
	if err != nil {
		t.Fatal(err)
		return
	}

	ciphertext, err := app.Encrypt(key, enc)
	if err != nil {
		t.Fatal(err)
		return
	}

	plaintext, err := app.Decrypt(key, ciphertext)
	if err != nil {
		t.Fatal(err)
		return
	}

	you := &testStruct{}
	err = app.UnmarshalCBOR(plaintext, you)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(me, you)
}

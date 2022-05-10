package common

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
)

func (app *App) SHA1(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	return h.Sum(nil)
}

func (app *App) SHA256(b ...[]byte) []byte {
	digest := sha256.Sum256(bytes.Join(b, nil))
	return digest[:]
}

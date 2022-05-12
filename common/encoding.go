package common

import (
	"encoding/ascii85"

	"github.com/fxamacker/cbor/v2"
)

func (app *App) MarshalCBOR(x interface{}) ([]byte, error) {
	return app.cbor.Marshal(x)
}

func (app *App) UnmarshalCBOR(b []byte, dst interface{}) error {
	return cbor.Unmarshal(b, dst)
}

func (app *App) CompactSerial(x interface{}) (string, error) {
	b, err := app.MarshalCBOR(x)
	if err != nil {
		return "", nil
	}
	dst := make([]byte, len(b))
	n := ascii85.Encode(dst, b)
	return string(dst[:n]), nil
}

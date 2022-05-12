package common

import (
	"bytes"
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
	buf := bytes.NewBuffer(nil)
	enc := ascii85.NewEncoder(buf)
	enc.Write(b)
	enc.Close()
	return string(buf.Bytes()), nil
}

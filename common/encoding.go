package common

import (
	"bytes"
	"encoding/ascii85"
	"encoding/json"

	"github.com/fxamacker/cbor/v2"
)

func (app *App) MarshalJSON(x interface{}) ([]byte, error) {
	return json.Marshal(x)
}

func (app *App) UnmarshalJSON(b []byte, dst interface{}) error {
	return json.Unmarshal(b, dst)
}

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

func (app *App) ExpandSerial(x interface{}) (string, error) {
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

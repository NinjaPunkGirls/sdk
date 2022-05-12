package common

import "github.com/fxamacker/cbor/v2"

func (app *App) MarshalCBOR(x interface{}) ([]byte, error) {
	return app.cbor.Marshal(x)
}

func (app *App) UnmarshalCBOR(b []byte, dst interface{}) error {
	return cbor.Unmarshal(b, dst)
}

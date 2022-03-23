package npgsdk

import (
	"github.com/golangdaddy/relysia-client"
)

func (sdk *SDK) CreateSwapOffer(walletID, tokenID, receiveType string, amount float64) (*relysia.OfferResponse, error) {

	return sdk.Relysia.Offer(walletID, tokenID, receiveType, amount)
}

func (sdk *SDK) AcceptSwapOffer(walletID, swapHex string) (*relysia.SwapResponse, error) {

	return sdk.Relysia.Swap(walletID, swapHex)
}

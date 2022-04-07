package npgsdk

import (
	"fmt"
	"log"

	"github.com/golangdaddy/relysia-client"
	"github.com/kr/pretty"
)

// WalletBalanceBSV gives you the BSV balance
func (sdk *SDK) WalletBalanceBSV(walletID, symbol string) (*relysia.BalanceResponse, error) {

	response, err := sdk.Relysia.Balance(walletID, "BSV", "")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// WalletBalanceSTAS gives you the stas tokens in a slice
func (sdk *SDK) WalletBalanceSTAS(walletID, symbol string) ([]map[string]interface{}, error) {

	response, err := sdk.Relysia.Balance(walletID, "STAS", "")
	if err != nil {
		return nil, err
	}

	pretty.Println(response)

	output := []map[string]interface{}{}
	for _, coin := range response.Coins {
		if coin.Symbol == symbol {

			tokenInfo, err := sdk.Relysia.GetToken(coin.ID())
			if err != nil {
				log.Println(fmt.Errorf("GetToken: %w", err))
				continue
			}

			output = append(output, tokenInfo)
		}
	}

	return output, nil
}

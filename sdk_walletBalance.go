package npgsdk

import "log"

func (sdk *SDK) WalletBalanceSTAS(walletID, symbol string) ([]map[string]interface{}, error) {

	response, err := sdk.Relysia.Balance(walletID, "STAS", "")
	if err != nil {
		return nil, err
	}

	output := []map[string]interface{}{}
	for _, coin := range response.Coins {
		if coin.Symbol == symbol {

			tokenInfo, err := sdk.Relysia.GetToken(coin.ID())
			if err != nil {
				log.Println(err)
				continue
			}

			output = append(output, tokenInfo)
		}
	}

	return output, nil
}

package npgsdk

import (
	"fmt"
	"strings"
)

func (sdk *SDK) GetWalletByTitle(title string) (string, error) {

	walletList, err := sdk.Relysia.Wallets()
	if err != nil {
		return "", err
	}

	for _, walletInfo := range walletList {
		if strings.ToLower(walletInfo.WalletTitle) == strings.ToLower(title) {
			return walletInfo.WalletID, nil
		}
	}

	return "", fmt.Errorf("no wallet found with title '%s'", title)
}

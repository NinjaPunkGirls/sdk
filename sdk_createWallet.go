package npgsdk

// CreateWallet makes a wallet for the user that has the given title
func (sdk *SDK) CreateWallet(title string) (string, error) {

	walletID, err := sdk.Relysia.CreateWallet(title)
	if err != nil {
		return "", err
	}

	return walletID, nil
}

package npgsdk

func (sdk *SDK) CreateWallet(title string) (string, error) {

	walletID, err := sdk.Relysia.CreateWallet(title)
	if err != nil {
		return "", err
	}

	return walletID, nil
}

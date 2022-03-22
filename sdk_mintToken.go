package npgsdk

import (
	"github.com/golangdaddy/relysia-client"
)

func BuildTokenRequest(name, description, image, symbol string, supply int) *relysia.IssueRequest {
	issueRequest := relysia.DemoTokenRequest()
	issueRequest.Splitable = false
	issueRequest.Name = name
	issueRequest.Description = description
	issueRequest.Image = image
	issueRequest.Symbol = symbol
	issueRequest.TokenSupply = supply
	issueRequest.Decimals = 0
	issueRequest.SatsPerToken = 1500
	issueRequest.Properties.Meta.Media = []*relysia.MetaMedia{
		&relysia.MetaMedia{
			URI:    "?",
			Type:   "?",
			AltURI: "?",
		},
	}
	return issueRequest
}

func (sdk *SDK) MintToken(walletID string, issueRequest *relysia.IssueRequest) (*relysia.IssueResponse, error) {

	response, err := sdk.Relysia.Issue(
		relysia.Headers{
			"protocol": "stas",
			"walletID": walletID,
		},
		issueRequest,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

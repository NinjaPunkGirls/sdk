# sdk

## NPG SDK

invoke a new SDK instance for each request

```
func (app *App) AuthenticatedUser(c *gin.Context) (*UserSession, error) {
	ctx := context.Background()

	token := c.Request.Header["Authorization"][0]

	doc, err := app.firestoreClient.Collection("sessions").Doc(
		hex.EncodeToString(app.Hash([]byte(token))),
	).Get(ctx)
	if err != nil {
		return nil, err
	}

	session := &UserSession{}
	if err := doc.DataTo(session); err != nil {
		return nil, err
	}
	session.sdk = npgsdk.InitSDK(CONST_PROJECT, token)

	return session, nil
}
```
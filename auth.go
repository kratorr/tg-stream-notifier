package main

import (
	"github.com/nicklaw5/helix/v2"
)

func Auth(clientID, clientSecret string) (string, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return "", err
	}

	resp, err := client.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		return "", err
	}
	// Set the access token on the client
	//client.SetAppAccessToken(resp.Data.AccessToken)
	return resp.Data.AccessToken, nil
}

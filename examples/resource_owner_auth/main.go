package main

////// Resource Owner Overview ///////////////////////////////////////////////////
//
// This example leverages a client_id, client_secret, resource owner username
// and resource owner password to authenticate to the SCIM Oauth2 Endpoint
// (https://<ScimUrl>/ouath2/token/<AppId>).
//
// If Authentication is successful an Oauth2 token with Refresh token is returned.
// The returned Oauth2 token may then be utilized to establish a Service
// based on the https client module to interact with the SCIM API.
//
///////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
	cybr_pam_scim "github.com/strick-j/cybr_pam_scim/pkg/cybr_pam_scim"
)

func main() {
	// Set the file name of the configuration file
	viper.SetConfigName("config")

	// Set the path to look for configuration file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Read environment variables from config.yml and set them with type assertion
	clientId, ok := viper.Get("IDENTITY.CLIENT_ID").(string)
	clientSecret, ok := viper.Get("IDENTITY.CLIENT_SECRET").(string)
	resourceUsername, ok := viper.Get("IDENTITY.USERNAME").(string)
	resourceSecret, ok := viper.Get("IDENTITY.SECRET").(string)
	clientAppId, ok := viper.Get("IDENTITY.APP_ID").(string)
	clientUrl, ok := viper.Get("IDENTITY.URL").(string)

	// If type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	// Obtain an auth token with the provided credentials and endpoint parameters
	// The Oauth2 Token format should be the following:
	// type Token struct {
	//     AccessToken string `json:"access_token"`
	//     TokenType string `json:"token_type,omitempty"`
	//     RefreshToken string `json:"refresh_token,omitempty"`
	//     Expiry time.Time `json:"expiry,omitempty"`
	// }
	authToken, err := cybr_pam_scim.OauthResourceOwner(clientId, clientSecret, clientAppId, clientUrl, resourceUsername, resourceSecret)
	if err != nil {
		log.Fatalf("Authentication Failed. %s", err)
	}

	authTokenJSON, err := json.MarshalIndent(authToken, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Auth Token JSON Response %s\n", string(authTokenJSON))
}

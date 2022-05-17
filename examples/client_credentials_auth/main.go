package main

////// Client Credentials Overview ///////////////////////////////////////////////////
//
// This example leverages a client_id and client_secret to authenticate
// to the SCIM Oauth2 Endpoint (https://<ScimUrl>/ouath2/token/<AppId>).
//
// If Authentication is successful a clientCredentials.Oauth2 token is returned.
// The returned Oauth2 token is then utilized to establish a Service
// based on thehttps client module to interact with the SCIM API.
//
/////////////////////////////////////////////////////////////////////////////

import (
	"context"
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
	clientAppId, ok := viper.Get("IDENTITY.APP_ID").(string)
	clientUrl, ok := viper.Get("IDENTITY.URL").(string)

	// If type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	// Obtain an auth token with the provided credentials and endpoint parameters
	// The Oauth2 Token format should be the following:
	// type Token struct {
	// 	   AccessToken string `json:"access_token"`
	//     TokenType string `json:"token_type,omitempty"`
	//     RefreshToken string `json:"refresh_token,omitempty"`
	//     Expiry time.Time `json:"expiry,omitempty"`
	// }
	// Note: Client Credentials token does not contain a refresh token
	authToken, err := cybr_pam_scim.OauthCredClient(clientId, clientSecret, clientAppId, clientUrl)
	if err != nil {
		log.Fatalf("Authentication Failed. %s", err)
	}

	// Marshal the authToken for display purposes
	authTokenJSON, err := json.MarshalIndent(authToken, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Print the auth Token
	fmt.Printf("Auth Token JSON Response %s\n", string(authTokenJSON))

	// Utilize the returned oauth2.Token to create a service that leverages the
	// the https client module
	s := cybr_pam_scim.NewService(clientUrl, "scim", "v2", false, authToken)

	// Utilize the returned service to interact with the SCIM API
	// In this example all users are being retrieved and the DisplayName of the
	// first user in the struct is being displayed
	Users, err := s.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error Retrieving users")
	}
	// Do something with the Users struct
	fmt.Printf(Users.Resources[1].DisplayName)
}

package main

////// User Auth Overview ///////////////////////////////////////////////////
//
// This example leverages a username and password to authenticate
// to the SCIM Oauth2 Endpoint (https://<ScimUrl>/ouath2/token/<AppId>).
// If Authentication is successful an Oauth2 token is returned.
// The returned Oauth2 token is then utilized to establish a Service
// based on thehttps client module to interact with the SCIM API.
//
/////////////////////////////////////////////////////////////////////////////

import (
	"context"
	"fmt"
	"log"
	"os"

	cybr_pam_scim "github.com/strick-j/cybr_pam_scim/pkg/cybr_pam_scim"
)

func main() {
	// Declare constants used in this example
	var (
		username = os.Getenv("IDENITY_USERNAME")
		secret   = os.Getenv("IDENTITY_SECRET")
		appId    = os.Getenv("IDENTITY_APP_ID")
		scimUrl  = os.Getenv("IDENTITY_URL")
	)

	// Obtain an auth token with the provided credentials and endpoint parameters
	// The Oauth2 Token format should be the following:
	// type Token struct {
	// 	   AccessToken string `json:"access_token"`
	//     TokenType string `json:"token_type,omitempty"`
	//     RefreshToken string `json:"refresh_token,omitempty"`
	//     Expiry time.Time `json:"expiry,omitempty"`
	// }
	authToken, err := cybr_pam_scim.OauthCredClient(username, secret, appId, scimUrl)
	if err != nil {
		log.Fatalf("Authentication Failed. %s", err)
	}

	// Utilize the returned oauth2.Token to create a service that leverages the
	// the https client module
	s := cybr_pam_scim.NewService(scimUrl, "scim", "v2", false, authToken)

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

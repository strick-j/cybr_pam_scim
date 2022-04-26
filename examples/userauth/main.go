package main

import (
	"context"
	"fmt"

	pamscim "github.com/strick-j/cybr-pam-scim/pkg/pamscim"
)

////// User auth Overview ///////////////////////////////////////////////////
//
// This example leverages a username and password to authenticate
// to the SCIM Oauth2 Endpoint (https://<ScimUrl>/ouath2/token/<AppId>).
// If Authentication is successful an Oauth2 token is returned.
// The returned Oauth2 token is then utilized to establish a Service
// based on thehttps client module to interact with the SCIM API.
//
/////////////////////////////////////////////////////////////////////////////

// Declare constants used in this example
const (
	User    = "identity-privilege-integration-user$@example.com"
	Secret  = "Sup3RS3cr3T123!"
	AppId   = "IdentityScimAppExample"
	ScimUrl = "example1234.my.idaptive.app"
)

func main() {
	// Obtain an auth token with the provided credentials and endpoint parameters
	// The Oauth2 Token format should be the following:
	// type Token struct {
	// 	   AccessToken string `json:"access_token"`
	//     TokenType string `json:"token_type,omitempty"`
	//     RefreshToken string `json:"refresh_token,omitempty"`
	//     Expiry time.Time `json:"expiry,omitempty"`
	// }
	authToken, err := pamscim.OauthCredClient(User, Secret, AppId, ScimUrl)
	if err != nil {
		fmt.Printf("Error obtaining auth token")
	}

	// Utilize the returned oauth2.Token to create a service that leverages the
	// the https client module
	s := pamscim.NewService(ScimUrl, "scim", "v2", false, authToken)

	// Utilize the service to interact with the SCIM API
	// In this example all users are being retrieved and the DisplayName of the
	// first user in the struct is being displayed
	Users, err := s.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error Retrieving users")
	}
	fmt.Printf(Users.Resources[1].DisplayName)
}

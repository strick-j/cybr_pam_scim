package main

import (
	"context"
	"fmt"
	"os"

	pamscim "github.com/strick-j/cybr-pam-scim/pkg/cybr_pam_scim"
	"golang.org/x/oauth2"
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

func main() {
	// Declare constants used in this example
	var (
		token   = os.Getenv("BEARER_TOKEN")
		scimUrl = os.Getenv("IDENTITY_URL")
	)

	// Create the token struct with the existing Bearer token retrieved
	// from the CyberArk Identity Portal
	authToken := &oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	// Utilize the returned oauth2.Token to create a service that leverages the
	// the https client module
	s := pamscim.NewService(scimUrl, "scim", "v2", false, authToken)

	// Utilize the service to interact with the SCIM API
	// In this example all users are being retrieved and the DisplayName of the
	// first user in the struct is being displayed
	Users, err := s.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error Retrieving users")
	}
	fmt.Printf(Users.Resources[1].DisplayName)
}

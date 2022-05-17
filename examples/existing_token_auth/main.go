package main

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	cybr_pam_scim "github.com/strick-j/cybr_pam_scim/pkg/cybr_pam_scim"
	"golang.org/x/oauth2"
)

////// Existing Token Overview ///////////////////////////////////////////////////
//
// This example leverages a an existing Bearer Token to authenticate.
// to the SCIM Oauth2 Endpoint (https://<ScimUrl>/ouath2/token/<AppId>).
//
//////////////////////////////////////////////////////////////////////////////////

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
	token, ok := viper.Get("IDENTITY.BEARER_TOKEN").(string)
	clientUrl, ok := viper.Get("IDENTITY.URL").(string)

	// If type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	// Create the token struct with the existing Bearer token retrieved
	// from the CyberArk Identity Portal
	authToken := &oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	// Utilize the returned oauth2.Token to create a service that leverages the
	// the https client module
	s := cybr_pam_scim.NewService(clientUrl, "scim", "v2", false, authToken)

	// Utilize the service to interact with the SCIM API
	// In this example all users are being retrieved and the DisplayName of the
	// first user in the struct is being displayed
	Users, err := s.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error Retrieving users")
	}
	fmt.Printf(Users.Resources[1].DisplayName)
}

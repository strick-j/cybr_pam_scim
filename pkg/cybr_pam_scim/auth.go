package cybr_pam_scim

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type tokenSource struct {
	ctx  context.Context
	conf *clientcredentials.Config
}

// OauthCredClient returns a validated Oauth2 Authentication Token based on the following provided information:
//   clientID - Username for the SCIM Application (e.g. "identity-privilege-integration-user$@example.com")
//   clientSecret - Password for the SCIM Application
//   clientAppID - ID for the SCIM Application
//   clientURL - URL for the SCIM Application (e.g. "example.my.idaptive.app")
func OauthCredClient(clientID, clientSecret, clientAppID, clientURL string) (*oauth2.Token, error) {
	// TODO: PARSE URL and add HTTPS if URL scheme not provided

	// Establish oauth2/clientcredentials config with user provided data
	var credentialConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle:    0,
		Scopes:       []string{"scim"},
	}

	// Create tokenSource with provided configuration info
	ts := &tokenSource{
		ctx:  context.Background(),
		conf: &credentialConfig,
	}

	// Request new token from SCIM server using Client Credentials
	authToken, err := ts.conf.Token(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain SCIM Oauth2 Token %w", err)
	}

	return authToken, nil
}

// OauthResourceOwner returns a validated Oauth2 Authentication Token with Refresh Token based on the following provided information:
//   clientID - Username for the SCIM Application (e.g. "identity-privilege-integration-user$@example.com")
//   clientSecret - Password for the SCIM Application
//   clientAppID - ID for the SCIM Application
//   clientURL - URL for the SCIM Application (e.g. "example.my.idaptive.app")
//   resourceUsername - Username for the Resource Owner
//   resourcePassword - Password for the Resource Owner
func OauthResourceOwner(clientID, clientSecret, clientAppID, clientURL, resourceUsername, resourcePassword string) (*oauth2.Token, error) {
	// TODO: PARSE URL and add HTTPS if URL scheme not provided

	endpoint := oauth2.Endpoint{
		AuthURL:   clientURL + "/oauth2/authorize/" + clientAppID,
		TokenURL:  clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle: 0,
	}

	// Establish oauth2/clientcredentials config with user provided data
	var resourceOwnerConfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		Scopes:       []string{"scim"},
	}

	ctx := context.Background()

	authToken, err := resourceOwnerConfig.PasswordCredentialsToken(ctx, resourceUsername, resourcePassword)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain SCIM Oauth2 Token %w", err)
	}

	return authToken, nil
}

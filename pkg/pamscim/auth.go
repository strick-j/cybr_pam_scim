package pamscim

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
func OauthCredClient(clientID string, clientSecret string, clientAppID string, clientURL string) (*oauth2.Token, error) {
	// Establish oauth2/clientcredentials config with user provided data
	var credentialConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + clientURL + "/oauth2/token/" + clientAppID,
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
		return nil, fmt.Errorf("Failed to obtain SCIM Oauth2 Token")
	}

	return authToken, nil
}

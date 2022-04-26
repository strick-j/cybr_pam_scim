package pamscim

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type Service struct {
	client *Client
}

type transport struct {
	token string
}

func NewService(clientURL string, clientApiEndpoint string, clientApiVersion string, verbose bool, authToken *oauth2.Token) *Service {
	t := transport{
		token: authToken.AccessToken,
	}

	return &Service{
		client: NewClient(
			&http.Client{Transport: &t},
			Options{
				ApiURL:  fmt.Sprintf("https://%s/%s/%s", clientURL, clientApiEndpoint, clientApiVersion),
				Verbose: verbose,
			},
		),
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", "Bearer "+t.token)

	return http.DefaultTransport.RoundTrip(r)
}

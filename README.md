# cybr-pam-scim <!-- omit in toc -->

Contains functions for interacting with the CyberArk Identity SCIM Interface for Privileged Access Management (Self Hosted or Privilege Cloud).

## Table of Contents <!-- omit in toc -->

- [Usage](#usage)
    - [Documentation](#Documentation)
- [Example Source Code](#example-source-code)
- [Security](#security)
- [Contributions](#contributions)
- [License](#license)

## Usage

All functions are documented with example usage in their respective go files. 

## Example Source Code

### Logon and GET Users

```go
package main

import (
	"fmt"
	"log"
	"os"

	cybr_pam_scim "github.com/strick-j/cybr_pam_scim"
)

func main() {
    var (
        username = os.Getenv("IDENITY_USERNAME")
        secret = os.Getenv("IDENTITY_SECRET")
        appId = os.Getenv("IDENTITY_APP_ID")
        scimUrl = os.Getenv("IDENTITY_URL")
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
		fmt.Fatalf("Authentication Failed. %s", err)
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
```

## Secrurity
If there is a security concern or bug discovered, please responsibly disclose all information to joe (dot) strickland (at) cyberark (dot) com.

## Contributions

Pull Requests are currently being accepted.  Please read and follow the guidelines laid out in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[Apache 2.0](LICENSE)
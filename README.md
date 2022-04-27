# cybr-pam-scim <!-- omit in toc -->

Contains functions for interacting with the CyberArk Identity SCIM Interface for Privileged Access Management (Self Hosted or Privilege Cloud).

## Table of Contents <!-- omit in toc -->

- [Usage](#usage)
    - [Authentication](#authentication)
	- [Users](#users)
	- [Groups](#groups)
	- [Containers (Safes)](#containers-safes)
	- [Privileged Data (Accounts)](#privileged-data-accounts)
- [Example Source Code](#example-source-code)
- [Security](#security)
- [Contributions](#contributions)
- [License](#license)

## Usage

All functions are documented with example usage in their respective go files. 

### Authentication

##### OauthCredClient
* Requires:
	* Username: The username (e.g. identity-integration-user@abc1234.my.idaptive.app)
	* Secret: The user secret for authentiction
	* Application Id: The SCIM Oauth Application Name (e.g. examplescimapp)
	* CyberArk Identity Url: The base URL (e.g. abc1234.my.idaptive.app)

* Returns:
	* Authentication Success: Authentication token in the format of [oauth2.token](https://pkg.go.dev/golang.org/x/oauth2#Token)
	* Authentication Failure: Error

### Service

#### NewService
* Requires:
	* CyberArk Identity Url - The base URL (e.g. abc1234.my.idaptive.app)
	* CyberArk Identity API Endpoint _ (e.g. scim)
	* CyberArk Identity API Version - (e.g. v2)
	* Authentication Token - (e.g. Token returned from OauthCredClient)

* Returns:
	* HTTP Client wrapped in a service with Roundtrip parameters set

### Users

| Function | Action | Input | Output | Notes |
| --- | --- | --- | --- | --- |
| GetUsers | GET | - | [types.Users](pkg\cybr_pam_scim\types\users.go) or error | - |
| GetUsersIndex | GET | Start Index and Count as strings | [types.Users](pkg\cybr_pam_scim\types\users.go) or error | - |
| GetUsersSort | GET |Sort By and Sord Order as strings | [types.Users](pkg\cybr_pam_scim\types\users.go) or error | - |
| GetUserById | GET | User Id as string | [types.User](pkg\cybr_pam_scim\types\users.go) or error | Requires PVWA 12.2+ |
| GetUserByFilter | GET | Filter Type and Filter Query as strings | [types.User](pkg\cybr_pam_scim\types\users.go) or error | Filter Query is Case Sensitive |
| AddUser | POST | [types.User](pkg\cybr_pam_scim\types\users.go) struct | [types.User](pkg\cybr_pam_scim\types\users.go) or error | - |
| UpdateUser | PUT | [types.User](pkg\cybr_pam_scim\types\users.go) struct | [types.User](pkg\cybr_pam_scim\types\users.go) or error | - |
| DeleteUser | DELETE | User Id as string | error | No response is returned on success |

### Groups

### Containers (Safes)

### Privileged Data (Accounts)


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
	//     AccessToken string `json:"access_token"`
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
# cybr-pam-scim <!-- omit in toc -->

Contains functions for interacting with the CyberArk Identity SCIM Interface for Privileged Access Management (Self Hosted or Privilege Cloud).

## Table of Contents <!-- omit in toc -->

- [Usage](#usage)
    - [Authentication](#authentication)
	- [Users](#users)
	- [Groups](#groups)
	- [Containers (Safes)](#containers-safes)
	- [Conatiner (Safe) Permissions](#container-safe-permissions)
	- [Privileged Data (Accounts)](#privileged-data-accounts)
- [Example Source Code](#example-source-code)
- [Security](#security)
- [Contributions](#contributions)
- [License](#license)

## Usage

```go
import (
	"github.com/strick-j/cybr_pam_scim"
	"github.com/strick-j/cybr_pam_scim/types"
)

```

All functions are documented with example usage in their respective go files. General flow for usage will be:
1. Obtain oauth2.token through authentication via Username and Password or by using existing Bearer Token
2. Establish Service with Oauth2 Token
3. Utilize Service to interact User, Group, Container, or Privileged Data functions

### Authentication

| Function | Input | Output |
|:--- |:--- |:--- |
| `OauthCredClient` | Username, Secret, Application Id, Identity URL | [oauth2.token](https://pkg.go.dev/golang.org/x/oauth2#Token) or error |

### Service

| Function | Input | Output |
|:--- |:--- |:--- |
| `NewService` | Identity URL, Identity API Endpoint, Identity API Version, Authentication Token | Service struct containing http.Client |

### Users

| Function | Input | Output | PVWA 12.2+ Required |
|:--- |:--- |:--- |:---:|
| `GetUsers` | - | [types.Users](pkg\cybr_pam_scim\types\users.go) or error | |
| `GetUsersIndex` | Start Index and Count | [types.Users](pkg\cybr_pam_scim\types\users.go) or error | X |
| `GetUsersSort` | Sort By and Sord Order | [types.Users](pkg\cybr_pam_scim\types\users.go) or error | X |
| `GetUserById` | User Id | [types.User](pkg\cybr_pam_scim\types\users.go) or error | X |
| `GetUserByFilter` | Filter Type and Filter Query | [types.User](pkg\cybr_pam_scim\types\users.go) or error | | 
| `AddUser` | [types.User](pkg\cybr_pam_scim\types\users.go) | [types.User](pkg\cybr_pam_scim\types\users.go) or error | X |
| `UpdateUser` | [types.User](pkg\cybr_pam_scim\types\users.go) | [types.User](pkg\cybr_pam_scim\types\users.go) or error | |
| `DeleteUser` | User Id | error |

1. GetUsersByFilter: Filter Query is case sensitive
2. UpdateUser: User Id must be included in the type.User struct for Update Safe permissions as the API endpoint is generated based on this info.

### Groups

| Function | Input | Output | PVWA 12.2+ Required |
|:--- |:--- |:--- |:---:|
| `GetGroups` | - | [types.Groups](pkg\cybr_pam_scim\types\groups.go) or error | |
| `GetGroupsIndex` | Start Index and Count | [types.Groups](pkg\cybr_pam_scim\types\groupss.go) or error | X |
| `GetGroupsSort` | Sort By and Sord Order | [types.Groups](pkg\cybr_pam_scim\types\groups.go) or error | X |
| `GetGroupById` | Group Id | [types.Group](pkg\cybr_pam_scim\types\groups.go) or error | X |
| `GetGroupByFilter` | Filter Type and Filter Query | [types.Group](pkg\cybr_pam_scim\types\groups.go) or error | |
| `AddGroup` | [types.Group](pkg\cybr_pam_scim\types\groups.go) | [types.Groupr](pkg\cybr_pam_scim\types\groups.go) or error | |
| `UpdateGroup` | [types.Group](pkg\cybr_pam_scim\types\groups.go) | [types.Group](pkg\cybr_pam_scim\types\groups.go) or error | X |
| `DeleteGroup` | Group Id | error |

1. GetGroupsByFilter: Filter Query is case sensitive
2. UpdateGroup: Group Id must be included in the type.Group struct for Update Safe permissions as the API endpoint is generated based on this info.

### Containers (Safes)

| Function | Input | Output | PVWA 12.2+ Required |
|:--- |:--- |:--- |:---:|
| `GetSafes` | - | [types.Containers](pkg\cybr_pam_scim\types\containers.go) or error | |
| `GetSafesIndex` | Start Index and Count | [types.Containers](pkg\cybr_pam_scim\types\containers.go) or error | X |
| `GetSafesSort` | Sort By and Sord Order | [types.Containers](pkg\cybr_pam_scim\types\containers.go) or error | X |
| `GetSafeByName` | Safe Name | [types.Container](pkg\cybr_pam_scim\types\containers.go) or error | X |
| `GetSafeByFilter` | Filter Type and Filter Query | [types.Container](pkg\cybr_pam_scim\types\containers.go) or error | |
| `AddSafe` | [types.Container](pkg\cybr_pam_scim\types\containers.go) | [types.Container](pkg\cybr_pam_scim\types\containers.go) or error | |
| `UpdateSafe` | [types.Container](pkg\cybr_pam_scim\types\containers.go) | [types.Container](pkg\cybr_pam_scim\types\containers.go) or error | X |
| `DeleteSafe` | Safe Name | error | |

Notes:
1. General: Safe functions utilize Safe Name instead of Id although both fields are typically the same in the returned types.Container struct.
2. GetSafeByFilter: Filter Query is case sensitive

### Container (Safe) Permissions

| Function | Input | Output | PVWA 12.2+ Required |
|:--- |:--- |:--- |:---:|
| `GetSafePermissions` | - | [types.ContainerPermissions](pkg\cybr_pam_scim\types\ContainerPermissions.go) or error | |
| `GetSafePermissionsIndex` | Start Index and Count | [types.ContainerPermissions](pkg\cybr_pam_scim\types\container_permissions.go) or error | X |
| `GetSafePermissionsSort` | Sort By and Sord Order | [types.ContainerPermissions](pkg\cybr_pam_scim\types\container_permissions.go) or error | X |
| `GetSafePermissionsByName` | Safe Name and User Name | [types.ContainerPermission](pkg\cybr_pam_scim\types\container_permissions.go) or error | X |
| `GetSafePermissionsByFilter` | Filter Type and Filter Query | [types.ContainerPermission](pkg\cybr_pam_scim\types\container_permissions.go) or error | |
| `AddSafePermissions` | [types.ContainerPermission](pkg\cybr_pam_scim\types\container_permissions.go) | [types.Container](pkg\cybr_pam_scim\types\container_permissions.go) or error | X |
| `UpdateSafePermissions` | [types.ContainerPermission](pkg\cybr_pam_scim\types\container_permissions.go) | [types.Container](pkg\cybr_pam_scim\types\container_permissions.go) or error | |
| `DeleteSafePermissions` | Safe Name and User or Group Name | error | |

Notes:
1. GetSafePermissionsByFilter: Filter Query is case sensitive
2. UpdateSafePermissions: User Display Name and Safe Name must be included in the type.ContainerPermissions struct for Update Safe permissions as the API endpoint is generated based on this info.
2. DeleteSafePermissions: Deletes a User or Group membership to a safe. You must provide either a User or Group Name in addition to the Safe Name.

### Privileged Data (Accounts)

| Function | Input | Output | PVWA 12.2+ Required |
|:--- |:--- |:--- |:---:|
| `GetPrivilegedData` | - | [types.PrivilegedDatas](pkg\cybr_pam_scim\types\privileged_data.go) or error | |
| `GetPrivilegedDataIndex` | Start Index and Count | [types.PrivilegedDatas](pkg\cybr_pam_scim\types\privileged_data.go) or error | X |
| `GetPrivilegedDataSort` | Sort By and Sord Order | [types.PrivilegedDatas](pkg\cybr_pam_scim\types\privileged_data.go) or error | X |
| `GetPrivilegedDataById` | Privileged Data Id | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) or error | X |
| `GetPrivilegedDataByFilter` | Filter Type and Filter Query | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) or error | |
| `AddPrivilegedData` | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) or error | X |
| `UpdatePrivilegedData` | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) or error | |
| `ModifyPrivilegedData` | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) | [types.PrivilegedData](pkg\cybr_pam_scim\types\privileged_data.go) or error | |
| `DeletePrivilegedData` | Privileged Data Id | error | |

Notes:
1. GetPrivilegedDataByFilter: Filter Query is case sensitive
2. UpdatePrivilegedData: The Privileged Data Id must be included in the types.PrivilegedData struct as the API endpoint is generated based on this info.
3. ModifyPrivilegedData: The Privileged Data Id must be included in the types.PrivilegedData struct as the API endpoint is generated based on this info.
4. ModifyPrivilegedData: The struct required to modify Privileged Data is uniqe in that it adds a nested Operations struct which contains the operations information (e.g. replace). Review the official CyberArk documentation for more info.

### General Usage Notes:
1. Filter Query is typically case sensitive.
2. Always include the object Id in structs when performing updates as it is frequently used in generating the API Endpoint.
3. Get, Get Index, Get Sort, and Update Object by Name or ID may not work with PVWA Versions below 12.2

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
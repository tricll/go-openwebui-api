# \AuthsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserAuthsAddPost**](AuthsAPI.md#AddUserAuthsAddPost) | **Post** /auths/add | Add User
[**DeleteApiKeyAuthsApiKeyDelete**](AuthsAPI.md#DeleteApiKeyAuthsApiKeyDelete) | **Delete** /auths/api_key | Delete Api Key
[**GenerateApiKeyAuthsApiKeyPost**](AuthsAPI.md#GenerateApiKeyAuthsApiKeyPost) | **Post** /auths/api_key | Generate Api Key
[**GetAdminConfigAuthsAdminConfigGet**](AuthsAPI.md#GetAdminConfigAuthsAdminConfigGet) | **Get** /auths/admin/config | Get Admin Config
[**GetAdminDetailsAuthsAdminDetailsGet**](AuthsAPI.md#GetAdminDetailsAuthsAdminDetailsGet) | **Get** /auths/admin/details | Get Admin Details
[**GetApiKeyAuthsApiKeyGet**](AuthsAPI.md#GetApiKeyAuthsApiKeyGet) | **Get** /auths/api_key | Get Api Key
[**GetLdapConfigAuthsAdminConfigLdapGet**](AuthsAPI.md#GetLdapConfigAuthsAdminConfigLdapGet) | **Get** /auths/admin/config/ldap | Get Ldap Config
[**GetLdapServerAuthsAdminConfigLdapServerGet**](AuthsAPI.md#GetLdapServerAuthsAdminConfigLdapServerGet) | **Get** /auths/admin/config/ldap/server | Get Ldap Server
[**GetSessionUserAuthsGet**](AuthsAPI.md#GetSessionUserAuthsGet) | **Get** /auths/ | Get Session User
[**LdapAuthAuthsLdapPost**](AuthsAPI.md#LdapAuthAuthsLdapPost) | **Post** /auths/ldap | Ldap Auth
[**SigninAuthsSigninPost**](AuthsAPI.md#SigninAuthsSigninPost) | **Post** /auths/signin | Signin
[**SignoutAuthsSignoutGet**](AuthsAPI.md#SignoutAuthsSignoutGet) | **Get** /auths/signout | Signout
[**SignupAuthsSignupPost**](AuthsAPI.md#SignupAuthsSignupPost) | **Post** /auths/signup | Signup
[**UpdateAdminConfigAuthsAdminConfigPost**](AuthsAPI.md#UpdateAdminConfigAuthsAdminConfigPost) | **Post** /auths/admin/config | Update Admin Config
[**UpdateLdapConfigAuthsAdminConfigLdapPost**](AuthsAPI.md#UpdateLdapConfigAuthsAdminConfigLdapPost) | **Post** /auths/admin/config/ldap | Update Ldap Config
[**UpdateLdapServerAuthsAdminConfigLdapServerPost**](AuthsAPI.md#UpdateLdapServerAuthsAdminConfigLdapServerPost) | **Post** /auths/admin/config/ldap/server | Update Ldap Server
[**UpdatePasswordAuthsUpdatePasswordPost**](AuthsAPI.md#UpdatePasswordAuthsUpdatePasswordPost) | **Post** /auths/update/password | Update Password
[**UpdateProfileAuthsUpdateProfilePost**](AuthsAPI.md#UpdateProfileAuthsUpdateProfilePost) | **Post** /auths/update/profile | Update Profile



## AddUserAuthsAddPost

> SigninResponse AddUserAuthsAddPost(ctx).AddUserForm(addUserForm).Execute()

Add User

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	addUserForm := *openapiclient.NewAddUserForm("Name_example", "Email_example", "Password_example") // AddUserForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.AddUserAuthsAddPost(context.Background()).AddUserForm(addUserForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.AddUserAuthsAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserAuthsAddPost`: SigninResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.AddUserAuthsAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserAuthsAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserForm** | [**AddUserForm**](AddUserForm.md) |  | 

### Return type

[**SigninResponse**](SigninResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKeyAuthsApiKeyDelete

> bool DeleteApiKeyAuthsApiKeyDelete(ctx).Execute()

Delete Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.DeleteApiKeyAuthsApiKeyDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.DeleteApiKeyAuthsApiKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiKeyAuthsApiKeyDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.DeleteApiKeyAuthsApiKeyDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyAuthsApiKeyDeleteRequest struct via the builder pattern


### Return type

**bool**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateApiKeyAuthsApiKeyPost

> ApiKey GenerateApiKeyAuthsApiKeyPost(ctx).Execute()

Generate Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GenerateApiKeyAuthsApiKeyPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GenerateApiKeyAuthsApiKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateApiKeyAuthsApiKeyPost`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GenerateApiKeyAuthsApiKeyPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateApiKeyAuthsApiKeyPostRequest struct via the builder pattern


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminConfigAuthsAdminConfigGet

> interface{} GetAdminConfigAuthsAdminConfigGet(ctx).Execute()

Get Admin Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GetAdminConfigAuthsAdminConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetAdminConfigAuthsAdminConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminConfigAuthsAdminConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetAdminConfigAuthsAdminConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminConfigAuthsAdminConfigGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminDetailsAuthsAdminDetailsGet

> interface{} GetAdminDetailsAuthsAdminDetailsGet(ctx).Execute()

Get Admin Details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GetAdminDetailsAuthsAdminDetailsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetAdminDetailsAuthsAdminDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminDetailsAuthsAdminDetailsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetAdminDetailsAuthsAdminDetailsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminDetailsAuthsAdminDetailsGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyAuthsApiKeyGet

> ApiKey GetApiKeyAuthsApiKeyGet(ctx).Execute()

Get Api Key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GetApiKeyAuthsApiKeyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetApiKeyAuthsApiKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeyAuthsApiKeyGet`: ApiKey
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetApiKeyAuthsApiKeyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyAuthsApiKeyGetRequest struct via the builder pattern


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapConfigAuthsAdminConfigLdapGet

> interface{} GetLdapConfigAuthsAdminConfigLdapGet(ctx).Execute()

Get Ldap Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GetLdapConfigAuthsAdminConfigLdapGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetLdapConfigAuthsAdminConfigLdapGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapConfigAuthsAdminConfigLdapGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetLdapConfigAuthsAdminConfigLdapGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapConfigAuthsAdminConfigLdapGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapServerAuthsAdminConfigLdapServerGet

> LdapServerConfig GetLdapServerAuthsAdminConfigLdapServerGet(ctx).Execute()

Get Ldap Server

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GetLdapServerAuthsAdminConfigLdapServerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetLdapServerAuthsAdminConfigLdapServerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapServerAuthsAdminConfigLdapServerGet`: LdapServerConfig
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetLdapServerAuthsAdminConfigLdapServerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapServerAuthsAdminConfigLdapServerGetRequest struct via the builder pattern


### Return type

[**LdapServerConfig**](LdapServerConfig.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionUserAuthsGet

> SessionUserResponse GetSessionUserAuthsGet(ctx).Execute()

Get Session User

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.GetSessionUserAuthsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.GetSessionUserAuthsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionUserAuthsGet`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.GetSessionUserAuthsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionUserAuthsGetRequest struct via the builder pattern


### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapAuthAuthsLdapPost

> SigninResponse LdapAuthAuthsLdapPost(ctx).LdapForm(ldapForm).Execute()

Ldap Auth

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	ldapForm := *openapiclient.NewLdapForm("User_example", "Password_example") // LdapForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.LdapAuthAuthsLdapPost(context.Background()).LdapForm(ldapForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.LdapAuthAuthsLdapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LdapAuthAuthsLdapPost`: SigninResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.LdapAuthAuthsLdapPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLdapAuthAuthsLdapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapForm** | [**LdapForm**](LdapForm.md) |  | 

### Return type

[**SigninResponse**](SigninResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SigninAuthsSigninPost

> SessionUserResponse SigninAuthsSigninPost(ctx).SigninForm(signinForm).Execute()

Signin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	signinForm := *openapiclient.NewSigninForm("Email_example", "Password_example") // SigninForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.SigninAuthsSigninPost(context.Background()).SigninForm(signinForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.SigninAuthsSigninPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SigninAuthsSigninPost`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.SigninAuthsSigninPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSigninAuthsSigninPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signinForm** | [**SigninForm**](SigninForm.md) |  | 

### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignoutAuthsSignoutGet

> interface{} SignoutAuthsSignoutGet(ctx).Execute()

Signout

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.SignoutAuthsSignoutGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.SignoutAuthsSignoutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignoutAuthsSignoutGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.SignoutAuthsSignoutGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSignoutAuthsSignoutGetRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupAuthsSignupPost

> SessionUserResponse SignupAuthsSignupPost(ctx).SignupForm(signupForm).Execute()

Signup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	signupForm := *openapiclient.NewSignupForm("Name_example", "Email_example", "Password_example") // SignupForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.SignupAuthsSignupPost(context.Background()).SignupForm(signupForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.SignupAuthsSignupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignupAuthsSignupPost`: SessionUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.SignupAuthsSignupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupAuthsSignupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupForm** | [**SignupForm**](SignupForm.md) |  | 

### Return type

[**SessionUserResponse**](SessionUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdminConfigAuthsAdminConfigPost

> interface{} UpdateAdminConfigAuthsAdminConfigPost(ctx).AdminConfig(adminConfig).Execute()

Update Admin Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	adminConfig := *openapiclient.NewAdminConfig(false, false, false, "DEFAULT_USER_ROLE_example", "JWT_EXPIRES_IN_example", false, false) // AdminConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateAdminConfigAuthsAdminConfigPost(context.Background()).AdminConfig(adminConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateAdminConfigAuthsAdminConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdminConfigAuthsAdminConfigPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateAdminConfigAuthsAdminConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdminConfigAuthsAdminConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminConfig** | [**AdminConfig**](AdminConfig.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapConfigAuthsAdminConfigLdapPost

> interface{} UpdateLdapConfigAuthsAdminConfigLdapPost(ctx).LdapConfigForm(ldapConfigForm).Execute()

Update Ldap Config

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	ldapConfigForm := *openapiclient.NewLdapConfigForm() // LdapConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateLdapConfigAuthsAdminConfigLdapPost(context.Background()).LdapConfigForm(ldapConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateLdapConfigAuthsAdminConfigLdapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLdapConfigAuthsAdminConfigLdapPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateLdapConfigAuthsAdminConfigLdapPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapConfigAuthsAdminConfigLdapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigForm** | [**LdapConfigForm**](LdapConfigForm.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapServerAuthsAdminConfigLdapServerPost

> interface{} UpdateLdapServerAuthsAdminConfigLdapServerPost(ctx).LdapServerConfig(ldapServerConfig).Execute()

Update Ldap Server

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	ldapServerConfig := *openapiclient.NewLdapServerConfig("Label_example", "Host_example", "AppDn_example", "AppDnPassword_example", "SearchBase_example") // LdapServerConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateLdapServerAuthsAdminConfigLdapServerPost(context.Background()).LdapServerConfig(ldapServerConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateLdapServerAuthsAdminConfigLdapServerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLdapServerAuthsAdminConfigLdapServerPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateLdapServerAuthsAdminConfigLdapServerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapServerAuthsAdminConfigLdapServerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapServerConfig** | [**LdapServerConfig**](LdapServerConfig.md) |  | 

### Return type

**interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordAuthsUpdatePasswordPost

> bool UpdatePasswordAuthsUpdatePasswordPost(ctx).UpdatePasswordForm(updatePasswordForm).Execute()

Update Password

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	updatePasswordForm := *openapiclient.NewUpdatePasswordForm("Password_example", "NewPassword_example") // UpdatePasswordForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdatePasswordAuthsUpdatePasswordPost(context.Background()).UpdatePasswordForm(updatePasswordForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdatePasswordAuthsUpdatePasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePasswordAuthsUpdatePasswordPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdatePasswordAuthsUpdatePasswordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordAuthsUpdatePasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePasswordForm** | [**UpdatePasswordForm**](UpdatePasswordForm.md) |  | 

### Return type

**bool**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfileAuthsUpdateProfilePost

> OpenWebuiAppsWebuiModelsAuthsUserResponse UpdateProfileAuthsUpdateProfilePost(ctx).UpdateProfileForm(updateProfileForm).Execute()

Update Profile

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tricll/go-openwebui-api"
)

func main() {
	updateProfileForm := *openapiclient.NewUpdateProfileForm("ProfileImageUrl_example", "Name_example") // UpdateProfileForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthsAPI.UpdateProfileAuthsUpdateProfilePost(context.Background()).UpdateProfileForm(updateProfileForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthsAPI.UpdateProfileAuthsUpdateProfilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfileAuthsUpdateProfilePost`: OpenWebuiAppsWebuiModelsAuthsUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthsAPI.UpdateProfileAuthsUpdateProfilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileAuthsUpdateProfilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileForm** | [**UpdateProfileForm**](UpdateProfileForm.md) |  | 

### Return type

[**OpenWebuiAppsWebuiModelsAuthsUserResponse**](OpenWebuiAppsWebuiModelsAuthsUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


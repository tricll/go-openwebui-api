# \UsersAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserByIdUsersUserIdDelete**](UsersAPI.md#DeleteUserByIdUsersUserIdDelete) | **Delete** /users/{user_id} | Delete User By Id
[**GetUserByIdUsersUserIdGet**](UsersAPI.md#GetUserByIdUsersUserIdGet) | **Get** /users/{user_id} | Get User By Id
[**GetUserGroupsUsersGroupsGet**](UsersAPI.md#GetUserGroupsUsersGroupsGet) | **Get** /users/groups | Get User Groups
[**GetUserInfoBySessionUserUsersUserInfoGet**](UsersAPI.md#GetUserInfoBySessionUserUsersUserInfoGet) | **Get** /users/user/info | Get User Info By Session User
[**GetUserPermissionsUsersDefaultPermissionsGet**](UsersAPI.md#GetUserPermissionsUsersDefaultPermissionsGet) | **Get** /users/default/permissions | Get User Permissions
[**GetUserPermissisionsUsersPermissionsGet**](UsersAPI.md#GetUserPermissisionsUsersPermissionsGet) | **Get** /users/permissions | Get User Permissisions
[**GetUserSettingsBySessionUserUsersUserSettingsGet**](UsersAPI.md#GetUserSettingsBySessionUserUsersUserSettingsGet) | **Get** /users/user/settings | Get User Settings By Session User
[**GetUsersUsersGet**](UsersAPI.md#GetUsersUsersGet) | **Get** /users/ | Get Users
[**UpdateUserByIdUsersUserIdUpdatePost**](UsersAPI.md#UpdateUserByIdUsersUserIdUpdatePost) | **Post** /users/{user_id}/update | Update User By Id
[**UpdateUserInfoBySessionUserUsersUserInfoUpdatePost**](UsersAPI.md#UpdateUserInfoBySessionUserUsersUserInfoUpdatePost) | **Post** /users/user/info/update | Update User Info By Session User
[**UpdateUserPermissionsUsersDefaultPermissionsPost**](UsersAPI.md#UpdateUserPermissionsUsersDefaultPermissionsPost) | **Post** /users/default/permissions | Update User Permissions
[**UpdateUserRoleUsersUpdateRolePost**](UsersAPI.md#UpdateUserRoleUsersUpdateRolePost) | **Post** /users/update/role | Update User Role
[**UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost**](UsersAPI.md#UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost) | **Post** /users/user/settings/update | Update User Settings By Session User



## DeleteUserByIdUsersUserIdDelete

> bool DeleteUserByIdUsersUserIdDelete(ctx, userId).Execute()

Delete User By Id

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserByIdUsersUserIdDelete(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserByIdUsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserByIdUsersUserIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserByIdUsersUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByIdUsersUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetUserByIdUsersUserIdGet

> OpenWebuiAppsWebuiRoutersUsersUserResponse GetUserByIdUsersUserIdGet(ctx, userId).Execute()

Get User By Id

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserByIdUsersUserIdGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserByIdUsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByIdUsersUserIdGet`: OpenWebuiAppsWebuiRoutersUsersUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserByIdUsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdUsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpenWebuiAppsWebuiRoutersUsersUserResponse**](OpenWebuiAppsWebuiRoutersUsersUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupsUsersGroupsGet

> interface{} GetUserGroupsUsersGroupsGet(ctx).Execute()

Get User Groups

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
	resp, r, err := apiClient.UsersAPI.GetUserGroupsUsersGroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupsUsersGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupsUsersGroupsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupsUsersGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsUsersGroupsGetRequest struct via the builder pattern


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


## GetUserInfoBySessionUserUsersUserInfoGet

> map[string]interface{} GetUserInfoBySessionUserUsersUserInfoGet(ctx).Execute()

Get User Info By Session User

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
	resp, r, err := apiClient.UsersAPI.GetUserInfoBySessionUserUsersUserInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserInfoBySessionUserUsersUserInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfoBySessionUserUsersUserInfoGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserInfoBySessionUserUsersUserInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoBySessionUserUsersUserInfoGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPermissionsUsersDefaultPermissionsGet

> interface{} GetUserPermissionsUsersDefaultPermissionsGet(ctx).Execute()

Get User Permissions

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
	resp, r, err := apiClient.UsersAPI.GetUserPermissionsUsersDefaultPermissionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPermissionsUsersDefaultPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissionsUsersDefaultPermissionsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPermissionsUsersDefaultPermissionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsUsersDefaultPermissionsGetRequest struct via the builder pattern


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


## GetUserPermissisionsUsersPermissionsGet

> interface{} GetUserPermissisionsUsersPermissionsGet(ctx).Execute()

Get User Permissisions

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
	resp, r, err := apiClient.UsersAPI.GetUserPermissisionsUsersPermissionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserPermissisionsUsersPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissisionsUsersPermissionsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserPermissisionsUsersPermissionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissisionsUsersPermissionsGetRequest struct via the builder pattern


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


## GetUserSettingsBySessionUserUsersUserSettingsGet

> UserSettings GetUserSettingsBySessionUserUsersUserSettingsGet(ctx).Execute()

Get User Settings By Session User

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
	resp, r, err := apiClient.UsersAPI.GetUserSettingsBySessionUserUsersUserSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSettingsBySessionUserUsersUserSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSettingsBySessionUserUsersUserSettingsGet`: UserSettings
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSettingsBySessionUserUsersUserSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsBySessionUserUsersUserSettingsGetRequest struct via the builder pattern


### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersUsersGet

> []UserModel GetUsersUsersGet(ctx).Skip(skip).Limit(limit).Execute()

Get Users

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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersUsersGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersUsersGet`: []UserModel
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]UserModel**](UserModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserByIdUsersUserIdUpdatePost

> UserModel UpdateUserByIdUsersUserIdUpdatePost(ctx, userId).UserUpdateForm(userUpdateForm).Execute()

Update User By Id

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
	userId := "userId_example" // string | 
	userUpdateForm := *openapiclient.NewUserUpdateForm("Name_example", "Email_example", "ProfileImageUrl_example") // UserUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserByIdUsersUserIdUpdatePost(context.Background(), userId).UserUpdateForm(userUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserByIdUsersUserIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserByIdUsersUserIdUpdatePost`: UserModel
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserByIdUsersUserIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserByIdUsersUserIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateForm** | [**UserUpdateForm**](UserUpdateForm.md) |  | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserInfoBySessionUserUsersUserInfoUpdatePost

> map[string]interface{} UpdateUserInfoBySessionUserUsersUserInfoUpdatePost(ctx).Body(body).Execute()

Update User Info By Session User

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserInfoBySessionUserUsersUserInfoUpdatePost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserInfoBySessionUserUsersUserInfoUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserInfoBySessionUserUsersUserInfoUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserInfoBySessionUserUsersUserInfoUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserInfoBySessionUserUsersUserInfoUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPermissionsUsersDefaultPermissionsPost

> interface{} UpdateUserPermissionsUsersDefaultPermissionsPost(ctx).UserPermissions(userPermissions).Execute()

Update User Permissions

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
	userPermissions := *openapiclient.NewUserPermissions(*openapiclient.NewWorkspacePermissions(false, false, false, false), *openapiclient.NewChatPermissions(false, false, false, false)) // UserPermissions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserPermissionsUsersDefaultPermissionsPost(context.Background()).UserPermissions(userPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserPermissionsUsersDefaultPermissionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserPermissionsUsersDefaultPermissionsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserPermissionsUsersDefaultPermissionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPermissionsUsersDefaultPermissionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPermissions** | [**UserPermissions**](UserPermissions.md) |  | 

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


## UpdateUserRoleUsersUpdateRolePost

> UserModel UpdateUserRoleUsersUpdateRolePost(ctx).UserRoleUpdateForm(userRoleUpdateForm).Execute()

Update User Role

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
	userRoleUpdateForm := *openapiclient.NewUserRoleUpdateForm("Id_example", "Role_example") // UserRoleUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserRoleUsersUpdateRolePost(context.Background()).UserRoleUpdateForm(userRoleUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserRoleUsersUpdateRolePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserRoleUsersUpdateRolePost`: UserModel
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserRoleUsersUpdateRolePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRoleUsersUpdateRolePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRoleUpdateForm** | [**UserRoleUpdateForm**](UserRoleUpdateForm.md) |  | 

### Return type

[**UserModel**](UserModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost

> UserSettings UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost(ctx).UserSettings(userSettings).Execute()

Update User Settings By Session User

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
	userSettings := *openapiclient.NewUserSettings() // UserSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost(context.Background()).UserSettings(userSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost`: UserSettings
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserSettingsBySessionUserUsersUserSettingsUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSettingsBySessionUserUsersUserSettingsUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userSettings** | [**UserSettings**](UserSettings.md) |  | 

### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


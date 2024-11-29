# \GroupsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewFunctionGroupsCreatePost**](GroupsAPI.md#CreateNewFunctionGroupsCreatePost) | **Post** /groups/create | Create New Function
[**DeleteGroupByIdGroupsIdIdDeleteDelete**](GroupsAPI.md#DeleteGroupByIdGroupsIdIdDeleteDelete) | **Delete** /groups/id/{id}/delete | Delete Group By Id
[**GetGroupByIdGroupsIdIdGet**](GroupsAPI.md#GetGroupByIdGroupsIdIdGet) | **Get** /groups/id/{id} | Get Group By Id
[**GetGroupsGroupsGet**](GroupsAPI.md#GetGroupsGroupsGet) | **Get** /groups/ | Get Groups
[**UpdateGroupByIdGroupsIdIdUpdatePost**](GroupsAPI.md#UpdateGroupByIdGroupsIdIdUpdatePost) | **Post** /groups/id/{id}/update | Update Group By Id



## CreateNewFunctionGroupsCreatePost

> GroupResponse CreateNewFunctionGroupsCreatePost(ctx).GroupForm(groupForm).Execute()

Create New Function

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
	groupForm := *openapiclient.NewGroupForm("Name_example", "Description_example") // GroupForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateNewFunctionGroupsCreatePost(context.Background()).GroupForm(groupForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateNewFunctionGroupsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewFunctionGroupsCreatePost`: GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateNewFunctionGroupsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewFunctionGroupsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupForm** | [**GroupForm**](GroupForm.md) |  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupByIdGroupsIdIdDeleteDelete

> bool DeleteGroupByIdGroupsIdIdDeleteDelete(ctx, id).Execute()

Delete Group By Id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.DeleteGroupByIdGroupsIdIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroupByIdGroupsIdIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupByIdGroupsIdIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.DeleteGroupByIdGroupsIdIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupByIdGroupsIdIdDeleteDeleteRequest struct via the builder pattern


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


## GetGroupByIdGroupsIdIdGet

> GroupResponse GetGroupByIdGroupsIdIdGet(ctx, id).Execute()

Get Group By Id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroupByIdGroupsIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupByIdGroupsIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupByIdGroupsIdIdGet`: GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupByIdGroupsIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByIdGroupsIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsGroupsGet

> []GroupResponse GetGroupsGroupsGet(ctx).Execute()

Get Groups

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
	resp, r, err := apiClient.GroupsAPI.GetGroupsGroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroupsGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsGroupsGet`: []GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroupsGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsGroupsGetRequest struct via the builder pattern


### Return type

[**[]GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupByIdGroupsIdIdUpdatePost

> GroupResponse UpdateGroupByIdGroupsIdIdUpdatePost(ctx, id).GroupUpdateForm(groupUpdateForm).Execute()

Update Group By Id

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
	id := "id_example" // string | 
	groupUpdateForm := *openapiclient.NewGroupUpdateForm("Name_example", "Description_example") // GroupUpdateForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroupByIdGroupsIdIdUpdatePost(context.Background(), id).GroupUpdateForm(groupUpdateForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroupByIdGroupsIdIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupByIdGroupsIdIdUpdatePost`: GroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroupByIdGroupsIdIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupByIdGroupsIdIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateForm** | [**GroupUpdateForm**](GroupUpdateForm.md) |  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \FoldersAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolderFoldersPost**](FoldersAPI.md#CreateFolderFoldersPost) | **Post** /folders/ | Create Folder
[**DeleteFolderByIdFoldersIdDelete**](FoldersAPI.md#DeleteFolderByIdFoldersIdDelete) | **Delete** /folders/{id} | Delete Folder By Id
[**GetFolderByIdFoldersIdGet**](FoldersAPI.md#GetFolderByIdFoldersIdGet) | **Get** /folders/{id} | Get Folder By Id
[**GetFoldersFoldersGet**](FoldersAPI.md#GetFoldersFoldersGet) | **Get** /folders/ | Get Folders
[**UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost**](FoldersAPI.md#UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost) | **Post** /folders/{id}/update/expanded | Update Folder Is Expanded By Id
[**UpdateFolderNameByIdFoldersIdUpdatePost**](FoldersAPI.md#UpdateFolderNameByIdFoldersIdUpdatePost) | **Post** /folders/{id}/update | Update Folder Name By Id
[**UpdateFolderParentIdByIdFoldersIdUpdateParentPost**](FoldersAPI.md#UpdateFolderParentIdByIdFoldersIdUpdateParentPost) | **Post** /folders/{id}/update/parent | Update Folder Parent Id By Id



## CreateFolderFoldersPost

> interface{} CreateFolderFoldersPost(ctx).FolderForm(folderForm).Execute()

Create Folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderForm := *openapiclient.NewFolderForm("Name_example") // FolderForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.CreateFolderFoldersPost(context.Background()).FolderForm(folderForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.CreateFolderFoldersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolderFoldersPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.CreateFolderFoldersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderFoldersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderForm** | [**FolderForm**](FolderForm.md) |  | 

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


## DeleteFolderByIdFoldersIdDelete

> interface{} DeleteFolderByIdFoldersIdDelete(ctx, id).Execute()

Delete Folder By Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.DeleteFolderByIdFoldersIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.DeleteFolderByIdFoldersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolderByIdFoldersIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.DeleteFolderByIdFoldersIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderByIdFoldersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetFolderByIdFoldersIdGet

> FolderModel GetFolderByIdFoldersIdGet(ctx, id).Execute()

Get Folder By Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFolderByIdFoldersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFolderByIdFoldersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderByIdFoldersIdGet`: FolderModel
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFolderByIdFoldersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByIdFoldersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderModel**](FolderModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFoldersFoldersGet

> []FolderModel GetFoldersFoldersGet(ctx).Execute()

Get Folders

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.GetFoldersFoldersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.GetFoldersFoldersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFoldersFoldersGet`: []FolderModel
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.GetFoldersFoldersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFoldersFoldersGetRequest struct via the builder pattern


### Return type

[**[]FolderModel**](FolderModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost

> interface{} UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost(ctx, id).FolderIsExpandedForm(folderIsExpandedForm).Execute()

Update Folder Is Expanded By Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	folderIsExpandedForm := *openapiclient.NewFolderIsExpandedForm(false) // FolderIsExpandedForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost(context.Background(), id).FolderIsExpandedForm(folderIsExpandedForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderIsExpandedByIdFoldersIdUpdateExpandedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderIsExpandedForm** | [**FolderIsExpandedForm**](FolderIsExpandedForm.md) |  | 

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


## UpdateFolderNameByIdFoldersIdUpdatePost

> interface{} UpdateFolderNameByIdFoldersIdUpdatePost(ctx, id).FolderForm(folderForm).Execute()

Update Folder Name By Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	folderForm := *openapiclient.NewFolderForm("Name_example") // FolderForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolderNameByIdFoldersIdUpdatePost(context.Background(), id).FolderForm(folderForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolderNameByIdFoldersIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderNameByIdFoldersIdUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolderNameByIdFoldersIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderNameByIdFoldersIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderForm** | [**FolderForm**](FolderForm.md) |  | 

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


## UpdateFolderParentIdByIdFoldersIdUpdateParentPost

> interface{} UpdateFolderParentIdByIdFoldersIdUpdateParentPost(ctx, id).FolderParentIdForm(folderParentIdForm).Execute()

Update Folder Parent Id By Id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	folderParentIdForm := *openapiclient.NewFolderParentIdForm() // FolderParentIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.UpdateFolderParentIdByIdFoldersIdUpdateParentPost(context.Background(), id).FolderParentIdForm(folderParentIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.UpdateFolderParentIdByIdFoldersIdUpdateParentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderParentIdByIdFoldersIdUpdateParentPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.UpdateFolderParentIdByIdFoldersIdUpdateParentPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderParentIdByIdFoldersIdUpdateParentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderParentIdForm** | [**FolderParentIdForm**](FolderParentIdForm.md) |  | 

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


# \KnowledgeAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToKnowledgeByIdKnowledgeIdFileAddPost**](KnowledgeAPI.md#AddFileToKnowledgeByIdKnowledgeIdFileAddPost) | **Post** /knowledge/{id}/file/add | Add File To Knowledge By Id
[**CreateNewKnowledgeKnowledgeCreatePost**](KnowledgeAPI.md#CreateNewKnowledgeKnowledgeCreatePost) | **Post** /knowledge/create | Create New Knowledge
[**DeleteKnowledgeByIdKnowledgeIdDeleteDelete**](KnowledgeAPI.md#DeleteKnowledgeByIdKnowledgeIdDeleteDelete) | **Delete** /knowledge/{id}/delete | Delete Knowledge By Id
[**GetKnowledgeByIdKnowledgeIdGet**](KnowledgeAPI.md#GetKnowledgeByIdKnowledgeIdGet) | **Get** /knowledge/{id} | Get Knowledge By Id
[**GetKnowledgeKnowledgeGet**](KnowledgeAPI.md#GetKnowledgeKnowledgeGet) | **Get** /knowledge/ | Get Knowledge
[**GetKnowledgeListKnowledgeListGet**](KnowledgeAPI.md#GetKnowledgeListKnowledgeListGet) | **Get** /knowledge/list | Get Knowledge List
[**RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost**](KnowledgeAPI.md#RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost) | **Post** /knowledge/{id}/file/remove | Remove File From Knowledge By Id
[**ResetKnowledgeByIdKnowledgeIdResetPost**](KnowledgeAPI.md#ResetKnowledgeByIdKnowledgeIdResetPost) | **Post** /knowledge/{id}/reset | Reset Knowledge By Id
[**UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost**](KnowledgeAPI.md#UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost) | **Post** /knowledge/{id}/file/update | Update File From Knowledge By Id
[**UpdateKnowledgeByIdKnowledgeIdUpdatePost**](KnowledgeAPI.md#UpdateKnowledgeByIdKnowledgeIdUpdatePost) | **Post** /knowledge/{id}/update | Update Knowledge By Id



## AddFileToKnowledgeByIdKnowledgeIdFileAddPost

> KnowledgeFilesResponse AddFileToKnowledgeByIdKnowledgeIdFileAddPost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Add File To Knowledge By Id

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
	knowledgeFileIdForm := *openapiclient.NewKnowledgeFileIdForm("FileId_example") // KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.AddFileToKnowledgeByIdKnowledgeIdFileAddPost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.AddFileToKnowledgeByIdKnowledgeIdFileAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFileToKnowledgeByIdKnowledgeIdFileAddPost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.AddFileToKnowledgeByIdKnowledgeIdFileAddPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToKnowledgeByIdKnowledgeIdFileAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewKnowledgeKnowledgeCreatePost

> KnowledgeResponse CreateNewKnowledgeKnowledgeCreatePost(ctx).KnowledgeForm(knowledgeForm).Execute()

Create New Knowledge

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
	knowledgeForm := *openapiclient.NewKnowledgeForm("Name_example", "Description_example") // KnowledgeForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.CreateNewKnowledgeKnowledgeCreatePost(context.Background()).KnowledgeForm(knowledgeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.CreateNewKnowledgeKnowledgeCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewKnowledgeKnowledgeCreatePost`: KnowledgeResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.CreateNewKnowledgeKnowledgeCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewKnowledgeKnowledgeCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knowledgeForm** | [**KnowledgeForm**](KnowledgeForm.md) |  | 

### Return type

[**KnowledgeResponse**](KnowledgeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKnowledgeByIdKnowledgeIdDeleteDelete

> bool DeleteKnowledgeByIdKnowledgeIdDeleteDelete(ctx, id).Execute()

Delete Knowledge By Id

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
	resp, r, err := apiClient.KnowledgeAPI.DeleteKnowledgeByIdKnowledgeIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.DeleteKnowledgeByIdKnowledgeIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKnowledgeByIdKnowledgeIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.DeleteKnowledgeByIdKnowledgeIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKnowledgeByIdKnowledgeIdDeleteDeleteRequest struct via the builder pattern


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


## GetKnowledgeByIdKnowledgeIdGet

> KnowledgeFilesResponse GetKnowledgeByIdKnowledgeIdGet(ctx, id).Execute()

Get Knowledge By Id

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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeByIdKnowledgeIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeByIdKnowledgeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeByIdKnowledgeIdGet`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeByIdKnowledgeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeByIdKnowledgeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeKnowledgeGet

> []KnowledgeUserResponse GetKnowledgeKnowledgeGet(ctx).Execute()

Get Knowledge

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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeKnowledgeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeKnowledgeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeKnowledgeGet`: []KnowledgeUserResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeKnowledgeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeKnowledgeGetRequest struct via the builder pattern


### Return type

[**[]KnowledgeUserResponse**](KnowledgeUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeListKnowledgeListGet

> []KnowledgeUserResponse GetKnowledgeListKnowledgeListGet(ctx).Execute()

Get Knowledge List

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
	resp, r, err := apiClient.KnowledgeAPI.GetKnowledgeListKnowledgeListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.GetKnowledgeListKnowledgeListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeListKnowledgeListGet`: []KnowledgeUserResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.GetKnowledgeListKnowledgeListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeListKnowledgeListGetRequest struct via the builder pattern


### Return type

[**[]KnowledgeUserResponse**](KnowledgeUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost

> KnowledgeFilesResponse RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Remove File From Knowledge By Id

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
	knowledgeFileIdForm := *openapiclient.NewKnowledgeFileIdForm("FileId_example") // KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.RemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFileFromKnowledgeByIdKnowledgeIdFileRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetKnowledgeByIdKnowledgeIdResetPost

> KnowledgeResponse ResetKnowledgeByIdKnowledgeIdResetPost(ctx, id).Execute()

Reset Knowledge By Id

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
	resp, r, err := apiClient.KnowledgeAPI.ResetKnowledgeByIdKnowledgeIdResetPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.ResetKnowledgeByIdKnowledgeIdResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetKnowledgeByIdKnowledgeIdResetPost`: KnowledgeResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.ResetKnowledgeByIdKnowledgeIdResetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetKnowledgeByIdKnowledgeIdResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KnowledgeResponse**](KnowledgeResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost

> KnowledgeFilesResponse UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost(ctx, id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()

Update File From Knowledge By Id

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
	knowledgeFileIdForm := *openapiclient.NewKnowledgeFileIdForm("FileId_example") // KnowledgeFileIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost(context.Background(), id).KnowledgeFileIdForm(knowledgeFileIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.UpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileFromKnowledgeByIdKnowledgeIdFileUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeFileIdForm** | [**KnowledgeFileIdForm**](KnowledgeFileIdForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKnowledgeByIdKnowledgeIdUpdatePost

> KnowledgeFilesResponse UpdateKnowledgeByIdKnowledgeIdUpdatePost(ctx, id).KnowledgeForm(knowledgeForm).Execute()

Update Knowledge By Id

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
	knowledgeForm := *openapiclient.NewKnowledgeForm("Name_example", "Description_example") // KnowledgeForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeAPI.UpdateKnowledgeByIdKnowledgeIdUpdatePost(context.Background(), id).KnowledgeForm(knowledgeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeAPI.UpdateKnowledgeByIdKnowledgeIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKnowledgeByIdKnowledgeIdUpdatePost`: KnowledgeFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeAPI.UpdateKnowledgeByIdKnowledgeIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKnowledgeByIdKnowledgeIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **knowledgeForm** | [**KnowledgeForm**](KnowledgeForm.md) |  | 

### Return type

[**KnowledgeFilesResponse**](KnowledgeFilesResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


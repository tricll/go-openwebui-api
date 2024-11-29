# \MemoriesAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemoryMemoriesAddPost**](MemoriesAPI.md#AddMemoryMemoriesAddPost) | **Post** /memories/add | Add Memory
[**DeleteMemoryByIdMemoriesMemoryIdDelete**](MemoriesAPI.md#DeleteMemoryByIdMemoriesMemoryIdDelete) | **Delete** /memories/{memory_id} | Delete Memory By Id
[**DeleteMemoryByUserIdMemoriesDeleteUserDelete**](MemoriesAPI.md#DeleteMemoryByUserIdMemoriesDeleteUserDelete) | **Delete** /memories/delete/user | Delete Memory By User Id
[**GetEmbeddingsMemoriesEfGet**](MemoriesAPI.md#GetEmbeddingsMemoriesEfGet) | **Get** /memories/ef | Get Embeddings
[**GetMemoriesMemoriesGet**](MemoriesAPI.md#GetMemoriesMemoriesGet) | **Get** /memories/ | Get Memories
[**QueryMemoryMemoriesQueryPost**](MemoriesAPI.md#QueryMemoryMemoriesQueryPost) | **Post** /memories/query | Query Memory
[**ResetMemoryFromVectorDbMemoriesResetPost**](MemoriesAPI.md#ResetMemoryFromVectorDbMemoriesResetPost) | **Post** /memories/reset | Reset Memory From Vector Db
[**UpdateMemoryByIdMemoriesMemoryIdUpdatePost**](MemoriesAPI.md#UpdateMemoryByIdMemoriesMemoryIdUpdatePost) | **Post** /memories/{memory_id}/update | Update Memory By Id



## AddMemoryMemoriesAddPost

> MemoryModel AddMemoryMemoriesAddPost(ctx).AddMemoryForm(addMemoryForm).Execute()

Add Memory

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
	addMemoryForm := *openapiclient.NewAddMemoryForm("Content_example") // AddMemoryForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.AddMemoryMemoriesAddPost(context.Background()).AddMemoryForm(addMemoryForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.AddMemoryMemoriesAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMemoryMemoriesAddPost`: MemoryModel
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.AddMemoryMemoriesAddPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMemoryMemoriesAddPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addMemoryForm** | [**AddMemoryForm**](AddMemoryForm.md) |  | 

### Return type

[**MemoryModel**](MemoryModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemoryByIdMemoriesMemoryIdDelete

> bool DeleteMemoryByIdMemoriesMemoryIdDelete(ctx, memoryId).Execute()

Delete Memory By Id

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
	memoryId := "memoryId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.DeleteMemoryByIdMemoriesMemoryIdDelete(context.Background(), memoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.DeleteMemoryByIdMemoriesMemoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMemoryByIdMemoriesMemoryIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.DeleteMemoryByIdMemoriesMemoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemoryByIdMemoriesMemoryIdDeleteRequest struct via the builder pattern


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


## DeleteMemoryByUserIdMemoriesDeleteUserDelete

> bool DeleteMemoryByUserIdMemoriesDeleteUserDelete(ctx).Execute()

Delete Memory By User Id

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
	resp, r, err := apiClient.MemoriesAPI.DeleteMemoryByUserIdMemoriesDeleteUserDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.DeleteMemoryByUserIdMemoriesDeleteUserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMemoryByUserIdMemoriesDeleteUserDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.DeleteMemoryByUserIdMemoriesDeleteUserDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemoryByUserIdMemoriesDeleteUserDeleteRequest struct via the builder pattern


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


## GetEmbeddingsMemoriesEfGet

> interface{} GetEmbeddingsMemoriesEfGet(ctx).Execute()

Get Embeddings

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
	resp, r, err := apiClient.MemoriesAPI.GetEmbeddingsMemoriesEfGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.GetEmbeddingsMemoriesEfGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmbeddingsMemoriesEfGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.GetEmbeddingsMemoriesEfGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmbeddingsMemoriesEfGetRequest struct via the builder pattern


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


## GetMemoriesMemoriesGet

> []MemoryModel GetMemoriesMemoriesGet(ctx).Execute()

Get Memories

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
	resp, r, err := apiClient.MemoriesAPI.GetMemoriesMemoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.GetMemoriesMemoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMemoriesMemoriesGet`: []MemoryModel
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.GetMemoriesMemoriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemoriesMemoriesGetRequest struct via the builder pattern


### Return type

[**[]MemoryModel**](MemoryModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMemoryMemoriesQueryPost

> interface{} QueryMemoryMemoriesQueryPost(ctx).QueryMemoryForm(queryMemoryForm).Execute()

Query Memory

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
	queryMemoryForm := *openapiclient.NewQueryMemoryForm("Content_example") // QueryMemoryForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.QueryMemoryMemoriesQueryPost(context.Background()).QueryMemoryForm(queryMemoryForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.QueryMemoryMemoriesQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryMemoryMemoriesQueryPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.QueryMemoryMemoriesQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMemoryMemoriesQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryMemoryForm** | [**QueryMemoryForm**](QueryMemoryForm.md) |  | 

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


## ResetMemoryFromVectorDbMemoriesResetPost

> bool ResetMemoryFromVectorDbMemoriesResetPost(ctx).Execute()

Reset Memory From Vector Db

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
	resp, r, err := apiClient.MemoriesAPI.ResetMemoryFromVectorDbMemoriesResetPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.ResetMemoryFromVectorDbMemoriesResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetMemoryFromVectorDbMemoriesResetPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.ResetMemoryFromVectorDbMemoriesResetPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetMemoryFromVectorDbMemoriesResetPostRequest struct via the builder pattern


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


## UpdateMemoryByIdMemoriesMemoryIdUpdatePost

> MemoryModel UpdateMemoryByIdMemoriesMemoryIdUpdatePost(ctx, memoryId).MemoryUpdateModel(memoryUpdateModel).Execute()

Update Memory By Id

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
	memoryId := "memoryId_example" // string | 
	memoryUpdateModel := *openapiclient.NewMemoryUpdateModel() // MemoryUpdateModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemoriesAPI.UpdateMemoryByIdMemoriesMemoryIdUpdatePost(context.Background(), memoryId).MemoryUpdateModel(memoryUpdateModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemoriesAPI.UpdateMemoryByIdMemoriesMemoryIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMemoryByIdMemoriesMemoryIdUpdatePost`: MemoryModel
	fmt.Fprintf(os.Stdout, "Response from `MemoriesAPI.UpdateMemoryByIdMemoriesMemoryIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemoryByIdMemoriesMemoryIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memoryUpdateModel** | [**MemoryUpdateModel**](MemoryUpdateModel.md) |  | 

### Return type

[**MemoryModel**](MemoryModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


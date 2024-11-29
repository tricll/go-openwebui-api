# \FilesAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllFilesFilesAllDelete**](FilesAPI.md#DeleteAllFilesFilesAllDelete) | **Delete** /files/all | Delete All Files
[**DeleteFileByIdFilesIdDelete**](FilesAPI.md#DeleteFileByIdFilesIdDelete) | **Delete** /files/{id} | Delete File By Id
[**GetFileByIdFilesIdGet**](FilesAPI.md#GetFileByIdFilesIdGet) | **Get** /files/{id} | Get File By Id
[**GetFileContentByIdFilesIdContentFileNameGet**](FilesAPI.md#GetFileContentByIdFilesIdContentFileNameGet) | **Get** /files/{id}/content/{file_name} | Get File Content By Id
[**GetFileContentByIdFilesIdContentGet**](FilesAPI.md#GetFileContentByIdFilesIdContentGet) | **Get** /files/{id}/content | Get File Content By Id
[**GetFileDataContentByIdFilesIdDataContentGet**](FilesAPI.md#GetFileDataContentByIdFilesIdDataContentGet) | **Get** /files/{id}/data/content | Get File Data Content By Id
[**GetHtmlFileContentByIdFilesIdContentHtmlGet**](FilesAPI.md#GetHtmlFileContentByIdFilesIdContentHtmlGet) | **Get** /files/{id}/content/html | Get Html File Content By Id
[**ListFilesFilesGet**](FilesAPI.md#ListFilesFilesGet) | **Get** /files/ | List Files
[**UpdateFileDataContentByIdFilesIdDataContentUpdatePost**](FilesAPI.md#UpdateFileDataContentByIdFilesIdDataContentUpdatePost) | **Post** /files/{id}/data/content/update | Update File Data Content By Id
[**UploadFileFilesPost**](FilesAPI.md#UploadFileFilesPost) | **Post** /files/ | Upload File



## DeleteAllFilesFilesAllDelete

> interface{} DeleteAllFilesFilesAllDelete(ctx).Execute()

Delete All Files

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
	resp, r, err := apiClient.FilesAPI.DeleteAllFilesFilesAllDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteAllFilesFilesAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllFilesFilesAllDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteAllFilesFilesAllDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllFilesFilesAllDeleteRequest struct via the builder pattern


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


## DeleteFileByIdFilesIdDelete

> interface{} DeleteFileByIdFilesIdDelete(ctx, id).Execute()

Delete File By Id

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
	resp, r, err := apiClient.FilesAPI.DeleteFileByIdFilesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFileByIdFilesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileByIdFilesIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFileByIdFilesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileByIdFilesIdDeleteRequest struct via the builder pattern


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


## GetFileByIdFilesIdGet

> FileModel GetFileByIdFilesIdGet(ctx, id).Execute()

Get File By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileByIdFilesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileByIdFilesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileByIdFilesIdGet`: FileModel
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileByIdFilesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileByIdFilesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileModel**](FileModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileContentByIdFilesIdContentFileNameGet

> interface{} GetFileContentByIdFilesIdContentFileNameGet(ctx, id).Execute()

Get File Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileContentByIdFilesIdContentFileNameGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileContentByIdFilesIdContentFileNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileContentByIdFilesIdContentFileNameGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileContentByIdFilesIdContentFileNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileContentByIdFilesIdContentFileNameGetRequest struct via the builder pattern


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


## GetFileContentByIdFilesIdContentGet

> interface{} GetFileContentByIdFilesIdContentGet(ctx, id).Execute()

Get File Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileContentByIdFilesIdContentGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileContentByIdFilesIdContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileContentByIdFilesIdContentGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileContentByIdFilesIdContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileContentByIdFilesIdContentGetRequest struct via the builder pattern


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


## GetFileDataContentByIdFilesIdDataContentGet

> interface{} GetFileDataContentByIdFilesIdDataContentGet(ctx, id).Execute()

Get File Data Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetFileDataContentByIdFilesIdDataContentGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileDataContentByIdFilesIdDataContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileDataContentByIdFilesIdDataContentGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileDataContentByIdFilesIdDataContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileDataContentByIdFilesIdDataContentGetRequest struct via the builder pattern


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


## GetHtmlFileContentByIdFilesIdContentHtmlGet

> interface{} GetHtmlFileContentByIdFilesIdContentHtmlGet(ctx, id).Execute()

Get Html File Content By Id

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
	resp, r, err := apiClient.FilesAPI.GetHtmlFileContentByIdFilesIdContentHtmlGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetHtmlFileContentByIdFilesIdContentHtmlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHtmlFileContentByIdFilesIdContentHtmlGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetHtmlFileContentByIdFilesIdContentHtmlGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHtmlFileContentByIdFilesIdContentHtmlGetRequest struct via the builder pattern


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


## ListFilesFilesGet

> []FileModelResponse ListFilesFilesGet(ctx).Execute()

List Files

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
	resp, r, err := apiClient.FilesAPI.ListFilesFilesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ListFilesFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFilesFilesGet`: []FileModelResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ListFilesFilesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesFilesGetRequest struct via the builder pattern


### Return type

[**[]FileModelResponse**](FileModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileDataContentByIdFilesIdDataContentUpdatePost

> interface{} UpdateFileDataContentByIdFilesIdDataContentUpdatePost(ctx, id).ContentForm(contentForm).Execute()

Update File Data Content By Id

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
	contentForm := *openapiclient.NewContentForm("Content_example") // ContentForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UpdateFileDataContentByIdFilesIdDataContentUpdatePost(context.Background(), id).ContentForm(contentForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UpdateFileDataContentByIdFilesIdDataContentUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileDataContentByIdFilesIdDataContentUpdatePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UpdateFileDataContentByIdFilesIdDataContentUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileDataContentByIdFilesIdDataContentUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentForm** | [**ContentForm**](ContentForm.md) |  | 

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


## UploadFileFilesPost

> FileModelResponse UploadFileFilesPost(ctx).File(file).Execute()

Upload File

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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadFileFilesPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadFileFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFileFilesPost`: FileModelResponse
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadFileFilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**FileModelResponse**](FileModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


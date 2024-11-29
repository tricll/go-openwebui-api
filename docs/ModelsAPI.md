# \ModelsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewModelModelsCreatePost**](ModelsAPI.md#CreateNewModelModelsCreatePost) | **Post** /models/create | Create New Model
[**DeleteAllModelsModelsDeleteAllDelete**](ModelsAPI.md#DeleteAllModelsModelsDeleteAllDelete) | **Delete** /models/delete/all | Delete All Models
[**DeleteModelByIdModelsModelDeleteDelete**](ModelsAPI.md#DeleteModelByIdModelsModelDeleteDelete) | **Delete** /models/model/delete | Delete Model By Id
[**GetBaseModelsModelsBaseGet**](ModelsAPI.md#GetBaseModelsModelsBaseGet) | **Get** /models/base | Get Base Models
[**GetModelByIdModelsModelGet**](ModelsAPI.md#GetModelByIdModelsModelGet) | **Get** /models/model | Get Model By Id
[**GetModelsModelsGet**](ModelsAPI.md#GetModelsModelsGet) | **Get** /models/ | Get Models
[**ToggleModelByIdModelsModelTogglePost**](ModelsAPI.md#ToggleModelByIdModelsModelTogglePost) | **Post** /models/model/toggle | Toggle Model By Id
[**UpdateModelByIdModelsModelUpdatePost**](ModelsAPI.md#UpdateModelByIdModelsModelUpdatePost) | **Post** /models/model/update | Update Model By Id



## CreateNewModelModelsCreatePost

> ModelModel CreateNewModelModelsCreatePost(ctx).ModelForm(modelForm).Execute()

Create New Model

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
	modelForm := *openapiclient.NewModelForm("Id_example", "Name_example", *openapiclient.NewModelMeta(), map[string]interface{}{"key": interface{}(123)}) // ModelForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.CreateNewModelModelsCreatePost(context.Background()).ModelForm(modelForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.CreateNewModelModelsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewModelModelsCreatePost`: ModelModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.CreateNewModelModelsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewModelModelsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelForm** | [**ModelForm**](ModelForm.md) |  | 

### Return type

[**ModelModel**](ModelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllModelsModelsDeleteAllDelete

> bool DeleteAllModelsModelsDeleteAllDelete(ctx).Execute()

Delete All Models

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
	resp, r, err := apiClient.ModelsAPI.DeleteAllModelsModelsDeleteAllDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.DeleteAllModelsModelsDeleteAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllModelsModelsDeleteAllDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.DeleteAllModelsModelsDeleteAllDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllModelsModelsDeleteAllDeleteRequest struct via the builder pattern


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


## DeleteModelByIdModelsModelDeleteDelete

> bool DeleteModelByIdModelsModelDeleteDelete(ctx).Id(id).Execute()

Delete Model By Id

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
	resp, r, err := apiClient.ModelsAPI.DeleteModelByIdModelsModelDeleteDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.DeleteModelByIdModelsModelDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModelByIdModelsModelDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.DeleteModelByIdModelsModelDeleteDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModelByIdModelsModelDeleteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

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


## GetBaseModelsModelsBaseGet

> []ModelResponse GetBaseModelsModelsBaseGet(ctx).Execute()

Get Base Models

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
	resp, r, err := apiClient.ModelsAPI.GetBaseModelsModelsBaseGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetBaseModelsModelsBaseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaseModelsModelsBaseGet`: []ModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetBaseModelsModelsBaseGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBaseModelsModelsBaseGetRequest struct via the builder pattern


### Return type

[**[]ModelResponse**](ModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelByIdModelsModelGet

> ModelResponse GetModelByIdModelsModelGet(ctx).Id(id).Execute()

Get Model By Id

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
	resp, r, err := apiClient.ModelsAPI.GetModelByIdModelsModelGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelByIdModelsModelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelByIdModelsModelGet`: ModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetModelByIdModelsModelGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelByIdModelsModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsModelsGet

> []ModelUserResponse GetModelsModelsGet(ctx).Id(id).Execute()

Get Models

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
	id := "id_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.GetModelsModelsGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelsModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsModelsGet`: []ModelUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GetModelsModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**[]ModelUserResponse**](ModelUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleModelByIdModelsModelTogglePost

> ModelResponse ToggleModelByIdModelsModelTogglePost(ctx).Id(id).Execute()

Toggle Model By Id

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
	resp, r, err := apiClient.ModelsAPI.ToggleModelByIdModelsModelTogglePost(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ToggleModelByIdModelsModelTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleModelByIdModelsModelTogglePost`: ModelResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ToggleModelByIdModelsModelTogglePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToggleModelByIdModelsModelTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**ModelResponse**](ModelResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelByIdModelsModelUpdatePost

> ModelModel UpdateModelByIdModelsModelUpdatePost(ctx).Id(id).ModelForm(modelForm).Execute()

Update Model By Id

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
	modelForm := *openapiclient.NewModelForm("Id_example", "Name_example", *openapiclient.NewModelMeta(), map[string]interface{}{"key": interface{}(123)}) // ModelForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.UpdateModelByIdModelsModelUpdatePost(context.Background()).Id(id).ModelForm(modelForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.UpdateModelByIdModelsModelUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateModelByIdModelsModelUpdatePost`: ModelModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.UpdateModelByIdModelsModelUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModelByIdModelsModelUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **modelForm** | [**ModelForm**](ModelForm.md) |  | 

### Return type

[**ModelModel**](ModelModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \PromptsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewPromptPromptsCreatePost**](PromptsAPI.md#CreateNewPromptPromptsCreatePost) | **Post** /prompts/create | Create New Prompt
[**DeletePromptByCommandPromptsCommandCommandDeleteDelete**](PromptsAPI.md#DeletePromptByCommandPromptsCommandCommandDeleteDelete) | **Delete** /prompts/command/{command}/delete | Delete Prompt By Command
[**GetPromptByCommandPromptsCommandCommandGet**](PromptsAPI.md#GetPromptByCommandPromptsCommandCommandGet) | **Get** /prompts/command/{command} | Get Prompt By Command
[**GetPromptListPromptsListGet**](PromptsAPI.md#GetPromptListPromptsListGet) | **Get** /prompts/list | Get Prompt List
[**GetPromptsPromptsGet**](PromptsAPI.md#GetPromptsPromptsGet) | **Get** /prompts/ | Get Prompts
[**UpdatePromptByCommandPromptsCommandCommandUpdatePost**](PromptsAPI.md#UpdatePromptByCommandPromptsCommandCommandUpdatePost) | **Post** /prompts/command/{command}/update | Update Prompt By Command



## CreateNewPromptPromptsCreatePost

> PromptModel CreateNewPromptPromptsCreatePost(ctx).PromptForm(promptForm).Execute()

Create New Prompt

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
	promptForm := *openapiclient.NewPromptForm("Command_example", "Title_example", "Content_example") // PromptForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.CreateNewPromptPromptsCreatePost(context.Background()).PromptForm(promptForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.CreateNewPromptPromptsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewPromptPromptsCreatePost`: PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.CreateNewPromptPromptsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewPromptPromptsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptForm** | [**PromptForm**](PromptForm.md) |  | 

### Return type

[**PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePromptByCommandPromptsCommandCommandDeleteDelete

> bool DeletePromptByCommandPromptsCommandCommandDeleteDelete(ctx, command).Execute()

Delete Prompt By Command

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
	command := "command_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.DeletePromptByCommandPromptsCommandCommandDeleteDelete(context.Background(), command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.DeletePromptByCommandPromptsCommandCommandDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePromptByCommandPromptsCommandCommandDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.DeletePromptByCommandPromptsCommandCommandDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePromptByCommandPromptsCommandCommandDeleteDeleteRequest struct via the builder pattern


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


## GetPromptByCommandPromptsCommandCommandGet

> PromptModel GetPromptByCommandPromptsCommandCommandGet(ctx, command).Execute()

Get Prompt By Command

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
	command := "command_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.GetPromptByCommandPromptsCommandCommandGet(context.Background(), command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptByCommandPromptsCommandCommandGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptByCommandPromptsCommandCommandGet`: PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptByCommandPromptsCommandCommandGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptByCommandPromptsCommandCommandGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptListPromptsListGet

> []PromptUserResponse GetPromptListPromptsListGet(ctx).Execute()

Get Prompt List

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
	resp, r, err := apiClient.PromptsAPI.GetPromptListPromptsListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptListPromptsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptListPromptsListGet`: []PromptUserResponse
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptListPromptsListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptListPromptsListGetRequest struct via the builder pattern


### Return type

[**[]PromptUserResponse**](PromptUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromptsPromptsGet

> []PromptModel GetPromptsPromptsGet(ctx).Execute()

Get Prompts

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
	resp, r, err := apiClient.PromptsAPI.GetPromptsPromptsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.GetPromptsPromptsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptsPromptsGet`: []PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.GetPromptsPromptsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptsPromptsGetRequest struct via the builder pattern


### Return type

[**[]PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePromptByCommandPromptsCommandCommandUpdatePost

> PromptModel UpdatePromptByCommandPromptsCommandCommandUpdatePost(ctx, command).PromptForm(promptForm).Execute()

Update Prompt By Command

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
	command := "command_example" // string | 
	promptForm := *openapiclient.NewPromptForm("Command_example", "Title_example", "Content_example") // PromptForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.UpdatePromptByCommandPromptsCommandCommandUpdatePost(context.Background(), command).PromptForm(promptForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.UpdatePromptByCommandPromptsCommandCommandUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePromptByCommandPromptsCommandCommandUpdatePost`: PromptModel
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.UpdatePromptByCommandPromptsCommandCommandUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePromptByCommandPromptsCommandCommandUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **promptForm** | [**PromptForm**](PromptForm.md) |  | 

### Return type

[**PromptModel**](PromptModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


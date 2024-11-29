# \ToolsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewToolsToolsCreatePost**](ToolsAPI.md#CreateNewToolsToolsCreatePost) | **Post** /tools/create | Create New Tools
[**DeleteToolsByIdToolsIdIdDeleteDelete**](ToolsAPI.md#DeleteToolsByIdToolsIdIdDeleteDelete) | **Delete** /tools/id/{id}/delete | Delete Tools By Id
[**ExportToolsToolsExportGet**](ToolsAPI.md#ExportToolsToolsExportGet) | **Get** /tools/export | Export Tools
[**GetToolListToolsListGet**](ToolsAPI.md#GetToolListToolsListGet) | **Get** /tools/list | Get Tool List
[**GetToolsByIdToolsIdIdGet**](ToolsAPI.md#GetToolsByIdToolsIdIdGet) | **Get** /tools/id/{id} | Get Tools By Id
[**GetToolsToolsGet**](ToolsAPI.md#GetToolsToolsGet) | **Get** /tools/ | Get Tools
[**GetToolsUserValvesByIdToolsIdIdValvesUserGet**](ToolsAPI.md#GetToolsUserValvesByIdToolsIdIdValvesUserGet) | **Get** /tools/id/{id}/valves/user | Get Tools User Valves By Id
[**GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet**](ToolsAPI.md#GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet) | **Get** /tools/id/{id}/valves/user/spec | Get Tools User Valves Spec By Id
[**GetToolsValvesByIdToolsIdIdValvesGet**](ToolsAPI.md#GetToolsValvesByIdToolsIdIdValvesGet) | **Get** /tools/id/{id}/valves | Get Tools Valves By Id
[**GetToolsValvesSpecByIdToolsIdIdValvesSpecGet**](ToolsAPI.md#GetToolsValvesSpecByIdToolsIdIdValvesSpecGet) | **Get** /tools/id/{id}/valves/spec | Get Tools Valves Spec By Id
[**UpdateToolsByIdToolsIdIdUpdatePost**](ToolsAPI.md#UpdateToolsByIdToolsIdIdUpdatePost) | **Post** /tools/id/{id}/update | Update Tools By Id
[**UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost**](ToolsAPI.md#UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost) | **Post** /tools/id/{id}/valves/user/update | Update Tools User Valves By Id
[**UpdateToolsValvesByIdToolsIdIdValvesUpdatePost**](ToolsAPI.md#UpdateToolsValvesByIdToolsIdIdValvesUpdatePost) | **Post** /tools/id/{id}/valves/update | Update Tools Valves By Id



## CreateNewToolsToolsCreatePost

> ToolResponse CreateNewToolsToolsCreatePost(ctx).ToolForm(toolForm).Execute()

Create New Tools

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
	toolForm := *openapiclient.NewToolForm("Id_example", "Name_example", "Content_example", *openapiclient.NewToolMeta()) // ToolForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.CreateNewToolsToolsCreatePost(context.Background()).ToolForm(toolForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.CreateNewToolsToolsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewToolsToolsCreatePost`: ToolResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.CreateNewToolsToolsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewToolsToolsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolForm** | [**ToolForm**](ToolForm.md) |  | 

### Return type

[**ToolResponse**](ToolResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToolsByIdToolsIdIdDeleteDelete

> bool DeleteToolsByIdToolsIdIdDeleteDelete(ctx, id).Execute()

Delete Tools By Id

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
	resp, r, err := apiClient.ToolsAPI.DeleteToolsByIdToolsIdIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.DeleteToolsByIdToolsIdIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToolsByIdToolsIdIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.DeleteToolsByIdToolsIdIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolsByIdToolsIdIdDeleteDeleteRequest struct via the builder pattern


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


## ExportToolsToolsExportGet

> []ToolModel ExportToolsToolsExportGet(ctx).Execute()

Export Tools

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
	resp, r, err := apiClient.ToolsAPI.ExportToolsToolsExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.ExportToolsToolsExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportToolsToolsExportGet`: []ToolModel
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.ExportToolsToolsExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportToolsToolsExportGetRequest struct via the builder pattern


### Return type

[**[]ToolModel**](ToolModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolListToolsListGet

> []ToolUserResponse GetToolListToolsListGet(ctx).Execute()

Get Tool List

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
	resp, r, err := apiClient.ToolsAPI.GetToolListToolsListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolListToolsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolListToolsListGet`: []ToolUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolListToolsListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolListToolsListGetRequest struct via the builder pattern


### Return type

[**[]ToolUserResponse**](ToolUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsByIdToolsIdIdGet

> ToolModel GetToolsByIdToolsIdIdGet(ctx, id).Execute()

Get Tools By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsByIdToolsIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsByIdToolsIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsByIdToolsIdIdGet`: ToolModel
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsByIdToolsIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsByIdToolsIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToolModel**](ToolModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsToolsGet

> []ToolUserResponse GetToolsToolsGet(ctx).Execute()

Get Tools

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
	resp, r, err := apiClient.ToolsAPI.GetToolsToolsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsToolsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsToolsGet`: []ToolUserResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsToolsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsToolsGetRequest struct via the builder pattern


### Return type

[**[]ToolUserResponse**](ToolUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolsUserValvesByIdToolsIdIdValvesUserGet

> map[string]interface{} GetToolsUserValvesByIdToolsIdIdValvesUserGet(ctx, id).Execute()

Get Tools User Valves By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsUserValvesByIdToolsIdIdValvesUserGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsUserValvesByIdToolsIdIdValvesUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsUserValvesByIdToolsIdIdValvesUserGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsUserValvesByIdToolsIdIdValvesUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsUserValvesByIdToolsIdIdValvesUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet

> map[string]interface{} GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet(ctx, id).Execute()

Get Tools User Valves Spec By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsUserValvesSpecByIdToolsIdIdValvesUserSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetToolsValvesByIdToolsIdIdValvesGet

> map[string]interface{} GetToolsValvesByIdToolsIdIdValvesGet(ctx, id).Execute()

Get Tools Valves By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsValvesByIdToolsIdIdValvesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsValvesByIdToolsIdIdValvesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsValvesByIdToolsIdIdValvesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsValvesByIdToolsIdIdValvesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsValvesByIdToolsIdIdValvesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetToolsValvesSpecByIdToolsIdIdValvesSpecGet

> map[string]interface{} GetToolsValvesSpecByIdToolsIdIdValvesSpecGet(ctx, id).Execute()

Get Tools Valves Spec By Id

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
	resp, r, err := apiClient.ToolsAPI.GetToolsValvesSpecByIdToolsIdIdValvesSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.GetToolsValvesSpecByIdToolsIdIdValvesSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolsValvesSpecByIdToolsIdIdValvesSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.GetToolsValvesSpecByIdToolsIdIdValvesSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsValvesSpecByIdToolsIdIdValvesSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateToolsByIdToolsIdIdUpdatePost

> ToolModel UpdateToolsByIdToolsIdIdUpdatePost(ctx, id).ToolForm(toolForm).Execute()

Update Tools By Id

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
	toolForm := *openapiclient.NewToolForm("Id_example", "Name_example", "Content_example", *openapiclient.NewToolMeta()) // ToolForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.UpdateToolsByIdToolsIdIdUpdatePost(context.Background(), id).ToolForm(toolForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.UpdateToolsByIdToolsIdIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolsByIdToolsIdIdUpdatePost`: ToolModel
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.UpdateToolsByIdToolsIdIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolsByIdToolsIdIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolForm** | [**ToolForm**](ToolForm.md) |  | 

### Return type

[**ToolModel**](ToolModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost

> map[string]interface{} UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost(ctx, id).Body(body).Execute()

Update Tools User Valves By Id

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.UpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolsUserValvesByIdToolsIdIdValvesUserUpdatePostRequest struct via the builder pattern


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


## UpdateToolsValvesByIdToolsIdIdValvesUpdatePost

> map[string]interface{} UpdateToolsValvesByIdToolsIdIdValvesUpdatePost(ctx, id).Body(body).Execute()

Update Tools Valves By Id

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.UpdateToolsValvesByIdToolsIdIdValvesUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.UpdateToolsValvesByIdToolsIdIdValvesUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToolsValvesByIdToolsIdIdValvesUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.UpdateToolsValvesByIdToolsIdIdValvesUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolsValvesByIdToolsIdIdValvesUpdatePostRequest struct via the builder pattern


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


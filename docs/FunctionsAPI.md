# \FunctionsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewFunctionFunctionsCreatePost**](FunctionsAPI.md#CreateNewFunctionFunctionsCreatePost) | **Post** /functions/create | Create New Function
[**DeleteFunctionByIdFunctionsIdIdDeleteDelete**](FunctionsAPI.md#DeleteFunctionByIdFunctionsIdIdDeleteDelete) | **Delete** /functions/id/{id}/delete | Delete Function By Id
[**GetFunctionByIdFunctionsIdIdGet**](FunctionsAPI.md#GetFunctionByIdFunctionsIdIdGet) | **Get** /functions/id/{id} | Get Function By Id
[**GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet**](FunctionsAPI.md#GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet) | **Get** /functions/id/{id}/valves/user | Get Function User Valves By Id
[**GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet**](FunctionsAPI.md#GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet) | **Get** /functions/id/{id}/valves/user/spec | Get Function User Valves Spec By Id
[**GetFunctionValvesByIdFunctionsIdIdValvesGet**](FunctionsAPI.md#GetFunctionValvesByIdFunctionsIdIdValvesGet) | **Get** /functions/id/{id}/valves | Get Function Valves By Id
[**GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet**](FunctionsAPI.md#GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet) | **Get** /functions/id/{id}/valves/spec | Get Function Valves Spec By Id
[**GetFunctionsFunctionsExportGet**](FunctionsAPI.md#GetFunctionsFunctionsExportGet) | **Get** /functions/export | Get Functions
[**GetFunctionsFunctionsGet**](FunctionsAPI.md#GetFunctionsFunctionsGet) | **Get** /functions/ | Get Functions
[**ToggleFunctionByIdFunctionsIdIdTogglePost**](FunctionsAPI.md#ToggleFunctionByIdFunctionsIdIdTogglePost) | **Post** /functions/id/{id}/toggle | Toggle Function By Id
[**ToggleGlobalByIdFunctionsIdIdToggleGlobalPost**](FunctionsAPI.md#ToggleGlobalByIdFunctionsIdIdToggleGlobalPost) | **Post** /functions/id/{id}/toggle/global | Toggle Global By Id
[**UpdateFunctionByIdFunctionsIdIdUpdatePost**](FunctionsAPI.md#UpdateFunctionByIdFunctionsIdIdUpdatePost) | **Post** /functions/id/{id}/update | Update Function By Id
[**UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost**](FunctionsAPI.md#UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost) | **Post** /functions/id/{id}/valves/user/update | Update Function User Valves By Id
[**UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost**](FunctionsAPI.md#UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost) | **Post** /functions/id/{id}/valves/update | Update Function Valves By Id



## CreateNewFunctionFunctionsCreatePost

> FunctionResponse CreateNewFunctionFunctionsCreatePost(ctx).FunctionForm(functionForm).Execute()

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
	functionForm := *openapiclient.NewFunctionForm("Id_example", "Name_example", "Content_example", *openapiclient.NewFunctionMeta()) // FunctionForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.CreateNewFunctionFunctionsCreatePost(context.Background()).FunctionForm(functionForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.CreateNewFunctionFunctionsCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewFunctionFunctionsCreatePost`: FunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.CreateNewFunctionFunctionsCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewFunctionFunctionsCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionForm** | [**FunctionForm**](FunctionForm.md) |  | 

### Return type

[**FunctionResponse**](FunctionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunctionByIdFunctionsIdIdDeleteDelete

> bool DeleteFunctionByIdFunctionsIdIdDeleteDelete(ctx, id).Execute()

Delete Function By Id

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
	resp, r, err := apiClient.FunctionsAPI.DeleteFunctionByIdFunctionsIdIdDeleteDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteFunctionByIdFunctionsIdIdDeleteDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFunctionByIdFunctionsIdIdDeleteDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.DeleteFunctionByIdFunctionsIdIdDeleteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionByIdFunctionsIdIdDeleteDeleteRequest struct via the builder pattern


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


## GetFunctionByIdFunctionsIdIdGet

> FunctionModel GetFunctionByIdFunctionsIdIdGet(ctx, id).Execute()

Get Function By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionByIdFunctionsIdIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionByIdFunctionsIdIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionByIdFunctionsIdIdGet`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionByIdFunctionsIdIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionByIdFunctionsIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet

> map[string]interface{} GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet(ctx, id).Execute()

Get Function User Valves By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionUserValvesByIdFunctionsIdIdValvesUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionUserValvesByIdFunctionsIdIdValvesUserGetRequest struct via the builder pattern


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


## GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet

> map[string]interface{} GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet(ctx, id).Execute()

Get Function User Valves Spec By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionUserValvesSpecByIdFunctionsIdIdValvesUserSpecGetRequest struct via the builder pattern


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


## GetFunctionValvesByIdFunctionsIdIdValvesGet

> map[string]interface{} GetFunctionValvesByIdFunctionsIdIdValvesGet(ctx, id).Execute()

Get Function Valves By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionValvesByIdFunctionsIdIdValvesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionValvesByIdFunctionsIdIdValvesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionValvesByIdFunctionsIdIdValvesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionValvesByIdFunctionsIdIdValvesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionValvesByIdFunctionsIdIdValvesGetRequest struct via the builder pattern


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


## GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet

> map[string]interface{} GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet(ctx, id).Execute()

Get Function Valves Spec By Id

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionValvesSpecByIdFunctionsIdIdValvesSpecGetRequest struct via the builder pattern


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


## GetFunctionsFunctionsExportGet

> []FunctionModel GetFunctionsFunctionsExportGet(ctx).Execute()

Get Functions

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsFunctionsExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsFunctionsExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsFunctionsExportGet`: []FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsFunctionsExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsFunctionsExportGetRequest struct via the builder pattern


### Return type

[**[]FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsFunctionsGet

> []FunctionResponse GetFunctionsFunctionsGet(ctx).Execute()

Get Functions

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsFunctionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsFunctionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsFunctionsGet`: []FunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsFunctionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsFunctionsGetRequest struct via the builder pattern


### Return type

[**[]FunctionResponse**](FunctionResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFunctionByIdFunctionsIdIdTogglePost

> FunctionModel ToggleFunctionByIdFunctionsIdIdTogglePost(ctx, id).Execute()

Toggle Function By Id

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
	resp, r, err := apiClient.FunctionsAPI.ToggleFunctionByIdFunctionsIdIdTogglePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ToggleFunctionByIdFunctionsIdIdTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleFunctionByIdFunctionsIdIdTogglePost`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ToggleFunctionByIdFunctionsIdIdTogglePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFunctionByIdFunctionsIdIdTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleGlobalByIdFunctionsIdIdToggleGlobalPost

> FunctionModel ToggleGlobalByIdFunctionsIdIdToggleGlobalPost(ctx, id).Execute()

Toggle Global By Id

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
	resp, r, err := apiClient.FunctionsAPI.ToggleGlobalByIdFunctionsIdIdToggleGlobalPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ToggleGlobalByIdFunctionsIdIdToggleGlobalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleGlobalByIdFunctionsIdIdToggleGlobalPost`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ToggleGlobalByIdFunctionsIdIdToggleGlobalPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleGlobalByIdFunctionsIdIdToggleGlobalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunctionByIdFunctionsIdIdUpdatePost

> FunctionModel UpdateFunctionByIdFunctionsIdIdUpdatePost(ctx, id).FunctionForm(functionForm).Execute()

Update Function By Id

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
	functionForm := *openapiclient.NewFunctionForm("Id_example", "Name_example", "Content_example", *openapiclient.NewFunctionMeta()) // FunctionForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunctionByIdFunctionsIdIdUpdatePost(context.Background(), id).FunctionForm(functionForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunctionByIdFunctionsIdIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionByIdFunctionsIdIdUpdatePost`: FunctionModel
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunctionByIdFunctionsIdIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionByIdFunctionsIdIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionForm** | [**FunctionForm**](FunctionForm.md) |  | 

### Return type

[**FunctionModel**](FunctionModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost

> map[string]interface{} UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost(ctx, id).Body(body).Execute()

Update Function User Valves By Id

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionUserValvesByIdFunctionsIdIdValvesUserUpdatePostRequest struct via the builder pattern


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


## UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost

> map[string]interface{} UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost(ctx, id).Body(body).Execute()

Update Function Valves By Id

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.UpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFunctionValvesByIdFunctionsIdIdValvesUpdatePostRequest struct via the builder pattern


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


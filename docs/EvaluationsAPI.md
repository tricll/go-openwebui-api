# \EvaluationsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedbackEvaluationsFeedbackPost**](EvaluationsAPI.md#CreateFeedbackEvaluationsFeedbackPost) | **Post** /evaluations/feedback | Create Feedback
[**DeleteAllFeedbacksEvaluationsFeedbacksAllDelete**](EvaluationsAPI.md#DeleteAllFeedbacksEvaluationsFeedbacksAllDelete) | **Delete** /evaluations/feedbacks/all | Delete All Feedbacks
[**DeleteFeedbackByIdEvaluationsFeedbackIdDelete**](EvaluationsAPI.md#DeleteFeedbackByIdEvaluationsFeedbackIdDelete) | **Delete** /evaluations/feedback/{id} | Delete Feedback By Id
[**DeleteFeedbacksEvaluationsFeedbacksDelete**](EvaluationsAPI.md#DeleteFeedbacksEvaluationsFeedbacksDelete) | **Delete** /evaluations/feedbacks | Delete Feedbacks
[**GetAllFeedbacksEvaluationsFeedbacksAllExportGet**](EvaluationsAPI.md#GetAllFeedbacksEvaluationsFeedbacksAllExportGet) | **Get** /evaluations/feedbacks/all/export | Get All Feedbacks
[**GetAllFeedbacksEvaluationsFeedbacksAllGet**](EvaluationsAPI.md#GetAllFeedbacksEvaluationsFeedbacksAllGet) | **Get** /evaluations/feedbacks/all | Get All Feedbacks
[**GetConfigEvaluationsConfigGet**](EvaluationsAPI.md#GetConfigEvaluationsConfigGet) | **Get** /evaluations/config | Get Config
[**GetFeedbackByIdEvaluationsFeedbackIdGet**](EvaluationsAPI.md#GetFeedbackByIdEvaluationsFeedbackIdGet) | **Get** /evaluations/feedback/{id} | Get Feedback By Id
[**GetFeedbacksEvaluationsFeedbacksUserGet**](EvaluationsAPI.md#GetFeedbacksEvaluationsFeedbacksUserGet) | **Get** /evaluations/feedbacks/user | Get Feedbacks
[**UpdateConfigEvaluationsConfigPost**](EvaluationsAPI.md#UpdateConfigEvaluationsConfigPost) | **Post** /evaluations/config | Update Config
[**UpdateFeedbackByIdEvaluationsFeedbackIdPost**](EvaluationsAPI.md#UpdateFeedbackByIdEvaluationsFeedbackIdPost) | **Post** /evaluations/feedback/{id} | Update Feedback By Id



## CreateFeedbackEvaluationsFeedbackPost

> FeedbackModel CreateFeedbackEvaluationsFeedbackPost(ctx).FeedbackForm(feedbackForm).Execute()

Create Feedback

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
	feedbackForm := *openapiclient.NewFeedbackForm("Type_example") // FeedbackForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.CreateFeedbackEvaluationsFeedbackPost(context.Background()).FeedbackForm(feedbackForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.CreateFeedbackEvaluationsFeedbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedbackEvaluationsFeedbackPost`: FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.CreateFeedbackEvaluationsFeedbackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackEvaluationsFeedbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackForm** | [**FeedbackForm**](FeedbackForm.md) |  | 

### Return type

[**FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllFeedbacksEvaluationsFeedbacksAllDelete

> interface{} DeleteAllFeedbacksEvaluationsFeedbacksAllDelete(ctx).Execute()

Delete All Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.DeleteAllFeedbacksEvaluationsFeedbacksAllDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.DeleteAllFeedbacksEvaluationsFeedbacksAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllFeedbacksEvaluationsFeedbacksAllDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.DeleteAllFeedbacksEvaluationsFeedbacksAllDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllFeedbacksEvaluationsFeedbacksAllDeleteRequest struct via the builder pattern


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


## DeleteFeedbackByIdEvaluationsFeedbackIdDelete

> interface{} DeleteFeedbackByIdEvaluationsFeedbackIdDelete(ctx, id).Execute()

Delete Feedback By Id

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
	resp, r, err := apiClient.EvaluationsAPI.DeleteFeedbackByIdEvaluationsFeedbackIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.DeleteFeedbackByIdEvaluationsFeedbackIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeedbackByIdEvaluationsFeedbackIdDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.DeleteFeedbackByIdEvaluationsFeedbackIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedbackByIdEvaluationsFeedbackIdDeleteRequest struct via the builder pattern


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


## DeleteFeedbacksEvaluationsFeedbacksDelete

> bool DeleteFeedbacksEvaluationsFeedbacksDelete(ctx).Execute()

Delete Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.DeleteFeedbacksEvaluationsFeedbacksDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.DeleteFeedbacksEvaluationsFeedbacksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeedbacksEvaluationsFeedbacksDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.DeleteFeedbacksEvaluationsFeedbacksDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedbacksEvaluationsFeedbacksDeleteRequest struct via the builder pattern


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


## GetAllFeedbacksEvaluationsFeedbacksAllExportGet

> []FeedbackModel GetAllFeedbacksEvaluationsFeedbacksAllExportGet(ctx).Execute()

Get All Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.GetAllFeedbacksEvaluationsFeedbacksAllExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetAllFeedbacksEvaluationsFeedbacksAllExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFeedbacksEvaluationsFeedbacksAllExportGet`: []FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetAllFeedbacksEvaluationsFeedbacksAllExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFeedbacksEvaluationsFeedbacksAllExportGetRequest struct via the builder pattern


### Return type

[**[]FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFeedbacksEvaluationsFeedbacksAllGet

> []FeedbackUserResponse GetAllFeedbacksEvaluationsFeedbacksAllGet(ctx).Execute()

Get All Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.GetAllFeedbacksEvaluationsFeedbacksAllGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetAllFeedbacksEvaluationsFeedbacksAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFeedbacksEvaluationsFeedbacksAllGet`: []FeedbackUserResponse
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetAllFeedbacksEvaluationsFeedbacksAllGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFeedbacksEvaluationsFeedbacksAllGetRequest struct via the builder pattern


### Return type

[**[]FeedbackUserResponse**](FeedbackUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigEvaluationsConfigGet

> interface{} GetConfigEvaluationsConfigGet(ctx).Execute()

Get Config

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
	resp, r, err := apiClient.EvaluationsAPI.GetConfigEvaluationsConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetConfigEvaluationsConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigEvaluationsConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetConfigEvaluationsConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigEvaluationsConfigGetRequest struct via the builder pattern


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


## GetFeedbackByIdEvaluationsFeedbackIdGet

> FeedbackModel GetFeedbackByIdEvaluationsFeedbackIdGet(ctx, id).Execute()

Get Feedback By Id

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
	resp, r, err := apiClient.EvaluationsAPI.GetFeedbackByIdEvaluationsFeedbackIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetFeedbackByIdEvaluationsFeedbackIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedbackByIdEvaluationsFeedbackIdGet`: FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetFeedbackByIdEvaluationsFeedbackIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackByIdEvaluationsFeedbackIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedbacksEvaluationsFeedbacksUserGet

> []FeedbackUserResponse GetFeedbacksEvaluationsFeedbacksUserGet(ctx).Execute()

Get Feedbacks

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
	resp, r, err := apiClient.EvaluationsAPI.GetFeedbacksEvaluationsFeedbacksUserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.GetFeedbacksEvaluationsFeedbacksUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedbacksEvaluationsFeedbacksUserGet`: []FeedbackUserResponse
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.GetFeedbacksEvaluationsFeedbacksUserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbacksEvaluationsFeedbacksUserGetRequest struct via the builder pattern


### Return type

[**[]FeedbackUserResponse**](FeedbackUserResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigEvaluationsConfigPost

> interface{} UpdateConfigEvaluationsConfigPost(ctx).UpdateConfigForm(updateConfigForm).Execute()

Update Config

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
	updateConfigForm := *openapiclient.NewUpdateConfigForm() // UpdateConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.UpdateConfigEvaluationsConfigPost(context.Background()).UpdateConfigForm(updateConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.UpdateConfigEvaluationsConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfigEvaluationsConfigPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.UpdateConfigEvaluationsConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigEvaluationsConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigForm** | [**UpdateConfigForm**](UpdateConfigForm.md) |  | 

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


## UpdateFeedbackByIdEvaluationsFeedbackIdPost

> FeedbackModel UpdateFeedbackByIdEvaluationsFeedbackIdPost(ctx, id).FeedbackForm(feedbackForm).Execute()

Update Feedback By Id

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
	feedbackForm := *openapiclient.NewFeedbackForm("Type_example") // FeedbackForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvaluationsAPI.UpdateFeedbackByIdEvaluationsFeedbackIdPost(context.Background(), id).FeedbackForm(feedbackForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluationsAPI.UpdateFeedbackByIdEvaluationsFeedbackIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeedbackByIdEvaluationsFeedbackIdPost`: FeedbackModel
	fmt.Fprintf(os.Stdout, "Response from `EvaluationsAPI.UpdateFeedbackByIdEvaluationsFeedbackIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedbackByIdEvaluationsFeedbackIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feedbackForm** | [**FeedbackForm**](FeedbackForm.md) |  | 

### Return type

[**FeedbackModel**](FeedbackModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


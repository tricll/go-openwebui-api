# \ConfigsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportConfigConfigsExportGet**](ConfigsAPI.md#ExportConfigConfigsExportGet) | **Get** /configs/export | Export Config
[**GetBannersConfigsBannersGet**](ConfigsAPI.md#GetBannersConfigsBannersGet) | **Get** /configs/banners | Get Banners
[**GetModelsConfigConfigsModelsGet**](ConfigsAPI.md#GetModelsConfigConfigsModelsGet) | **Get** /configs/models | Get Models Config
[**ImportConfigConfigsImportPost**](ConfigsAPI.md#ImportConfigConfigsImportPost) | **Post** /configs/import | Import Config
[**SetBannersConfigsBannersPost**](ConfigsAPI.md#SetBannersConfigsBannersPost) | **Post** /configs/banners | Set Banners
[**SetDefaultSuggestionsConfigsSuggestionsPost**](ConfigsAPI.md#SetDefaultSuggestionsConfigsSuggestionsPost) | **Post** /configs/suggestions | Set Default Suggestions
[**SetModelsConfigConfigsModelsPost**](ConfigsAPI.md#SetModelsConfigConfigsModelsPost) | **Post** /configs/models | Set Models Config



## ExportConfigConfigsExportGet

> map[string]interface{} ExportConfigConfigsExportGet(ctx).Execute()

Export Config

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
	resp, r, err := apiClient.ConfigsAPI.ExportConfigConfigsExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ExportConfigConfigsExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportConfigConfigsExportGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ExportConfigConfigsExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportConfigConfigsExportGetRequest struct via the builder pattern


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


## GetBannersConfigsBannersGet

> []BannerModel GetBannersConfigsBannersGet(ctx).Execute()

Get Banners

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
	resp, r, err := apiClient.ConfigsAPI.GetBannersConfigsBannersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetBannersConfigsBannersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBannersConfigsBannersGet`: []BannerModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetBannersConfigsBannersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannersConfigsBannersGetRequest struct via the builder pattern


### Return type

[**[]BannerModel**](BannerModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsConfigConfigsModelsGet

> ModelsConfigForm GetModelsConfigConfigsModelsGet(ctx).Execute()

Get Models Config

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
	resp, r, err := apiClient.ConfigsAPI.GetModelsConfigConfigsModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetModelsConfigConfigsModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsConfigConfigsModelsGet`: ModelsConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetModelsConfigConfigsModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsConfigConfigsModelsGetRequest struct via the builder pattern


### Return type

[**ModelsConfigForm**](ModelsConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportConfigConfigsImportPost

> map[string]interface{} ImportConfigConfigsImportPost(ctx).ImportConfigForm(importConfigForm).Execute()

Import Config

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
	importConfigForm := *openapiclient.NewImportConfigForm(map[string]interface{}(123)) // ImportConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.ImportConfigConfigsImportPost(context.Background()).ImportConfigForm(importConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.ImportConfigConfigsImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportConfigConfigsImportPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.ImportConfigConfigsImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportConfigConfigsImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importConfigForm** | [**ImportConfigForm**](ImportConfigForm.md) |  | 

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


## SetBannersConfigsBannersPost

> []BannerModel SetBannersConfigsBannersPost(ctx).SetBannersForm(setBannersForm).Execute()

Set Banners

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
	setBannersForm := *openapiclient.NewSetBannersForm([]openapiclient.BannerModel{*openapiclient.NewBannerModel("Id_example", "Type_example", "Content_example", false, int32(123))}) // SetBannersForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetBannersConfigsBannersPost(context.Background()).SetBannersForm(setBannersForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetBannersConfigsBannersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBannersConfigsBannersPost`: []BannerModel
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetBannersConfigsBannersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBannersConfigsBannersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setBannersForm** | [**SetBannersForm**](SetBannersForm.md) |  | 

### Return type

[**[]BannerModel**](BannerModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultSuggestionsConfigsSuggestionsPost

> []PromptSuggestion SetDefaultSuggestionsConfigsSuggestionsPost(ctx).SetDefaultSuggestionsForm(setDefaultSuggestionsForm).Execute()

Set Default Suggestions

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
	setDefaultSuggestionsForm := *openapiclient.NewSetDefaultSuggestionsForm([]openapiclient.PromptSuggestion{*openapiclient.NewPromptSuggestion([]string{"Title_example"}, "Content_example")}) // SetDefaultSuggestionsForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetDefaultSuggestionsConfigsSuggestionsPost(context.Background()).SetDefaultSuggestionsForm(setDefaultSuggestionsForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetDefaultSuggestionsConfigsSuggestionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultSuggestionsConfigsSuggestionsPost`: []PromptSuggestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetDefaultSuggestionsConfigsSuggestionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultSuggestionsConfigsSuggestionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultSuggestionsForm** | [**SetDefaultSuggestionsForm**](SetDefaultSuggestionsForm.md) |  | 

### Return type

[**[]PromptSuggestion**](PromptSuggestion.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetModelsConfigConfigsModelsPost

> ModelsConfigForm SetModelsConfigConfigsModelsPost(ctx).ModelsConfigForm(modelsConfigForm).Execute()

Set Models Config

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
	modelsConfigForm := *openapiclient.NewModelsConfigForm("DEFAULT_MODELS_example", []string{"MODEL_ORDER_LIST_example"}) // ModelsConfigForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigsAPI.SetModelsConfigConfigsModelsPost(context.Background()).ModelsConfigForm(modelsConfigForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.SetModelsConfigConfigsModelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetModelsConfigConfigsModelsPost`: ModelsConfigForm
	fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.SetModelsConfigConfigsModelsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetModelsConfigConfigsModelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelsConfigForm** | [**ModelsConfigForm**](ModelsConfigForm.md) |  | 

### Return type

[**ModelsConfigForm**](ModelsConfigForm.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


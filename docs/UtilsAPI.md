# \UtilsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadChatAsPdfUtilsPdfPost**](UtilsAPI.md#DownloadChatAsPdfUtilsPdfPost) | **Post** /utils/pdf | Download Chat As Pdf
[**DownloadDbUtilsDbDownloadGet**](UtilsAPI.md#DownloadDbUtilsDbDownloadGet) | **Get** /utils/db/download | Download Db
[**DownloadLitellmConfigYamlUtilsLitellmConfigGet**](UtilsAPI.md#DownloadLitellmConfigYamlUtilsLitellmConfigGet) | **Get** /utils/litellm/config | Download Litellm Config Yaml
[**FormatCodeUtilsCodeFormatPost**](UtilsAPI.md#FormatCodeUtilsCodeFormatPost) | **Post** /utils/code/format | Format Code
[**GetGravatarUtilsGravatarGet**](UtilsAPI.md#GetGravatarUtilsGravatarGet) | **Get** /utils/gravatar | Get Gravatar
[**GetHtmlFromMarkdownUtilsMarkdownPost**](UtilsAPI.md#GetHtmlFromMarkdownUtilsMarkdownPost) | **Post** /utils/markdown | Get Html From Markdown



## DownloadChatAsPdfUtilsPdfPost

> interface{} DownloadChatAsPdfUtilsPdfPost(ctx).ChatTitleMessagesForm(chatTitleMessagesForm).Execute()

Download Chat As Pdf

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
	chatTitleMessagesForm := *openapiclient.NewChatTitleMessagesForm("Title_example", []map[string]interface{}{map[string]interface{}(123)}) // ChatTitleMessagesForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.DownloadChatAsPdfUtilsPdfPost(context.Background()).ChatTitleMessagesForm(chatTitleMessagesForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.DownloadChatAsPdfUtilsPdfPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadChatAsPdfUtilsPdfPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.DownloadChatAsPdfUtilsPdfPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadChatAsPdfUtilsPdfPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatTitleMessagesForm** | [**ChatTitleMessagesForm**](ChatTitleMessagesForm.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDbUtilsDbDownloadGet

> interface{} DownloadDbUtilsDbDownloadGet(ctx).Execute()

Download Db

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
	resp, r, err := apiClient.UtilsAPI.DownloadDbUtilsDbDownloadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.DownloadDbUtilsDbDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDbUtilsDbDownloadGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.DownloadDbUtilsDbDownloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDbUtilsDbDownloadGetRequest struct via the builder pattern


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


## DownloadLitellmConfigYamlUtilsLitellmConfigGet

> interface{} DownloadLitellmConfigYamlUtilsLitellmConfigGet(ctx).Execute()

Download Litellm Config Yaml

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
	resp, r, err := apiClient.UtilsAPI.DownloadLitellmConfigYamlUtilsLitellmConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.DownloadLitellmConfigYamlUtilsLitellmConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLitellmConfigYamlUtilsLitellmConfigGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.DownloadLitellmConfigYamlUtilsLitellmConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLitellmConfigYamlUtilsLitellmConfigGetRequest struct via the builder pattern


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


## FormatCodeUtilsCodeFormatPost

> interface{} FormatCodeUtilsCodeFormatPost(ctx).CodeFormatRequest(codeFormatRequest).Execute()

Format Code

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
	codeFormatRequest := *openapiclient.NewCodeFormatRequest("Code_example") // CodeFormatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.FormatCodeUtilsCodeFormatPost(context.Background()).CodeFormatRequest(codeFormatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.FormatCodeUtilsCodeFormatPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FormatCodeUtilsCodeFormatPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.FormatCodeUtilsCodeFormatPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFormatCodeUtilsCodeFormatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeFormatRequest** | [**CodeFormatRequest**](CodeFormatRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGravatarUtilsGravatarGet

> interface{} GetGravatarUtilsGravatarGet(ctx).Email(email).Execute()

Get Gravatar

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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.GetGravatarUtilsGravatarGet(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.GetGravatarUtilsGravatarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGravatarUtilsGravatarGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.GetGravatarUtilsGravatarGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGravatarUtilsGravatarGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 

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


## GetHtmlFromMarkdownUtilsMarkdownPost

> interface{} GetHtmlFromMarkdownUtilsMarkdownPost(ctx).MarkdownForm(markdownForm).Execute()

Get Html From Markdown

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
	markdownForm := *openapiclient.NewMarkdownForm("Md_example") // MarkdownForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilsAPI.GetHtmlFromMarkdownUtilsMarkdownPost(context.Background()).MarkdownForm(markdownForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.GetHtmlFromMarkdownUtilsMarkdownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHtmlFromMarkdownUtilsMarkdownPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.GetHtmlFromMarkdownUtilsMarkdownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHtmlFromMarkdownUtilsMarkdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markdownForm** | [**MarkdownForm**](MarkdownForm.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


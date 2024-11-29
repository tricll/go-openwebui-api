# \ChatsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagByIdAndTagNameChatsIdTagsPost**](ChatsAPI.md#AddTagByIdAndTagNameChatsIdTagsPost) | **Post** /chats/{id}/tags | Add Tag By Id And Tag Name
[**ArchiveAllChatsChatsArchiveAllPost**](ChatsAPI.md#ArchiveAllChatsChatsArchiveAllPost) | **Post** /chats/archive/all | Archive All Chats
[**ArchiveChatByIdChatsIdArchivePost**](ChatsAPI.md#ArchiveChatByIdChatsIdArchivePost) | **Post** /chats/{id}/archive | Archive Chat By Id
[**CloneChatByIdChatsIdClonePost**](ChatsAPI.md#CloneChatByIdChatsIdClonePost) | **Post** /chats/{id}/clone | Clone Chat By Id
[**CreateNewChatChatsNewPost**](ChatsAPI.md#CreateNewChatChatsNewPost) | **Post** /chats/new | Create New Chat
[**DeleteAllTagsByIdChatsIdTagsAllDelete**](ChatsAPI.md#DeleteAllTagsByIdChatsIdTagsAllDelete) | **Delete** /chats/{id}/tags/all | Delete All Tags By Id
[**DeleteAllUserChatsChatsDelete**](ChatsAPI.md#DeleteAllUserChatsChatsDelete) | **Delete** /chats/ | Delete All User Chats
[**DeleteChatByIdChatsIdDelete**](ChatsAPI.md#DeleteChatByIdChatsIdDelete) | **Delete** /chats/{id} | Delete Chat By Id
[**DeleteSharedChatByIdChatsIdShareDelete**](ChatsAPI.md#DeleteSharedChatByIdChatsIdShareDelete) | **Delete** /chats/{id}/share | Delete Shared Chat By Id
[**DeleteTagByIdAndTagNameChatsIdTagsDelete**](ChatsAPI.md#DeleteTagByIdAndTagNameChatsIdTagsDelete) | **Delete** /chats/{id}/tags | Delete Tag By Id And Tag Name
[**GetAllUserChatsInDbChatsAllDbGet**](ChatsAPI.md#GetAllUserChatsInDbChatsAllDbGet) | **Get** /chats/all/db | Get All User Chats In Db
[**GetAllUserTagsChatsAllTagsGet**](ChatsAPI.md#GetAllUserTagsChatsAllTagsGet) | **Get** /chats/all/tags | Get All User Tags
[**GetArchivedSessionUserChatListChatsArchivedGet**](ChatsAPI.md#GetArchivedSessionUserChatListChatsArchivedGet) | **Get** /chats/archived | Get Archived Session User Chat List
[**GetChatByIdChatsIdGet**](ChatsAPI.md#GetChatByIdChatsIdGet) | **Get** /chats/{id} | Get Chat By Id
[**GetChatTagsByIdChatsIdTagsGet**](ChatsAPI.md#GetChatTagsByIdChatsIdTagsGet) | **Get** /chats/{id}/tags | Get Chat Tags By Id
[**GetChatsByFolderIdChatsFolderFolderIdGet**](ChatsAPI.md#GetChatsByFolderIdChatsFolderFolderIdGet) | **Get** /chats/folder/{folder_id} | Get Chats By Folder Id
[**GetPinnedStatusByIdChatsIdPinnedGet**](ChatsAPI.md#GetPinnedStatusByIdChatsIdPinnedGet) | **Get** /chats/{id}/pinned | Get Pinned Status By Id
[**GetSessionUserChatListChatsGet**](ChatsAPI.md#GetSessionUserChatListChatsGet) | **Get** /chats/ | Get Session User Chat List
[**GetSessionUserChatListChatsListGet**](ChatsAPI.md#GetSessionUserChatListChatsListGet) | **Get** /chats/list | Get Session User Chat List
[**GetSharedChatByIdChatsShareShareIdGet**](ChatsAPI.md#GetSharedChatByIdChatsShareShareIdGet) | **Get** /chats/share/{share_id} | Get Shared Chat By Id
[**GetUserArchivedChatsChatsAllArchivedGet**](ChatsAPI.md#GetUserArchivedChatsChatsAllArchivedGet) | **Get** /chats/all/archived | Get User Archived Chats
[**GetUserChatListByTagNameChatsTagsPost**](ChatsAPI.md#GetUserChatListByTagNameChatsTagsPost) | **Post** /chats/tags | Get User Chat List By Tag Name
[**GetUserChatListByUserIdChatsListUserUserIdGet**](ChatsAPI.md#GetUserChatListByUserIdChatsListUserUserIdGet) | **Get** /chats/list/user/{user_id} | Get User Chat List By User Id
[**GetUserChatsChatsAllGet**](ChatsAPI.md#GetUserChatsChatsAllGet) | **Get** /chats/all | Get User Chats
[**GetUserPinnedChatsChatsPinnedGet**](ChatsAPI.md#GetUserPinnedChatsChatsPinnedGet) | **Get** /chats/pinned | Get User Pinned Chats
[**ImportChatChatsImportPost**](ChatsAPI.md#ImportChatChatsImportPost) | **Post** /chats/import | Import Chat
[**PinChatByIdChatsIdPinPost**](ChatsAPI.md#PinChatByIdChatsIdPinPost) | **Post** /chats/{id}/pin | Pin Chat By Id
[**SearchUserChatsChatsSearchGet**](ChatsAPI.md#SearchUserChatsChatsSearchGet) | **Get** /chats/search | Search User Chats
[**ShareChatByIdChatsIdSharePost**](ChatsAPI.md#ShareChatByIdChatsIdSharePost) | **Post** /chats/{id}/share | Share Chat By Id
[**UpdateChatByIdChatsIdPost**](ChatsAPI.md#UpdateChatByIdChatsIdPost) | **Post** /chats/{id} | Update Chat By Id
[**UpdateChatFolderIdByIdChatsIdFolderPost**](ChatsAPI.md#UpdateChatFolderIdByIdChatsIdFolderPost) | **Post** /chats/{id}/folder | Update Chat Folder Id By Id



## AddTagByIdAndTagNameChatsIdTagsPost

> []TagModel AddTagByIdAndTagNameChatsIdTagsPost(ctx, id).TagForm(tagForm).Execute()

Add Tag By Id And Tag Name

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
	tagForm := *openapiclient.NewTagForm("Name_example") // TagForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.AddTagByIdAndTagNameChatsIdTagsPost(context.Background(), id).TagForm(tagForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.AddTagByIdAndTagNameChatsIdTagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTagByIdAndTagNameChatsIdTagsPost`: []TagModel
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.AddTagByIdAndTagNameChatsIdTagsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagByIdAndTagNameChatsIdTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagForm** | [**TagForm**](TagForm.md) |  | 

### Return type

[**[]TagModel**](TagModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveAllChatsChatsArchiveAllPost

> bool ArchiveAllChatsChatsArchiveAllPost(ctx).Execute()

Archive All Chats

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
	resp, r, err := apiClient.ChatsAPI.ArchiveAllChatsChatsArchiveAllPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.ArchiveAllChatsChatsArchiveAllPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveAllChatsChatsArchiveAllPost`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.ArchiveAllChatsChatsArchiveAllPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAllChatsChatsArchiveAllPostRequest struct via the builder pattern


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


## ArchiveChatByIdChatsIdArchivePost

> ChatResponse ArchiveChatByIdChatsIdArchivePost(ctx, id).Execute()

Archive Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.ArchiveChatByIdChatsIdArchivePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.ArchiveChatByIdChatsIdArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveChatByIdChatsIdArchivePost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.ArchiveChatByIdChatsIdArchivePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveChatByIdChatsIdArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneChatByIdChatsIdClonePost

> ChatResponse CloneChatByIdChatsIdClonePost(ctx, id).Execute()

Clone Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.CloneChatByIdChatsIdClonePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.CloneChatByIdChatsIdClonePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneChatByIdChatsIdClonePost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.CloneChatByIdChatsIdClonePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneChatByIdChatsIdClonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewChatChatsNewPost

> ChatResponse CreateNewChatChatsNewPost(ctx).ChatForm(chatForm).Execute()

Create New Chat

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
	chatForm := *openapiclient.NewChatForm(map[string]interface{}(123)) // ChatForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.CreateNewChatChatsNewPost(context.Background()).ChatForm(chatForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.CreateNewChatChatsNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewChatChatsNewPost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.CreateNewChatChatsNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewChatChatsNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatForm** | [**ChatForm**](ChatForm.md) |  | 

### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllTagsByIdChatsIdTagsAllDelete

> bool DeleteAllTagsByIdChatsIdTagsAllDelete(ctx, id).Execute()

Delete All Tags By Id

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
	resp, r, err := apiClient.ChatsAPI.DeleteAllTagsByIdChatsIdTagsAllDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.DeleteAllTagsByIdChatsIdTagsAllDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllTagsByIdChatsIdTagsAllDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.DeleteAllTagsByIdChatsIdTagsAllDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllTagsByIdChatsIdTagsAllDeleteRequest struct via the builder pattern


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


## DeleteAllUserChatsChatsDelete

> bool DeleteAllUserChatsChatsDelete(ctx).Execute()

Delete All User Chats

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
	resp, r, err := apiClient.ChatsAPI.DeleteAllUserChatsChatsDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.DeleteAllUserChatsChatsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllUserChatsChatsDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.DeleteAllUserChatsChatsDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllUserChatsChatsDeleteRequest struct via the builder pattern


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


## DeleteChatByIdChatsIdDelete

> bool DeleteChatByIdChatsIdDelete(ctx, id).Execute()

Delete Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.DeleteChatByIdChatsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.DeleteChatByIdChatsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteChatByIdChatsIdDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.DeleteChatByIdChatsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChatByIdChatsIdDeleteRequest struct via the builder pattern


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


## DeleteSharedChatByIdChatsIdShareDelete

> bool DeleteSharedChatByIdChatsIdShareDelete(ctx, id).Execute()

Delete Shared Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.DeleteSharedChatByIdChatsIdShareDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.DeleteSharedChatByIdChatsIdShareDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSharedChatByIdChatsIdShareDelete`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.DeleteSharedChatByIdChatsIdShareDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedChatByIdChatsIdShareDeleteRequest struct via the builder pattern


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


## DeleteTagByIdAndTagNameChatsIdTagsDelete

> []TagModel DeleteTagByIdAndTagNameChatsIdTagsDelete(ctx, id).TagForm(tagForm).Execute()

Delete Tag By Id And Tag Name

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
	tagForm := *openapiclient.NewTagForm("Name_example") // TagForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.DeleteTagByIdAndTagNameChatsIdTagsDelete(context.Background(), id).TagForm(tagForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.DeleteTagByIdAndTagNameChatsIdTagsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTagByIdAndTagNameChatsIdTagsDelete`: []TagModel
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.DeleteTagByIdAndTagNameChatsIdTagsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagByIdAndTagNameChatsIdTagsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagForm** | [**TagForm**](TagForm.md) |  | 

### Return type

[**[]TagModel**](TagModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUserChatsInDbChatsAllDbGet

> []ChatResponse GetAllUserChatsInDbChatsAllDbGet(ctx).Execute()

Get All User Chats In Db

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
	resp, r, err := apiClient.ChatsAPI.GetAllUserChatsInDbChatsAllDbGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetAllUserChatsInDbChatsAllDbGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllUserChatsInDbChatsAllDbGet`: []ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetAllUserChatsInDbChatsAllDbGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUserChatsInDbChatsAllDbGetRequest struct via the builder pattern


### Return type

[**[]ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUserTagsChatsAllTagsGet

> []TagModel GetAllUserTagsChatsAllTagsGet(ctx).Execute()

Get All User Tags

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
	resp, r, err := apiClient.ChatsAPI.GetAllUserTagsChatsAllTagsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetAllUserTagsChatsAllTagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllUserTagsChatsAllTagsGet`: []TagModel
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetAllUserTagsChatsAllTagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUserTagsChatsAllTagsGetRequest struct via the builder pattern


### Return type

[**[]TagModel**](TagModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedSessionUserChatListChatsArchivedGet

> []ChatTitleIdResponse GetArchivedSessionUserChatListChatsArchivedGet(ctx).Skip(skip).Limit(limit).Execute()

Get Archived Session User Chat List

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
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetArchivedSessionUserChatListChatsArchivedGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetArchivedSessionUserChatListChatsArchivedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArchivedSessionUserChatListChatsArchivedGet`: []ChatTitleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetArchivedSessionUserChatListChatsArchivedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedSessionUserChatListChatsArchivedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]ChatTitleIdResponse**](ChatTitleIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatByIdChatsIdGet

> ChatResponse GetChatByIdChatsIdGet(ctx, id).Execute()

Get Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.GetChatByIdChatsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetChatByIdChatsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatByIdChatsIdGet`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetChatByIdChatsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatByIdChatsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatTagsByIdChatsIdTagsGet

> []TagModel GetChatTagsByIdChatsIdTagsGet(ctx, id).Execute()

Get Chat Tags By Id

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
	resp, r, err := apiClient.ChatsAPI.GetChatTagsByIdChatsIdTagsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetChatTagsByIdChatsIdTagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatTagsByIdChatsIdTagsGet`: []TagModel
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetChatTagsByIdChatsIdTagsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatTagsByIdChatsIdTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TagModel**](TagModel.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatsByFolderIdChatsFolderFolderIdGet

> []ChatResponse GetChatsByFolderIdChatsFolderFolderIdGet(ctx, folderId).Execute()

Get Chats By Folder Id

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
	folderId := "folderId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetChatsByFolderIdChatsFolderFolderIdGet(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetChatsByFolderIdChatsFolderFolderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatsByFolderIdChatsFolderFolderIdGet`: []ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetChatsByFolderIdChatsFolderFolderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatsByFolderIdChatsFolderFolderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPinnedStatusByIdChatsIdPinnedGet

> bool GetPinnedStatusByIdChatsIdPinnedGet(ctx, id).Execute()

Get Pinned Status By Id

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
	resp, r, err := apiClient.ChatsAPI.GetPinnedStatusByIdChatsIdPinnedGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetPinnedStatusByIdChatsIdPinnedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPinnedStatusByIdChatsIdPinnedGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetPinnedStatusByIdChatsIdPinnedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPinnedStatusByIdChatsIdPinnedGetRequest struct via the builder pattern


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


## GetSessionUserChatListChatsGet

> []ChatTitleIdResponse GetSessionUserChatListChatsGet(ctx).Page(page).Execute()

Get Session User Chat List

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
	page := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetSessionUserChatListChatsGet(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetSessionUserChatListChatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionUserChatListChatsGet`: []ChatTitleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetSessionUserChatListChatsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionUserChatListChatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 

### Return type

[**[]ChatTitleIdResponse**](ChatTitleIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionUserChatListChatsListGet

> []ChatTitleIdResponse GetSessionUserChatListChatsListGet(ctx).Page(page).Execute()

Get Session User Chat List

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
	page := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetSessionUserChatListChatsListGet(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetSessionUserChatListChatsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionUserChatListChatsListGet`: []ChatTitleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetSessionUserChatListChatsListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionUserChatListChatsListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 

### Return type

[**[]ChatTitleIdResponse**](ChatTitleIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedChatByIdChatsShareShareIdGet

> ChatResponse GetSharedChatByIdChatsShareShareIdGet(ctx, shareId).Execute()

Get Shared Chat By Id

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
	shareId := "shareId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetSharedChatByIdChatsShareShareIdGet(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetSharedChatByIdChatsShareShareIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedChatByIdChatsShareShareIdGet`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetSharedChatByIdChatsShareShareIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedChatByIdChatsShareShareIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserArchivedChatsChatsAllArchivedGet

> []ChatResponse GetUserArchivedChatsChatsAllArchivedGet(ctx).Execute()

Get User Archived Chats

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
	resp, r, err := apiClient.ChatsAPI.GetUserArchivedChatsChatsAllArchivedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetUserArchivedChatsChatsAllArchivedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserArchivedChatsChatsAllArchivedGet`: []ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetUserArchivedChatsChatsAllArchivedGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserArchivedChatsChatsAllArchivedGetRequest struct via the builder pattern


### Return type

[**[]ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserChatListByTagNameChatsTagsPost

> []ChatTitleIdResponse GetUserChatListByTagNameChatsTagsPost(ctx).TagFilterForm(tagFilterForm).Execute()

Get User Chat List By Tag Name

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
	tagFilterForm := *openapiclient.NewTagFilterForm("Name_example") // TagFilterForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetUserChatListByTagNameChatsTagsPost(context.Background()).TagFilterForm(tagFilterForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetUserChatListByTagNameChatsTagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserChatListByTagNameChatsTagsPost`: []ChatTitleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetUserChatListByTagNameChatsTagsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserChatListByTagNameChatsTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagFilterForm** | [**TagFilterForm**](TagFilterForm.md) |  | 

### Return type

[**[]ChatTitleIdResponse**](ChatTitleIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserChatListByUserIdChatsListUserUserIdGet

> []ChatTitleIdResponse GetUserChatListByUserIdChatsListUserUserIdGet(ctx, userId).Skip(skip).Limit(limit).Execute()

Get User Chat List By User Id

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
	userId := "userId_example" // string | 
	skip := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.GetUserChatListByUserIdChatsListUserUserIdGet(context.Background(), userId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetUserChatListByUserIdChatsListUserUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserChatListByUserIdChatsListUserUserIdGet`: []ChatTitleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetUserChatListByUserIdChatsListUserUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserChatListByUserIdChatsListUserUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]ChatTitleIdResponse**](ChatTitleIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserChatsChatsAllGet

> []ChatResponse GetUserChatsChatsAllGet(ctx).Execute()

Get User Chats

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
	resp, r, err := apiClient.ChatsAPI.GetUserChatsChatsAllGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetUserChatsChatsAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserChatsChatsAllGet`: []ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetUserChatsChatsAllGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserChatsChatsAllGetRequest struct via the builder pattern


### Return type

[**[]ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPinnedChatsChatsPinnedGet

> []ChatResponse GetUserPinnedChatsChatsPinnedGet(ctx).Execute()

Get User Pinned Chats

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
	resp, r, err := apiClient.ChatsAPI.GetUserPinnedChatsChatsPinnedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.GetUserPinnedChatsChatsPinnedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPinnedChatsChatsPinnedGet`: []ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.GetUserPinnedChatsChatsPinnedGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPinnedChatsChatsPinnedGetRequest struct via the builder pattern


### Return type

[**[]ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportChatChatsImportPost

> ChatResponse ImportChatChatsImportPost(ctx).ChatImportForm(chatImportForm).Execute()

Import Chat

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
	chatImportForm := *openapiclient.NewChatImportForm(map[string]interface{}(123)) // ChatImportForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.ImportChatChatsImportPost(context.Background()).ChatImportForm(chatImportForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.ImportChatChatsImportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportChatChatsImportPost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.ImportChatChatsImportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportChatChatsImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatImportForm** | [**ChatImportForm**](ChatImportForm.md) |  | 

### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinChatByIdChatsIdPinPost

> ChatResponse PinChatByIdChatsIdPinPost(ctx, id).Execute()

Pin Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.PinChatByIdChatsIdPinPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.PinChatByIdChatsIdPinPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PinChatByIdChatsIdPinPost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.PinChatByIdChatsIdPinPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinChatByIdChatsIdPinPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUserChatsChatsSearchGet

> []ChatTitleIdResponse SearchUserChatsChatsSearchGet(ctx).Text(text).Page(page).Execute()

Search User Chats

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
	text := "text_example" // string | 
	page := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.SearchUserChatsChatsSearchGet(context.Background()).Text(text).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.SearchUserChatsChatsSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUserChatsChatsSearchGet`: []ChatTitleIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.SearchUserChatsChatsSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUserChatsChatsSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  | 
 **page** | **int32** |  | 

### Return type

[**[]ChatTitleIdResponse**](ChatTitleIdResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareChatByIdChatsIdSharePost

> ChatResponse ShareChatByIdChatsIdSharePost(ctx, id).Execute()

Share Chat By Id

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
	resp, r, err := apiClient.ChatsAPI.ShareChatByIdChatsIdSharePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.ShareChatByIdChatsIdSharePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareChatByIdChatsIdSharePost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.ShareChatByIdChatsIdSharePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareChatByIdChatsIdSharePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChatByIdChatsIdPost

> ChatResponse UpdateChatByIdChatsIdPost(ctx, id).ChatForm(chatForm).Execute()

Update Chat By Id

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
	chatForm := *openapiclient.NewChatForm(map[string]interface{}(123)) // ChatForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.UpdateChatByIdChatsIdPost(context.Background(), id).ChatForm(chatForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.UpdateChatByIdChatsIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChatByIdChatsIdPost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.UpdateChatByIdChatsIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChatByIdChatsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatForm** | [**ChatForm**](ChatForm.md) |  | 

### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChatFolderIdByIdChatsIdFolderPost

> ChatResponse UpdateChatFolderIdByIdChatsIdFolderPost(ctx, id).ChatFolderIdForm(chatFolderIdForm).Execute()

Update Chat Folder Id By Id

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
	chatFolderIdForm := *openapiclient.NewChatFolderIdForm() // ChatFolderIdForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatsAPI.UpdateChatFolderIdByIdChatsIdFolderPost(context.Background(), id).ChatFolderIdForm(chatFolderIdForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatsAPI.UpdateChatFolderIdByIdChatsIdFolderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChatFolderIdByIdChatsIdFolderPost`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatsAPI.UpdateChatFolderIdByIdChatsIdFolderPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChatFolderIdByIdChatsIdFolderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatFolderIdForm** | [**ChatFolderIdForm**](ChatFolderIdForm.md) |  | 

### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[HTTPBearer](../README.md#HTTPBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


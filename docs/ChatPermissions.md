# ChatPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUpload** | **bool** |  | 
**Delete** | **bool** |  | 
**Edit** | **bool** |  | 
**Temporary** | **bool** |  | 

## Methods

### NewChatPermissions

`func NewChatPermissions(fileUpload bool, delete bool, edit bool, temporary bool, ) *ChatPermissions`

NewChatPermissions instantiates a new ChatPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPermissionsWithDefaults

`func NewChatPermissionsWithDefaults() *ChatPermissions`

NewChatPermissionsWithDefaults instantiates a new ChatPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUpload

`func (o *ChatPermissions) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *ChatPermissions) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *ChatPermissions) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.


### GetDelete

`func (o *ChatPermissions) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ChatPermissions) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ChatPermissions) SetDelete(v bool)`

SetDelete sets Delete field to given value.


### GetEdit

`func (o *ChatPermissions) GetEdit() bool`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *ChatPermissions) GetEditOk() (*bool, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *ChatPermissions) SetEdit(v bool)`

SetEdit sets Edit field to given value.


### GetTemporary

`func (o *ChatPermissions) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *ChatPermissions) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *ChatPermissions) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



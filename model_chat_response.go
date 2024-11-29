/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ChatResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatResponse{}

// ChatResponse struct for ChatResponse
type ChatResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Title string `json:"title"`
	Chat map[string]interface{} `json:"chat"`
	UpdatedAt int32 `json:"updated_at"`
	CreatedAt int32 `json:"created_at"`
	ShareId NullableString `json:"share_id,omitempty"`
	Archived bool `json:"archived"`
	Pinned NullableBool `json:"pinned,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	FolderId NullableString `json:"folder_id,omitempty"`
}

type _ChatResponse ChatResponse

// NewChatResponse instantiates a new ChatResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatResponse(id string, userId string, title string, chat map[string]interface{}, updatedAt int32, createdAt int32, archived bool) *ChatResponse {
	this := ChatResponse{}
	this.Id = id
	this.UserId = userId
	this.Title = title
	this.Chat = chat
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.Archived = archived
	return &this
}

// NewChatResponseWithDefaults instantiates a new ChatResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatResponseWithDefaults() *ChatResponse {
	this := ChatResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ChatResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChatResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *ChatResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ChatResponse) SetUserId(v string) {
	o.UserId = v
}

// GetTitle returns the Title field value
func (o *ChatResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ChatResponse) SetTitle(v string) {
	o.Title = v
}

// GetChat returns the Chat field value
func (o *ChatResponse) GetChat() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Chat
}

// GetChatOk returns a tuple with the Chat field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetChatOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Chat, true
}

// SetChat sets field value
func (o *ChatResponse) SetChat(v map[string]interface{}) {
	o.Chat = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ChatResponse) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ChatResponse) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChatResponse) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChatResponse) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetShareId returns the ShareId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatResponse) GetShareId() string {
	if o == nil || IsNil(o.ShareId.Get()) {
		var ret string
		return ret
	}
	return *o.ShareId.Get()
}

// GetShareIdOk returns a tuple with the ShareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatResponse) GetShareIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareId.Get(), o.ShareId.IsSet()
}

// HasShareId returns a boolean if a field has been set.
func (o *ChatResponse) HasShareId() bool {
	if o != nil && o.ShareId.IsSet() {
		return true
	}

	return false
}

// SetShareId gets a reference to the given NullableString and assigns it to the ShareId field.
func (o *ChatResponse) SetShareId(v string) {
	o.ShareId.Set(&v)
}
// SetShareIdNil sets the value for ShareId to be an explicit nil
func (o *ChatResponse) SetShareIdNil() {
	o.ShareId.Set(nil)
}

// UnsetShareId ensures that no value is present for ShareId, not even an explicit nil
func (o *ChatResponse) UnsetShareId() {
	o.ShareId.Unset()
}

// GetArchived returns the Archived field value
func (o *ChatResponse) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *ChatResponse) SetArchived(v bool) {
	o.Archived = v
}

// GetPinned returns the Pinned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatResponse) GetPinned() bool {
	if o == nil || IsNil(o.Pinned.Get()) {
		var ret bool
		return ret
	}
	return *o.Pinned.Get()
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatResponse) GetPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pinned.Get(), o.Pinned.IsSet()
}

// HasPinned returns a boolean if a field has been set.
func (o *ChatResponse) HasPinned() bool {
	if o != nil && o.Pinned.IsSet() {
		return true
	}

	return false
}

// SetPinned gets a reference to the given NullableBool and assigns it to the Pinned field.
func (o *ChatResponse) SetPinned(v bool) {
	o.Pinned.Set(&v)
}
// SetPinnedNil sets the value for Pinned to be an explicit nil
func (o *ChatResponse) SetPinnedNil() {
	o.Pinned.Set(nil)
}

// UnsetPinned ensures that no value is present for Pinned, not even an explicit nil
func (o *ChatResponse) UnsetPinned() {
	o.Pinned.Unset()
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ChatResponse) GetMeta() map[string]interface{} {
	if o == nil || IsNil(o.Meta) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ChatResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ChatResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatResponse) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatResponse) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *ChatResponse) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *ChatResponse) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *ChatResponse) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *ChatResponse) UnsetFolderId() {
	o.FolderId.Unset()
}

func (o ChatResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["title"] = o.Title
	toSerialize["chat"] = o.Chat
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	if o.ShareId.IsSet() {
		toSerialize["share_id"] = o.ShareId.Get()
	}
	toSerialize["archived"] = o.Archived
	if o.Pinned.IsSet() {
		toSerialize["pinned"] = o.Pinned.Get()
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if o.FolderId.IsSet() {
		toSerialize["folder_id"] = o.FolderId.Get()
	}
	return toSerialize, nil
}

func (o *ChatResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"title",
		"chat",
		"updated_at",
		"created_at",
		"archived",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChatResponse := _ChatResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatResponse)

	if err != nil {
		return err
	}

	*o = ChatResponse(varChatResponse)

	return err
}

type NullableChatResponse struct {
	value *ChatResponse
	isSet bool
}

func (v NullableChatResponse) Get() *ChatResponse {
	return v.value
}

func (v *NullableChatResponse) Set(val *ChatResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatResponse(val *ChatResponse) *NullableChatResponse {
	return &NullableChatResponse{value: val, isSet: true}
}

func (v NullableChatResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



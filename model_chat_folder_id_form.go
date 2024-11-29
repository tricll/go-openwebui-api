/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openwebuiapi

import (
	"encoding/json"
)

// checks if the ChatFolderIdForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatFolderIdForm{}

// ChatFolderIdForm struct for ChatFolderIdForm
type ChatFolderIdForm struct {
	FolderId NullableString `json:"folder_id,omitempty"`
}

// NewChatFolderIdForm instantiates a new ChatFolderIdForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatFolderIdForm() *ChatFolderIdForm {
	this := ChatFolderIdForm{}
	return &this
}

// NewChatFolderIdFormWithDefaults instantiates a new ChatFolderIdForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatFolderIdFormWithDefaults() *ChatFolderIdForm {
	this := ChatFolderIdForm{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatFolderIdForm) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatFolderIdForm) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *ChatFolderIdForm) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *ChatFolderIdForm) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *ChatFolderIdForm) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *ChatFolderIdForm) UnsetFolderId() {
	o.FolderId.Unset()
}

func (o ChatFolderIdForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatFolderIdForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderId.IsSet() {
		toSerialize["folder_id"] = o.FolderId.Get()
	}
	return toSerialize, nil
}

type NullableChatFolderIdForm struct {
	value *ChatFolderIdForm
	isSet bool
}

func (v NullableChatFolderIdForm) Get() *ChatFolderIdForm {
	return v.value
}

func (v *NullableChatFolderIdForm) Set(val *ChatFolderIdForm) {
	v.value = val
	v.isSet = true
}

func (v NullableChatFolderIdForm) IsSet() bool {
	return v.isSet
}

func (v *NullableChatFolderIdForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatFolderIdForm(val *ChatFolderIdForm) *NullableChatFolderIdForm {
	return &NullableChatFolderIdForm{value: val, isSet: true}
}

func (v NullableChatFolderIdForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatFolderIdForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



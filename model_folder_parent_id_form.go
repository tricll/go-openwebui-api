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

// checks if the FolderParentIdForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderParentIdForm{}

// FolderParentIdForm struct for FolderParentIdForm
type FolderParentIdForm struct {
	ParentId NullableString `json:"parent_id,omitempty"`
}

// NewFolderParentIdForm instantiates a new FolderParentIdForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderParentIdForm() *FolderParentIdForm {
	this := FolderParentIdForm{}
	return &this
}

// NewFolderParentIdFormWithDefaults instantiates a new FolderParentIdForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderParentIdFormWithDefaults() *FolderParentIdForm {
	this := FolderParentIdForm{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderParentIdForm) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderParentIdForm) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *FolderParentIdForm) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *FolderParentIdForm) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *FolderParentIdForm) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *FolderParentIdForm) UnsetParentId() {
	o.ParentId.Unset()
}

func (o FolderParentIdForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderParentIdForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	return toSerialize, nil
}

type NullableFolderParentIdForm struct {
	value *FolderParentIdForm
	isSet bool
}

func (v NullableFolderParentIdForm) Get() *FolderParentIdForm {
	return v.value
}

func (v *NullableFolderParentIdForm) Set(val *FolderParentIdForm) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderParentIdForm) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderParentIdForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderParentIdForm(val *FolderParentIdForm) *NullableFolderParentIdForm {
	return &NullableFolderParentIdForm{value: val, isSet: true}
}

func (v NullableFolderParentIdForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderParentIdForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



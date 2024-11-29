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

// checks if the FileMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileMeta{}

// FileMeta struct for FileMeta
type FileMeta struct {
	Name *Name `json:"name,omitempty"`
	ContentType *ContentType `json:"content_type,omitempty"`
	Size *Size `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FileMeta FileMeta

// NewFileMeta instantiates a new FileMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileMeta() *FileMeta {
	this := FileMeta{}
	return &this
}

// NewFileMetaWithDefaults instantiates a new FileMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileMetaWithDefaults() *FileMeta {
	this := FileMeta{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileMeta) GetName() Name {
	if o == nil || IsNil(o.Name) {
		var ret Name
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMeta) GetNameOk() (*Name, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name and assigns it to the Name field.
func (o *FileMeta) SetName(v Name) {
	o.Name = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *FileMeta) GetContentType() ContentType {
	if o == nil || IsNil(o.ContentType) {
		var ret ContentType
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMeta) GetContentTypeOk() (*ContentType, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *FileMeta) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given ContentType and assigns it to the ContentType field.
func (o *FileMeta) SetContentType(v ContentType) {
	o.ContentType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileMeta) GetSize() Size {
	if o == nil || IsNil(o.Size) {
		var ret Size
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileMeta) GetSizeOk() (*Size, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileMeta) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given Size and assigns it to the Size field.
func (o *FileMeta) SetSize(v Size) {
	o.Size = &v
}

func (o FileMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileMeta) UnmarshalJSON(data []byte) (err error) {
	varFileMeta := _FileMeta{}

	err = json.Unmarshal(data, &varFileMeta)

	if err != nil {
		return err
	}

	*o = FileMeta(varFileMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileMeta struct {
	value *FileMeta
	isSet bool
}

func (v NullableFileMeta) Get() *FileMeta {
	return v.value
}

func (v *NullableFileMeta) Set(val *FileMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableFileMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableFileMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileMeta(val *FileMeta) *NullableFileMeta {
	return &NullableFileMeta{value: val, isSet: true}
}

func (v NullableFileMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



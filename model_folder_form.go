/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the FolderForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderForm{}

// FolderForm struct for FolderForm
type FolderForm struct {
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _FolderForm FolderForm

// NewFolderForm instantiates a new FolderForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderForm(name string) *FolderForm {
	this := FolderForm{}
	this.Name = name
	return &this
}

// NewFolderFormWithDefaults instantiates a new FolderForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderFormWithDefaults() *FolderForm {
	this := FolderForm{}
	return &this
}

// GetName returns the Name field value
func (o *FolderForm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FolderForm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FolderForm) SetName(v string) {
	o.Name = v
}

func (o FolderForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FolderForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varFolderForm := _FolderForm{}

	err = json.Unmarshal(data, &varFolderForm)

	if err != nil {
		return err
	}

	*o = FolderForm(varFolderForm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFolderForm struct {
	value *FolderForm
	isSet bool
}

func (v NullableFolderForm) Get() *FolderForm {
	return v.value
}

func (v *NullableFolderForm) Set(val *FolderForm) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderForm) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderForm(val *FolderForm) *NullableFolderForm {
	return &NullableFolderForm{value: val, isSet: true}
}

func (v NullableFolderForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



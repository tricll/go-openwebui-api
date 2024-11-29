/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openwebuiapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateProfileForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfileForm{}

// UpdateProfileForm struct for UpdateProfileForm
type UpdateProfileForm struct {
	ProfileImageUrl string `json:"profile_image_url"`
	Name string `json:"name"`
}

type _UpdateProfileForm UpdateProfileForm

// NewUpdateProfileForm instantiates a new UpdateProfileForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfileForm(profileImageUrl string, name string) *UpdateProfileForm {
	this := UpdateProfileForm{}
	this.ProfileImageUrl = profileImageUrl
	this.Name = name
	return &this
}

// NewUpdateProfileFormWithDefaults instantiates a new UpdateProfileForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfileFormWithDefaults() *UpdateProfileForm {
	this := UpdateProfileForm{}
	return &this
}

// GetProfileImageUrl returns the ProfileImageUrl field value
func (o *UpdateProfileForm) GetProfileImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileImageUrl
}

// GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateProfileForm) GetProfileImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileImageUrl, true
}

// SetProfileImageUrl sets field value
func (o *UpdateProfileForm) SetProfileImageUrl(v string) {
	o.ProfileImageUrl = v
}

// GetName returns the Name field value
func (o *UpdateProfileForm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProfileForm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProfileForm) SetName(v string) {
	o.Name = v
}

func (o UpdateProfileForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfileForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile_image_url"] = o.ProfileImageUrl
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateProfileForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile_image_url",
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

	varUpdateProfileForm := _UpdateProfileForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateProfileForm)

	if err != nil {
		return err
	}

	*o = UpdateProfileForm(varUpdateProfileForm)

	return err
}

type NullableUpdateProfileForm struct {
	value *UpdateProfileForm
	isSet bool
}

func (v NullableUpdateProfileForm) Get() *UpdateProfileForm {
	return v.value
}

func (v *NullableUpdateProfileForm) Set(val *UpdateProfileForm) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfileForm) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfileForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfileForm(val *UpdateProfileForm) *NullableUpdateProfileForm {
	return &NullableUpdateProfileForm{value: val, isSet: true}
}

func (v NullableUpdateProfileForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfileForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



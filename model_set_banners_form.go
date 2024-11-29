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

// checks if the SetBannersForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetBannersForm{}

// SetBannersForm struct for SetBannersForm
type SetBannersForm struct {
	Banners []BannerModel `json:"banners"`
}

type _SetBannersForm SetBannersForm

// NewSetBannersForm instantiates a new SetBannersForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBannersForm(banners []BannerModel) *SetBannersForm {
	this := SetBannersForm{}
	this.Banners = banners
	return &this
}

// NewSetBannersFormWithDefaults instantiates a new SetBannersForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBannersFormWithDefaults() *SetBannersForm {
	this := SetBannersForm{}
	return &this
}

// GetBanners returns the Banners field value
func (o *SetBannersForm) GetBanners() []BannerModel {
	if o == nil {
		var ret []BannerModel
		return ret
	}

	return o.Banners
}

// GetBannersOk returns a tuple with the Banners field value
// and a boolean to check if the value has been set.
func (o *SetBannersForm) GetBannersOk() ([]BannerModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Banners, true
}

// SetBanners sets field value
func (o *SetBannersForm) SetBanners(v []BannerModel) {
	o.Banners = v
}

func (o SetBannersForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetBannersForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["banners"] = o.Banners
	return toSerialize, nil
}

func (o *SetBannersForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"banners",
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

	varSetBannersForm := _SetBannersForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetBannersForm)

	if err != nil {
		return err
	}

	*o = SetBannersForm(varSetBannersForm)

	return err
}

type NullableSetBannersForm struct {
	value *SetBannersForm
	isSet bool
}

func (v NullableSetBannersForm) Get() *SetBannersForm {
	return v.value
}

func (v *NullableSetBannersForm) Set(val *SetBannersForm) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBannersForm) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBannersForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBannersForm(val *SetBannersForm) *NullableSetBannersForm {
	return &NullableSetBannersForm{value: val, isSet: true}
}

func (v NullableSetBannersForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBannersForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



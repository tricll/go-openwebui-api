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

// checks if the TagFilterForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagFilterForm{}

// TagFilterForm struct for TagFilterForm
type TagFilterForm struct {
	Name string `json:"name"`
	Skip NullableInt32 `json:"skip,omitempty"`
	Limit NullableInt32 `json:"limit,omitempty"`
}

type _TagFilterForm TagFilterForm

// NewTagFilterForm instantiates a new TagFilterForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagFilterForm(name string) *TagFilterForm {
	this := TagFilterForm{}
	this.Name = name
	return &this
}

// NewTagFilterFormWithDefaults instantiates a new TagFilterForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagFilterFormWithDefaults() *TagFilterForm {
	this := TagFilterForm{}
	return &this
}

// GetName returns the Name field value
func (o *TagFilterForm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagFilterForm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagFilterForm) SetName(v string) {
	o.Name = v
}

// GetSkip returns the Skip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagFilterForm) GetSkip() int32 {
	if o == nil || IsNil(o.Skip.Get()) {
		var ret int32
		return ret
	}
	return *o.Skip.Get()
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagFilterForm) GetSkipOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skip.Get(), o.Skip.IsSet()
}

// HasSkip returns a boolean if a field has been set.
func (o *TagFilterForm) HasSkip() bool {
	if o != nil && o.Skip.IsSet() {
		return true
	}

	return false
}

// SetSkip gets a reference to the given NullableInt32 and assigns it to the Skip field.
func (o *TagFilterForm) SetSkip(v int32) {
	o.Skip.Set(&v)
}
// SetSkipNil sets the value for Skip to be an explicit nil
func (o *TagFilterForm) SetSkipNil() {
	o.Skip.Set(nil)
}

// UnsetSkip ensures that no value is present for Skip, not even an explicit nil
func (o *TagFilterForm) UnsetSkip() {
	o.Skip.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagFilterForm) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagFilterForm) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *TagFilterForm) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *TagFilterForm) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *TagFilterForm) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *TagFilterForm) UnsetLimit() {
	o.Limit.Unset()
}

func (o TagFilterForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagFilterForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Skip.IsSet() {
		toSerialize["skip"] = o.Skip.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	return toSerialize, nil
}

func (o *TagFilterForm) UnmarshalJSON(data []byte) (err error) {
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

	varTagFilterForm := _TagFilterForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagFilterForm)

	if err != nil {
		return err
	}

	*o = TagFilterForm(varTagFilterForm)

	return err
}

type NullableTagFilterForm struct {
	value *TagFilterForm
	isSet bool
}

func (v NullableTagFilterForm) Get() *TagFilterForm {
	return v.value
}

func (v *NullableTagFilterForm) Set(val *TagFilterForm) {
	v.value = val
	v.isSet = true
}

func (v NullableTagFilterForm) IsSet() bool {
	return v.isSet
}

func (v *NullableTagFilterForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagFilterForm(val *TagFilterForm) *NullableTagFilterForm {
	return &NullableTagFilterForm{value: val, isSet: true}
}

func (v NullableTagFilterForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagFilterForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



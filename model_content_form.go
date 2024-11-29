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

// checks if the ContentForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentForm{}

// ContentForm struct for ContentForm
type ContentForm struct {
	Content string `json:"content"`
}

type _ContentForm ContentForm

// NewContentForm instantiates a new ContentForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentForm(content string) *ContentForm {
	this := ContentForm{}
	this.Content = content
	return &this
}

// NewContentFormWithDefaults instantiates a new ContentForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentFormWithDefaults() *ContentForm {
	this := ContentForm{}
	return &this
}

// GetContent returns the Content field value
func (o *ContentForm) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ContentForm) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ContentForm) SetContent(v string) {
	o.Content = v
}

func (o ContentForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *ContentForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
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

	varContentForm := _ContentForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentForm)

	if err != nil {
		return err
	}

	*o = ContentForm(varContentForm)

	return err
}

type NullableContentForm struct {
	value *ContentForm
	isSet bool
}

func (v NullableContentForm) Get() *ContentForm {
	return v.value
}

func (v *NullableContentForm) Set(val *ContentForm) {
	v.value = val
	v.isSet = true
}

func (v NullableContentForm) IsSet() bool {
	return v.isSet
}

func (v *NullableContentForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentForm(val *ContentForm) *NullableContentForm {
	return &NullableContentForm{value: val, isSet: true}
}

func (v NullableContentForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



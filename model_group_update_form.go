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

// checks if the GroupUpdateForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUpdateForm{}

// GroupUpdateForm struct for GroupUpdateForm
type GroupUpdateForm struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	UserIds []string `json:"user_ids,omitempty"`
	AdminIds []string `json:"admin_ids,omitempty"`
}

type _GroupUpdateForm GroupUpdateForm

// NewGroupUpdateForm instantiates a new GroupUpdateForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdateForm(name string, description string) *GroupUpdateForm {
	this := GroupUpdateForm{}
	this.Name = name
	this.Description = description
	return &this
}

// NewGroupUpdateFormWithDefaults instantiates a new GroupUpdateForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateFormWithDefaults() *GroupUpdateForm {
	this := GroupUpdateForm{}
	return &this
}

// GetName returns the Name field value
func (o *GroupUpdateForm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupUpdateForm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupUpdateForm) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *GroupUpdateForm) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GroupUpdateForm) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GroupUpdateForm) SetDescription(v string) {
	o.Description = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUpdateForm) GetPermissions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUpdateForm) GetPermissionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Permissions) {
		return map[string]interface{}{}, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupUpdateForm) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]interface{} and assigns it to the Permissions field.
func (o *GroupUpdateForm) SetPermissions(v map[string]interface{}) {
	o.Permissions = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUpdateForm) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUpdateForm) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *GroupUpdateForm) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *GroupUpdateForm) SetUserIds(v []string) {
	o.UserIds = v
}

// GetAdminIds returns the AdminIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUpdateForm) GetAdminIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AdminIds
}

// GetAdminIdsOk returns a tuple with the AdminIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUpdateForm) GetAdminIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdminIds) {
		return nil, false
	}
	return o.AdminIds, true
}

// HasAdminIds returns a boolean if a field has been set.
func (o *GroupUpdateForm) HasAdminIds() bool {
	if o != nil && !IsNil(o.AdminIds) {
		return true
	}

	return false
}

// SetAdminIds gets a reference to the given []string and assigns it to the AdminIds field.
func (o *GroupUpdateForm) SetAdminIds(v []string) {
	o.AdminIds = v
}

func (o GroupUpdateForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUpdateForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.UserIds != nil {
		toSerialize["user_ids"] = o.UserIds
	}
	if o.AdminIds != nil {
		toSerialize["admin_ids"] = o.AdminIds
	}
	return toSerialize, nil
}

func (o *GroupUpdateForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
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

	varGroupUpdateForm := _GroupUpdateForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupUpdateForm)

	if err != nil {
		return err
	}

	*o = GroupUpdateForm(varGroupUpdateForm)

	return err
}

type NullableGroupUpdateForm struct {
	value *GroupUpdateForm
	isSet bool
}

func (v NullableGroupUpdateForm) Get() *GroupUpdateForm {
	return v.value
}

func (v *NullableGroupUpdateForm) Set(val *GroupUpdateForm) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateForm) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateForm(val *GroupUpdateForm) *NullableGroupUpdateForm {
	return &NullableGroupUpdateForm{value: val, isSet: true}
}

func (v NullableGroupUpdateForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



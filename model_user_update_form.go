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

// checks if the UserUpdateForm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdateForm{}

// UserUpdateForm struct for UserUpdateForm
type UserUpdateForm struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ProfileImageUrl string `json:"profile_image_url"`
	Password NullableString `json:"password,omitempty"`
}

type _UserUpdateForm UserUpdateForm

// NewUserUpdateForm instantiates a new UserUpdateForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateForm(name string, email string, profileImageUrl string) *UserUpdateForm {
	this := UserUpdateForm{}
	this.Name = name
	this.Email = email
	this.ProfileImageUrl = profileImageUrl
	return &this
}

// NewUserUpdateFormWithDefaults instantiates a new UserUpdateForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateFormWithDefaults() *UserUpdateForm {
	this := UserUpdateForm{}
	return &this
}

// GetName returns the Name field value
func (o *UserUpdateForm) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserUpdateForm) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserUpdateForm) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *UserUpdateForm) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserUpdateForm) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserUpdateForm) SetEmail(v string) {
	o.Email = v
}

// GetProfileImageUrl returns the ProfileImageUrl field value
func (o *UserUpdateForm) GetProfileImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileImageUrl
}

// GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field value
// and a boolean to check if the value has been set.
func (o *UserUpdateForm) GetProfileImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileImageUrl, true
}

// SetProfileImageUrl sets field value
func (o *UserUpdateForm) SetProfileImageUrl(v string) {
	o.ProfileImageUrl = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdateForm) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdateForm) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UserUpdateForm) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UserUpdateForm) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UserUpdateForm) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UserUpdateForm) UnsetPassword() {
	o.Password.Unset()
}

func (o UserUpdateForm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdateForm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["profile_image_url"] = o.ProfileImageUrl
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

func (o *UserUpdateForm) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"email",
		"profile_image_url",
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

	varUserUpdateForm := _UserUpdateForm{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserUpdateForm)

	if err != nil {
		return err
	}

	*o = UserUpdateForm(varUserUpdateForm)

	return err
}

type NullableUserUpdateForm struct {
	value *UserUpdateForm
	isSet bool
}

func (v NullableUserUpdateForm) Get() *UserUpdateForm {
	return v.value
}

func (v *NullableUserUpdateForm) Set(val *UserUpdateForm) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateForm) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateForm(val *UserUpdateForm) *NullableUserUpdateForm {
	return &NullableUserUpdateForm{value: val, isSet: true}
}

func (v NullableUserUpdateForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



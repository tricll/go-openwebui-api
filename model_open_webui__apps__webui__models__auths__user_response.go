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

// checks if the OpenWebuiAppsWebuiModelsAuthsUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenWebuiAppsWebuiModelsAuthsUserResponse{}

// OpenWebuiAppsWebuiModelsAuthsUserResponse struct for OpenWebuiAppsWebuiModelsAuthsUserResponse
type OpenWebuiAppsWebuiModelsAuthsUserResponse struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Role string `json:"role"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type _OpenWebuiAppsWebuiModelsAuthsUserResponse OpenWebuiAppsWebuiModelsAuthsUserResponse

// NewOpenWebuiAppsWebuiModelsAuthsUserResponse instantiates a new OpenWebuiAppsWebuiModelsAuthsUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenWebuiAppsWebuiModelsAuthsUserResponse(id string, email string, name string, role string, profileImageUrl string) *OpenWebuiAppsWebuiModelsAuthsUserResponse {
	this := OpenWebuiAppsWebuiModelsAuthsUserResponse{}
	this.Id = id
	this.Email = email
	this.Name = name
	this.Role = role
	this.ProfileImageUrl = profileImageUrl
	return &this
}

// NewOpenWebuiAppsWebuiModelsAuthsUserResponseWithDefaults instantiates a new OpenWebuiAppsWebuiModelsAuthsUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenWebuiAppsWebuiModelsAuthsUserResponseWithDefaults() *OpenWebuiAppsWebuiModelsAuthsUserResponse {
	this := OpenWebuiAppsWebuiModelsAuthsUserResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) SetRole(v string) {
	o.Role = v
}

// GetProfileImageUrl returns the ProfileImageUrl field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetProfileImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileImageUrl
}

// GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field value
// and a boolean to check if the value has been set.
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) GetProfileImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileImageUrl, true
}

// SetProfileImageUrl sets field value
func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) SetProfileImageUrl(v string) {
	o.ProfileImageUrl = v
}

func (o OpenWebuiAppsWebuiModelsAuthsUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenWebuiAppsWebuiModelsAuthsUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["profile_image_url"] = o.ProfileImageUrl
	return toSerialize, nil
}

func (o *OpenWebuiAppsWebuiModelsAuthsUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"email",
		"name",
		"role",
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

	varOpenWebuiAppsWebuiModelsAuthsUserResponse := _OpenWebuiAppsWebuiModelsAuthsUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenWebuiAppsWebuiModelsAuthsUserResponse)

	if err != nil {
		return err
	}

	*o = OpenWebuiAppsWebuiModelsAuthsUserResponse(varOpenWebuiAppsWebuiModelsAuthsUserResponse)

	return err
}

type NullableOpenWebuiAppsWebuiModelsAuthsUserResponse struct {
	value *OpenWebuiAppsWebuiModelsAuthsUserResponse
	isSet bool
}

func (v NullableOpenWebuiAppsWebuiModelsAuthsUserResponse) Get() *OpenWebuiAppsWebuiModelsAuthsUserResponse {
	return v.value
}

func (v *NullableOpenWebuiAppsWebuiModelsAuthsUserResponse) Set(val *OpenWebuiAppsWebuiModelsAuthsUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenWebuiAppsWebuiModelsAuthsUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenWebuiAppsWebuiModelsAuthsUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenWebuiAppsWebuiModelsAuthsUserResponse(val *OpenWebuiAppsWebuiModelsAuthsUserResponse) *NullableOpenWebuiAppsWebuiModelsAuthsUserResponse {
	return &NullableOpenWebuiAppsWebuiModelsAuthsUserResponse{value: val, isSet: true}
}

func (v NullableOpenWebuiAppsWebuiModelsAuthsUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenWebuiAppsWebuiModelsAuthsUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



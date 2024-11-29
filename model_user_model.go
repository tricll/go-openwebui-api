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

// checks if the UserModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModel{}

// UserModel struct for UserModel
type UserModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role *string `json:"role,omitempty"`
	ProfileImageUrl string `json:"profile_image_url"`
	LastActiveAt int32 `json:"last_active_at"`
	UpdatedAt int32 `json:"updated_at"`
	CreatedAt int32 `json:"created_at"`
	ApiKey NullableString `json:"api_key,omitempty"`
	Settings NullableUserSettings `json:"settings,omitempty"`
	Info map[string]interface{} `json:"info,omitempty"`
	OauthSub NullableString `json:"oauth_sub,omitempty"`
}

type _UserModel UserModel

// NewUserModel instantiates a new UserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModel(id string, name string, email string, profileImageUrl string, lastActiveAt int32, updatedAt int32, createdAt int32) *UserModel {
	this := UserModel{}
	this.Id = id
	this.Name = name
	this.Email = email
	var role string = "pending"
	this.Role = &role
	this.ProfileImageUrl = profileImageUrl
	this.LastActiveAt = lastActiveAt
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewUserModelWithDefaults instantiates a new UserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelWithDefaults() *UserModel {
	this := UserModel{}
	var role string = "pending"
	this.Role = &role
	return &this
}

// GetId returns the Id field value
func (o *UserModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *UserModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserModel) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *UserModel) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserModel) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserModel) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModel) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserModel) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserModel) SetRole(v string) {
	o.Role = &v
}

// GetProfileImageUrl returns the ProfileImageUrl field value
func (o *UserModel) GetProfileImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileImageUrl
}

// GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetProfileImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileImageUrl, true
}

// SetProfileImageUrl sets field value
func (o *UserModel) SetProfileImageUrl(v string) {
	o.ProfileImageUrl = v
}

// GetLastActiveAt returns the LastActiveAt field value
func (o *UserModel) GetLastActiveAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastActiveAt
}

// GetLastActiveAtOk returns a tuple with the LastActiveAt field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetLastActiveAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastActiveAt, true
}

// SetLastActiveAt sets field value
func (o *UserModel) SetLastActiveAt(v int32) {
	o.LastActiveAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserModel) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserModel) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserModel) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserModel) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserModel) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserModel) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey.Get()) {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserModel) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *UserModel) HasApiKey() bool {
	if o != nil && o.ApiKey.IsSet() {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given NullableString and assigns it to the ApiKey field.
func (o *UserModel) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}
// SetApiKeyNil sets the value for ApiKey to be an explicit nil
func (o *UserModel) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
func (o *UserModel) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserModel) GetSettings() UserSettings {
	if o == nil || IsNil(o.Settings.Get()) {
		var ret UserSettings
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserModel) GetSettingsOk() (*UserSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *UserModel) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableUserSettings and assigns it to the Settings field.
func (o *UserModel) SetSettings(v UserSettings) {
	o.Settings.Set(&v)
}
// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *UserModel) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *UserModel) UnsetSettings() {
	o.Settings.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserModel) GetInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserModel) GetInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Info) {
		return map[string]interface{}{}, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *UserModel) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]interface{} and assigns it to the Info field.
func (o *UserModel) SetInfo(v map[string]interface{}) {
	o.Info = v
}

// GetOauthSub returns the OauthSub field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserModel) GetOauthSub() string {
	if o == nil || IsNil(o.OauthSub.Get()) {
		var ret string
		return ret
	}
	return *o.OauthSub.Get()
}

// GetOauthSubOk returns a tuple with the OauthSub field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserModel) GetOauthSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OauthSub.Get(), o.OauthSub.IsSet()
}

// HasOauthSub returns a boolean if a field has been set.
func (o *UserModel) HasOauthSub() bool {
	if o != nil && o.OauthSub.IsSet() {
		return true
	}

	return false
}

// SetOauthSub gets a reference to the given NullableString and assigns it to the OauthSub field.
func (o *UserModel) SetOauthSub(v string) {
	o.OauthSub.Set(&v)
}
// SetOauthSubNil sets the value for OauthSub to be an explicit nil
func (o *UserModel) SetOauthSubNil() {
	o.OauthSub.Set(nil)
}

// UnsetOauthSub ensures that no value is present for OauthSub, not even an explicit nil
func (o *UserModel) UnsetOauthSub() {
	o.OauthSub.Unset()
}

func (o UserModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	toSerialize["profile_image_url"] = o.ProfileImageUrl
	toSerialize["last_active_at"] = o.LastActiveAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	if o.ApiKey.IsSet() {
		toSerialize["api_key"] = o.ApiKey.Get()
	}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.OauthSub.IsSet() {
		toSerialize["oauth_sub"] = o.OauthSub.Get()
	}
	return toSerialize, nil
}

func (o *UserModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"email",
		"profile_image_url",
		"last_active_at",
		"updated_at",
		"created_at",
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

	varUserModel := _UserModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserModel)

	if err != nil {
		return err
	}

	*o = UserModel(varUserModel)

	return err
}

type NullableUserModel struct {
	value *UserModel
	isSet bool
}

func (v NullableUserModel) Get() *UserModel {
	return v.value
}

func (v *NullableUserModel) Set(val *UserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModel(val *UserModel) *NullableUserModel {
	return &NullableUserModel{value: val, isSet: true}
}

func (v NullableUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



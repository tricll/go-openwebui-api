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

// checks if the KnowledgeUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KnowledgeUserResponse{}

// KnowledgeUserResponse struct for KnowledgeUserResponse
type KnowledgeUserResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Data map[string]interface{} `json:"data,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	AccessControl map[string]interface{} `json:"access_control,omitempty"`
	CreatedAt int32 `json:"created_at"`
	UpdatedAt int32 `json:"updated_at"`
	User NullableOpenWebuiAppsWebuiModelsUsersUserResponse `json:"user,omitempty"`
	Files []KnowledgeResponseFilesInner `json:"files,omitempty"`
}

type _KnowledgeUserResponse KnowledgeUserResponse

// NewKnowledgeUserResponse instantiates a new KnowledgeUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnowledgeUserResponse(id string, userId string, name string, description string, createdAt int32, updatedAt int32) *KnowledgeUserResponse {
	this := KnowledgeUserResponse{}
	this.Id = id
	this.UserId = userId
	this.Name = name
	this.Description = description
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewKnowledgeUserResponseWithDefaults instantiates a new KnowledgeUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnowledgeUserResponseWithDefaults() *KnowledgeUserResponse {
	this := KnowledgeUserResponse{}
	return &this
}

// GetId returns the Id field value
func (o *KnowledgeUserResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KnowledgeUserResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KnowledgeUserResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *KnowledgeUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *KnowledgeUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *KnowledgeUserResponse) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *KnowledgeUserResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KnowledgeUserResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KnowledgeUserResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *KnowledgeUserResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *KnowledgeUserResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *KnowledgeUserResponse) SetDescription(v string) {
	o.Description = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeUserResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeUserResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KnowledgeUserResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *KnowledgeUserResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeUserResponse) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeUserResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *KnowledgeUserResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *KnowledgeUserResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeUserResponse) GetAccessControl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeUserResponse) GetAccessControlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return map[string]interface{}{}, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *KnowledgeUserResponse) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given map[string]interface{} and assigns it to the AccessControl field.
func (o *KnowledgeUserResponse) SetAccessControl(v map[string]interface{}) {
	o.AccessControl = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *KnowledgeUserResponse) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *KnowledgeUserResponse) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *KnowledgeUserResponse) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *KnowledgeUserResponse) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *KnowledgeUserResponse) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *KnowledgeUserResponse) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeUserResponse) GetUser() OpenWebuiAppsWebuiModelsUsersUserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret OpenWebuiAppsWebuiModelsUsersUserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeUserResponse) GetUserOk() (*OpenWebuiAppsWebuiModelsUsersUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *KnowledgeUserResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableOpenWebuiAppsWebuiModelsUsersUserResponse and assigns it to the User field.
func (o *KnowledgeUserResponse) SetUser(v OpenWebuiAppsWebuiModelsUsersUserResponse) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *KnowledgeUserResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *KnowledgeUserResponse) UnsetUser() {
	o.User.Unset()
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnowledgeUserResponse) GetFiles() []KnowledgeResponseFilesInner {
	if o == nil {
		var ret []KnowledgeResponseFilesInner
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnowledgeUserResponse) GetFilesOk() ([]KnowledgeResponseFilesInner, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *KnowledgeUserResponse) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []KnowledgeResponseFilesInner and assigns it to the Files field.
func (o *KnowledgeUserResponse) SetFiles(v []KnowledgeResponseFilesInner) {
	o.Files = v
}

func (o KnowledgeUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KnowledgeUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.AccessControl != nil {
		toSerialize["access_control"] = o.AccessControl
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

func (o *KnowledgeUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"name",
		"description",
		"created_at",
		"updated_at",
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

	varKnowledgeUserResponse := _KnowledgeUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKnowledgeUserResponse)

	if err != nil {
		return err
	}

	*o = KnowledgeUserResponse(varKnowledgeUserResponse)

	return err
}

type NullableKnowledgeUserResponse struct {
	value *KnowledgeUserResponse
	isSet bool
}

func (v NullableKnowledgeUserResponse) Get() *KnowledgeUserResponse {
	return v.value
}

func (v *NullableKnowledgeUserResponse) Set(val *KnowledgeUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKnowledgeUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKnowledgeUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnowledgeUserResponse(val *KnowledgeUserResponse) *NullableKnowledgeUserResponse {
	return &NullableKnowledgeUserResponse{value: val, isSet: true}
}

func (v NullableKnowledgeUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnowledgeUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



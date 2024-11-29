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

// checks if the PromptUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromptUserResponse{}

// PromptUserResponse struct for PromptUserResponse
type PromptUserResponse struct {
	Command string `json:"command"`
	UserId string `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Timestamp int32 `json:"timestamp"`
	AccessControl map[string]interface{} `json:"access_control,omitempty"`
	User NullableOpenWebuiAppsWebuiModelsUsersUserResponse `json:"user,omitempty"`
}

type _PromptUserResponse PromptUserResponse

// NewPromptUserResponse instantiates a new PromptUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptUserResponse(command string, userId string, title string, content string, timestamp int32) *PromptUserResponse {
	this := PromptUserResponse{}
	this.Command = command
	this.UserId = userId
	this.Title = title
	this.Content = content
	this.Timestamp = timestamp
	return &this
}

// NewPromptUserResponseWithDefaults instantiates a new PromptUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptUserResponseWithDefaults() *PromptUserResponse {
	this := PromptUserResponse{}
	return &this
}

// GetCommand returns the Command field value
func (o *PromptUserResponse) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *PromptUserResponse) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *PromptUserResponse) SetCommand(v string) {
	o.Command = v
}

// GetUserId returns the UserId field value
func (o *PromptUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PromptUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PromptUserResponse) SetUserId(v string) {
	o.UserId = v
}

// GetTitle returns the Title field value
func (o *PromptUserResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PromptUserResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PromptUserResponse) SetTitle(v string) {
	o.Title = v
}

// GetContent returns the Content field value
func (o *PromptUserResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PromptUserResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *PromptUserResponse) SetContent(v string) {
	o.Content = v
}

// GetTimestamp returns the Timestamp field value
func (o *PromptUserResponse) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PromptUserResponse) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PromptUserResponse) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PromptUserResponse) GetAccessControl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PromptUserResponse) GetAccessControlOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return map[string]interface{}{}, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *PromptUserResponse) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given map[string]interface{} and assigns it to the AccessControl field.
func (o *PromptUserResponse) SetAccessControl(v map[string]interface{}) {
	o.AccessControl = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PromptUserResponse) GetUser() OpenWebuiAppsWebuiModelsUsersUserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret OpenWebuiAppsWebuiModelsUsersUserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PromptUserResponse) GetUserOk() (*OpenWebuiAppsWebuiModelsUsersUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PromptUserResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableOpenWebuiAppsWebuiModelsUsersUserResponse and assigns it to the User field.
func (o *PromptUserResponse) SetUser(v OpenWebuiAppsWebuiModelsUsersUserResponse) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *PromptUserResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PromptUserResponse) UnsetUser() {
	o.User.Unset()
}

func (o PromptUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromptUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["command"] = o.Command
	toSerialize["user_id"] = o.UserId
	toSerialize["title"] = o.Title
	toSerialize["content"] = o.Content
	toSerialize["timestamp"] = o.Timestamp
	if o.AccessControl != nil {
		toSerialize["access_control"] = o.AccessControl
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *PromptUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"command",
		"user_id",
		"title",
		"content",
		"timestamp",
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

	varPromptUserResponse := _PromptUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPromptUserResponse)

	if err != nil {
		return err
	}

	*o = PromptUserResponse(varPromptUserResponse)

	return err
}

type NullablePromptUserResponse struct {
	value *PromptUserResponse
	isSet bool
}

func (v NullablePromptUserResponse) Get() *PromptUserResponse {
	return v.value
}

func (v *NullablePromptUserResponse) Set(val *PromptUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptUserResponse(val *PromptUserResponse) *NullablePromptUserResponse {
	return &NullablePromptUserResponse{value: val, isSet: true}
}

func (v NullablePromptUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



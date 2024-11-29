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

// checks if the FunctionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionModel{}

// FunctionModel struct for FunctionModel
type FunctionModel struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Content string `json:"content"`
	Meta FunctionMeta `json:"meta"`
	IsActive *bool `json:"is_active,omitempty"`
	IsGlobal *bool `json:"is_global,omitempty"`
	UpdatedAt int32 `json:"updated_at"`
	CreatedAt int32 `json:"created_at"`
}

type _FunctionModel FunctionModel

// NewFunctionModel instantiates a new FunctionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionModel(id string, userId string, name string, type_ string, content string, meta FunctionMeta, updatedAt int32, createdAt int32) *FunctionModel {
	this := FunctionModel{}
	this.Id = id
	this.UserId = userId
	this.Name = name
	this.Type = type_
	this.Content = content
	this.Meta = meta
	var isActive bool = false
	this.IsActive = &isActive
	var isGlobal bool = false
	this.IsGlobal = &isGlobal
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewFunctionModelWithDefaults instantiates a new FunctionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionModelWithDefaults() *FunctionModel {
	this := FunctionModel{}
	var isActive bool = false
	this.IsActive = &isActive
	var isGlobal bool = false
	this.IsGlobal = &isGlobal
	return &this
}

// GetId returns the Id field value
func (o *FunctionModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FunctionModel) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *FunctionModel) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FunctionModel) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value
func (o *FunctionModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FunctionModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FunctionModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FunctionModel) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *FunctionModel) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *FunctionModel) SetContent(v string) {
	o.Content = v
}

// GetMeta returns the Meta field value
func (o *FunctionModel) GetMeta() FunctionMeta {
	if o == nil {
		var ret FunctionMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetMetaOk() (*FunctionMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FunctionModel) SetMeta(v FunctionMeta) {
	o.Meta = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *FunctionModel) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *FunctionModel) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *FunctionModel) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise.
func (o *FunctionModel) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal) {
		var ret bool
		return ret
	}
	return *o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetIsGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobal) {
		return nil, false
	}
	return o.IsGlobal, true
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *FunctionModel) HasIsGlobal() bool {
	if o != nil && !IsNil(o.IsGlobal) {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given bool and assigns it to the IsGlobal field.
func (o *FunctionModel) SetIsGlobal(v bool) {
	o.IsGlobal = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FunctionModel) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FunctionModel) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FunctionModel) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FunctionModel) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FunctionModel) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

func (o FunctionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	toSerialize["meta"] = o.Meta
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsGlobal) {
		toSerialize["is_global"] = o.IsGlobal
	}
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *FunctionModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"name",
		"type",
		"content",
		"meta",
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

	varFunctionModel := _FunctionModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionModel)

	if err != nil {
		return err
	}

	*o = FunctionModel(varFunctionModel)

	return err
}

type NullableFunctionModel struct {
	value *FunctionModel
	isSet bool
}

func (v NullableFunctionModel) Get() *FunctionModel {
	return v.value
}

func (v *NullableFunctionModel) Set(val *FunctionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionModel(val *FunctionModel) *NullableFunctionModel {
	return &NullableFunctionModel{value: val, isSet: true}
}

func (v NullableFunctionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



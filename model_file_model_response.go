/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openwebuiapi

import (
	"encoding/json"
	"fmt"
)

// checks if the FileModelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileModelResponse{}

// FileModelResponse struct for FileModelResponse
type FileModelResponse struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Hash *Hash `json:"hash,omitempty"`
	Filename string `json:"filename"`
	Data *Data `json:"data,omitempty"`
	Meta FileMeta `json:"meta"`
	CreatedAt int32 `json:"created_at"`
	UpdatedAt int32 `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _FileModelResponse FileModelResponse

// NewFileModelResponse instantiates a new FileModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileModelResponse(id string, userId string, filename string, meta FileMeta, createdAt int32, updatedAt int32) *FileModelResponse {
	this := FileModelResponse{}
	this.Id = id
	this.UserId = userId
	this.Filename = filename
	this.Meta = meta
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewFileModelResponseWithDefaults instantiates a new FileModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileModelResponseWithDefaults() *FileModelResponse {
	this := FileModelResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FileModelResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileModelResponse) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *FileModelResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FileModelResponse) SetUserId(v string) {
	o.UserId = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *FileModelResponse) GetHash() Hash {
	if o == nil || IsNil(o.Hash) {
		var ret Hash
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetHashOk() (*Hash, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *FileModelResponse) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given Hash and assigns it to the Hash field.
func (o *FileModelResponse) SetHash(v Hash) {
	o.Hash = &v
}

// GetFilename returns the Filename field value
func (o *FileModelResponse) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *FileModelResponse) SetFilename(v string) {
	o.Filename = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FileModelResponse) GetData() Data {
	if o == nil || IsNil(o.Data) {
		var ret Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetDataOk() (*Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FileModelResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Data and assigns it to the Data field.
func (o *FileModelResponse) SetData(v Data) {
	o.Data = &v
}

// GetMeta returns the Meta field value
func (o *FileModelResponse) GetMeta() FileMeta {
	if o == nil {
		var ret FileMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetMetaOk() (*FileMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FileModelResponse) SetMeta(v FileMeta) {
	o.Meta = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FileModelResponse) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FileModelResponse) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FileModelResponse) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FileModelResponse) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FileModelResponse) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o FileModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileModelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	toSerialize["filename"] = o.Filename
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["meta"] = o.Meta
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileModelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"filename",
		"meta",
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

	varFileModelResponse := _FileModelResponse{}

	err = json.Unmarshal(data, &varFileModelResponse)

	if err != nil {
		return err
	}

	*o = FileModelResponse(varFileModelResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "hash")
		delete(additionalProperties, "filename")
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileModelResponse struct {
	value *FileModelResponse
	isSet bool
}

func (v NullableFileModelResponse) Get() *FileModelResponse {
	return v.value
}

func (v *NullableFileModelResponse) Set(val *FileModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileModelResponse(val *FileModelResponse) *NullableFileModelResponse {
	return &NullableFileModelResponse{value: val, isSet: true}
}

func (v NullableFileModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



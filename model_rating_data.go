/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RatingData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatingData{}

// RatingData struct for RatingData
type RatingData struct {
	Rating *Rating `json:"rating,omitempty"`
	ModelId *ModelId `json:"model_id,omitempty"`
	SiblingModelIds *SiblingModelIds `json:"sibling_model_ids,omitempty"`
	Reason *Reason `json:"reason,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RatingData RatingData

// NewRatingData instantiates a new RatingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatingData() *RatingData {
	this := RatingData{}
	return &this
}

// NewRatingDataWithDefaults instantiates a new RatingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingDataWithDefaults() *RatingData {
	this := RatingData{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *RatingData) GetRating() Rating {
	if o == nil || IsNil(o.Rating) {
		var ret Rating
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingData) GetRatingOk() (*Rating, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *RatingData) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given Rating and assigns it to the Rating field.
func (o *RatingData) SetRating(v Rating) {
	o.Rating = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *RatingData) GetModelId() ModelId {
	if o == nil || IsNil(o.ModelId) {
		var ret ModelId
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingData) GetModelIdOk() (*ModelId, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *RatingData) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given ModelId and assigns it to the ModelId field.
func (o *RatingData) SetModelId(v ModelId) {
	o.ModelId = &v
}

// GetSiblingModelIds returns the SiblingModelIds field value if set, zero value otherwise.
func (o *RatingData) GetSiblingModelIds() SiblingModelIds {
	if o == nil || IsNil(o.SiblingModelIds) {
		var ret SiblingModelIds
		return ret
	}
	return *o.SiblingModelIds
}

// GetSiblingModelIdsOk returns a tuple with the SiblingModelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingData) GetSiblingModelIdsOk() (*SiblingModelIds, bool) {
	if o == nil || IsNil(o.SiblingModelIds) {
		return nil, false
	}
	return o.SiblingModelIds, true
}

// HasSiblingModelIds returns a boolean if a field has been set.
func (o *RatingData) HasSiblingModelIds() bool {
	if o != nil && !IsNil(o.SiblingModelIds) {
		return true
	}

	return false
}

// SetSiblingModelIds gets a reference to the given SiblingModelIds and assigns it to the SiblingModelIds field.
func (o *RatingData) SetSiblingModelIds(v SiblingModelIds) {
	o.SiblingModelIds = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RatingData) GetReason() Reason {
	if o == nil || IsNil(o.Reason) {
		var ret Reason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingData) GetReasonOk() (*Reason, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RatingData) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given Reason and assigns it to the Reason field.
func (o *RatingData) SetReason(v Reason) {
	o.Reason = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RatingData) GetComment() Comment {
	if o == nil || IsNil(o.Comment) {
		var ret Comment
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingData) GetCommentOk() (*Comment, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RatingData) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given Comment and assigns it to the Comment field.
func (o *RatingData) SetComment(v Comment) {
	o.Comment = &v
}

func (o RatingData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.ModelId) {
		toSerialize["model_id"] = o.ModelId
	}
	if !IsNil(o.SiblingModelIds) {
		toSerialize["sibling_model_ids"] = o.SiblingModelIds
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RatingData) UnmarshalJSON(data []byte) (err error) {
	varRatingData := _RatingData{}

	err = json.Unmarshal(data, &varRatingData)

	if err != nil {
		return err
	}

	*o = RatingData(varRatingData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rating")
		delete(additionalProperties, "model_id")
		delete(additionalProperties, "sibling_model_ids")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRatingData struct {
	value *RatingData
	isSet bool
}

func (v NullableRatingData) Get() *RatingData {
	return v.value
}

func (v *NullableRatingData) Set(val *RatingData) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingData) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingData(val *RatingData) *NullableRatingData {
	return &NullableRatingData{value: val, isSet: true}
}

func (v NullableRatingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



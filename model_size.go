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


// Size struct for Size
type Size struct {
	Int32 *int32
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Size) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Size)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Size) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableSize struct {
	value *Size
	isSet bool
}

func (v NullableSize) Get() *Size {
	return v.value
}

func (v *NullableSize) Set(val *Size) {
	v.value = val
	v.isSet = true
}

func (v NullableSize) IsSet() bool {
	return v.isSet
}

func (v *NullableSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSize(val *Size) *NullableSize {
	return &NullableSize{value: val, isSet: true}
}

func (v NullableSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



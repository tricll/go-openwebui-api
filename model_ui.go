/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// Ui struct for Ui
type Ui struct {
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Ui) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny);
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.MapmapOfStringAny, return on the first match
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Ui)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Ui) MarshalJSON() ([]byte, error) {
	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableUi struct {
	value *Ui
	isSet bool
}

func (v NullableUi) Get() *Ui {
	return v.value
}

func (v *NullableUi) Set(val *Ui) {
	v.value = val
	v.isSet = true
}

func (v NullableUi) IsSet() bool {
	return v.isSet
}

func (v *NullableUi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUi(val *Ui) *NullableUi {
	return &NullableUi{value: val, isSet: true}
}

func (v NullableUi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



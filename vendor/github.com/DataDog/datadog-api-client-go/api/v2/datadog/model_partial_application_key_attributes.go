/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// PartialApplicationKeyAttributes Attributes of a partial application key.
type PartialApplicationKeyAttributes struct {
	// Creation date of the application key.
	CreatedAt *string `json:"created_at,omitempty"`
	// The last four characters of the application key.
	Last4 *string `json:"last4,omitempty"`
	// Name of the application key.
	Name *string `json:"name,omitempty"`
}

// NewPartialApplicationKeyAttributes instantiates a new PartialApplicationKeyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialApplicationKeyAttributes() *PartialApplicationKeyAttributes {
	this := PartialApplicationKeyAttributes{}
	return &this
}

// NewPartialApplicationKeyAttributesWithDefaults instantiates a new PartialApplicationKeyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialApplicationKeyAttributesWithDefaults() *PartialApplicationKeyAttributes {
	this := PartialApplicationKeyAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PartialApplicationKeyAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialApplicationKeyAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PartialApplicationKeyAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PartialApplicationKeyAttributes) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialApplicationKeyAttributes) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PartialApplicationKeyAttributes) SetLast4(v string) {
	o.Last4 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialApplicationKeyAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialApplicationKeyAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialApplicationKeyAttributes) SetName(v string) {
	o.Name = &v
}

func (o PartialApplicationKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePartialApplicationKeyAttributes struct {
	value *PartialApplicationKeyAttributes
	isSet bool
}

func (v NullablePartialApplicationKeyAttributes) Get() *PartialApplicationKeyAttributes {
	return v.value
}

func (v *NullablePartialApplicationKeyAttributes) Set(val *PartialApplicationKeyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialApplicationKeyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialApplicationKeyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialApplicationKeyAttributes(val *PartialApplicationKeyAttributes) *NullablePartialApplicationKeyAttributes {
	return &NullablePartialApplicationKeyAttributes{value: val, isSet: true}
}

func (v NullablePartialApplicationKeyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialApplicationKeyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
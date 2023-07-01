/*
Robinhood API

Robinhood API Documentation

API version: 3.0.1
Contact: austin.millan@protonmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PaginatedOptionChain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedOptionChain{}

// PaginatedOptionChain struct for PaginatedOptionChain
type PaginatedOptionChain struct {
	Count *string `json:"count,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results []OptionChain `json:"results,omitempty"`
}

// NewPaginatedOptionChain instantiates a new PaginatedOptionChain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOptionChain() *PaginatedOptionChain {
	this := PaginatedOptionChain{}
	return &this
}

// NewPaginatedOptionChainWithDefaults instantiates a new PaginatedOptionChain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOptionChainWithDefaults() *PaginatedOptionChain {
	this := PaginatedOptionChain{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedOptionChain) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOptionChain) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedOptionChain) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *PaginatedOptionChain) SetCount(v string) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedOptionChain) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOptionChain) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedOptionChain) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedOptionChain) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedOptionChain) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOptionChain) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedOptionChain) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedOptionChain) SetPrevious(v string) {
	o.Previous = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedOptionChain) GetResults() []OptionChain {
	if o == nil || IsNil(o.Results) {
		var ret []OptionChain
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedOptionChain) GetResultsOk() ([]OptionChain, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedOptionChain) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OptionChain and assigns it to the Results field.
func (o *PaginatedOptionChain) SetResults(v []OptionChain) {
	o.Results = v
}

func (o PaginatedOptionChain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedOptionChain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedOptionChain struct {
	value *PaginatedOptionChain
	isSet bool
}

func (v NullablePaginatedOptionChain) Get() *PaginatedOptionChain {
	return v.value
}

func (v *NullablePaginatedOptionChain) Set(val *PaginatedOptionChain) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOptionChain) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOptionChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOptionChain(val *PaginatedOptionChain) *NullablePaginatedOptionChain {
	return &NullablePaginatedOptionChain{value: val, isSet: true}
}

func (v NullablePaginatedOptionChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOptionChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


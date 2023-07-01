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

// checks if the PaginatedInstrumentSplit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedInstrumentSplit{}

// PaginatedInstrumentSplit struct for PaginatedInstrumentSplit
type PaginatedInstrumentSplit struct {
	Count *string `json:"count,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results []InstrumentSplit `json:"results,omitempty"`
}

// NewPaginatedInstrumentSplit instantiates a new PaginatedInstrumentSplit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedInstrumentSplit() *PaginatedInstrumentSplit {
	this := PaginatedInstrumentSplit{}
	return &this
}

// NewPaginatedInstrumentSplitWithDefaults instantiates a new PaginatedInstrumentSplit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedInstrumentSplitWithDefaults() *PaginatedInstrumentSplit {
	this := PaginatedInstrumentSplit{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedInstrumentSplit) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentSplit) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedInstrumentSplit) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *PaginatedInstrumentSplit) SetCount(v string) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedInstrumentSplit) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentSplit) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedInstrumentSplit) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedInstrumentSplit) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedInstrumentSplit) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentSplit) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedInstrumentSplit) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedInstrumentSplit) SetPrevious(v string) {
	o.Previous = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedInstrumentSplit) GetResults() []InstrumentSplit {
	if o == nil || IsNil(o.Results) {
		var ret []InstrumentSplit
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedInstrumentSplit) GetResultsOk() ([]InstrumentSplit, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedInstrumentSplit) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InstrumentSplit and assigns it to the Results field.
func (o *PaginatedInstrumentSplit) SetResults(v []InstrumentSplit) {
	o.Results = v
}

func (o PaginatedInstrumentSplit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedInstrumentSplit) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedInstrumentSplit struct {
	value *PaginatedInstrumentSplit
	isSet bool
}

func (v NullablePaginatedInstrumentSplit) Get() *PaginatedInstrumentSplit {
	return v.value
}

func (v *NullablePaginatedInstrumentSplit) Set(val *PaginatedInstrumentSplit) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedInstrumentSplit) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedInstrumentSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedInstrumentSplit(val *PaginatedInstrumentSplit) *NullablePaginatedInstrumentSplit {
	return &NullablePaginatedInstrumentSplit{value: val, isSet: true}
}

func (v NullablePaginatedInstrumentSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedInstrumentSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


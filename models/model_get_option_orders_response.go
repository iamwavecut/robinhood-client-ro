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

// checks if the GetOptionOrdersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOptionOrdersResponse{}

// GetOptionOrdersResponse struct for GetOptionOrdersResponse
type GetOptionOrdersResponse struct {
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results []OptionOrder `json:"results,omitempty"`
}

// NewGetOptionOrdersResponse instantiates a new GetOptionOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOptionOrdersResponse() *GetOptionOrdersResponse {
	this := GetOptionOrdersResponse{}
	return &this
}

// NewGetOptionOrdersResponseWithDefaults instantiates a new GetOptionOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOptionOrdersResponseWithDefaults() *GetOptionOrdersResponse {
	this := GetOptionOrdersResponse{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *GetOptionOrdersResponse) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionOrdersResponse) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *GetOptionOrdersResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *GetOptionOrdersResponse) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *GetOptionOrdersResponse) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionOrdersResponse) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *GetOptionOrdersResponse) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *GetOptionOrdersResponse) SetPrevious(v string) {
	o.Previous = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetOptionOrdersResponse) GetResults() []OptionOrder {
	if o == nil || IsNil(o.Results) {
		var ret []OptionOrder
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptionOrdersResponse) GetResultsOk() ([]OptionOrder, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetOptionOrdersResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OptionOrder and assigns it to the Results field.
func (o *GetOptionOrdersResponse) SetResults(v []OptionOrder) {
	o.Results = v
}

func (o GetOptionOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOptionOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableGetOptionOrdersResponse struct {
	value *GetOptionOrdersResponse
	isSet bool
}

func (v NullableGetOptionOrdersResponse) Get() *GetOptionOrdersResponse {
	return v.value
}

func (v *NullableGetOptionOrdersResponse) Set(val *GetOptionOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOptionOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOptionOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOptionOrdersResponse(val *GetOptionOrdersResponse) *NullableGetOptionOrdersResponse {
	return &NullableGetOptionOrdersResponse{value: val, isSet: true}
}

func (v NullableGetOptionOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOptionOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


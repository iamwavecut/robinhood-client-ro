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

// checks if the UnderlyingInstrument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnderlyingInstrument{}

// UnderlyingInstrument struct for UnderlyingInstrument
type UnderlyingInstrument struct {
	Id *string `json:"id,omitempty"`
	Instrument *string `json:"instrument,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
}

// NewUnderlyingInstrument instantiates a new UnderlyingInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnderlyingInstrument() *UnderlyingInstrument {
	this := UnderlyingInstrument{}
	return &this
}

// NewUnderlyingInstrumentWithDefaults instantiates a new UnderlyingInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnderlyingInstrumentWithDefaults() *UnderlyingInstrument {
	this := UnderlyingInstrument{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnderlyingInstrument) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderlyingInstrument) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnderlyingInstrument) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UnderlyingInstrument) SetId(v string) {
	o.Id = &v
}

// GetInstrument returns the Instrument field value if set, zero value otherwise.
func (o *UnderlyingInstrument) GetInstrument() string {
	if o == nil || IsNil(o.Instrument) {
		var ret string
		return ret
	}
	return *o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderlyingInstrument) GetInstrumentOk() (*string, bool) {
	if o == nil || IsNil(o.Instrument) {
		return nil, false
	}
	return o.Instrument, true
}

// HasInstrument returns a boolean if a field has been set.
func (o *UnderlyingInstrument) HasInstrument() bool {
	if o != nil && !IsNil(o.Instrument) {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given string and assigns it to the Instrument field.
func (o *UnderlyingInstrument) SetInstrument(v string) {
	o.Instrument = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnderlyingInstrument) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnderlyingInstrument) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnderlyingInstrument) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *UnderlyingInstrument) SetQuantity(v float32) {
	o.Quantity = &v
}

func (o UnderlyingInstrument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnderlyingInstrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Instrument) {
		toSerialize["instrument"] = o.Instrument
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableUnderlyingInstrument struct {
	value *UnderlyingInstrument
	isSet bool
}

func (v NullableUnderlyingInstrument) Get() *UnderlyingInstrument {
	return v.value
}

func (v *NullableUnderlyingInstrument) Set(val *UnderlyingInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableUnderlyingInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableUnderlyingInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnderlyingInstrument(val *UnderlyingInstrument) *NullableUnderlyingInstrument {
	return &NullableUnderlyingInstrument{value: val, isSet: true}
}

func (v NullableUnderlyingInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnderlyingInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



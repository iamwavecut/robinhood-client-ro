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

// checks if the OptionOrderLeg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionOrderLeg{}

// OptionOrderLeg struct for OptionOrderLeg
type OptionOrderLeg struct {
	Executions []OptionOrderLegExecutionsInner `json:"executions,omitempty"`
	Id *string `json:"id,omitempty"`
	Option *string `json:"option,omitempty"`
	PositionEffect *PositionEffect `json:"position_effect,omitempty"`
	RatioQuantity *float32 `json:"ratio_quantity,omitempty"`
	Side *Side `json:"side,omitempty"`
}

// NewOptionOrderLeg instantiates a new OptionOrderLeg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionOrderLeg() *OptionOrderLeg {
	this := OptionOrderLeg{}
	return &this
}

// NewOptionOrderLegWithDefaults instantiates a new OptionOrderLeg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionOrderLegWithDefaults() *OptionOrderLeg {
	this := OptionOrderLeg{}
	return &this
}

// GetExecutions returns the Executions field value if set, zero value otherwise.
func (o *OptionOrderLeg) GetExecutions() []OptionOrderLegExecutionsInner {
	if o == nil || IsNil(o.Executions) {
		var ret []OptionOrderLegExecutionsInner
		return ret
	}
	return o.Executions
}

// GetExecutionsOk returns a tuple with the Executions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrderLeg) GetExecutionsOk() ([]OptionOrderLegExecutionsInner, bool) {
	if o == nil || IsNil(o.Executions) {
		return nil, false
	}
	return o.Executions, true
}

// HasExecutions returns a boolean if a field has been set.
func (o *OptionOrderLeg) HasExecutions() bool {
	if o != nil && !IsNil(o.Executions) {
		return true
	}

	return false
}

// SetExecutions gets a reference to the given []OptionOrderLegExecutionsInner and assigns it to the Executions field.
func (o *OptionOrderLeg) SetExecutions(v []OptionOrderLegExecutionsInner) {
	o.Executions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionOrderLeg) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrderLeg) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionOrderLeg) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptionOrderLeg) SetId(v string) {
	o.Id = &v
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *OptionOrderLeg) GetOption() string {
	if o == nil || IsNil(o.Option) {
		var ret string
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrderLeg) GetOptionOk() (*string, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *OptionOrderLeg) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given string and assigns it to the Option field.
func (o *OptionOrderLeg) SetOption(v string) {
	o.Option = &v
}

// GetPositionEffect returns the PositionEffect field value if set, zero value otherwise.
func (o *OptionOrderLeg) GetPositionEffect() PositionEffect {
	if o == nil || IsNil(o.PositionEffect) {
		var ret PositionEffect
		return ret
	}
	return *o.PositionEffect
}

// GetPositionEffectOk returns a tuple with the PositionEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrderLeg) GetPositionEffectOk() (*PositionEffect, bool) {
	if o == nil || IsNil(o.PositionEffect) {
		return nil, false
	}
	return o.PositionEffect, true
}

// HasPositionEffect returns a boolean if a field has been set.
func (o *OptionOrderLeg) HasPositionEffect() bool {
	if o != nil && !IsNil(o.PositionEffect) {
		return true
	}

	return false
}

// SetPositionEffect gets a reference to the given PositionEffect and assigns it to the PositionEffect field.
func (o *OptionOrderLeg) SetPositionEffect(v PositionEffect) {
	o.PositionEffect = &v
}

// GetRatioQuantity returns the RatioQuantity field value if set, zero value otherwise.
func (o *OptionOrderLeg) GetRatioQuantity() float32 {
	if o == nil || IsNil(o.RatioQuantity) {
		var ret float32
		return ret
	}
	return *o.RatioQuantity
}

// GetRatioQuantityOk returns a tuple with the RatioQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrderLeg) GetRatioQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.RatioQuantity) {
		return nil, false
	}
	return o.RatioQuantity, true
}

// HasRatioQuantity returns a boolean if a field has been set.
func (o *OptionOrderLeg) HasRatioQuantity() bool {
	if o != nil && !IsNil(o.RatioQuantity) {
		return true
	}

	return false
}

// SetRatioQuantity gets a reference to the given float32 and assigns it to the RatioQuantity field.
func (o *OptionOrderLeg) SetRatioQuantity(v float32) {
	o.RatioQuantity = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OptionOrderLeg) GetSide() Side {
	if o == nil || IsNil(o.Side) {
		var ret Side
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrderLeg) GetSideOk() (*Side, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OptionOrderLeg) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given Side and assigns it to the Side field.
func (o *OptionOrderLeg) SetSide(v Side) {
	o.Side = &v
}

func (o OptionOrderLeg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionOrderLeg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Executions) {
		toSerialize["executions"] = o.Executions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !IsNil(o.PositionEffect) {
		toSerialize["position_effect"] = o.PositionEffect
	}
	if !IsNil(o.RatioQuantity) {
		toSerialize["ratio_quantity"] = o.RatioQuantity
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	return toSerialize, nil
}

type NullableOptionOrderLeg struct {
	value *OptionOrderLeg
	isSet bool
}

func (v NullableOptionOrderLeg) Get() *OptionOrderLeg {
	return v.value
}

func (v *NullableOptionOrderLeg) Set(val *OptionOrderLeg) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionOrderLeg) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionOrderLeg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionOrderLeg(val *OptionOrderLeg) *NullableOptionOrderLeg {
	return &NullableOptionOrderLeg{value: val, isSet: true}
}

func (v NullableOptionOrderLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionOrderLeg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



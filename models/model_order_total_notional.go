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

// checks if the OrderTotalNotional type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTotalNotional{}

// OrderTotalNotional struct for OrderTotalNotional
type OrderTotalNotional struct {
	Amount *string `json:"amount,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencyId *string `json:"currency_id,omitempty"`
}

// NewOrderTotalNotional instantiates a new OrderTotalNotional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTotalNotional() *OrderTotalNotional {
	this := OrderTotalNotional{}
	return &this
}

// NewOrderTotalNotionalWithDefaults instantiates a new OrderTotalNotional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTotalNotionalWithDefaults() *OrderTotalNotional {
	this := OrderTotalNotional{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderTotalNotional) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotalNotional) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderTotalNotional) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *OrderTotalNotional) SetAmount(v string) {
	o.Amount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *OrderTotalNotional) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotalNotional) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *OrderTotalNotional) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *OrderTotalNotional) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *OrderTotalNotional) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotalNotional) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *OrderTotalNotional) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *OrderTotalNotional) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

func (o OrderTotalNotional) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTotalNotional) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	return toSerialize, nil
}

type NullableOrderTotalNotional struct {
	value *OrderTotalNotional
	isSet bool
}

func (v NullableOrderTotalNotional) Get() *OrderTotalNotional {
	return v.value
}

func (v *NullableOrderTotalNotional) Set(val *OrderTotalNotional) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTotalNotional) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTotalNotional) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTotalNotional(val *OrderTotalNotional) *NullableOrderTotalNotional {
	return &NullableOrderTotalNotional{value: val, isSet: true}
}

func (v NullableOrderTotalNotional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTotalNotional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



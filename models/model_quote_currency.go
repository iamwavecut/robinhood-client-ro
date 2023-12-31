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

// checks if the QuoteCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteCurrency{}

// QuoteCurrency struct for QuoteCurrency
type QuoteCurrency struct {
	Code *string `json:"code,omitempty"`
	Id *string `json:"id,omitempty"`
	Increment *string `json:"increment,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewQuoteCurrency instantiates a new QuoteCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteCurrency() *QuoteCurrency {
	this := QuoteCurrency{}
	return &this
}

// NewQuoteCurrencyWithDefaults instantiates a new QuoteCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteCurrencyWithDefaults() *QuoteCurrency {
	this := QuoteCurrency{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *QuoteCurrency) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteCurrency) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *QuoteCurrency) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *QuoteCurrency) SetCode(v string) {
	o.Code = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuoteCurrency) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteCurrency) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuoteCurrency) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QuoteCurrency) SetId(v string) {
	o.Id = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *QuoteCurrency) GetIncrement() string {
	if o == nil || IsNil(o.Increment) {
		var ret string
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteCurrency) GetIncrementOk() (*string, bool) {
	if o == nil || IsNil(o.Increment) {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *QuoteCurrency) HasIncrement() bool {
	if o != nil && !IsNil(o.Increment) {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given string and assigns it to the Increment field.
func (o *QuoteCurrency) SetIncrement(v string) {
	o.Increment = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QuoteCurrency) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteCurrency) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QuoteCurrency) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QuoteCurrency) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QuoteCurrency) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteCurrency) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QuoteCurrency) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QuoteCurrency) SetType(v string) {
	o.Type = &v
}

func (o QuoteCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Increment) {
		toSerialize["increment"] = o.Increment
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableQuoteCurrency struct {
	value *QuoteCurrency
	isSet bool
}

func (v NullableQuoteCurrency) Get() *QuoteCurrency {
	return v.value
}

func (v *NullableQuoteCurrency) Set(val *QuoteCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteCurrency(val *QuoteCurrency) *NullableQuoteCurrency {
	return &NullableQuoteCurrency{value: val, isSet: true}
}

func (v NullableQuoteCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



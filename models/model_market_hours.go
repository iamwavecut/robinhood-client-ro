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
	"time"
)

// checks if the MarketHours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketHours{}

// MarketHours struct for MarketHours
type MarketHours struct {
	ClosesAt *time.Time `json:"closes_at,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	ExtendedClosesAt *time.Time `json:"extended_closes_at,omitempty"`
	ExtendedOpensAt *time.Time `json:"extended_opens_at,omitempty"`
	IsOpen *bool `json:"is_open,omitempty"`
	NextOpenHours *string `json:"next_open_hours,omitempty"`
	OpensAt *time.Time `json:"opens_at,omitempty"`
	PreviousOpenHours *string `json:"previous_open_hours,omitempty"`
}

// NewMarketHours instantiates a new MarketHours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketHours() *MarketHours {
	this := MarketHours{}
	return &this
}

// NewMarketHoursWithDefaults instantiates a new MarketHours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketHoursWithDefaults() *MarketHours {
	this := MarketHours{}
	return &this
}

// GetClosesAt returns the ClosesAt field value if set, zero value otherwise.
func (o *MarketHours) GetClosesAt() time.Time {
	if o == nil || IsNil(o.ClosesAt) {
		var ret time.Time
		return ret
	}
	return *o.ClosesAt
}

// GetClosesAtOk returns a tuple with the ClosesAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetClosesAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosesAt) {
		return nil, false
	}
	return o.ClosesAt, true
}

// HasClosesAt returns a boolean if a field has been set.
func (o *MarketHours) HasClosesAt() bool {
	if o != nil && !IsNil(o.ClosesAt) {
		return true
	}

	return false
}

// SetClosesAt gets a reference to the given time.Time and assigns it to the ClosesAt field.
func (o *MarketHours) SetClosesAt(v time.Time) {
	o.ClosesAt = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MarketHours) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MarketHours) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *MarketHours) SetDate(v time.Time) {
	o.Date = &v
}

// GetExtendedClosesAt returns the ExtendedClosesAt field value if set, zero value otherwise.
func (o *MarketHours) GetExtendedClosesAt() time.Time {
	if o == nil || IsNil(o.ExtendedClosesAt) {
		var ret time.Time
		return ret
	}
	return *o.ExtendedClosesAt
}

// GetExtendedClosesAtOk returns a tuple with the ExtendedClosesAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetExtendedClosesAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExtendedClosesAt) {
		return nil, false
	}
	return o.ExtendedClosesAt, true
}

// HasExtendedClosesAt returns a boolean if a field has been set.
func (o *MarketHours) HasExtendedClosesAt() bool {
	if o != nil && !IsNil(o.ExtendedClosesAt) {
		return true
	}

	return false
}

// SetExtendedClosesAt gets a reference to the given time.Time and assigns it to the ExtendedClosesAt field.
func (o *MarketHours) SetExtendedClosesAt(v time.Time) {
	o.ExtendedClosesAt = &v
}

// GetExtendedOpensAt returns the ExtendedOpensAt field value if set, zero value otherwise.
func (o *MarketHours) GetExtendedOpensAt() time.Time {
	if o == nil || IsNil(o.ExtendedOpensAt) {
		var ret time.Time
		return ret
	}
	return *o.ExtendedOpensAt
}

// GetExtendedOpensAtOk returns a tuple with the ExtendedOpensAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetExtendedOpensAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExtendedOpensAt) {
		return nil, false
	}
	return o.ExtendedOpensAt, true
}

// HasExtendedOpensAt returns a boolean if a field has been set.
func (o *MarketHours) HasExtendedOpensAt() bool {
	if o != nil && !IsNil(o.ExtendedOpensAt) {
		return true
	}

	return false
}

// SetExtendedOpensAt gets a reference to the given time.Time and assigns it to the ExtendedOpensAt field.
func (o *MarketHours) SetExtendedOpensAt(v time.Time) {
	o.ExtendedOpensAt = &v
}

// GetIsOpen returns the IsOpen field value if set, zero value otherwise.
func (o *MarketHours) GetIsOpen() bool {
	if o == nil || IsNil(o.IsOpen) {
		var ret bool
		return ret
	}
	return *o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetIsOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpen) {
		return nil, false
	}
	return o.IsOpen, true
}

// HasIsOpen returns a boolean if a field has been set.
func (o *MarketHours) HasIsOpen() bool {
	if o != nil && !IsNil(o.IsOpen) {
		return true
	}

	return false
}

// SetIsOpen gets a reference to the given bool and assigns it to the IsOpen field.
func (o *MarketHours) SetIsOpen(v bool) {
	o.IsOpen = &v
}

// GetNextOpenHours returns the NextOpenHours field value if set, zero value otherwise.
func (o *MarketHours) GetNextOpenHours() string {
	if o == nil || IsNil(o.NextOpenHours) {
		var ret string
		return ret
	}
	return *o.NextOpenHours
}

// GetNextOpenHoursOk returns a tuple with the NextOpenHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetNextOpenHoursOk() (*string, bool) {
	if o == nil || IsNil(o.NextOpenHours) {
		return nil, false
	}
	return o.NextOpenHours, true
}

// HasNextOpenHours returns a boolean if a field has been set.
func (o *MarketHours) HasNextOpenHours() bool {
	if o != nil && !IsNil(o.NextOpenHours) {
		return true
	}

	return false
}

// SetNextOpenHours gets a reference to the given string and assigns it to the NextOpenHours field.
func (o *MarketHours) SetNextOpenHours(v string) {
	o.NextOpenHours = &v
}

// GetOpensAt returns the OpensAt field value if set, zero value otherwise.
func (o *MarketHours) GetOpensAt() time.Time {
	if o == nil || IsNil(o.OpensAt) {
		var ret time.Time
		return ret
	}
	return *o.OpensAt
}

// GetOpensAtOk returns a tuple with the OpensAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetOpensAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpensAt) {
		return nil, false
	}
	return o.OpensAt, true
}

// HasOpensAt returns a boolean if a field has been set.
func (o *MarketHours) HasOpensAt() bool {
	if o != nil && !IsNil(o.OpensAt) {
		return true
	}

	return false
}

// SetOpensAt gets a reference to the given time.Time and assigns it to the OpensAt field.
func (o *MarketHours) SetOpensAt(v time.Time) {
	o.OpensAt = &v
}

// GetPreviousOpenHours returns the PreviousOpenHours field value if set, zero value otherwise.
func (o *MarketHours) GetPreviousOpenHours() string {
	if o == nil || IsNil(o.PreviousOpenHours) {
		var ret string
		return ret
	}
	return *o.PreviousOpenHours
}

// GetPreviousOpenHoursOk returns a tuple with the PreviousOpenHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketHours) GetPreviousOpenHoursOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousOpenHours) {
		return nil, false
	}
	return o.PreviousOpenHours, true
}

// HasPreviousOpenHours returns a boolean if a field has been set.
func (o *MarketHours) HasPreviousOpenHours() bool {
	if o != nil && !IsNil(o.PreviousOpenHours) {
		return true
	}

	return false
}

// SetPreviousOpenHours gets a reference to the given string and assigns it to the PreviousOpenHours field.
func (o *MarketHours) SetPreviousOpenHours(v string) {
	o.PreviousOpenHours = &v
}

func (o MarketHours) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketHours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClosesAt) {
		toSerialize["closes_at"] = o.ClosesAt
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.ExtendedClosesAt) {
		toSerialize["extended_closes_at"] = o.ExtendedClosesAt
	}
	if !IsNil(o.ExtendedOpensAt) {
		toSerialize["extended_opens_at"] = o.ExtendedOpensAt
	}
	if !IsNil(o.IsOpen) {
		toSerialize["is_open"] = o.IsOpen
	}
	if !IsNil(o.NextOpenHours) {
		toSerialize["next_open_hours"] = o.NextOpenHours
	}
	if !IsNil(o.OpensAt) {
		toSerialize["opens_at"] = o.OpensAt
	}
	if !IsNil(o.PreviousOpenHours) {
		toSerialize["previous_open_hours"] = o.PreviousOpenHours
	}
	return toSerialize, nil
}

type NullableMarketHours struct {
	value *MarketHours
	isSet bool
}

func (v NullableMarketHours) Get() *MarketHours {
	return v.value
}

func (v *NullableMarketHours) Set(val *MarketHours) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketHours) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketHours(val *MarketHours) *NullableMarketHours {
	return &NullableMarketHours{value: val, isSet: true}
}

func (v NullableMarketHours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



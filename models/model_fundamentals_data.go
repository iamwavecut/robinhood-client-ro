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

// checks if the FundamentalsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FundamentalsData{}

// FundamentalsData struct for FundamentalsData
type FundamentalsData struct {
	AverageVolume *string `json:"average_volume,omitempty"`
	Description *string `json:"description,omitempty"`
	DividendYield *string `json:"dividend_yield,omitempty"`
	High *string `json:"high,omitempty"`
	High52Weeks *string `json:"high_52_weeks,omitempty"`
	Instrument *string `json:"instrument,omitempty"`
	Low *string `json:"low,omitempty"`
	Low52Weeks *string `json:"low_52_weeks,omitempty"`
	MarketCap *string `json:"market_cap,omitempty"`
	Open *string `json:"open,omitempty"`
	PeRatio *string `json:"pe_ratio,omitempty"`
	Volume *string `json:"volume,omitempty"`
}

// NewFundamentalsData instantiates a new FundamentalsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsData() *FundamentalsData {
	this := FundamentalsData{}
	return &this
}

// NewFundamentalsDataWithDefaults instantiates a new FundamentalsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsDataWithDefaults() *FundamentalsData {
	this := FundamentalsData{}
	return &this
}

// GetAverageVolume returns the AverageVolume field value if set, zero value otherwise.
func (o *FundamentalsData) GetAverageVolume() string {
	if o == nil || IsNil(o.AverageVolume) {
		var ret string
		return ret
	}
	return *o.AverageVolume
}

// GetAverageVolumeOk returns a tuple with the AverageVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetAverageVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.AverageVolume) {
		return nil, false
	}
	return o.AverageVolume, true
}

// HasAverageVolume returns a boolean if a field has been set.
func (o *FundamentalsData) HasAverageVolume() bool {
	if o != nil && !IsNil(o.AverageVolume) {
		return true
	}

	return false
}

// SetAverageVolume gets a reference to the given string and assigns it to the AverageVolume field.
func (o *FundamentalsData) SetAverageVolume(v string) {
	o.AverageVolume = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FundamentalsData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FundamentalsData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FundamentalsData) SetDescription(v string) {
	o.Description = &v
}

// GetDividendYield returns the DividendYield field value if set, zero value otherwise.
func (o *FundamentalsData) GetDividendYield() string {
	if o == nil || IsNil(o.DividendYield) {
		var ret string
		return ret
	}
	return *o.DividendYield
}

// GetDividendYieldOk returns a tuple with the DividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetDividendYieldOk() (*string, bool) {
	if o == nil || IsNil(o.DividendYield) {
		return nil, false
	}
	return o.DividendYield, true
}

// HasDividendYield returns a boolean if a field has been set.
func (o *FundamentalsData) HasDividendYield() bool {
	if o != nil && !IsNil(o.DividendYield) {
		return true
	}

	return false
}

// SetDividendYield gets a reference to the given string and assigns it to the DividendYield field.
func (o *FundamentalsData) SetDividendYield(v string) {
	o.DividendYield = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *FundamentalsData) GetHigh() string {
	if o == nil || IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetHighOk() (*string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *FundamentalsData) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *FundamentalsData) SetHigh(v string) {
	o.High = &v
}

// GetHigh52Weeks returns the High52Weeks field value if set, zero value otherwise.
func (o *FundamentalsData) GetHigh52Weeks() string {
	if o == nil || IsNil(o.High52Weeks) {
		var ret string
		return ret
	}
	return *o.High52Weeks
}

// GetHigh52WeeksOk returns a tuple with the High52Weeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetHigh52WeeksOk() (*string, bool) {
	if o == nil || IsNil(o.High52Weeks) {
		return nil, false
	}
	return o.High52Weeks, true
}

// HasHigh52Weeks returns a boolean if a field has been set.
func (o *FundamentalsData) HasHigh52Weeks() bool {
	if o != nil && !IsNil(o.High52Weeks) {
		return true
	}

	return false
}

// SetHigh52Weeks gets a reference to the given string and assigns it to the High52Weeks field.
func (o *FundamentalsData) SetHigh52Weeks(v string) {
	o.High52Weeks = &v
}

// GetInstrument returns the Instrument field value if set, zero value otherwise.
func (o *FundamentalsData) GetInstrument() string {
	if o == nil || IsNil(o.Instrument) {
		var ret string
		return ret
	}
	return *o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetInstrumentOk() (*string, bool) {
	if o == nil || IsNil(o.Instrument) {
		return nil, false
	}
	return o.Instrument, true
}

// HasInstrument returns a boolean if a field has been set.
func (o *FundamentalsData) HasInstrument() bool {
	if o != nil && !IsNil(o.Instrument) {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given string and assigns it to the Instrument field.
func (o *FundamentalsData) SetInstrument(v string) {
	o.Instrument = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *FundamentalsData) GetLow() string {
	if o == nil || IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetLowOk() (*string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *FundamentalsData) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *FundamentalsData) SetLow(v string) {
	o.Low = &v
}

// GetLow52Weeks returns the Low52Weeks field value if set, zero value otherwise.
func (o *FundamentalsData) GetLow52Weeks() string {
	if o == nil || IsNil(o.Low52Weeks) {
		var ret string
		return ret
	}
	return *o.Low52Weeks
}

// GetLow52WeeksOk returns a tuple with the Low52Weeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetLow52WeeksOk() (*string, bool) {
	if o == nil || IsNil(o.Low52Weeks) {
		return nil, false
	}
	return o.Low52Weeks, true
}

// HasLow52Weeks returns a boolean if a field has been set.
func (o *FundamentalsData) HasLow52Weeks() bool {
	if o != nil && !IsNil(o.Low52Weeks) {
		return true
	}

	return false
}

// SetLow52Weeks gets a reference to the given string and assigns it to the Low52Weeks field.
func (o *FundamentalsData) SetLow52Weeks(v string) {
	o.Low52Weeks = &v
}

// GetMarketCap returns the MarketCap field value if set, zero value otherwise.
func (o *FundamentalsData) GetMarketCap() string {
	if o == nil || IsNil(o.MarketCap) {
		var ret string
		return ret
	}
	return *o.MarketCap
}

// GetMarketCapOk returns a tuple with the MarketCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetMarketCapOk() (*string, bool) {
	if o == nil || IsNil(o.MarketCap) {
		return nil, false
	}
	return o.MarketCap, true
}

// HasMarketCap returns a boolean if a field has been set.
func (o *FundamentalsData) HasMarketCap() bool {
	if o != nil && !IsNil(o.MarketCap) {
		return true
	}

	return false
}

// SetMarketCap gets a reference to the given string and assigns it to the MarketCap field.
func (o *FundamentalsData) SetMarketCap(v string) {
	o.MarketCap = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *FundamentalsData) GetOpen() string {
	if o == nil || IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetOpenOk() (*string, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *FundamentalsData) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *FundamentalsData) SetOpen(v string) {
	o.Open = &v
}

// GetPeRatio returns the PeRatio field value if set, zero value otherwise.
func (o *FundamentalsData) GetPeRatio() string {
	if o == nil || IsNil(o.PeRatio) {
		var ret string
		return ret
	}
	return *o.PeRatio
}

// GetPeRatioOk returns a tuple with the PeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetPeRatioOk() (*string, bool) {
	if o == nil || IsNil(o.PeRatio) {
		return nil, false
	}
	return o.PeRatio, true
}

// HasPeRatio returns a boolean if a field has been set.
func (o *FundamentalsData) HasPeRatio() bool {
	if o != nil && !IsNil(o.PeRatio) {
		return true
	}

	return false
}

// SetPeRatio gets a reference to the given string and assigns it to the PeRatio field.
func (o *FundamentalsData) SetPeRatio(v string) {
	o.PeRatio = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *FundamentalsData) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsData) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *FundamentalsData) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *FundamentalsData) SetVolume(v string) {
	o.Volume = &v
}

func (o FundamentalsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundamentalsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageVolume) {
		toSerialize["average_volume"] = o.AverageVolume
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DividendYield) {
		toSerialize["dividend_yield"] = o.DividendYield
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.High52Weeks) {
		toSerialize["high_52_weeks"] = o.High52Weeks
	}
	if !IsNil(o.Instrument) {
		toSerialize["instrument"] = o.Instrument
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Low52Weeks) {
		toSerialize["low_52_weeks"] = o.Low52Weeks
	}
	if !IsNil(o.MarketCap) {
		toSerialize["market_cap"] = o.MarketCap
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.PeRatio) {
		toSerialize["pe_ratio"] = o.PeRatio
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableFundamentalsData struct {
	value *FundamentalsData
	isSet bool
}

func (v NullableFundamentalsData) Get() *FundamentalsData {
	return v.value
}

func (v *NullableFundamentalsData) Set(val *FundamentalsData) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsData) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsData(val *FundamentalsData) *NullableFundamentalsData {
	return &NullableFundamentalsData{value: val, isSet: true}
}

func (v NullableFundamentalsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



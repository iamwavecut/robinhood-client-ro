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
	"fmt"
)

// InvestmentObjective the model 'InvestmentObjective'
type InvestmentObjective string

// List of InvestmentObjective
const (
	CAP_PRESERVE_INVEST_OBJ InvestmentObjective = "cap_preserve_invest_obj"
	INCOME_INVEST_OBJ InvestmentObjective = "income_invest_obj"
	GROWTH_INVEST_OBJ InvestmentObjective = "growth_invest_obj"
	SPECULATION_INVEST_OBJ InvestmentObjective = "speculation_invest_obj"
	OTHER_INVEST_OBJ InvestmentObjective = "other_invest_obj"
)

// All allowed values of InvestmentObjective enum
var AllowedInvestmentObjectiveEnumValues = []InvestmentObjective{
	"cap_preserve_invest_obj",
	"income_invest_obj",
	"growth_invest_obj",
	"speculation_invest_obj",
	"other_invest_obj",
}

func (v *InvestmentObjective) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvestmentObjective(value)
	for _, existing := range AllowedInvestmentObjectiveEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvestmentObjective", value)
}

// NewInvestmentObjectiveFromValue returns a pointer to a valid InvestmentObjective
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvestmentObjectiveFromValue(v string) (*InvestmentObjective, error) {
	ev := InvestmentObjective(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvestmentObjective: valid values are %v", v, AllowedInvestmentObjectiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvestmentObjective) IsValid() bool {
	for _, existing := range AllowedInvestmentObjectiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvestmentObjective value
func (v InvestmentObjective) Ptr() *InvestmentObjective {
	return &v
}

type NullableInvestmentObjective struct {
	value *InvestmentObjective
	isSet bool
}

func (v NullableInvestmentObjective) Get() *InvestmentObjective {
	return v.value
}

func (v *NullableInvestmentObjective) Set(val *InvestmentObjective) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentObjective) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentObjective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentObjective(val *InvestmentObjective) *NullableInvestmentObjective {
	return &NullableInvestmentObjective{value: val, isSet: true}
}

func (v NullableInvestmentObjective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentObjective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


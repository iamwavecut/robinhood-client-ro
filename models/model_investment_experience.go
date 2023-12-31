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

// InvestmentExperience the model 'InvestmentExperience'
type InvestmentExperience string

// List of InvestmentExperience
const (
	EXTENSIVE_INVESTMENT_EXP InvestmentExperience = "extensive_investment_exp"
	GOOD_INVESTMENT_EXP InvestmentExperience = "good_investment_exp"
	LIMITED_INVESTMENT_EXP InvestmentExperience = "limited_investment_exp"
	NO_INVESTMENT_EXP InvestmentExperience = "no_investment_exp"
)

// All allowed values of InvestmentExperience enum
var AllowedInvestmentExperienceEnumValues = []InvestmentExperience{
	"extensive_investment_exp",
	"good_investment_exp",
	"limited_investment_exp",
	"no_investment_exp",
}

func (v *InvestmentExperience) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InvestmentExperience(value)
	for _, existing := range AllowedInvestmentExperienceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InvestmentExperience", value)
}

// NewInvestmentExperienceFromValue returns a pointer to a valid InvestmentExperience
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvestmentExperienceFromValue(v string) (*InvestmentExperience, error) {
	ev := InvestmentExperience(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InvestmentExperience: valid values are %v", v, AllowedInvestmentExperienceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvestmentExperience) IsValid() bool {
	for _, existing := range AllowedInvestmentExperienceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvestmentExperience value
func (v InvestmentExperience) Ptr() *InvestmentExperience {
	return &v
}

type NullableInvestmentExperience struct {
	value *InvestmentExperience
	isSet bool
}

func (v NullableInvestmentExperience) Get() *InvestmentExperience {
	return v.value
}

func (v *NullableInvestmentExperience) Set(val *InvestmentExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentExperience(val *InvestmentExperience) *NullableInvestmentExperience {
	return &NullableInvestmentExperience{value: val, isSet: true}
}

func (v NullableInvestmentExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


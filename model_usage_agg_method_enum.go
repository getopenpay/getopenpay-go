/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"fmt"
)

// UsageAggMethodEnum the model 'UsageAggMethodEnum'
type UsageAggMethodEnum string

// List of UsageAggMethodEnum
const (
	USAGEAGGMETHODENUM_SUM UsageAggMethodEnum = "sum"
	USAGEAGGMETHODENUM_LAST_DURING_PERIOD UsageAggMethodEnum = "last_during_period"
	USAGEAGGMETHODENUM_LAST_EVER UsageAggMethodEnum = "last_ever"
	USAGEAGGMETHODENUM_MAX UsageAggMethodEnum = "max"
)

// All allowed values of UsageAggMethodEnum enum
var AllowedUsageAggMethodEnumEnumValues = []UsageAggMethodEnum{
	"sum",
	"last_during_period",
	"last_ever",
	"max",
}

func (v *UsageAggMethodEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageAggMethodEnum(value)
	for _, existing := range AllowedUsageAggMethodEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageAggMethodEnum", value)
}

// NewUsageAggMethodEnumFromValue returns a pointer to a valid UsageAggMethodEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageAggMethodEnumFromValue(v string) (*UsageAggMethodEnum, error) {
	ev := UsageAggMethodEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageAggMethodEnum: valid values are %v", v, AllowedUsageAggMethodEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageAggMethodEnum) IsValid() bool {
	for _, existing := range AllowedUsageAggMethodEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageAggMethodEnum value
func (v UsageAggMethodEnum) Ptr() *UsageAggMethodEnum {
	return &v
}

type NullableUsageAggMethodEnum struct {
	value *UsageAggMethodEnum
	isSet bool
}

func (v NullableUsageAggMethodEnum) Get() *UsageAggMethodEnum {
	return v.value
}

func (v *NullableUsageAggMethodEnum) Set(val *UsageAggMethodEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAggMethodEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAggMethodEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAggMethodEnum(val *UsageAggMethodEnum) *NullableUsageAggMethodEnum {
	return &NullableUsageAggMethodEnum{value: val, isSet: true}
}

func (v NullableUsageAggMethodEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAggMethodEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

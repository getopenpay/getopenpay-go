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

// BillingSchemeEnum the model 'BillingSchemeEnum'
type BillingSchemeEnum string

// List of BillingSchemeEnum
const (
	BILLINGSCHEMEENUM_PER_UNIT BillingSchemeEnum = "per_unit"
	BILLINGSCHEMEENUM_TIERED BillingSchemeEnum = "tiered"
)

// All allowed values of BillingSchemeEnum enum
var AllowedBillingSchemeEnumEnumValues = []BillingSchemeEnum{
	"per_unit",
	"tiered",
}

func (v *BillingSchemeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingSchemeEnum(value)
	for _, existing := range AllowedBillingSchemeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingSchemeEnum", value)
}

// NewBillingSchemeEnumFromValue returns a pointer to a valid BillingSchemeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingSchemeEnumFromValue(v string) (*BillingSchemeEnum, error) {
	ev := BillingSchemeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingSchemeEnum: valid values are %v", v, AllowedBillingSchemeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingSchemeEnum) IsValid() bool {
	for _, existing := range AllowedBillingSchemeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingSchemeEnum value
func (v BillingSchemeEnum) Ptr() *BillingSchemeEnum {
	return &v
}

type NullableBillingSchemeEnum struct {
	value *BillingSchemeEnum
	isSet bool
}

func (v NullableBillingSchemeEnum) Get() *BillingSchemeEnum {
	return v.value
}

func (v *NullableBillingSchemeEnum) Set(val *BillingSchemeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingSchemeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingSchemeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingSchemeEnum(val *BillingSchemeEnum) *NullableBillingSchemeEnum {
	return &NullableBillingSchemeEnum{value: val, isSet: true}
}

func (v NullableBillingSchemeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingSchemeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

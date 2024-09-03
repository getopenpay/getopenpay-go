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

// NonPciIntegrationEnum the model 'NonPciIntegrationEnum'
type NonPciIntegrationEnum string

// List of NonPciIntegrationEnum
const (
	NONPCIINTEGRATIONENUM_CHARTMOGUL NonPciIntegrationEnum = "chartmogul"
	NONPCIINTEGRATIONENUM_INTERCOM NonPciIntegrationEnum = "intercom"
)

// All allowed values of NonPciIntegrationEnum enum
var AllowedNonPciIntegrationEnumEnumValues = []NonPciIntegrationEnum{
	"chartmogul",
	"intercom",
}

func (v *NonPciIntegrationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NonPciIntegrationEnum(value)
	for _, existing := range AllowedNonPciIntegrationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NonPciIntegrationEnum", value)
}

// NewNonPciIntegrationEnumFromValue returns a pointer to a valid NonPciIntegrationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNonPciIntegrationEnumFromValue(v string) (*NonPciIntegrationEnum, error) {
	ev := NonPciIntegrationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NonPciIntegrationEnum: valid values are %v", v, AllowedNonPciIntegrationEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NonPciIntegrationEnum) IsValid() bool {
	for _, existing := range AllowedNonPciIntegrationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NonPciIntegrationEnum value
func (v NonPciIntegrationEnum) Ptr() *NonPciIntegrationEnum {
	return &v
}

type NullableNonPciIntegrationEnum struct {
	value *NonPciIntegrationEnum
	isSet bool
}

func (v NullableNonPciIntegrationEnum) Get() *NonPciIntegrationEnum {
	return v.value
}

func (v *NullableNonPciIntegrationEnum) Set(val *NonPciIntegrationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableNonPciIntegrationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableNonPciIntegrationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonPciIntegrationEnum(val *NonPciIntegrationEnum) *NullableNonPciIntegrationEnum {
	return &NullableNonPciIntegrationEnum{value: val, isSet: true}
}

func (v NullableNonPciIntegrationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonPciIntegrationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// CheckoutMode the model 'CheckoutMode'
type CheckoutMode string

// List of CheckoutMode
const (
	CHECKOUTMODE_PAYMENT CheckoutMode = "payment"
	CHECKOUTMODE_SETUP CheckoutMode = "setup"
	CHECKOUTMODE_SUBSCRIPTION CheckoutMode = "subscription"
)

// All allowed values of CheckoutMode enum
var AllowedCheckoutModeEnumValues = []CheckoutMode{
	"payment",
	"setup",
	"subscription",
}

func (v *CheckoutMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CheckoutMode(value)
	for _, existing := range AllowedCheckoutModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CheckoutMode", value)
}

// NewCheckoutModeFromValue returns a pointer to a valid CheckoutMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCheckoutModeFromValue(v string) (*CheckoutMode, error) {
	ev := CheckoutMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CheckoutMode: valid values are %v", v, AllowedCheckoutModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CheckoutMode) IsValid() bool {
	for _, existing := range AllowedCheckoutModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CheckoutMode value
func (v CheckoutMode) Ptr() *CheckoutMode {
	return &v
}

type NullableCheckoutMode struct {
	value *CheckoutMode
	isSet bool
}

func (v NullableCheckoutMode) Get() *CheckoutMode {
	return v.value
}

func (v *NullableCheckoutMode) Set(val *CheckoutMode) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutMode) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutMode(val *CheckoutMode) *NullableCheckoutMode {
	return &NullableCheckoutMode{value: val, isSet: true}
}

func (v NullableCheckoutMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


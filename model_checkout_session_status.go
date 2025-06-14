/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionStatus the model 'CheckoutSessionStatus'
type CheckoutSessionStatus string

// List of CheckoutSessionStatus
const (
	CHECKOUTSESSIONSTATUS_OPEN CheckoutSessionStatus = "open"
	CHECKOUTSESSIONSTATUS_EXPIRED CheckoutSessionStatus = "expired"
	CHECKOUTSESSIONSTATUS_COMPLETE CheckoutSessionStatus = "complete"
)

// All allowed values of CheckoutSessionStatus enum
var AllowedCheckoutSessionStatusEnumValues = []CheckoutSessionStatus{
	"open",
	"expired",
	"complete",
}

func (v *CheckoutSessionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CheckoutSessionStatus(value)
	for _, existing := range AllowedCheckoutSessionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CheckoutSessionStatus", value)
}

// NewCheckoutSessionStatusFromValue returns a pointer to a valid CheckoutSessionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCheckoutSessionStatusFromValue(v string) (*CheckoutSessionStatus, error) {
	ev := CheckoutSessionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CheckoutSessionStatus: valid values are %v", v, AllowedCheckoutSessionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CheckoutSessionStatus) IsValid() bool {
	for _, existing := range AllowedCheckoutSessionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CheckoutSessionStatus value
func (v CheckoutSessionStatus) Ptr() *CheckoutSessionStatus {
	return &v
}

type NullableCheckoutSessionStatus struct {
	value *CheckoutSessionStatus
	isSet bool
}

func (v NullableCheckoutSessionStatus) Get() *CheckoutSessionStatus {
	return v.value
}

func (v *NullableCheckoutSessionStatus) Set(val *CheckoutSessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionStatus(val *CheckoutSessionStatus) *NullableCheckoutSessionStatus {
	return &NullableCheckoutSessionStatus{value: val, isSet: true}
}

func (v NullableCheckoutSessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// CustomerStatus the model 'CustomerStatus'
type CustomerStatus string

// List of CustomerStatus
const (
	CUSTOMERSTATUS_ACTIVE CustomerStatus = "active"
	CUSTOMERSTATUS_INACTIVE CustomerStatus = "inactive"
	CUSTOMERSTATUS_TRIALING CustomerStatus = "trialing"
	CUSTOMERSTATUS_CANCELLED CustomerStatus = "cancelled"
)

// All allowed values of CustomerStatus enum
var AllowedCustomerStatusEnumValues = []CustomerStatus{
	"active",
	"inactive",
	"trialing",
	"cancelled",
}

func (v *CustomerStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerStatus(value)
	for _, existing := range AllowedCustomerStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerStatus", value)
}

// NewCustomerStatusFromValue returns a pointer to a valid CustomerStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerStatusFromValue(v string) (*CustomerStatus, error) {
	ev := CustomerStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerStatus: valid values are %v", v, AllowedCustomerStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerStatus) IsValid() bool {
	for _, existing := range AllowedCustomerStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerStatus value
func (v CustomerStatus) Ptr() *CustomerStatus {
	return &v
}

type NullableCustomerStatus struct {
	value *CustomerStatus
	isSet bool
}

func (v NullableCustomerStatus) Get() *CustomerStatus {
	return v.value
}

func (v *NullableCustomerStatus) Set(val *CustomerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerStatus(val *CustomerStatus) *NullableCustomerStatus {
	return &NullableCustomerStatus{value: val, isSet: true}
}

func (v NullableCustomerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// BillingReasonEnum the model 'BillingReasonEnum'
type BillingReasonEnum string

// List of BillingReasonEnum
const (
	BILLINGREASONENUM_MANUAL BillingReasonEnum = "manual"
	BILLINGREASONENUM_SUBSCRIPTION_CREATE BillingReasonEnum = "subscription_create"
	BILLINGREASONENUM_SUBSCRIPTION_CYCLE BillingReasonEnum = "subscription_cycle"
	BILLINGREASONENUM_SUBSCRIPTION_THRESHOLD BillingReasonEnum = "subscription_threshold"
	BILLINGREASONENUM_SUBSCRIPTION_UPDATE BillingReasonEnum = "subscription_update"
	BILLINGREASONENUM_UPCOMING BillingReasonEnum = "upcoming"
)

// All allowed values of BillingReasonEnum enum
var AllowedBillingReasonEnumEnumValues = []BillingReasonEnum{
	"manual",
	"subscription_create",
	"subscription_cycle",
	"subscription_threshold",
	"subscription_update",
	"upcoming",
}

func (v *BillingReasonEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingReasonEnum(value)
	for _, existing := range AllowedBillingReasonEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingReasonEnum", value)
}

// NewBillingReasonEnumFromValue returns a pointer to a valid BillingReasonEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingReasonEnumFromValue(v string) (*BillingReasonEnum, error) {
	ev := BillingReasonEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingReasonEnum: valid values are %v", v, AllowedBillingReasonEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingReasonEnum) IsValid() bool {
	for _, existing := range AllowedBillingReasonEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingReasonEnum value
func (v BillingReasonEnum) Ptr() *BillingReasonEnum {
	return &v
}

type NullableBillingReasonEnum struct {
	value *BillingReasonEnum
	isSet bool
}

func (v NullableBillingReasonEnum) Get() *BillingReasonEnum {
	return v.value
}

func (v *NullableBillingReasonEnum) Set(val *BillingReasonEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingReasonEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingReasonEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingReasonEnum(val *BillingReasonEnum) *NullableBillingReasonEnum {
	return &NullableBillingReasonEnum{value: val, isSet: true}
}

func (v NullableBillingReasonEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingReasonEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


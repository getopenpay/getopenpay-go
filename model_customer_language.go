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

// CustomerLanguage the model 'CustomerLanguage'
type CustomerLanguage string

// List of CustomerLanguage
const (
	CUSTOMERLANGUAGE_EN CustomerLanguage = "en"
	CUSTOMERLANGUAGE_EN_US CustomerLanguage = "en_us"
	CUSTOMERLANGUAGE_PT CustomerLanguage = "pt"
	CUSTOMERLANGUAGE_ES CustomerLanguage = "es"
	CUSTOMERLANGUAGE_DE CustomerLanguage = "de"
	CUSTOMERLANGUAGE_FR CustomerLanguage = "fr"
	CUSTOMERLANGUAGE_IT CustomerLanguage = "it"
	CUSTOMERLANGUAGE_NL CustomerLanguage = "nl"
	CUSTOMERLANGUAGE_RO CustomerLanguage = "ro"
	CUSTOMERLANGUAGE_ZH CustomerLanguage = "zh"
	CUSTOMERLANGUAGE_ZH_HK CustomerLanguage = "zh_hk"
	CUSTOMERLANGUAGE_ZH_TW CustomerLanguage = "zh_tw"
	CUSTOMERLANGUAGE_RU CustomerLanguage = "ru"
	CUSTOMERLANGUAGE_JA CustomerLanguage = "ja"
	CUSTOMERLANGUAGE_KO CustomerLanguage = "ko"
	CUSTOMERLANGUAGE_AR CustomerLanguage = "ar"
	CUSTOMERLANGUAGE_TR CustomerLanguage = "tr"
	CUSTOMERLANGUAGE_HI CustomerLanguage = "hi"
)

// All allowed values of CustomerLanguage enum
var AllowedCustomerLanguageEnumValues = []CustomerLanguage{
	"en",
	"en_us",
	"pt",
	"es",
	"de",
	"fr",
	"it",
	"nl",
	"ro",
	"zh",
	"zh_hk",
	"zh_tw",
	"ru",
	"ja",
	"ko",
	"ar",
	"tr",
	"hi",
}

func (v *CustomerLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerLanguage(value)
	for _, existing := range AllowedCustomerLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerLanguage", value)
}

// NewCustomerLanguageFromValue returns a pointer to a valid CustomerLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerLanguageFromValue(v string) (*CustomerLanguage, error) {
	ev := CustomerLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerLanguage: valid values are %v", v, AllowedCustomerLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerLanguage) IsValid() bool {
	for _, existing := range AllowedCustomerLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerLanguage value
func (v CustomerLanguage) Ptr() *CustomerLanguage {
	return &v
}

type NullableCustomerLanguage struct {
	value *CustomerLanguage
	isSet bool
}

func (v NullableCustomerLanguage) Get() *CustomerLanguage {
	return v.value
}

func (v *NullableCustomerLanguage) Set(val *CustomerLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerLanguage(val *CustomerLanguage) *NullableCustomerLanguage {
	return &NullableCustomerLanguage{value: val, isSet: true}
}

func (v NullableCustomerLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


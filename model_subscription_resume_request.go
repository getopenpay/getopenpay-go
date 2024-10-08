/*
OpenPay API

super charge your subscription management.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"encoding/json"
)

// checks if the SubscriptionResumeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionResumeRequest{}

// SubscriptionResumeRequest struct for SubscriptionResumeRequest
type SubscriptionResumeRequest struct {
	ProrationBehavior *ProrationEnum `json:"proration_behavior,omitempty"`
}

// NewSubscriptionResumeRequest instantiates a new SubscriptionResumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionResumeRequest() *SubscriptionResumeRequest {
	this := SubscriptionResumeRequest{}
	var prorationBehavior ProrationEnum = PRORATIONENUM_ALWAYS_INVOICE
	this.ProrationBehavior = &prorationBehavior
	return &this
}

// NewSubscriptionResumeRequestWithDefaults instantiates a new SubscriptionResumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionResumeRequestWithDefaults() *SubscriptionResumeRequest {
	this := SubscriptionResumeRequest{}
	var prorationBehavior ProrationEnum = PRORATIONENUM_ALWAYS_INVOICE
	this.ProrationBehavior = &prorationBehavior
	return &this
}

// GetProrationBehavior returns the ProrationBehavior field value if set, zero value otherwise.
func (o *SubscriptionResumeRequest) GetProrationBehavior() ProrationEnum {
	if o == nil || IsNil(o.ProrationBehavior) {
		var ret ProrationEnum
		return ret
	}
	return *o.ProrationBehavior
}

// GetProrationBehaviorOk returns a tuple with the ProrationBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResumeRequest) GetProrationBehaviorOk() (*ProrationEnum, bool) {
	if o == nil || IsNil(o.ProrationBehavior) {
		return nil, false
	}
	return o.ProrationBehavior, true
}

// HasProrationBehavior returns a boolean if a field has been set.
func (o *SubscriptionResumeRequest) HasProrationBehavior() bool {
	if o != nil && !IsNil(o.ProrationBehavior) {
		return true
	}

	return false
}

// SetProrationBehavior gets a reference to the given ProrationEnum and assigns it to the ProrationBehavior field.
func (o *SubscriptionResumeRequest) SetProrationBehavior(v ProrationEnum) {
	o.ProrationBehavior = &v
}

func (o SubscriptionResumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionResumeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProrationBehavior) {
		toSerialize["proration_behavior"] = o.ProrationBehavior
	}
	return toSerialize, nil
}

type NullableSubscriptionResumeRequest struct {
	value *SubscriptionResumeRequest
	isSet bool
}

func (v NullableSubscriptionResumeRequest) Get() *SubscriptionResumeRequest {
	return v.value
}

func (v *NullableSubscriptionResumeRequest) Set(val *SubscriptionResumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResumeRequest(val *SubscriptionResumeRequest) *NullableSubscriptionResumeRequest {
	return &NullableSubscriptionResumeRequest{value: val, isSet: true}
}

func (v NullableSubscriptionResumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


